package models

import (
	"math"
)

// PropertyDAO объекты недвижимости.
type PropertyDAO struct {
	GlobalID       int64              `db:"global_id" bson:"global_id"`
	Address        string             `db:"address" bson:"ADDRESS"`
	InMoscow       bool               `db:"in_moscow" bson:"OnTerritoryOfMoscow"`
	Type           string             `db:"type" bson:"OBJ_TYPE"`
	Unom           int64              `db:"unom" bson:"UNOM"`
	Kladr          int64              `db:"kladr" bson:"KLADR"`
	GeoData        interface{}        `db:"geo_data" bson:"geoData"`
	GeoDataCenter  interface{}        `db:"geodata_center" bson:"geodata_center"`
	Disconnections []DisconnectionDAO `db:"disconnections,omitempty" bson:"disconnections,omitempty"`
}

// ToDTO преобразование объекта в DTO.
func (property *PropertyDAO) ToDTO() *PropertyDTO {
	disconnections := make([]DisconnectionDTO, 0, len(property.Disconnections))
	for _, disconnection := range property.Disconnections {
		disconnections = append(disconnections, *disconnection.ToDTO())
	}
	return &PropertyDTO{
		GlobalID:       property.GlobalID,
		GeoData:        property.GeoData,
		GeoDataCenter:  property.GeoDataCenter,
		Disconnections: disconnections,
	}
}

// DisconnectionDAO - информация об отключениях из таблицы 6.
type DisconnectionDAO struct {
	Reason                   string      `bson:"Причина"`
	Source                   string      `bson:"Источник"`
	RegistrationDate         string      `bson:"Дата регистрации отключения"`
	PlannedDisconnectionDate interface{} `bson:"Планируемая дата отключения"`
	PlannedReconnectionDate  interface{} `bson:"Планируемая дата включения"`
	ActualDisconnectionDate  interface{} `bson:"Фактическая дата отключения"`
	ActualReconnectionDate   interface{} `bson:"Фактическая дата включения"`
	DisconnectionType        string      `bson:"Вид отключения"`
}

// DisconnectionDTO представление данных об отключении для отображения.
type DisconnectionDTO struct {
	Reason                   string      `json:"reason"`
	Source                   string      `json:"source"`
	RegistrationDate         string      `json:"registrationDate"`
	PlannedDisconnectionDate interface{} `json:"plannedDisconnectionDate"`
	PlannedReconnectionDate  interface{} `json:"plannedReconnectionDate"`
	ActualDisconnectionDate  interface{} `json:"actualDisconnectionDate"`
	ActualReconnectionDate   interface{} `json:"actualReconnectionDate"`
	DisconnectionType        string      `json:"disconnectionType"`
}

// ToDTO преобразование объекта в DTO.
func (disconnection *DisconnectionDAO) ToDTO() *DisconnectionDTO {
	// check if PlannedDisconnectionDate is not NaN
	plannedDate := ""
	if value, ok := disconnection.PlannedDisconnectionDate.(float64); ok {
		if math.IsNaN(value) {
			disconnection.PlannedReconnectionDate = ""
		}
	} else {
		plannedDate = disconnection.PlannedDisconnectionDate.(string)
	}

	return &DisconnectionDTO{
		Reason:                   disconnection.Reason,
		Source:                   disconnection.Source,
		RegistrationDate:         disconnection.RegistrationDate,
		PlannedDisconnectionDate: plannedDate,
		PlannedReconnectionDate:  disconnection.PlannedReconnectionDate,
		ActualDisconnectionDate:  disconnection.ActualDisconnectionDate,
		ActualReconnectionDate:   disconnection.ActualReconnectionDate,
		DisconnectionType:        disconnection.DisconnectionType,
	}
}

// PropertyDTO представления объекта в json.
type PropertyDTO struct {
	GlobalID       int64              `json:"globalID"`
	GeoData        interface{}        `json:"polygon"`
	GeoDataCenter  interface{}        `json:"point"`
	Disconnections []DisconnectionDTO `json:"disconnections,omitempty"`
}
