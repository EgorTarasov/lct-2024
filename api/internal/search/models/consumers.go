package models

// Filter фильтр для поиска по объектам.
type Filter struct {
	Name   string   `json:"filterName"`
	Values []string `json:"values"`
}

type HeatingPoint struct {
	BalanceHolder     string      `bson:"balance_holder" json:"balanceHolder"`
	Address           string      `bson:"heating_point_address" json:"-"`
	ConsumerAddress   string      `bson:"consumer_address" json:"consumerAddress"`
	Source            string      `bson:"heating_point_src" json:""`
	Number            string      `bson:"heating_point_number" json:"number"`
	Type              string      `bson:"heating_point_type" json:"type"`
	LocationType      string      `bson:"heating_point_location_type" json:"locationType"`
	District          string      `bson:"municipal_district" json:"district"`
	CommissioningDate interface{} `bson:"commissioning_date" json:"commissioningDate"`
}

type HeatingPointDTO struct {
	HeatingPoint
	// models.Address
}
