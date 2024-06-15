package models

import (
	"time"
)

type PredictionResult struct {
	Unom          int64     `json:"unom"`          // (unom;date)
	Date          time.Time `json:"date"`          // (unom;date)
	P1            float64   `json:"p1"`            // P1 <= 0
	P2            float64   `json:"p2"`            // P2 <= 0
	T1            float64   `json:"t1"`            // T1 < min
	T2            float64   `json:"t2"`            // T1 > max
	No            float64   `json:"no"`            // Нет
	NoHeating     float64   `json:"noHeating"`     // Отсутствие отопления в доме
	Leak          float64   `json:"leak"`          // Протечка труб в подъезде
	StrongLeak    float64   `json:"strongLeak"`    // Сильная течь в системе отопления
	TempLow       float64   `json:"tempLow"`       // Температура в квартире ниже нормативной
	TempLowCommon float64   `json:"tempLowCommon"` // Температура в помещении общего пользования ниже нормативной
	LeakSystem    float64   `json:"leakSystem"`    // Течь в системе отопления
}
