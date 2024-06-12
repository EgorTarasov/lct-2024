package models

import (
	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
)

// Filter фильтр для поиска по объектам.
type Filter struct {
	Name   string   `json:"filterName"`
	Values []string `json:"values"`
}

type HeatingPoint struct {
	BalanceHolder       string         `bson:"balance_holder" json:"balanceHolder"`
	Address             string         `bson:"heating_point_address" json:"-"`
	ConsumerAddress     shared.Address `bson:"consumer_full_address" json:"consumerAddress"`
	Source              string         `bson:"heating_point_src" json:""`
	Number              string         `bson:"heating_point_number" json:"number"`
	Type                string         `bson:"heating_point_type" json:"type"`
	LocationType        string         `bson:"heating_point_location_type" json:"locationType"`
	District            string         `bson:"municipal_district" json:"district"`
	CommissioningDate   interface{}    `bson:"commissioning_date" json:"commissioningDate"`
	HeatingPointAddress shared.Address `bson:"heating_point_full_address" json:"heatingPointAddress"`
}

// Event событие которое произошло с объектом.
type Event struct {
}

// HeatingPointDTO DTO для объекта теплоснабжения.
type HeatingPointDTO struct {
	HeatingPoint

	// models.Address
}

type StateConsumer struct {
	AdmDistrict       string      `bson:"administrative_district" json:"admDistrict"`
	Material          string      `bson:"material" json:"material"`
	MunicupalDistrict string      `bson:"municipal_district" json:"municupalDistrict"`
	Purpose           string      `bson:"purpose" json:"purpose"`
	Type              string      `bson:"type" json:"type"`
	Area              interface{} `bson:"property_total_area" json:"area"`
	Floors            interface{} `bson:"property_floors" json:"floors"`
	PropertyClass     string      `bson:"property_class" json:"propertyClass"`
	Unom              int64       `bson:"unom" json:"unom"`
}

// MKDConsumer тип потребителя многоквартирного дома (МКД).
//
// 0   address                 4372 non-null   object
// 1   external_id             4372 non-null   int64
// 2   bti_address             4372 non-null   object
// 3   unom                    4372 non-null   int64
// 4   county                  4372 non-null   object
// 5   district                4372 non-null   object
// 6   project_series          4372 non-null   object
// 7   floors                  4372 non-null   int64
// 8   entrances               4372 non-null   int64
// 9   apartments              4372 non-null   int64
// 10  total_area              4372 non-null   float64
// 11  total_living_area       4372 non-null   float64
// 12  total_non_living_area   4372 non-null   object
// 13  Depreciation            4372 non-null   float64
// 14  wall_material           4372 non-null   int64
// 15  emergency_feature       4372 non-null   int64
// 16  passenger_elevator      4372 non-null   int64
// 17  service_elevator        4372 non-null   int64
// 18  roof_cleaning_sequence  4372 non-null   int64
// 19  roof_materials          4372 non-null   int64
// 20  types_of_housing_stock  4372 non-null   int64
// 21  mkd_state               4372 non-null   int64

type MKDConsumer struct {
	Deprecation          float64     `bson:"Depreciation" json:"deprecation"`
	Address              string      `bson:"address" json:"address"`
	Apartments           int64       `bson:"apartments" json:"apartments"`
	FullAddress          string      `bson:"bti_address" json:"fullAddress"`
	Country              string      `bson:"county" json:"country"`
	District             string      `bson:"district" json:"district"`
	Feature              int64       `bson:"emergency_feature" json:"feature"`
	Entrances            int64       `bson:"entrances" json:"entrances"`
	ExternalID           int64       `bson:"external_id" json:"externalID"`
	Floors               int64       `bson:"floors" json:"floors"`
	MKDState             int64       `bson:"mkd_state" json:"state"`
	PassengerElevator    int64       `bson:"passenger_elevator" json:"PassengerElevator"`
	ProjectSeries        interface{} `bson:"project_series" json:"projectSeries"`
	RoofCleaningSequence int64       `bson:"roof_cleaning_sequence" json:"roofCleaningSequence"`
	TotalArea            float64     `bson:"total_area" json:"totalArea"`
	Unom                 int64       `bson:"unom" json:"unom"`
	WallMaterial         int64       `bson:"wall_material" json:"wallMaterial"`
	RoofMaterials        int64       `bson:"roof_materials" json:"roofMaterials"`
	TotalLivingArea      float64     `bson:"total_living_area" json:"totalLivingArea"`
	ServiceElevator      int64       `bson:"service_elevator" json:"serviceElevator"`
	TotalNonLivingArea   interface{} `bson:"total_non_living_area" json:"totalNonLivingArea"`
	TypesOfHousingStock  int64       `bson:"types_of_housing_stock" json:"typesOfHousingStock"`
}

// DispatchServices данные о диспетчерской службе.
type DispatchServices struct {
	Address              string `bson:"address" json:"address"`
	HeatDispatcherNumber string `bson:"chp" json:"heatDispatcherNumber"`
	ConsumerName         string `bson:"consumer" json:"consumer"`
	ConsumerGroup        string `bson:"consumer_group" json:"consumerGroup"`
	Country              string `bson:"country" json:"country"`
	DispatchAddress      string `bson:"dispatch_address" json:"centerAddress"`
	DispatchNumber       string `bson:"dispatch_number" json:"dispatchNumber"`
	ExternalID           string `bson:"external_id" json:"externalID"`
	Unom                 int64  `bson:"unom" json:"unom"`
	ConsumerFullAddress  string `bson:"Полный адрес" json:"consumerFullAddress"`
}
