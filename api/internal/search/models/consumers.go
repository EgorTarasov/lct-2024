package models

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

// Filter фильтр для поиска по объектам.
type Filter struct {
	Name   string   `json:"filterName"`
	Values []string `json:"values"`
}

type HeatingPoint struct {
	BalanceHolder       string      `bson:"balance_holder" json:"Балансодержатель"`
	Address             string      `bson:"heating_point_address" json:"-"`
	ConsumerAddress     Address     `bson:"consumer_full_address" json:"Адрес строения"`
	Source              string      `bson:"heating_point_src" json:"Источник теплоснабжения"`
	Number              string      `bson:"heating_point_number" json:"Номер ТП (тепловой пункт)"`
	Type                string      `bson:"heating_point_type" json:"Вид ТП"`
	LocationType        string      `bson:"heating_point_location_type" json:"Тип по размещению"`
	District            string      `bson:"municipal_district" json:"Муниципальный район"`
	CommissioningDate   interface{} `bson:"commissioning_date" json:"Дата ввода в эксплуатацию"`
	HeatingPointAddress Address     `bson:"heating_point_full_address" json:"Адрес потребителя"`
}

// Event событие которое произошло с объектом.
type Event struct {
	Address           string `bson:"address" json:"Адрес"`
	Country           string `bson:"country" json:"Округ"`
	EndedAt           string `bson:"ended_at" json:"Дата и время завершения события"`
	ExternalClosedAt  string `bson:"external_closed_at" json:"Дата закрытия"`
	Source            string `bson:"src" json:"Источник"`
	Title             string `bson:"title" json:"Наименование"`
	Unom              int64  `bson:"unom" json:"УНОМ"`
	ExternalCreatedAt string `bson:"external_created_at" json:"Дата создания во внешней системе"`
	ExternalEndedAt   string `bson:"external_ended_at" json:"Дата и время завершения события во внешней системе"`
}

// HeatingPointDTO DTO для объекта теплоснабжения.
type HeatingPointDTO struct {
	HeatingPoint

	// models.Address
}

type StateConsumer struct {
	AdmDistrict       string      `bson:"administrative_district" json:"Административный округ"`
	Material          string      `bson:"material" json:"Материал"`
	MunicupalDistrict string      `bson:"municipal_district" json:"Муниципальный округ"`
	Purpose           string      `bson:"purpose" json:"Назначение"`
	Type              string      `bson:"type" json:"Тип"`
	Area              interface{} `bson:"property_total_area" json:"Общая площадь"`
	Floors            interface{} `bson:"property_floors" json:"Этажность"`
	PropertyClass     string      `bson:"property_class" json:"Класс"`
	Unom              int64       `bson:"unom" json:"Уном"`
	Events            []Event     `json:"События"`
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
	Deprecation          float64     `bson:"Depreciation" json:"Износ объекта (по БТИ)"`
	Address              string      `bson:"address" json:"Адрес из сторонней системы"`
	Apartments           int64       `bson:"apartments" json:"Количество квартир"`
	FullAddress          string      `bson:"bti_address" json:"Адрес по БТИ"`
	Country              string      `bson:"county" json:"Округ"`
	District             string      `bson:"district" json:"Район"`
	Feature              int64       `bson:"emergency_feature" json:"Признак аварийности здания"`
	Entrances            int64       `bson:"entrances" json:"Количество подъездов"`
	ExternalID           int64       `bson:"external_id" json:"Идентификатор из сторонней системы"`
	Floors               int64       `bson:"floors" json:"Количество этажей"`
	MKDState             int64       `bson:"mkd_state" json:"Статусы МКД"`
	PassengerElevator    int64       `bson:"passenger_elevator" json:"Пассажирские лифты"`
	ProjectSeries        interface{} `bson:"project_series" json:"Серии проектов"`
	RoofCleaningSequence int64       `bson:"roof_cleaning_sequence" json:"Очередность уборки кровли"`
	TotalArea            float64     `bson:"total_area" json:"Общая площадь"`
	Unom                 int64       `bson:"unom" json:"УНОМ"`
	WallMaterial         int64       `bson:"wall_material" json:"Материалы стен"`
	RoofMaterials        int64       `bson:"roof_materials" json:"Материалы кровли по БТИ"`
	TotalLivingArea      float64     `bson:"total_living_area" json:"Общая жилая площадь"`
	ServiceElevator      int64       `bson:"service_elevator" json:"Грузовые лифты"`
	TotalNonLivingArea   interface{} `bson:"total_non_living_area" json:"Общая нежилая площадь"`
	TypesOfHousingStock  int64       `bson:"types_of_housing_stock" json:"Типы жилищного фонда"`
	Events               []Event     `json:"События"`
}

// DispatchServices данные о диспетчерской службе.
type DispatchServices struct {
	Address              string  `bson:"address" json:"Адрес потребителя"`
	HeatDispatcherNumber string  `bson:"chp" json:"ЦТП"`
	ConsumerName         string  `bson:"consumer" json:"Потребитель (или УК)"`
	ConsumerGroup        string  `bson:"consumer_group" json:"Группа"`
	Country              string  `bson:"country" json:"Округ"`
	DispatchAddress      string  `bson:"dispatch_address" json:"Адрес ОДС"`
	DispatchNumber       string  `bson:"dispatch_number" json:"№ ОДС"`
	ExternalID           string  `bson:"external_id" json:"ID УУ"`
	Unom                 int64   `bson:"unom" json:"УНОМ"`
	ConsumerFullAddress  string  `bson:"Полный адрес" json:"consumerFullAddress"`
	Events               []Event `json:"События"`
}
