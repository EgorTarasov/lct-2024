# Установка
## 1. Склонировать репозиторий
```bash
git clone https://github.com/EgorTarasov/lct-2024
```

## 2. Провести настройку окружения
В папке config находится .env.example с переменными окружения для запуска docker compose 
его требуется скопировать в .env и заполнить значения переменных окружения
```bash
cp config/.env.example config/.env
```
После чего требуется экспортировать переменные окружения
```bash
source config/.env
```

## 3. Настройки приложений
В папке config находятся файлы конфигурации для каждого приложения: 
- api.yaml для сервиса на go

> После запуска приложения так же требуется создание bucket в s3 
> и получение credentials для работы функционала с отчетностью и файлами 

## 4. Запуск приложений
Для запуска приложений требуется запустить docker compose 
и выполнить миграции для бд с использованием [goose](https://github.com/pressly/goose) 
инструкцию по установки можно найти тут [интсрукция](https://github.com/pressly/goose?tab=readme-ov-file#install)  

В папке `api` расположен Makefile запустив который с командой `migration-up` 
вы сопоставите версию бд с требуемой для работы сервиса *можно указывать параметры подключения к бд через переменные окружения (см `api/Makefile`)*  

В папке api выполните команду
```bash
make migration-up
```
После чего получите вывод от goose о выполнении миграций:
```
➜  api git:(dev-backend) ✗ make migration-up 
goose -dir "/Users/user/lct-2024/api/cmd/migrations" postgres "user=Dino password=Dino-misos2024 dbname=dev host=192.168.1.70 port=54000 sslmode=disable" up
2024/06/15 15:17:46 goose: no migrations to run. current version: 20240613181516
```

## 5. Запуск docker-compose
Для запуска приложений требуется перейти в директорию `docker` и выполнить команду
```bash
cd docker
docker-compose up
```
После чего все сервисы будут запущены и веб интерфейс будет доступен по адресу `http://localhost:3000`

## 6. Локальный запуск
Замечение версия веб-приложения использует запросы к нашему api в случае развертки 
на другом сетевом адресе требуется заново собрать образ для `ui` указав другое название тэка при сборке в `frontend/Makefile` и дальнейшего использования в docker-compose  