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
            "name": "API Support",
            "url": "http://cloud-barista.github.io",
            "email": "contact-to-cloud-barista@googlegroups.com"
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
        "/health": {
            "get": {
                "description": "Check API server is running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[System] Utility"
                ],
                "summary": "Check API server is running",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/httpVersion": {
            "get": {
                "description": "Checks and logs the HTTP version of the incoming request to the server console.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[System] Utility"
                ],
                "summary": "Check HTTP version of incoming request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/sample/users": {
            "get": {
                "description": "Get information of all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Get a list of users",
                "responses": {
                    "200": {
                        "description": "(sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUsersResponse"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/sample/users/{id}": {
            "get": {
                "description": "Get information of a user with a specific ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Get specific user information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUserResponse"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information to update",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deletion successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "patch": {
                "description": "Patch a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Patch a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information to update",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.PatchUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.PatchUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/tofu/config/vpn-tunnels": {
            "post": {
                "description": "Create configurations for VPN tunnels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Tofu] Commands"
                ],
                "summary": "Create configurations for VPN tunnels",
                "parameters": [
                    {
                        "description": "Create configurations for VPN tunnels",
                        "name": "ConfigVPNTunnels",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TofuConfigVPNTunnelsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/tofu/init": {
            "post": {
                "description": "Prepare your working directory for other commands",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Tofu] Commands"
                ],
                "summary": "Prepare your working directory for other commands",
                "parameters": [
                    {
                        "description": "TofuInitRequest",
                        "name": "TofuInitRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TofuInitRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/tofu/show/{namespaceId}": {
            "get": {
                "description": "Show the current state of a saved plan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Tofu] Commands"
                ],
                "summary": "Show the current state of a saved plan",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/tofu/version": {
            "get": {
                "description": "Check Tofu version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Tofu] Commands"
                ],
                "summary": "Check Tofu version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.GetUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.GetUsersResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MyUser"
                    }
                }
            }
        },
        "handlers.PatchUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.PatchUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.TfVarsVPNTunnels": {
            "type": "object",
            "properties": {
                "my-imported-aws-subnet-id": {
                    "type": "string"
                },
                "my-imported-aws-vpc-id": {
                    "type": "string"
                },
                "my-imported-gcp-subnet-id": {
                    "type": "string"
                },
                "my-imported-gcp-vpc-id": {
                    "type": "string"
                }
            }
        },
        "handlers.TofuConfigVPNTunnelsRequest": {
            "type": "object",
            "properties": {
                "namespaceId": {
                    "type": "string"
                },
                "tfVars": {
                    "$ref": "#/definitions/handlers.TfVarsVPNTunnels"
                }
            }
        },
        "handlers.TofuInitRequest": {
            "type": "object",
            "properties": {
                "namespaceId": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.MyUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "text": {
                    "type": "string",
                    "example": "Any text"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "latest",
	Host:             "",
	BasePath:         "/mc-net",
	Schemes:          []string{},
	Title:            "POC-MC-Net-TF REST API",
	Description:      "POC-MC-Net-TF REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
