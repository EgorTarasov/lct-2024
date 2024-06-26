package reports

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type ReportState string

const (
	ReportStateInProgress ReportState = "in_progress"
	ReportStateReady      ReportState = "ready"
)

// Report of incidents for given time period.
type Report struct {
	ID          int64       `db:"id" json:"id"`
	UserID      int64       `db:"user_id" json:"userID"`
	Title       string      `db:"title" json:"title"`
	StartDate   time.Time   `db:"start_date" json:"startDate"`
	EndDate     time.Time   `db:"end_date" json:"endDate"`
	S3Key       string      `db:"s3_key" json:"-"`
	Status      ReportState `db:"status" json:"status"`
	CreatedAt   time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time   `db:"updated_at" json:"-"`
	DeletedAt   null.Time   `db:"deleted_at" json:"-"`
	DownloadUrl string      `json:"downloadUrl"`
}

// CreateReport request for creating report.
type CreateReport struct {
	Title     string    `json:"title" `
	UserID    int64     `json:"-"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type PredictionRecord struct {
	ID                                        int64     `db:"id" json:"id"`
	Unom                                      int64     `db:"unom" json:"unom"`
	PredictionDate                            time.Time `db:"prediction_date" json:"predictionDate"`
	P1LessThanOrEqualTo0                      float64   `db:"p1_less_than_or_equal_to_0" json:"P1 <= 0"`
	P2LessThanOrEqualTo0                      float64   `db:"p2_less_than_or_equal_to_0" json:"P2 <= 0"`
	T1LessThanMin                             float64   `db:"t1_less_than_min" json:"T1 < min"`
	T1GreaterThanMax                          float64   `db:"t1_greater_than_max" json:"T1 > max"`
	No                                        float64   `db:"no" json:"Нет"`
	LackOfHeatingInTheHouse                   float64   `db:"lack_of_heating_in_the_house" json:"Отсутствие отопления в доме"`
	PipeLeakInTheEntrance                     float64   `db:"pipe_leak_in_the_entrance" json:"Протечка труб в подъезде"`
	StrongLeakInTheHeatingSystem              float64   `db:"strong_leak_in_the_heating_system" json:"Сильная течь в системе отопления"`
	TemperatureInTheApartmentBelowTheStandard float64   `db:"temperature_in_the_apartment_below_the_standard" json:"Температура в квартире ниже нормативной"`
	TemperatureInPublicAreasBelowTheStandard  float64   `db:"temperature_in_public_areas_below_the_standard" json:"Температура в помещении общего пользования ниже нормативной"`
	LeakInTheHeatingSystem                    float64   `db:"leak_in_the_heating_system" json:"Течь в системе отопления"`
}
