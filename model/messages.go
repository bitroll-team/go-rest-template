package model

import "github.com/gin-gonic/gin"

func MsgBadRequest() gin.H {
	return gin.H{"msg": "Bad request"}
}

func MsgIntServerErr() gin.H {
	return gin.H{"msg": "Internal server error"}
}

func MsgValidationErr(err string) gin.H {
	return gin.H{
		"message": "Validation error",
		"errors":  err,
	}
}
