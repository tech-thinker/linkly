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
        },
        "/api/v1/domains": {
            "get": {
                "description": "Get all domains",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "domains"
                ],
                "summary": "Get all domains",
                "operationId": "get-domains",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Domain"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/links": {
            "get": {
                "description": "Get list of links",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Get all links",
                "operationId": "get-all-links",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Link"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Add a new link",
                "operationId": "add-new-link",
                "parameters": [
                    {
                        "description": "Link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LinkBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/links/{id}": {
            "get": {
                "description": "Get a link",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Get a link",
                "operationId": "get-a-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Update a link",
                "operationId": "update-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LinkBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a link",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Delete a link",
                "operationId": "delete-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/links/{id}/qrcode": {
            "get": {
                "description": "Generate a qr code for a link",
                "produces": [
                    "application/json",
                    "data:image/png"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Generate a qr code for a link",
                "operationId": "generate-qr-code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QRCode"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/links/{id}/stats": {
            "get": {
                "description": "Get stats of a link",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Get stats of a link",
                "operationId": "get-link-stats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Stat"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/trackers": {
            "get": {
                "description": "Get all trackers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Get all trackers",
                "operationId": "get-trackers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Tracker"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/trackers/gen": {
            "get": {
                "description": "Generate a new tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Generate a new tracker",
                "operationId": "generate-tracker",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tracker"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/trackers/{id}": {
            "get": {
                "description": "Get a tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Get a tracker",
                "operationId": "get-tracker",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tracker"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Delete a tracker",
                "operationId": "delete-tracker",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/trackers/{id}/qr.png": {
            "get": {
                "description": "Get a qr code png for a tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "data:image/png"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Get a qr code png for a tracker",
                "operationId": "get-qr-code",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tracker"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/trackers/{id}/status": {
            "get": {
                "description": "Get the status of the tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trackers"
                ],
                "summary": "Get the status of the tracker",
                "operationId": "get-tracker-status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TrackerStatus"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/{link}": {
            "get": {
                "description": "Redirect to the target url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL Shortener"
                ],
                "summary": "Redirect to the target url",
                "operationId": "redirect",
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/{link}/qrcode": {
            "get": {
                "description": "Generate a QR code for the short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL Shortener"
                ],
                "summary": "Generate a QR code for the short url",
                "operationId": "generate-qr",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QRCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/{link}/track": {
            "get": {
                "description": "Track a url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL Shortener"
                ],
                "summary": "Track a url",
                "operationId": "track",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
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
        },
        "models.Browser": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name is the name of the browser",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the number of times the browser was visited",
                    "type": "integer"
                }
            }
        },
        "models.Country": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name is the name of the country",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the number of times the country was visited",
                    "type": "integer"
                }
            }
        },
        "models.Domain": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address is the address of the domain",
                    "type": "string"
                },
                "banned": {
                    "description": "Banned is a boolean that determines if the domain is banned",
                    "type": "boolean"
                },
                "created_at": {
                    "description": "CreatedAt is the time the domain was created",
                    "type": "string"
                },
                "homepage": {
                    "description": "Homepage is the homepage of the domain",
                    "type": "string"
                },
                "id": {
                    "description": "ID is the primary key for the domain. generates uuid using gorm",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt is the time the domain was updated",
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.ServiceError"
                }
            }
        },
        "models.Link": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address is the address of the link",
                    "type": "string"
                },
                "banned": {
                    "description": "Banned is a boolean that determines if the link is banned",
                    "type": "boolean"
                },
                "created_at": {
                    "description": "Stats is the stats for the link\nStats Stat ` + "`" + `json:\"stats,omitempty\"` + "`" + `\n CreatedAt is the time the link was created",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description for the link",
                    "type": "string"
                },
                "expire_at": {
                    "description": "ExipreAt is the time when the link expires",
                    "type": "string"
                },
                "id": {
                    "description": "ID is the primary key for the link. generates uuid using gorm",
                    "type": "string"
                },
                "ip": {
                    "description": "IP is the ip address of the user who created the link [security,spam]",
                    "type": "string"
                },
                "link": {
                    "description": "Link is the unique address that is being stored",
                    "type": "string"
                },
                "password": {
                    "description": "Password is the password for the link",
                    "type": "string"
                },
                "reusable": {
                    "description": "Reusable is a boolean that determines if the link is reusable",
                    "type": "boolean"
                },
                "target": {
                    "description": "Target is the target for the link",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt is the time the link was last updated",
                    "type": "string"
                },
                "user_id": {
                    "description": "UserID is the user who created the link",
                    "type": "string"
                },
                "visit_count": {
                    "description": "VisitCount is the number of times the link has been visited",
                    "type": "integer"
                }
            }
        },
        "models.LinkBody": {
            "type": "object",
            "properties": {
                "customurl": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "expire_in": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "reusable": {
                    "type": "boolean"
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.OS": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name is the name of the operating system",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the number of times the operating system was visited",
                    "type": "integer"
                }
            }
        },
        "models.QRCode": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Referrer": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name is the name of the referrer",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the number of times the referrer was visited",
                    "type": "integer"
                }
            }
        },
        "models.ServiceError": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Stat": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address is the address of the link",
                    "type": "string"
                },
                "all_time": {
                    "description": "AllTime is the number of times the link has been visited",
                    "$ref": "#/definitions/models.StatItem"
                },
                "banned": {
                    "description": "Banned is a boolean that determines if the link is banned",
                    "type": "boolean"
                },
                "created_at": {
                    "description": "CreatedAt is the time the link was created",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description for the link",
                    "type": "string"
                },
                "expire_at": {
                    "description": "ExipreAt is the time when the link expires",
                    "type": "string"
                },
                "ip": {
                    "description": "IP is the ip address of the user who created the link [security,spam]",
                    "type": "string"
                },
                "last_day": {
                    "description": "LastDay is the number of times the link has been visited in the last day",
                    "$ref": "#/definitions/models.StatItem"
                },
                "last_month": {
                    "description": "LastMonth is the number of times the link has been visited in the last month",
                    "$ref": "#/definitions/models.StatItem"
                },
                "last_week": {
                    "description": "LastWeek is the number of times the link has been visited in the last week",
                    "$ref": "#/definitions/models.StatItem"
                },
                "link": {
                    "description": "Link is the unique address that is being stored",
                    "type": "string"
                },
                "password": {
                    "description": "Password is the password for the link",
                    "type": "string"
                },
                "reusable": {
                    "description": "Reusable is a boolean that determines if the link is reusable",
                    "type": "boolean"
                },
                "target": {
                    "description": "Target is the target for the link",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt is the time the link was last updated",
                    "type": "string"
                },
                "user_id": {
                    "description": "UserID is the user who created the link",
                    "type": "string"
                }
            }
        },
        "models.StatItem": {
            "type": "object",
            "properties": {
                "browser": {
                    "description": "Browser is the browser the link was visited from",
                    "$ref": "#/definitions/models.Browser"
                },
                "country": {
                    "description": "Country is the country the link was visited from",
                    "$ref": "#/definitions/models.Country"
                },
                "os": {
                    "description": "OS is the operating system the link was visited from",
                    "$ref": "#/definitions/models.OS"
                },
                "referrer": {
                    "description": "Referrer is the referrer the link was visited from",
                    "$ref": "#/definitions/models.Referrer"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "models.Tracker": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "visit_count": {
                    "type": "integer"
                }
            }
        },
        "models.TrackerStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "seen": {
                    "type": "boolean"
                },
                "url": {
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
