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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book": {
            "get": {
                "description": "Get List Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get List Book",
                "operationId": "get_list_book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetListBook",
                        "schema": {
                            "$ref": "#/definitions/models.GetListBookResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create Book",
                "operationId": "create_book",
                "parameters": [
                    {
                        "description": "CreateBookRequestBody",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBook"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetBookBody",
                        "schema": {
                            "$ref": "#/definitions/models.CreateBook"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/book/{id}": {
            "get": {
                "description": "Get By Id Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get By Id Book",
                "operationId": "get_by_id_book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetIdBookBody",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update Book",
                "operationId": "update_book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateBookRequestBody",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBookSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UpdateBookBody",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBook"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Delete Book",
                "operationId": "delete_book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DeleteBookBody",
                        "schema": {
                            "$ref": "#/definitions/models.Empty"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/category": {
            "get": {
                "description": "Get List Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get List Category",
                "operationId": "get_list_category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategoryListBody",
                        "schema": {
                            "$ref": "#/definitions/models.GetListCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create category",
                "operationId": "create_category",
                "parameters": [
                    {
                        "description": "CreateCategoryRequestBody",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategory"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "get": {
                "description": "Get By ID Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get By ID Category",
                "operationId": "get_by_id_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category",
                "operationId": "update_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateCategoryRequestBody",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCategorySwag"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "UpdateCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete Category",
                "operationId": "delete_category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "DeleteCategoryBody",
                        "schema": {
                            "$ref": "#/definitions/models.EmptyCat"
                        }
                    },
                    "400": {
                        "description": "Invalid Argumant",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Book": {
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
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateBook": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Empty": {
            "type": "object"
        },
        "models.EmptyCat": {
            "type": "object"
        },
        "models.GetListBookResponse": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Book"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.GetListCategoryResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.UpdateBook": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.UpdateBookSwag": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.UpdateCategorySwag": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
