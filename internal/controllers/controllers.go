package controllers

import "github.com/gin-gonic/gin"

var (
	errBadInput      = gin.H{"outcome": false, "message": "Bad input"}
	errSlugInUse     = gin.H{"outcome": false, "message": "Slug in use"}
	errRequestFailed = gin.H{"outcome": false, "message": "Request Failed"}
	errNotFound      = gin.H{"outcome": false, "message": "Not Found"}
	errNone          = gin.H{"outcome": true, "message": nil}
)
