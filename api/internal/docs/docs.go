// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "BSD 3-Clause License",
            "url": "https://raw.githubusercontent.com/EgorTarasov/true-tech/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/consumers/filters": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search consumers"
                ],
                "summary": "получение списка всех фильтров для поиска по объектам",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Filter"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/consumers/info/unoms": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumers"
                ],
                "summary": "получение информации о потребителях по unoms",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "уникальные номера объектов",
                        "name": "unoms",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/consumers/q": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search consumers"
                ],
                "summary": "поиск по объектам consumers с учетом фильтров",
                "parameters": [
                    {
                        "description": "фильтры для поиска",
                        "name": "filters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Filter"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.HeatingPointDTO"
                            }
                        }
                    }
                }
            }
        },
        "/geo/location/unom": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "получение гео данных по unom",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "уникальный номер объекта",
                        "name": "unom",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Address"
                        }
                    }
                }
            }
        },
        "/geo/location/unoms": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "получение гео данных по unoms",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "уникальные номера объектов",
                        "name": "unoms",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.Address"
                            }
                        }
                    }
                }
            }
        },
        "/geo/moek": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "третья тестовая ручка",
                "parameters": [
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "radius",
                        "name": "radius",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MoekDTO"
                            }
                        }
                    }
                }
            }
        },
        "/geo/property": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "вторая тестовая ручка",
                "parameters": [
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "radius",
                        "name": "radius",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PropertyDTO"
                            }
                        }
                    }
                }
            }
        },
        "/geo/property/id/:object": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "тестовая ручка на проверку postgis",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PropertyDTO"
                        }
                    }
                }
            }
        },
        "/map/events": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "map"
                ],
                "summary": "получение аварийных ситуаций и информации об объекте",
                "parameters": [
                    {
                        "description": "emergencyRequest",
                        "name": "emergencyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.emergencyRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/search/objects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "поиск по объектам (потребителям) тепло энергии",
                "parameters": [
                    {
                        "type": "string",
                        "description": "поисковой запрос",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.StatePropertySearchResult"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/users/file/list": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "list uploads",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "list uploads",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Upload"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        },
        "/users/file/status/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "check file processing",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "check file processing",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Upload"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "users with email + password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Auth with email creds",
                "parameters": [
                    {
                        "description": "user creds",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.emailCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.accessTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get user data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/token.UserPayload"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "creating email account with FirstName and LastName",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "creating email account",
                "parameters": [
                    {
                        "description": "User Email",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RegisterData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.accessTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        },
        "/users/upload": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "upload file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.fileUploadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "constants.Role": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "User",
                "Admin",
                "Moderator"
            ]
        },
        "handler.Address": {
            "type": "object",
            "required": [
                "address",
                "municipalDistrict"
            ],
            "properties": {
                "address": {
                    "description": "Полный Адрес в реестре.",
                    "type": "string"
                },
                "border": {
                    "description": "Граница объекта на карте."
                },
                "center": {
                    "description": "Центр объекта на карте.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/handler.Point"
                        }
                    ]
                },
                "municipalDistrict": {
                    "description": "Район округ.",
                    "type": "string"
                },
                "unom": {
                    "description": "Уникальный номер объекта недвижимости.",
                    "type": "integer"
                }
            }
        },
        "handler.Point": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handler.RegisterData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.accessTokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "handler.emailCredentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.emergencyRequest": {
            "type": "object",
            "required": [
                "distance",
                "latitude",
                "longitude"
            ],
            "properties": {
                "distance": {
                    "type": "integer",
                    "maximum": 10000,
                    "minimum": 0
                },
                "latitude": {
                    "type": "number",
                    "maximum": 90,
                    "minimum": -90
                },
                "longitude": {
                    "type": "number",
                    "maximum": 180,
                    "minimum": -180
                }
            }
        },
        "handler.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "handler.fileUploadResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "internal_search_models.Address": {
            "type": "object",
            "required": [
                "address",
                "municipalDistrict"
            ],
            "properties": {
                "address": {
                    "description": "Полный Адрес в реестре.",
                    "type": "string"
                },
                "border": {
                    "description": "Граница объекта на карте."
                },
                "center": {
                    "description": "Центр объекта на карте.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/internal_search_models.Point"
                        }
                    ]
                },
                "municipalDistrict": {
                    "description": "Район округ.",
                    "type": "string"
                },
                "unom": {
                    "description": "Уникальный номер объекта недвижимости.",
                    "type": "integer"
                }
            }
        },
        "internal_search_models.Point": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.DisconnectionDTO": {
            "type": "object",
            "properties": {
                "actualDisconnectionDate": {},
                "actualReconnectionDate": {},
                "disconnectionType": {
                    "type": "string"
                },
                "plannedDisconnectionDate": {},
                "plannedReconnectionDate": {},
                "reason": {
                    "type": "string"
                },
                "registrationDate": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "models.Filter": {
            "type": "object",
            "properties": {
                "filterName": {
                    "type": "string"
                },
                "values": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.HeatingPointDTO": {
            "type": "object",
            "properties": {
                "Адрес потребителя": {
                    "$ref": "#/definitions/internal_search_models.Address"
                },
                "Адрес строения": {
                    "$ref": "#/definitions/internal_search_models.Address"
                },
                "Балансодержатель": {
                    "type": "string"
                },
                "Вид ТП": {
                    "type": "string"
                },
                "Дата ввода в эксплуатацию": {},
                "Источник теплоснабжения": {
                    "type": "string"
                },
                "Муниципальный район": {
                    "type": "string"
                },
                "Номер ТП (тепловой пункт)": {
                    "type": "string"
                },
                "Тип по размещению": {
                    "type": "string"
                }
            }
        },
        "models.MoekDTO": {
            "type": "object",
            "properties": {
                "loadActual": {
                    "type": "number"
                },
                "loadAvg": {
                    "type": "number"
                },
                "loadHeating": {
                    "type": "number"
                },
                "loadVent": {
                    "type": "number"
                },
                "number": {
                    "type": "string"
                },
                "point": {},
                "polygon": {},
                "src": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.PropertyDTO": {
            "type": "object",
            "properties": {
                "disconnections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DisconnectionDTO"
                    }
                },
                "globalID": {
                    "type": "integer"
                },
                "point": {},
                "polygon": {}
            }
        },
        "models.StatePropertySearchResult": {
            "type": "object",
            "required": [
                "address",
                "municipalDistrict"
            ],
            "properties": {
                "address": {
                    "description": "Полный Адрес в реестре.",
                    "type": "string"
                },
                "area": {},
                "border": {
                    "description": "Граница объекта на карте."
                },
                "center": {
                    "description": "Центр объекта на карте.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/internal_search_models.Point"
                        }
                    ]
                },
                "class": {
                    "type": "string"
                },
                "feature": {
                    "type": "string"
                },
                "floors": {},
                "municipalDistrict": {
                    "description": "Район округ.",
                    "type": "string"
                },
                "purpose": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "unom": {
                    "description": "Уникальный номер объекта недвижимости.",
                    "type": "integer"
                }
            }
        },
        "models.Upload": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "filename": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "token.UserPayload": {
            "type": "object",
            "properties": {
                "auth_type": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/constants.Role"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.lct.larek.tech",
	BasePath:         "/ //no-lint.",
	Schemes:          []string{},
	Title:            "lct api",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
