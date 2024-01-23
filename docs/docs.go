// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "iostapovychweb@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/table/{sheet}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "returns cell value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sheet name",
                        "name": "sheet",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.ValueResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        },
        "/table/{sheet}/{cell}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "returns cell value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sheet name",
                        "name": "sheet",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cell address",
                        "name": "cell",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/api.ValueResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new pet to the store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "adds value to cell",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sheet name",
                        "name": "sheet",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cell address",
                        "name": "cell",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Set Cell Value",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CellValue"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CellValue": {
            "type": "object",
            "required": [
                "value"
            ],
            "properties": {
                "value": {
                    "type": "string"
                }
            }
        },
        "api.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "api.ValueResponse": {
            "type": "object",
            "required": [
                "cell",
                "sheet"
            ],
            "properties": {
                "cell": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "sheet": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ITSUXEL",
	Description:      "Educational Excel-like API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
