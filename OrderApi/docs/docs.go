// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "zulkarnaen",
            "email": "premiumforspot@gmail.com"
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
        "/orders": {
            "post": {
                "description": "Create an order including its items, if provided.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create an order",
                "parameters": [
                    {
                        "description": "JSON of the order to be made.",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OrderBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/orders/{orderID}": {
            "get": {
                "description": "get order by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get an order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID number of the order",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "update order by ID including its items. Previous items are discarded.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update an order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID number of the order to be updated.",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON of the order to be updated.",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OrderBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessH"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete order by ID including its items.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete an order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID number of the order to be deleted.",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessH"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorH"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ErrorH": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string",
                    "example": "The error is explained here."
                }
            }
        },
        "controllers.SuccessH": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Operation successfull."
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Some description."
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "itemCode": {
                    "type": "string",
                    "example": "SOMECODE"
                },
                "orderID": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.ItemBody": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Some description."
                },
                "itemCode": {
                    "type": "string",
                    "example": "SOMECODE"
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "contoh"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                }
            }
        },
        "models.OrderBody": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "contoh"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ItemBody"
                    }
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Order API",
	Description:      "API server for orders in \"Scalable Webservice with Golang\" course from Hacktiv8 ?? Kominfo.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
