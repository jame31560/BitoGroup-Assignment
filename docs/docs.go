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
        "/add_user": {
            "post": {
                "description": "Add a new user to the matching system and find any possible matches for the new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add a person and match",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AddSinglePersonAndMatchReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.AddSinglePersonAndMatchRes"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/api.Err"
                        }
                    }
                }
            }
        },
        "/query_single_user": {
            "get": {
                "description": "Find the most N possible matched single people, where N is a request parameter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Query Single People.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "number to query",
                        "name": "num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.QuerySinglePeopleRes"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/api.Err"
                        }
                    }
                }
            }
        },
        "/remove_user": {
            "delete": {
                "description": "Remove a user from the matching system so that the user cannot be matched anymore.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Remove a user.",
                "parameters": [
                    {
                        "description": "Remove user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.RemoveSinglePersonReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RemoveSinglePersonRes"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/api.Err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AddSinglePersonAndMatchReq": {
            "type": "object",
            "properties": {
                "date_num": {
                    "type": "integer"
                },
                "gender": {
                    "type": "integer"
                },
                "height": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.AddSinglePersonAndMatchRes": {
            "type": "object",
            "properties": {
                "match_user_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "api.Err": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "api.QuerySinglePeopleRes": {
            "type": "object",
            "properties": {
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.UserRes"
                    }
                }
            }
        },
        "api.RemoveSinglePersonReq": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "api.RemoveSinglePersonRes": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "api.UserRes": {
            "type": "object",
            "properties": {
                "date_num": {
                    "type": "integer"
                },
                "gender": {
                    "type": "integer"
                },
                "height": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
