// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todo": {
            "get": {
                "description": "Getting all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Get all todos",
                "operationId": "get-all-todos",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "todo's todo id",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.GetAllTodosResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "post": {
                "description": "creating a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Create a todo",
                "operationId": "create-todo",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_datas.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.CreateTodoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            }
        },
        "/todo/{todoId}": {
            "get": {
                "description": "Getting a todo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Get todo by ID",
                "operationId": "get-todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "todo's todo id",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.GetTodoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "put": {
                "description": "Updating a todo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Update todo",
                "operationId": "update-todo",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_datas.UpdateTodoRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "todo's todo id",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.UpdateTodoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleting a todo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Delete todo by ID",
                "operationId": "delete-todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "todo's todo id",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.DeleteTodoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "doc_datas.CreateTodoRequest": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "description": {
                    "type": "string",
                    "example": "Cook fried rice with egg and chicken"
                },
                "title": {
                    "type": "string",
                    "example": "Make Dinner"
                }
            }
        },
        "doc_datas.CreateTodoResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "description": {
                    "type": "string",
                    "example": "Cook fried rice with egg and chicken"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Make Dinner"
                }
            }
        },
        "doc_datas.DeleteTodoResponse": {
            "type": "object",
            "properties": {
                "affected_row": {
                    "type": "integer",
                    "example": 1
                },
                "status_delete": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "doc_datas.GetAllTodosResponse": {
            "type": "object",
            "properties": {
                "todos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/doc_datas.GetTodoResponse"
                    }
                }
            }
        },
        "doc_datas.GetTodoResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "description": {
                    "type": "string",
                    "example": "Cook fried chicken with spicy sauce"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Make Delicious Dinner"
                }
            }
        },
        "doc_datas.UpdateTodoRequest": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "description": {
                    "type": "string",
                    "example": "Cook fried chicken with spicy sauce"
                },
                "title": {
                    "type": "string",
                    "example": "Make Delicious Dinner"
                }
            }
        },
        "doc_datas.UpdateTodoResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "description": {
                    "type": "string",
                    "example": "Cook fried chicken with spicy sauce"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Make Delicious Dinner"
                }
            }
        },
        "error_utils.MessageErrData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
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
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
