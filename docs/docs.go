// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/newsletter/subscribers": {
            "post": {
                "description": "Add a new subscriber to the newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscribers"
                ],
                "summary": "Add a new subscriber",
                "parameters": [
                    {
                        "description": "Subscriber object",
                        "name": "subscriber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Subscriber"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request Error message",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Error message",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/newsletter/subscribers/{id}": {
            "delete": {
                "description": "Delete a subscriber from the newsletter",
                "tags": [
                    "subscribers"
                ],
                "summary": "Delete a subscriber",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscriber ID or email",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request Error message",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": " Internal Error message",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BaseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Subscriber": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://newsletter-test-4aaa4eezza-ew.a.run.app",
	BasePath:         "/v1",
	Schemes:          []string{"https"},
	Title:            "Newsletter API",
	Description:      "This is Solidform's Newsletter API to handle subscriptions and sending emails",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}