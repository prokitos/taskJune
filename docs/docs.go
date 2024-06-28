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
            "email": "soberkoder@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/GetOrderBook": {
            "get": {
                "description": "get DepthOrder by name or pair",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DepthOrder"
                ],
                "summary": "get DepthOrder by name or pair",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Show by exchange_name",
                        "name": "exchange_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Show by pair",
                        "name": "pair",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/GetOrderHistory": {
            "get": {
                "description": "get order History",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DepthOrder"
                ],
                "summary": "get order History",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Artem",
                        "name": "client_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "",
                        "name": "exchange_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "",
                        "name": "label",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "",
                        "name": "pair",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/SaveOrder": {
            "post": {
                "description": "save order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DepthOrder"
                ],
                "summary": "save order",
                "parameters": [
                    {
                        "description": "insert client and history",
                        "name": "clientHistory",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ClientAndHistorySwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/SaveOrderBook": {
            "post": {
                "description": "insert orderBook",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DepthOrder"
                ],
                "summary": "insert orderBook",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert order",
                        "name": "exchange_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "insert order",
                        "name": "pair",
                        "in": "query"
                    },
                    {
                        "description": "Insert order",
                        "name": "orderBook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ArrayDepthOrderSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ArrayDepthOrderSwag": {
            "type": "object",
            "properties": {
                "array": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DepthOrderSwag"
                    }
                }
            }
        },
        "models.ClientAndHistorySwag": {
            "type": "object",
            "properties": {
                "client": {
                    "$ref": "#/definitions/models.ClientSwag"
                },
                "history": {
                    "$ref": "#/definitions/models.HistoryOrderSwag"
                }
            }
        },
        "models.ClientSwag": {
            "type": "object",
            "properties": {
                "client_name": {
                    "type": "string",
                    "example": "Artem"
                },
                "exchange_name": {
                    "type": "string",
                    "example": ""
                },
                "label": {
                    "type": "string",
                    "example": ""
                },
                "pair": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "models.DepthOrderSwag": {
            "type": "object",
            "properties": {
                "baseQty": {
                    "type": "number",
                    "example": 95
                },
                "price": {
                    "type": "number",
                    "example": 120
                }
            }
        },
        "models.HistoryOrderSwag": {
            "type": "object",
            "properties": {
                "algorithm_name_placed": {
                    "type": "string",
                    "example": ""
                },
                "base_qty": {
                    "type": "number",
                    "example": 10
                },
                "client_name": {
                    "type": "string",
                    "example": "Artem"
                },
                "commission_quote_qty": {
                    "type": "number",
                    "example": 12
                },
                "exchange_name": {
                    "type": "string",
                    "example": ""
                },
                "highest_buy_prc": {
                    "type": "number",
                    "example": 150
                },
                "label": {
                    "type": "string",
                    "example": ""
                },
                "lowest_sell_prc": {
                    "type": "number",
                    "example": 100
                },
                "pair": {
                    "type": "string",
                    "example": "rub-eur"
                },
                "price": {
                    "type": "number",
                    "example": 112.7
                },
                "side": {
                    "type": "string",
                    "example": ""
                },
                "time_placed": {
                    "type": "string",
                    "example": "2017-07-21T17:32:28Z"
                },
                "type": {
                    "type": "string",
                    "example": ""
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "User API",
	Description:      "This is a sample service for managing users",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}