package models

// MoekDAO объекты Московской общественной энергетики.
type MoekDAO struct {
	Type          string      `db:"type" bson:"Вид ТП"`
	Src           string      `db:"src" bson:"Источник теплоснабжения"`
	Number        string      `db:"number" bson:"Номер ТП"`
	LoadAvg       float64     `db:"load_avg" bson:"Тепловая нагрузка ГВС ср."`
	LoadActual    float64     `db:"load_actual" bson:"Тепловая нагрузка ГВС факт."`
	LoadVent      float64     `db:"load_vent" bson:"Тепловая нагрузка Вентиляция"`
	LoadHeating   float64     `db:"load_heating" bson:"Тепловая нагрузка Отопление"`
	GeoData       interface{} `db:"geo_data" bson:"geoData"`
	GeoDataCenter interface{} `db:"geodata_center" bson:"geoData_center"`
}

// MoekDTO представления объекта в json.
type MoekDTO struct {
	Type          string      `json:"type"`
	Src           string      `json:"src"`
	Number        string      `json:"number"`
	LoadAvg       float64     `json:"loadAvg"`
	LoadActual    float64     `json:"loadActual"`
	LoadVent      float64     `json:"loadVent"`
	LoadHeating   float64     `json:"loadHeating"`
	GeoData       interface{} `json:"polygon"`
	GeoDataCenter interface{} `json:"point"`
}

// ToDTO преобразование объекта в DTO.
func (moek *MoekDAO) ToDTO() *MoekDTO {
	return &MoekDTO{
		Type:          moek.Type,
		Src:           moek.Src,
		Number:        moek.Number,
		LoadAvg:       moek.LoadAvg,
		LoadActual:    moek.LoadActual,
		LoadVent:      moek.LoadVent,
		LoadHeating:   moek.LoadHeating,
		GeoData:       moek.GeoData,
		GeoDataCenter: moek.GeoDataCenter,
	}
}
