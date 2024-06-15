package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// https://recharts.org/en-US/examples

type GraphDataPoint struct {
	Temp       float64 `json:"temp"`
	TimeString int     `json:"time"` // время (1 час, 3 часа и т.п.)
}

type Graph struct {
	Name   string           `json:"name"`
	Points []GraphDataPoint `json:"values"`
}

type Incident struct {
	ID int64 `db:"id" json:"id"`

	OpenedAt time.Time `db:"opened_at" json:"openedAt"`
	ClosedAt null.Time `db:"closed_at" json:"closedAt"`
	Title    string    `db:"title" json:"title"`
	Status   string    `db:"status" json:"status"`
	Priority string    `db:"priority" json:"priority"`
	Unom     int64     `db:"unom" json:"unom"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`

	Consumer         *MKDConsumer      `json:"mkdConsumer"`
	DispatchServices *DispatchServices `json:"dispatchServices"`
	HeatingPoint     *HeatingPoint     `json:"heatingPoint"`

	HeatingGraph   Graph         `json:"graph"`
	Measurements   []Measurement `json:"measurements"`
	RelatedObjects UnomResult    `json:"relatedObjects"`
}

type Measurement struct {
	Unom                      int64     `bson:"unom" json:"unom"`
	EnergyConsumption         float64   `bson:"thermal_energy_consumption" json:"energyConsumption"`
	Datetime                  time.Time `bson:"datetime" json:"datetime"`
	Address                   string    `bson:"address" json:"address"`
	HeatingCircuit            string    `bson:"central_heating_circuit" json:"heatingPoint"`
	Consumer                  string    `bson:"consumer" json:"consumer"` // Бюджет / МКД
	CoolantInput              float64   `bson:"coolant_returned_temperature" json:"coolantInput"`
	CoolantOutput             float64   `bson:"coolant_returned_to_central_system" json:"coolantOutput"`
	CoolantSuppliedDifference float64   `bson:"coolant_supplied_difference_leakage" json:"coolantSuppliedDifferenceLeakage"`
	MeasurementDate           string    `bson:"measurment_datetime" json:"measurementDate"`
	MeasuringDeviceHours      float64   `bson:"measuring_device_hours" json:"measuringDeviceHours"`
	MeasuringDeviceBrand      string    `bson:"metering_device_brand" json:"measuringDeviceBrand"`
}
