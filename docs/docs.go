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
        "/classwork": {
            "post": {
                "description": "Returns classwork for the marking periods specified.\nIf no marking periods are specified, the classwork for the current marking period is returned.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classwork"
                ],
                "parameters": [
                    {
                        "description": "Body Params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ClassworkRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ClassworkResponse"
                        }
                    }
                }
            }
        },
        "/ipr": {
            "post": {
                "description": "Returns the IPR(s) for the user. If the date parameter is not passed into the body or is invalid, the most recent IPR is returned.\nIt is important the format of the date follows the format \"01/02/2006\" (01 = month, 02 = day, 2006 = year), with leading zeros like shown in the format.\nFor all possible dates, refer to the \"/ipr/all\" endpoint.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ipr"
                ],
                "parameters": [
                    {
                        "description": "Body Params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.IprRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.IPRResponse"
                        }
                    }
                }
            }
        },
        "/ipr/all": {
            "post": {
                "description": "Returns all the IPRs for the user, or just the dates depending on the DatesOnly parameter's value in the body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ipr"
                ],
                "parameters": [
                    {
                        "description": "Body Params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.IprAllRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.IPRResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Pre-registers the user with the API by logging them into HAC early, and caching the cookies.\nSubsequent requests using the same credentials will use these stored cookies, leading to faster response times for other endpoints.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Body Params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    }
                }
            }
        },
        "/reportcard": {
            "post": {
                "description": "Returns report card data for the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reportcard"
                ],
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ReportCardRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ReportCardResponse"
                        }
                    }
                }
            }
        },
        "/schedule": {
            "post": {
                "description": "Returns the schedule for the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ScheduleRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ScheduleResponse"
                        }
                    }
                }
            }
        },
        "/transcript": {
            "post": {
                "description": "Returns the transcript for the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transcript"
                ],
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.TranscriptRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TranscriptResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Absences": {
            "type": "object",
            "properties": {
                "excusedAbsence": {
                    "description": "The amount of excused absences for the class",
                    "type": "string"
                },
                "excusedTardy": {
                    "description": "The amount of excused tardies for the class",
                    "type": "string"
                },
                "unexcusedAbsence": {
                    "description": "The amount of unexcused absences for the class",
                    "type": "string"
                },
                "unexcusedTardy": {
                    "description": "The amount of unexcused tardies for the class",
                    "type": "string"
                }
            }
        },
        "models.Assignment": {
            "type": "object",
            "properties": {
                "assignedDate": {
                    "description": "The date the assignment was assigned",
                    "type": "string"
                },
                "category": {
                    "description": "The category of the assignment (major, minor, other, etc)",
                    "type": "string"
                },
                "dropped": {
                    "description": "Whether the assignment was dropped or not",
                    "type": "boolean"
                },
                "dueDate": {
                    "description": "The date the assignment is due",
                    "type": "string"
                },
                "grade": {
                    "description": "What grade the user got on the assignment",
                    "type": "string"
                },
                "name": {
                    "description": "The name of the assignment",
                    "type": "string"
                },
                "totalPoints": {
                    "description": "The total points that could be earned on the assignment",
                    "type": "string"
                }
            }
        },
        "models.Class": {
            "type": "object",
            "properties": {
                "course": {
                    "description": "The course ID of the class",
                    "type": "string"
                },
                "name": {
                    "description": "The name of the class",
                    "type": "string"
                },
                "period": {
                    "description": "What period the class is for the student, relative to the current schedule",
                    "type": "string"
                },
                "room": {
                    "description": "The room number of the class",
                    "type": "string"
                },
                "teacher": {
                    "description": "The name of the teacher of the class",
                    "type": "string"
                }
            }
        },
        "models.Classwork": {
            "type": "object",
            "properties": {
                "entries": {
                    "description": "An array of ClassworkEntry structs containing classwork for each class",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ClassworkEntry"
                    }
                },
                "sixWeeks": {
                    "description": "The marking period the classwork is for",
                    "type": "integer"
                }
            }
        },
        "models.ClassworkEntry": {
            "type": "object",
            "properties": {
                "assignments": {
                    "description": "All the assignments currently entered for the class",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Assignment"
                    }
                },
                "average": {
                    "description": "The average grade for that class",
                    "type": "string"
                },
                "class": {
                    "description": "Class information about the entry",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Class"
                        }
                    ]
                },
                "position": {
                    "description": "The position of the class, used for ordering",
                    "type": "integer"
                }
            }
        },
        "models.ClassworkRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "markingPeriods": {
                    "description": "The marking period to pull data from",
                    "type": "array",
                    "maxItems": 6,
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.ClassworkResponse": {
            "type": "object",
            "properties": {
                "classwork": {
                    "description": "The resulting classwork",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Classwork"
                    }
                },
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                }
            }
        },
        "models.IPR": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "The date the IPR was submitted",
                    "type": "string"
                },
                "entries": {
                    "description": "An array representing all the IPR entries",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.IPREntry"
                    }
                }
            }
        },
        "models.IPREntry": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Information about the class related to the IPREntry",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Class"
                        }
                    ]
                },
                "grade": {
                    "description": "The average at the moment the progress report was submitted",
                    "type": "string"
                }
            }
        },
        "models.IPRResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "ipr": {
                    "description": "The resulting IPR(s)",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.IPR"
                    }
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                }
            }
        },
        "models.IprAllRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "datesOnly": {
                    "description": "Whether to return only dates or all the IPRs",
                    "type": "boolean",
                    "default": false,
                    "example": true
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.IprRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "date": {
                    "description": "The date of the IPR to return",
                    "type": "string",
                    "example": "09/06/2022"
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "base": {
                    "description": "The base URL signed in to",
                    "type": "string"
                },
                "password": {
                    "description": "The password used to sign in with",
                    "type": "string"
                },
                "username": {
                    "description": "The username used to sign in with",
                    "type": "string"
                }
            }
        },
        "models.LoginRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "login": {
                    "description": "Data about the login",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Login"
                        }
                    ]
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                }
            }
        },
        "models.ReportCard": {
            "type": "object",
            "properties": {
                "entries": {
                    "description": "All the report card entries",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ReportCardEntry"
                    }
                }
            }
        },
        "models.ReportCardEntry": {
            "type": "object",
            "properties": {
                "absences": {
                    "description": "Data about absences",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Absences"
                        }
                    ]
                },
                "attemptedCredit": {
                    "description": "The amount of credit attempted",
                    "type": "string"
                },
                "averages": {
                    "description": "Data about grades",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SixWeeksGrades"
                        }
                    ]
                },
                "class": {
                    "description": "Information about the class for the entry",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Class"
                        }
                    ]
                },
                "comments": {
                    "description": "Data about comments",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SixWeeksOther"
                        }
                    ]
                },
                "conduct": {
                    "description": "Data about conduct",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SixWeeksOther"
                        }
                    ]
                },
                "earnedCredit": {
                    "description": "The amount of credit earned",
                    "type": "string"
                }
            }
        },
        "models.ReportCardRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.ReportCardResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                },
                "reportCard": {
                    "description": "The resulting report card",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.ReportCard"
                        }
                    ]
                }
            }
        },
        "models.Schedule": {
            "type": "object",
            "properties": {
                "entries": {
                    "description": "An array containing all the schedule entries",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ScheduleEntry"
                    }
                }
            }
        },
        "models.ScheduleEntry": {
            "type": "object",
            "properties": {
                "active": {
                    "description": "Whether the class is active or not",
                    "type": "boolean"
                },
                "building": {
                    "description": "The building the class is in",
                    "type": "string"
                },
                "class": {
                    "description": "Information about the Class related to the Schedule",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Class"
                        }
                    ]
                },
                "days": {
                    "description": "The days the class is active for",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "markingPeriods": {
                    "description": "The marking periods the class is active for",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.ScheduleRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.ScheduleResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                },
                "schedule": {
                    "description": "The resulting schedule",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Schedule"
                        }
                    ]
                }
            }
        },
        "models.SixWeeksGrades": {
            "type": "object",
            "properties": {
                "exam1": {
                    "type": "string"
                },
                "exam2": {
                    "type": "string"
                },
                "fifth": {
                    "type": "string"
                },
                "first": {
                    "type": "string"
                },
                "fourth": {
                    "type": "string"
                },
                "second": {
                    "type": "string"
                },
                "sem1": {
                    "type": "string"
                },
                "sem2": {
                    "type": "string"
                },
                "sixth": {
                    "type": "string"
                },
                "third": {
                    "type": "string"
                }
            }
        },
        "models.SixWeeksOther": {
            "type": "object",
            "properties": {
                "fifth": {
                    "type": "string"
                },
                "first": {
                    "type": "string"
                },
                "fourth": {
                    "type": "string"
                },
                "second": {
                    "type": "string"
                },
                "sixth": {
                    "type": "string"
                },
                "third": {
                    "type": "string"
                }
            }
        },
        "models.Transcript": {
            "type": "object",
            "properties": {
                "entries": {
                    "description": "All the transcript groups",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TranscriptGroup"
                    }
                },
                "unweighted": {
                    "description": "The unweighted GPA",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.TranscriptGPA"
                        }
                    ]
                },
                "weighted": {
                    "description": "The weighted GPA",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.TranscriptGPA"
                        }
                    ]
                }
            }
        },
        "models.TranscriptGPA": {
            "type": "object",
            "properties": {
                "gpa": {
                    "description": "The GPA",
                    "type": "string"
                },
                "quartile": {
                    "description": "The class quartile",
                    "type": "string"
                },
                "rank": {
                    "description": "The class rank",
                    "type": "string"
                },
                "type": {
                    "description": "The type of GPA (unweighted/weighted)",
                    "type": "string"
                }
            }
        },
        "models.TranscriptGroup": {
            "type": "object",
            "properties": {
                "building": {
                    "description": "The building in which the classes in the group weree taken",
                    "type": "string"
                },
                "entries": {
                    "description": "The individual class entries for this group",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TranscriptGroupEntry"
                    }
                },
                "gradeLevel": {
                    "description": "The grade level the group represents data for",
                    "type": "string"
                },
                "semester": {
                    "description": "The semester the group represents data for",
                    "type": "string"
                },
                "totalCredit": {
                    "description": "The total credit earned for all the entries in this group",
                    "type": "string"
                },
                "year": {
                    "description": "The school year the group represents data for",
                    "type": "string"
                }
            }
        },
        "models.TranscriptGroupEntry": {
            "type": "object",
            "properties": {
                "average": {
                    "description": "The average grade for the class",
                    "type": "string"
                },
                "class": {
                    "description": "The class related to the entry",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Class"
                        }
                    ]
                },
                "credit": {
                    "description": "The credit earned for the class",
                    "type": "string"
                }
            }
        },
        "models.TranscriptRequestBody": {
            "type": "object",
            "required": [
                "base",
                "password",
                "username"
            ],
            "properties": {
                "base": {
                    "description": "The base URL for the PowerSchool HAC service",
                    "type": "string",
                    "minLength": 1,
                    "example": "https://homeaccess.katyisd.org"
                },
                "password": {
                    "description": "The password to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j382704"
                },
                "username": {
                    "description": "The username to log in with",
                    "type": "string",
                    "minLength": 1,
                    "example": "j1732901"
                }
            }
        },
        "models.TranscriptResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "If there was an error",
                    "type": "boolean"
                },
                "msg": {
                    "description": "The associated message",
                    "type": "string"
                },
                "transcript": {
                    "description": "The resulting transcript",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Transcript"
                        }
                    ]
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
