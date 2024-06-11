package models

// TODO: move shared to one module

// Address godoc
//
//	минимальные данные об адресе для отображения на карте
//
// Общие данные об адресе.
type Address struct {
	// Уникальный номер объекта недвижимости.
	Unom int64 `json:"unom" bson:"unom" validate:"required max=20000000"`
	// Полный Адрес в реестре.
	Address string `json:"address" bson:"address" validate:"required"`
	// Район округ.
	MunicipalDistrict string `json:"municipalDistrict" bson:"municipal_district" validate:"required"`
	// Граница объекта на карте.
	Polygon interface{} `json:"border" bson:"geo_data"`
	// Центр объекта на карте.
	Center Point `json:"center" bson:"geo_center"`
}

// Point Репрезентация точки в формате geoJson.
type Point struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
