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
        "/account/": {
            "get": {
                "description": "will respond user account info when user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "get login user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Account"
                        }
                    }
                }
            }
        },
        "/account/login": {
            "get": {
                "description": "login an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "email password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginCredential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.tokenSample"
                        }
                    }
                }
            }
        },
        "/account/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "register a user",
                "parameters": [
                    {
                        "description": "email password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginCredential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Account"
                        }
                    }
                }
            }
        },
        "/card/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "get cards of user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Card"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "create a card",
                "parameters": [
                    {
                        "description": "create card",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/card.cardBinding"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/ent.Card"
                        }
                    }
                }
            }
        },
        "/card/:id": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "edit a card",
                "parameters": [
                    {
                        "description": "edit card",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/card.cardBinding"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Card"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "delete a card",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/management/car": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "add a car",
                "parameters": [
                    {
                        "description": "car info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Car"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/ent.Car"
                        }
                    }
                }
            }
        },
        "/management/car/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "update a car",
                "parameters": [
                    {
                        "type": "string",
                        "name": "brand",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "carType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "color",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "model",
                        "in": "query"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "update a car",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/management/car/:id": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "update a car",
                "parameters": [
                    {
                        "description": "car info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Car"
                        }
                    }
                }
            }
        },
        "/management/location": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "add a location",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Location"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "add a location",
                "parameters": [
                    {
                        "description": "location info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Location"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/ent.Location"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "delete a location",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/management/location/:id": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "management"
                ],
                "summary": "update a location",
                "parameters": [
                    {
                        "description": "location info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Location"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Location"
                        }
                    }
                }
            }
        },
        "/user/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "read a user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update a user info",
                "parameters": [
                    {
                        "description": "updated info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginCredential": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.tokenSample": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "card.cardBinding": {
            "type": "object",
            "required": [
                "cardholder_name",
                "number",
                "valid_until"
            ],
            "properties": {
                "cardholder_name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "valid_until": {
                    "type": "string"
                }
            }
        },
        "ent.Account": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AccountQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.AccountEdges"
                },
                "email": {
                    "description": "Email holds the value of the \"email\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "is_admin": {
                    "description": "IsAdmin holds the value of the \"is_admin\" field.",
                    "type": "boolean"
                }
            }
        },
        "ent.AccountEdges": {
            "type": "object",
            "properties": {
                "user": {
                    "description": "User holds the value of the user edge.",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Billing": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BillingQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.BillingEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "status": {
                    "description": "Status holds the value of the \"status\" field.",
                    "type": "string"
                }
            }
        },
        "ent.BillingEdges": {
            "type": "object",
            "properties": {
                "booking": {
                    "description": "Booking holds the value of the booking edge.",
                    "$ref": "#/definitions/ent.Booking"
                },
                "card": {
                    "description": "Card holds the value of the card edge.",
                    "$ref": "#/definitions/ent.Card"
                },
                "user": {
                    "description": "User holds the value of the user edge.",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Booking": {
            "type": "object",
            "properties": {
                "booking_status": {
                    "description": "BookingStatus holds the value of the \"booking_status\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BookingQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.BookingEdges"
                },
                "end_at": {
                    "description": "EndAt holds the value of the \"end_at\" field.",
                    "type": "string"
                },
                "fuel_level_at_begin": {
                    "description": "FuelLevelAtBegin holds the value of the \"fuel_level_at_begin\" field.",
                    "type": "number"
                },
                "fuel_level_at_end": {
                    "description": "FuelLevelAtEnd holds the value of the \"fuel_level_at_end\" field.",
                    "type": "number"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "mileage_begin": {
                    "description": "MileageBegin holds the value of the \"mileage_begin\" field.",
                    "type": "integer"
                },
                "mileage_end": {
                    "description": "MileageEnd holds the value of the \"mileage_end\" field.",
                    "type": "integer"
                },
                "return_car_at": {
                    "description": "ReturnCarAt holds the value of the \"return_car_at\" field.",
                    "type": "string"
                },
                "start_at": {
                    "description": "StartAt holds the value of the \"start_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.BookingEdges": {
            "type": "object",
            "properties": {
                "billing": {
                    "description": "Billing holds the value of the billing edge.",
                    "$ref": "#/definitions/ent.Billing"
                },
                "car": {
                    "description": "Car holds the value of the car edge.",
                    "$ref": "#/definitions/ent.Car"
                },
                "user": {
                    "description": "User holds the value of the user edge.",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Car": {
            "type": "object",
            "properties": {
                "brand": {
                    "description": "Brand holds the value of the \"brand\" field.",
                    "type": "string"
                },
                "car_type": {
                    "description": "CarType holds the value of the \"car_type\" field.",
                    "type": "string"
                },
                "color": {
                    "description": "Color holds the value of the \"color\" field.",
                    "type": "string"
                },
                "deposit": {
                    "description": "Deposit holds the value of the \"deposit\" field.",
                    "type": "number"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CarQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.CarEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "mileage": {
                    "description": "Mileage holds the value of the \"mileage\" field.",
                    "type": "integer"
                },
                "model": {
                    "description": "Model holds the value of the \"model\" field.",
                    "type": "string"
                },
                "plate_country": {
                    "description": "PlateCountry holds the value of the \"plate_country\" field.",
                    "type": "string"
                },
                "plate_number": {
                    "description": "PlateNumber holds the value of the \"plate_number\" field.",
                    "type": "string"
                },
                "price": {
                    "description": "Price holds the value of the \"price\" field.",
                    "type": "number"
                },
                "status": {
                    "description": "Status holds the value of the \"status\" field.",
                    "type": "string"
                },
                "unit_price": {
                    "description": "UnitPrice holds the value of the \"unit_price\" field.",
                    "type": "number"
                },
                "year": {
                    "description": "Year holds the value of the \"year\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.CarEdges": {
            "type": "object",
            "properties": {
                "booking": {
                    "description": "Booking holds the value of the booking edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Booking"
                    }
                },
                "location": {
                    "description": "Location holds the value of the location edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Location"
                    }
                }
            }
        },
        "ent.Card": {
            "type": "object",
            "properties": {
                "cardholder_name": {
                    "description": "CardholderName holds the value of the \"cardholder_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CardQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.CardEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "number": {
                    "description": "Number holds the value of the \"number\" field.",
                    "type": "string"
                },
                "valid_until": {
                    "description": "ValidUntil holds the value of the \"valid_until\" field.",
                    "type": "string"
                }
            }
        },
        "ent.CardEdges": {
            "type": "object",
            "properties": {
                "owner": {
                    "description": "Owner holds the value of the owner edge.",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Flaw": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.Location": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the LocationQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.LocationEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "latitude": {
                    "description": "Latitude holds the value of the \"latitude\" field.",
                    "type": "number"
                },
                "longitude": {
                    "description": "Longitude holds the value of the \"longitude\" field.",
                    "type": "number"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.LocationEdges": {
            "type": "object",
            "properties": {
                "cars": {
                    "description": "Cars holds the value of the cars edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Car"
                    }
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address holds the value of the \"address\" field.",
                    "type": "string"
                },
                "birthday": {
                    "description": "Birthday holds the value of the \"birthday\" field.",
                    "type": "string"
                },
                "driver_license_country": {
                    "description": "DriverLicenseCountry holds the value of the \"driver_license_country\" field.",
                    "type": "string"
                },
                "driver_license_id": {
                    "description": "DriverLicenseID holds the value of the \"driver_license_id\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.UserEdges"
                },
                "first_name": {
                    "description": "FirstName holds the value of the \"first_name\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.",
                    "type": "string"
                },
                "postal_code": {
                    "description": "PostalCode holds the value of the \"postal_code\" field.",
                    "type": "string"
                },
                "tel": {
                    "description": "Tel holds the value of the \"tel\" field.",
                    "type": "string"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "Account holds the value of the account edge.",
                    "$ref": "#/definitions/ent.Account"
                },
                "booking": {
                    "description": "Booking holds the value of the booking edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Booking"
                    }
                },
                "cards": {
                    "description": "Card holds the value of the card edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Card"
                    }
                },
                "flaws": {
                    "description": "NoteFlaws holds the value of the note_flaws edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Flaw"
                    }
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
