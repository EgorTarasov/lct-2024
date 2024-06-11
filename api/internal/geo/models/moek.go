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

// JDSDAO Joint dispatch services запись о (Объединенные диспетчерские службы) в БД.
type JDSDAO struct {
	Number        string      `bson:"odsNumber" json:"odsNumber"`
	Address       string      `bson:"odsAddress" json:"odsAddress"`
	GeoData       interface{} `bson:"geoData" json:"geoData"`
	GeoDataCenter interface{} `bson:"geoDataCenter" json:"geoDataCenter"`
}

// ConsumerDAO потребитель
type ConsumerDAO struct {
	Group   string `bson:"group" json:"group"`
	Unom    string `bson:"consumerUnom" json:"unom"`
	Address string `bson:"consumerAddress" json:"address"`
}

// ConsumerDTO потребитель
type ConsumerDTO struct {
	Group   string `json:"group"`
	Unom    string `json:"unom"`
	Address string `json:"address"`
}

func (consumer *ConsumerDAO) ToDTO() *ConsumerDTO {
	return &ConsumerDTO{
		Group:   consumer.Group,
		Unom:    consumer.Unom,
		Address: consumer.Address,
	}
}

// CHP Combined heat and power plant
type CHP struct {
	Name      string        `bson:"ctp" json:"chp"`
	Consumers []ConsumerDTO `bson:"consumers" json:"consumers"`
}
