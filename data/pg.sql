-- 13. Адресный реестр объектов недвижимости города Москвы.xlsx
create table if not exists address_registry(
    unom bigint primary key,
    -- UNOM Учётный номер объекта адресации в БД БТИ (кроме помещения и комнаты)
    global_id bigint,

    longitude numeric,
    latitude numeric,

    geo_data jsonb,
    -- geodata
    geo_center jsonb,
    -- geodata_center
    object_type text,
    -- OBJ_TYPE
    object_in_moscow boolean,
    -- OnTerritoryOfMoscow
    address text,
    -- ADDRESS Полное юридическое написание адреса или описание местоположения
    address_simple text,
    -- SIMPLE_ADDRESS Упрощённое написание адреса или описание местоположения
    adm_area text,
    -- ADM_AREA Административный округ
    municipal_district text,
    -- DISTRICT Муниципальный округ, поселение
    nreg bigint,
    -- Уникальный номер адреса в Адресном реестре
    dreg date,
    -- Дата регистрации адреса в Адресном реестре
    n_fias uuid,
    -- Уникальный номер адреса в государственном адресном реестре
    d_fias date,
    -- Дата регистрации адреса в государственном адресном реестре
    n_kad text,
    -- KAD_N Кадастровый номер объекта недвижимости (кроме земельного участка)
    zu_kad text,
    -- Кадастровый номер земельного участка (для ОКС)
    kladr bigint,
    -- KLADR Код КЛАДР для адресообразующего элемента нижнего уровня
    t_doc text,
    -- TDOC Документ-основание регистрационных действий
    n_doc text,
    -- NDOC Номер документа о регистрации адреса
    d_doc text,
    -- Дата документа о регистрации адреса
    address_type text,
    -- ADR_TYPE Тип адреса
    address_vid text,
    -- Вид адреса
    address_sostad text,
    -- SOSTAD Состояние адреса
    address_status text,
    -- STATUS Внесён в ГКН
    p0 text,
    -- Страна
    p1 text,
    -- Субъект РФ
    p2 text,
    -- Внутригородская территория
    p3 text,
    -- поселение
    p4 text,
    -- город
    p5 text,
    -- Муниципальный округ
    p6 text,
    -- Населённый пункт
    p7 text,
    -- Наименование элемента планировочной структуры или улично-дорожной сети
    p90 text,
    -- Наименование элемента планировочной структуры или улично-дорожной сети
    p91 text,
    -- Уточнение дополнительного адресообразующего элемента
    l1_type text,
    -- L1_TYPE Тип номера дома, владения, участка
    l1_value text,
    -- L1_VALUE Номер дома, владения, участка
    l2_type text,
    -- L2_TYPE Тип номера корпуса
    l2_value text,
    -- L2_VALUE Номер корпуса
    l3_type text,
    -- L3_TYPE Тип номера строения, сооружения
    l3_value text,
    -- L3_VALUE 	Номер строения, сооружения
    l4_type text,
    -- Тип номера помещения
    l4_value text,
    -- Номер помещения
    l5_type text,
    -- Тип номера комнаты
    l5_value text -- Номер комнаты
);

CREATE INDEX idx_latitude ON address_registry(latitude);
CREATE INDEX idx_longitude ON address_registry(longitude);



-- 6.Плановые-Внеплановые отключения 01.10.2023-30.04.2023.xlsx
create table if not exists disconnections(
    id bigserial primary key,
    cause text,
    -- причина
    src text,
    -- источник
    registered_at timestamp,
    -- Дата регистрации отключения
    planned_shutdown_at timestamp,
    -- Планируемая дата отключения
    planned_start_at timestamp,
    -- Планируемая дата включения
    actual_shutdown_at timestamp,
    -- Фактическая дата отключения
    actual_start_at timestamp,
    -- Фактическая дата включения
    shutdown_type text,
    -- Вид отключения
    unom bigint references address_registry(unom) on delete cascade,
    -- УНОМ
    address text -- Адрес
);

-- 7. Схема подключений МОЭК.xlsx
create table if not exists consumers(
    id bigserial primary key,
    external_id bigint,
    -- № п/п
    heating_point_number text,
    -- Номер ТП (тепловой пункт) 04-06-0601/010
    heating_point_address text,
    -- Адрес ТП
    heating_point_type text,
    -- Вид ТП
    heating_point_location_type text,
    -- Тип по размещению
    heating_point_src text,
    -- Источник теплоснабжения
    county text,
    --	Административный округ (ТП)
    municipal_district text,
    -- Муниципальный район Municipal district
    commissioning_date timestamp,
    -- Дата ввода в эксплуатацию Commissioning date
    balance_holder text,
    -- Балансодержатель
    consumer_address text,
    -- Адрес строения
    heating_load_avg numeric,
    -- Тепловая нагрузка ГВС ср.
    heating_load_actual numeric,
    -- Тепловая нагрузка ГВС факт.
    consumer_heating_building_load numeric,
    -- Тепловая нагрузка отопления строения
    consumer_heating_ventilation_load numeric,
    dispatching boolean -- Диспетчеризация
);



-- 6.Плановые-Внеплановые отключения 01.10.2023-30.04.2023.xlsx
create table if not exists events(
    id bigserial primary key,
    title text,
    src text,
    external_created_at timestamp,
    -- Дата создания во внешней системе
    external_closed_at timestamp,
    -- Дата закрытия
    county text,
    -- округ нужен индекс, много запросов на ==
    unom bigint references address_registry(unom) on delete cascade,
    -- УНОМ - для мапинга с адресами и т.п.
    actual_start_at timestamp,
    -- Фактическая дата включения
    address text,
    -- Адрес
    ended_at timestamp -- Дата и время завершения события
);
-- 8.Данные АСУПР с диспетчерскими ОДС.xlsx
create table if not exists dispatch_services(
    id bigserial primary key,
    external_id bigint,
    -- ID YY
    address text,
    -- Адрес
    country text,
    -- Округ
    unom bigint references address_registry(unom)  on delete cascade,
    -- UNOM
    consumer_group text,
    -- Группа
    dispatch_number text,
    -- № ОДС
    dispatch_address text,
    -- Адрес ОДС
    consumer text,
    -- Потребитель (или УК)
    chp text -- ЦТП (центральная ТП)04-05-0604/185
);
-- 9.Выгрузка БТИ.xlsx
-- Государственный технический учёт и техническая инвентаризацию объектов недвижимости
create table if not exists state_property(
    id bigserial primary key,
    external_id bigint,
    -- № п/п
    city text,
    -- Город
    administrative_district text,
    -- Административный округ	Administrative District
    municipal_district text,
    -- Муниципальный округ
    locality text,
    -- Населенный пункт (p6 в 13 таблице)
    street text,
    -- Улица
    property_number_type text,
    --Тип номера дома
    property_number text,
    -- Номер дома
    property_case_number text,
    -- Номер корпуса
    property_building_number_type text,
    -- Тип номера строения/сооружения
    property_building_number text,
    -- Номер строения
    unom bigint references address_registry(unom)  on delete cascade,
    -- UNOM
    unad text,
    -- UNAD
    material text,
    -- Материал
    purpose text,
    -- Назначение
    property_class text,
    -- Класс
    property_type text,
    -- Тип
    property_floors bigint,
    -- Этажность
    property_feature text,
    -- Признак
    property_total_area numeric -- Общая площадь
);
-- 11.Выгрузка_ОДПУ_отопление_ВАО_20240522
-- Общедомовые приборы учета
-- Common house metering devices
create table if not exists property_metering_devices(
    id bigserial primary key,
    external_id bigint,
    -- ID YY
    external_ty_id bigint,
    -- ID ТУ
    country text,
    -- Округ
    district text,
    -- Район
    consumer text,
    -- Потребители
    consumer_group text,
    -- Группа МКД
    unom bigint references address_registry(unom)  on delete cascade,
    -- UNOM
    address text,
    -- Адрес
    central_heating_circuit text,
    -- Центральное отопление(контур)
    metering_device_brand text,
    -- Марка счетчика
    metering_device_serial_number text,
    -- Серия/Номер счетчика
    metering_device_hours numeric,
    -- Наработка часов счётчика
    measurment_date date,
    -- Дата
    measurment_datetime timestamp,
    -- Месяц/Год
    measurment_units text,
    -- Unit
    -- Объём поданого теплоносителя в систему ЦО (float64)
    coolant_supplied_to_central_system numeric,
    -- Volume of coolant supplied to the central heating system
    -- Объём обратного теплоносителя из системы ЦО
    coolant_returned_to_central_system numeric,
    -- Разница между подачей и обраткой(Подмес)
    coolant_supplied_difference_mixture numeric,
    -- Разница между подачей и обраткой(Утечка)
    coolant_supplied_difference_leakage numeric,
    -- Температура подачи
    coolant_supplied_temperature numeric,
    -- Температура обратки
    coolant_returned_temperature numeric,
    -- Расход тепловой энергии
    thermal_energy_consumption numeric,
    -- Расход тепловой энергии
    errors text -- Ошибки
);
-- 12. Класс энергоэффективности соцобъектов.xlsx
create table if not exists social_property_energy_rating(
    id bigserial primary key,
    department text,
    -- Департамент
    buildings_cnt int,
    -- Всего строений
    buildings_area numeric,
    -- Общая площадь, м²
    buildings_heated_area numeric,
    -- Отапливаемая площадь, м²
    buildings_type text,
    -- Тип учреждения / Тип строения
    buildings_energy_efficiency_class text,
    -- Класс энергоэффективности здания
    buildings_floors int,
    -- Этажность
    buildings_elevators_cnt int,
    -- Количество лифтов
    buildings_wear_percentage float,
    -- Фактический износ здания %
    buildings_enterance_cnt int,
    -- Количество входов
    buildings_commissioning_year int,
    -- Год ввода здания в эксплуатацию
    employee_avg_cnt int -- Среднее количество работников, чел.
);
