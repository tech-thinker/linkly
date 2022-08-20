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
        "termsOfService": "/terms/",
        "contact": {
            "name": "API Support",
            "url": "/support"
        },
        "license": {
            "name": "MIT License",
            "url": "https://github.com/tech-thinker/linkly/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/health": {
            "get": {
                "description": "checks the health of the system.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Show the status of the system.",
                "operationId": "healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Check"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Check": {
            "type": "object",
            "properties": {
                "alloc_bytes": {
                    "description": "TotalAllocBytes is the bytes allocated and not yet freed.",
                    "type": "integer"
                },
                "failures": {
                    "description": "Failures holds the failed checks along with their messages.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "goroutines_count": {
                    "description": "GoroutinesCount is the number of the current goroutines.",
                    "type": "integer"
                },
                "heap_objects_count": {
                    "description": "HeapObjectsCount is the number of objects in the go heap.",
                    "type": "integer"
                },
                "startup": {
                    "description": "StartUp is the time to boot up the system.",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the check status.",
                    "type": "string"
                },
                "timestamp": {
                    "description": "Timestamp is the time in which the check occurred.",
                    "type": "string"
                },
                "total_alloc_bytes": {
                    "description": "TotalAllocBytes is the total bytes allocated.",
                    "type": "integer"
                },
                "uptime": {
                    "description": "Uptime is the time in which the check occurred.",
                    "type": "string"
                },
                "version": {
                    "description": "Version is the go version.",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Linkly API",
	Description:      "URL shortener API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
