package router

import (
	"com/example/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const EXP_ACCESS_COOKIE = 60 * 4

func (r *Router) Login(ctx *gin.Context) {
	// validate

	var req model.ReqLogin
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}

	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	userId, err := r.ctrl.Login(req)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	// TODO: Create token

	ctx.SetCookie("accessToken", userId, EXP_ACCESS_COOKIE, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "User logged in"})
}
