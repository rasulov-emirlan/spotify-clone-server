// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rasulov Emirlan",
            "email": "rasulov-emirlan@gmail.com"
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
        "/auth/login": {
            "post": {
                "description": "Returns a json web token if user is registered in database and enters correct data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Authorization request",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.authRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "json web token"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registers a new user and returns his token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Authorization request",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.authRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "json web token"
                    }
                }
            }
        },
        "/genres": {
            "get": {
                "description": "Lists all the genres in our database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "genres"
                ],
                "summary": "List genres",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            },
            "put": {
                "description": "Adds a song to a genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "genres"
                ],
                "summary": "Add a song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for a genre",
                        "name": "genre",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id for a song",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we added a new song to the genre"
                    }
                }
            },
            "post": {
                "description": "Creates a new genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "genres"
                ],
                "summary": "Create a new genre",
                "parameters": [
                    {
                        "description": "A name for new genre",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.genresCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we created your genre"
                    }
                }
            }
        },
        "/genres/{genre}": {
            "get": {
                "description": "Adds a song to a genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "genres"
                ],
                "summary": "Add a song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id for a genre",
                        "name": "genre",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we added a new song to the genre"
                    }
                }
            }
        },
        "/playlists": {
            "get": {
                "description": "Lists all the playlists in our database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "List playlists",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            },
            "put": {
                "description": "Adds a song to whatever playlist you want to. But it has to be your playlist that you created",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "Add a song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWToken for auth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The id for the playlist",
                        "name": "playlist",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The id for the song",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we created your playlist"
                    }
                }
            },
            "post": {
                "description": "Creates a new playlist that can be accesed by anyone but only you can edit it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "Create a playlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWToken for auth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The name of the playlist",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "The name of the playlist",
                        "name": "cover",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we created your playlist"
                    }
                }
            }
        },
        "/playlists/{id}": {
            "get": {
                "description": "Gives you an array of json with songs from a playlist you want",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "Get Songs from playlist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The id for the playlist",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we created your playlist"
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Returns songs from some id to some id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get songs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "from which id to start",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "at which id to end",
                        "name": "to",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "here your songs"
                    }
                }
            },
            "post": {
                "description": "Uploads a song and its cover with all the info about that song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Upload a song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWToken for auth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "The actual audiofile",
                        "name": "audio",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "The cover for the song",
                        "name": "cover",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The name for that song",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "we uploaded your song"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.authRequest": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "2000-01-01"
                },
                "email": {
                    "type": "string",
                    "example": "john@gmai.com"
                },
                "full_name": {
                    "type": "string",
                    "example": "Johny Cash"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "Johnny"
                }
            }
        },
        "handlers.authResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "handlers.genresCreateRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.1",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Spotify Clone Server",
	Description: "This is a backend server for spotify clone.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
