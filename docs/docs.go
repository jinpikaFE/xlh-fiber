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
        "/v1/captcha": {
            "post": {
                "description": "获取验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取验证码"
                ],
                "summary": "获取验证码",
                "parameters": [
                    {
                        "description": "mobile",
                        "name": "mobile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginMobile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/region": {
            "get": {
                "description": "获取省市区",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取省市区"
                ],
                "summary": "获取省市区",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/sort": {
            "get": {
                "description": "获取分类详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sort"
                ],
                "summary": "获取分类详情",
                "parameters": [
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            },
            "post": {
                "description": "添加sort",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sort"
                ],
                "summary": "添加sort",
                "parameters": [
                    {
                        "description": "shop_sort",
                        "name": "shop_sort",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ShopSort"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/sort/{id}": {
            "put": {
                "description": "编辑Sort",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sort"
                ],
                "summary": "编辑Sort",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ShopSort ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "shop_sort",
                        "name": "shop_sort",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ShopSort"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除sort",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sort"
                ],
                "summary": "删除sort",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ShopSort ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/sorts": {
            "get": {
                "description": "获取分类列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sort"
                ],
                "summary": "获取分类列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/upload": {
            "post": {
                "description": "文件上传",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件上传"
                ],
                "summary": "文件上传",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "description": "获取User详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取User详情",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatarUrl",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "index",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "nickName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "openid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "region",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "unionid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "query tag是query参数别名，json xml，form适合post",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            },
            "post": {
                "description": "添加user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "添加user",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "put": {
                "description": "编辑用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "编辑用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "删除user",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "获取User列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取User列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatarUrl",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "index",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "nickName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "openid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "region",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "unionid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "query tag是query参数别名，json xml，form适合post",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseHTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResponseHTTP": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "required": [
                "appid",
                "appsecret",
                "code",
                "loginType",
                "password",
                "username"
            ],
            "properties": {
                "appid": {
                    "type": "string"
                },
                "appsecret": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "loginType": {
                    "type": "string",
                    "enum": [
                        "1",
                        "2",
                        "3"
                    ]
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LoginMobile": {
            "type": "object",
            "required": [
                "captcha",
                "mobile"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                }
            }
        },
        "models.ShopSort": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "shopSubSort": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ShopSubSort"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ShopSubSort": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "shopSortID": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "index": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "openid": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "unionid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "description": "query tag是query参数别名，json xml，form适合post",
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
