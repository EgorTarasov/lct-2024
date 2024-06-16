package models

// StateProperty  отображение записей в state_property в mongodb.
type StateProperty struct {
	Unom              int64       `bson:"unom" `
	PropertyClass     string      `bson:"property_class" json:"class"`
	PropertyFloors    interface{} `bson:"property_floors" json:"floors"`
	PropertyType      string      `bson:"property_type" json:"type"`
	PropertyTotalArea interface{} `bson:"property_total_area" json:"area"`
	PropertyFeature   string      `bson:"property_feature" json:"feature"`
	Purpose           string      `bson:"purpose" json:"purpose"`
}

// StatePropertySearchResult результат поиска с данными для отображения на карте.
type StatePropertySearchResult struct {
	StateProperty
	Address
}
