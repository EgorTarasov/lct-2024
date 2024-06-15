package models

import (
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/search/models"
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

	Consumer         *models.MKDConsumer      `json:"mkdConsumer"`
	DispatchServices *models.DispatchServices `json:"dispatchServices"`
	HeatingPoint     *models.HeatingPoint     `json:"heatingPoint"`

	HeatingGraph Graph `json:"graph"`
}
