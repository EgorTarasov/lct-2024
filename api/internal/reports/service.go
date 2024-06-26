package reports

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	pkgs3 "github.com/EgorTarasov/lct-2024/api/pkg/s3"
	"github.com/minio/minio-go/v7"
	"github.com/xuri/excelize/v2"
	"go.opentelemetry.io/otel/trace"
)

type service struct {
	reports  ReportRepo
	s3       *pkgs3.S3
	s3Bucket string
	tracer   trace.Tracer
}

type ReportRepo interface {
	CreateReport(ctx context.Context, reportCreate CreateReport) (int64, error)
	GetReportByID(ctx context.Context, id int64) (Report, error)
	GetUsersReports(ctx context.Context, userID int64) ([]Report, error)
	GetPredictionCount(ctx context.Context, startDate, endDate time.Time) (int64, error)
	GetPredictionRecords(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]PredictionRecord, error)
	AddS3Key(ctx context.Context, s3Key string, id int64) error
}

// NewService конструктор сервиса для работы с отчетами.
func NewService(reports ReportRepo, s3 *pkgs3.S3, tracer trace.Tracer) *service {
	return &service{
		reports:  reports,
		tracer:   tracer,
		s3:       s3,
		s3Bucket: "uploads",
	}
}

func (s *service) CreateReport(ctx context.Context, reportCreate CreateReport) (int64, error) {
	ctx, span := s.tracer.Start(ctx, "reports.CreateReport")
	defer span.End()
	id, err := s.reports.CreateReport(ctx, reportCreate)
	if err != nil {
		return 0, err
	}
	// start go routine for creating and uploading report in excel to s3
	go s.assembleReport(ctx, reportCreate, id)

	return id, nil
}

func (s *service) GetReportByID(ctx context.Context, id int64) (Report, error) {
	ctx, span := s.tracer.Start(ctx, "reports.GetReportStatus")
	defer span.End()
	report, err := s.reports.GetReportByID(ctx, id)
	if err != nil {
		return Report{}, err
	}
	url, err := s.s3.PresignedGetObject(ctx, s.s3Bucket, report.S3Key, time.Hour*72, nil)
	if err != nil {
		return Report{}, err
	}
	url.Host = "s3.larek.tech"
	url.Scheme = "https"
	report.DownloadUrl = url.String()

	return report, nil
}

func (s *service) GetReports(ctx context.Context, userID int64) ([]Report, error) {
	ctx, span := s.tracer.Start(ctx, "reports.GetReports")
	defer span.End()
	reports, err := s.reports.GetUsersReports(ctx, userID)
	if err != nil {
		return nil, err
	}
	for i := range reports {
		url, err := s.s3.PresignedGetObject(ctx, s.s3Bucket, reports[i].S3Key, time.Hour*72, nil)
		if err != nil {
			return nil, err
		}
		url.Host = "s3.larek.tech"
		url.Scheme = "https"
		reports[i].DownloadUrl = url.String()

	}
	return reports, nil
}

func (s *service) assembleReport(ctx context.Context, reportCreate CreateReport, id int64) {
	// create excel file
	f := excelize.NewFile()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return
	}
	overalCnt, err := s.reports.GetPredictionCount(ctx, reportCreate.StartDate, reportCreate.EndDate)
	if err != nil {
		return
	}

	// put data to excel by bathes of 1000 records
	startYear, startMonth, startDay := reportCreate.StartDate.Date()
	endYear, endMonth, endDay := reportCreate.EndDate.Date()
	sheetName := fmt.Sprintf("Предсказания_%v.%v.%v-%v.%v.%v_%d",
		startYear, startMonth, startDay,
		endYear, endMonth, endDay,
		id)
	var batchSize int64 = 1000
	// set headers as json tag in PredictionRecord
	f.SetActiveSheet(index)
	_ = f.SetCellValue(sheetName, "A1", "ID")
	_ = f.SetCellValue(sheetName, "B1", "Дата предсказания")
	_ = f.SetCellValue(sheetName, "C1", "p2 <= 0")
	_ = f.SetCellValue(sheetName, "D1", "P2LessThanOrEqualTo0")
	_ = f.SetCellValue(sheetName, "E1", "T1LessThanMin")

	totalBatches := (overalCnt + batchSize - 1) / batchSize
	for batch := int64(0); batch < totalBatches; batch++ {
		start := batch * batchSize
		records, err := s.reports.GetPredictionRecords(ctx, reportCreate.StartDate, reportCreate.EndDate, int(batchSize), int(start))
		if err != nil {
			return
		}

		for j, record := range records {
			row := start + int64(j) + 2
			_ = f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), record.ID)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.PredictionDate)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.Unom)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.P1LessThanOrEqualTo0)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("3%d", row), record.P2LessThanOrEqualTo0)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), record.T1LessThanMin)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), record.T1GreaterThanMax)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), record.No)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), record.LackOfHeatingInTheHouse)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), record.PipeLeakInTheEntrance)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), record.StrongLeakInTheHeatingSystem)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), record.TemperatureInTheApartmentBelowTheStandard)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), record.TemperatureInPublicAreasBelowTheStandard)
			_ = f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), record.LeakInTheHeatingSystem)
		}
	}

	fileName := fmt.Sprintf("%s.xlsx", sheetName)
	if err := f.SaveAs(fileName); err != nil {
		return
	}
	_ = f.Close()
	// upload to s3
	s3Key := fmt.Sprintf("reports/%s.xlsx", sheetName)

	rawBytes, err := os.ReadFile(fileName)
	defer func() {
		_ = os.Remove(fileName)
	}()
	if err != nil {
		return
	}
	reader := bytes.NewReader(rawBytes)

	_, err = s.s3.PutObject(ctx, s.s3Bucket, s3Key, reader, int64(len(rawBytes)), minio.PutObjectOptions{ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"})
	if err != nil {
		return
	}
	if err := s.reports.AddS3Key(ctx, s3Key, id); err != nil {
		return
	}
	// upload to s3
}
