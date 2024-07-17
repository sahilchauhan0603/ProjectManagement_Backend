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
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
                ],
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
        "/api/v1/project": {
            "get": {
                "description": "Get all projects",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get all projects",
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Project"
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
                "description": "Create a new project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "description": "Project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
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
        "/api/v1/project/{id}": {
            "get": {
                "description": "Get project by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get project by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
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
                "description": "Update project by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Update project by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "description": "Project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
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
                "description": "Delete project by ID",
                "tags": [
                    "project"
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Delete project by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
                ],
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
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
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
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
                    {
                        "name": "Authorization",
                        "in": "header",
                        "required": true,
                        "type": "string"
                    }
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
    },
    "definitions": {
        "models.Admin": {
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
                "project": {
                    "description": "Projects associated with the Admin",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Project"
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
        "models.Project": {
            "description": "Project model",
            "type": "object",
            "properties": {
                "admin_id": {
                    "description": "AdminID is the foreign key to Admin\nrequired: true",
                    "type": "integer"
                },
                "approach": {
                    "description": "Approach used in the project\nrequired: true",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt timestamp",
                    "type": "string"
                },
                "description": {
                    "description": "Description of the project\nrequired: true",
                    "type": "string"
                },
                "project_id": {
                    "description": "ProjectID is the primary key\nrequired: true",
                    "type": "integer"
                },
                "proof": {
                    "description": "Proof related to the project\nrequired: true",
                    "type": "string"
                },
                "supervisor": {
                    "description": "Supervisor of the project\nrequired: true",
                    "type": "string"
                },
                "title": {
                    "description": "Title of the project\nrequired: true",
                    "type": "string"
                },
                "topic": {
                    "description": "Topic of the project\nrequired: true",
                    "type": "string"
                },
                "uploaders": {
                    "description": "Uploaders associated with the project",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Uploader"
                    }
                }
            }
        },
        "models.Uploader": {
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
                "project_id": {
                    "description": "ProjectID is the foreign key to Project\nrequired: true",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
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
