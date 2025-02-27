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
        "/brands": {
            "get": {
                "description": "Retrieve a list of all brands in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brands"
                ],
                "summary": "Get all brands",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Brand"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new brand and store it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brands"
                ],
                "summary": "Create a new brand",
                "parameters": [
                    {
                        "description": "Brand Data",
                        "name": "brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Brand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.Brand"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/brands/{id}": {
            "get": {
                "description": "Get the details of a brand by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brands"
                ],
                "summary": "Get a brand by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Brand"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the brand with the given ID from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brands"
                ],
                "summary": "Delete a brand by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Retrieve a list of all products in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product and store it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Product Data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get the details of a product by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the product with the given ID from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entities.Product"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "description": "Get a list of all transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get all transactions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new transaction and store it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "description": "Transaction Data",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    }
                }
            }
        },
        "/transactions/user/{user_id}": {
            "get": {
                "description": "Retrieve all transactions for a specific user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get Transaction Details by User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Transaction"
                            }
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
        "/transactions/{id}": {
            "get": {
                "description": "Get transaction details by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get a transaction by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a transaction by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Delete a transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entities.Transaction"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieve all users from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new user to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "Update user details such as name, points, and balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data to be updated",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to update user",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/users/{id}/update": {
            "put": {
                "description": "Update the points of a user based on their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update User Points",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Points to be added",
                        "name": "points",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    }
                }
            }
        },
        "/users/{id}/vouchers": {
            "get": {
                "description": "Fetch all vouchers that a user has enough points to redeem and that are not expired",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get available vouchers for a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of available vouchers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Voucher"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve vouchers",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/vouchers": {
            "get": {
                "description": "Retrieve all vouchers from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vouchers"
                ],
                "summary": "Get all vouchers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Voucher"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new voucher to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vouchers"
                ],
                "summary": "Create a voucher",
                "parameters": [
                    {
                        "description": "Voucher data",
                        "name": "voucher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Voucher"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.Voucher"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Brand": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Product": {
            "type": "object",
            "properties": {
                "brandID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "entities.Transaction": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "productID": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "userID": {
                    "type": "string"
                },
                "voucherID": {
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "Balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                }
            }
        },
        "entities.Voucher": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "cost_in_point": {
                    "type": "integer"
                },
                "expiration": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "GoRestAPI Swagger Example API",
	Description:      "This is a simple Go REST API with Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
