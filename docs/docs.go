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
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "paths": {
        "/api/v1/admin": {
            "get": {
                "description": "Get all admins",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get all admins",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Admin"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Create a new admin",
                "parameters": [
                    {
                        "description": "Admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    },
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/{id}": {
            "get": {
                "description": "Get admin by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get admin by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Update admin by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    },
                    "400": {
                        "description": "Invalid ID or Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete admin by ID",
                "tags": [
                    "admin"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Delete admin by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/work": {
            "get": {
                "description": "Get all works",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get all works",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Work"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new work",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Create a new work",
                "parameters": [
                    {
                        "description": "Project",
                        "name": "work",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Work"
                        }
                    },
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Work"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/work/{id}": {
            "get": {
                "description": "Get work by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get work by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Work ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Work"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update work by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Update work by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Work ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Work",
                        "name": "work",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Work"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Work"
                        }
                    },
                    "400": {
                        "description": "Invalid ID or Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete work by ID",
                "tags": [
                    "work"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Delete work by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Work ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/upload": {
            "get": {
                "description": "Get all uploaders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploader"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get all uploaders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Uploader"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new uploader",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploader"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Create a new uploader",
                "parameters": [
                    {
                        "description": "Uploader",
                        "name": "uploader",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Uploader"
                        }
                    },
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Uploader"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/upload/{id}": {
            "get": {
                "description": "Get uploader by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploader"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get uploader by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Uploader ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Uploader"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update uploader by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploader"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Update uploader by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Uploader ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Uploader",
                        "name": "uploader",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Uploader"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Uploader"
                        }
                    },
                    "400": {
                        "description": "Invalid ID or Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete uploader by ID",
                "tags": [
                    "uploader"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Delete uploader by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Uploader ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sendEmail": {
            "post": {
                "description": "Sends an Email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Sends an email to reset password",
                "parameters": [
                    {
                        "description": "Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmailRequest"
                        }
                    },
                ],
                "responses": {
                    "201": {
                        "description": "Successfully Send",
                        "schema": {
                            "$ref": "#/definitions/models.EmailRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
    },
    "definitions": {
        "Admin": {
            "description": "Admin model",
            "type": "object",
            "properties": {
                "admin_id": {
                    "description": "ID of the admin\nrequired: true",
                    "type": "integer"
                },
                "admin_name": {
                    "description": "Name of the Admin\nrequired: true",
                    "type": "string"
                },
                "createdAt": {
                    "description": "CreatedAt timestamp",
                    "type": "string"
                },
                "password": {
                    "description": "Password of the Admin\nrequired: true",
                    "type": "string"
                },
                "work": {
                    "description": "Works associated with the Admin",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Work"
                    }
                },
                "uploader": {
                    "description": "Uploader associated with the Admin",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Uploader"
                    }
                },
                "username": {
                    "description": "UerName of the Admin\nrequired: true",
                    "type": "string"
                }
            }
        },
        "Work": {
            "description": "Work model",
            "type": "object",
            "properties": {
                "admin_id": {
                    "description": "AdminID is the foreign key to Admin\nrequired: true",
                    "type": "integer"
                },
                "approach": {
                    "description": "Approach used in the work\nrequired: true",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt timestamp",
                    "type": "string"
                },
                "description": {
                    "description": "Description of the work\nrequired: true",
                    "type": "string"
                },
                "work_id": {
                    "description": "WorkID is the primary key\nrequired: true",
                    "type": "integer"
                },
                "proof": {
                    "description": "Proof related to the work\nrequired: true",
                    "type": "string"
                },
                "supervisor": {
                    "description": "Supervisor of the work\nrequired: true",
                    "type": "string"
                },
                "title": {
                    "description": "Title of the work\nrequired: true",
                    "type": "string"
                },
                "topic": {
                    "description": "Topic of the work\nrequired: true",
                    "type": "string"
                },
                "type": {
                    "description": "Type of the work\nrequired: true",
                    "type": "string"
                },
                "uploaders": {
                    "description": "Uploaders associated with the work",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Uploader"
                    }
                }
            }
        },
        "Uploader": {
            "description": "Uploader model",
            "type": "object",
            "properties": {
                "admin_id": {
                    "description": "AdminID is the foreign key to Admin\nrequired: true",
                    "type": "integer"
                },
                "college_name": {
                    "description": "CollegeName of the uploader\nrequired: true",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt timestamp",
                    "type": "string"
                },
                "email": {
                    "description": "Email of the uploader\nrequired: true",
                    "type": "string"
                },
                "enrollment_no": {
                    "description": "EnrollmentNo is the primary key\nrequired: true",
                    "type": "integer"
                },
                "first_name": {
                    "description": "FirstName of the uploader\nrequired: true",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName of the uploader\nrequired: true",
                    "type": "string"
                },
                "work_id": {
                    "description": "WorkID is the foreign key to Work\nrequired: true",
                    "type": "integer"
                }
            }
        },
        "EmailRequest": {
            "description": "Email model",
            "type": "object",
            "properties": {
                "Email": {
                    "description": "Email\nrequired: true",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "projectmanagement-backend-tysm.onrender.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Project Management APIs",
	Description:      "This is Project Management API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
