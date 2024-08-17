// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/result/{task_id}": {
            "get": {
                "description": "Get the result of the task by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get Result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the task",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/status/{task_id}": {
            "get": {
                "description": "Get the status of the ongoing task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the task",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/task": {
            "post": {
                "description": "Creates a task",
                "tags": [
                    "Task"
                ],
                "summary": "Post task",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Homework1",
	Description:      "this is my second homework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
