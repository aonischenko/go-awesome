// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-27 13:46:46.9875614 +0300 EEST m=+0.027040101

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/div": {
            "get": {
                "description": "Should return status 200 with an division operation result",
                "produces": [
                    "application/json"
                ],
                "summary": "Division using request url params",
                "operationId": "v1GetDiv",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "division operation numerator",
                        "name": "x",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "division operation denominator",
                        "name": "y",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.OpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            },
            "put": {
                "description": "Should return status 200 with an division operation result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Division using request body",
                "operationId": "v1PutDiv",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.OpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/div": {
            "get": {
                "description": "Should return status 200 with an division operation result",
                "produces": [
                    "application/json"
                ],
                "summary": "Division using request url params",
                "operationId": "v2GetDiv",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "division operation numerator",
                        "name": "x",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "division operation denominator",
                        "name": "y",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.OpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            },
            "put": {
                "description": "Should return status 200 with an division operation result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Division using request body",
                "operationId": "v2PutDiv",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.OpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApiError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "ts": {
                    "type": "string"
                }
            }
        },
        "model.OpResult": {
            "type": "object",
            "properties": {
                "operation": {
                    "type": "object"
                },
                "result": {
                    "type": "object"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
