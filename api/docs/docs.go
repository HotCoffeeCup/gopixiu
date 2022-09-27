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
        "termsOfService": "https://github.com/caoyingjunz/gopixiu",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/caoyingjunz/gopixiu",
            "email": "support@gopixiu.io"
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
        "/clouds/v1/{cloud_name}/kubeconfigs": {
            "get": {
                "description": "get by cloud kubeConfig ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kubeConfigs"
                ],
                "summary": "get a cloud custom kubeConfig",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "cloud name",
                        "name": "cloud_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/types.KubeConfigOptions"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update by cloud kubeConfig",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kubeConfigs"
                ],
                "summary": "Update a cloud custom kubeConfig",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "cloud name",
                        "name": "cloud_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Cloud ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/types.KubeConfigOptions"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create by cloud kubeConfig",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kubeConfigs"
                ],
                "summary": "Create a cloud custom kubeConfig",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "cloud name",
                        "name": "cloud_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "service_account, cluster_role",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.KubeConfigOptions"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/types.KubeConfigOptions"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            }
        },
        "/clouds/v1/{cloud_name}/kubeconfigs/{id}": {
            "get": {
                "description": "get by cloud kubeConfig ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kubeConfigs"
                ],
                "summary": "get a cloud custom kubeConfig",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "cloud name",
                        "name": "cloud_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "kubeConfig ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/types.KubeConfigOptions"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by cloud kubeConfig ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kubeConfigs"
                ],
                "summary": "Delete a cloud custom kubeConfig",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "cloud name",
                        "name": "cloud_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Cloud ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            }
        },
        "/clouds/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clouds"
                ],
                "summary": "Get a cloud",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Cloud ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by cloud ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clouds"
                ],
                "summary": "Delete a cloud",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Cloud ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "返回的状态码",
                    "type": "integer"
                },
                "message": {
                    "description": "异常返回时的错误信息",
                    "type": "string"
                },
                "result": {
                    "description": "正常返回时的数据，可以为任意数据结构"
                }
            }
        },
        "types.KubeConfigOptions": {
            "type": "object",
            "properties": {
                "cloud_name": {
                    "type": "string"
                },
                "cluster_role": {
                    "type": "string"
                },
                "config": {
                    "type": "string"
                },
                "expiration_timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "service_account": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Pixiu API Documentation",
	Description:      "Use the Pixiu APIs to your cloud",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}