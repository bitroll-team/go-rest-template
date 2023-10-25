package router

import (
	"com/example/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Register(ctx *gin.Context) {
	// validate

	var req model.ReqRegister
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}

	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	if err := r.ctrl.Register(req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User registered"})
}
