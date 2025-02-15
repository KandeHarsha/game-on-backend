package handler

import (
	"KandeHarsha/logic"
	"KandeHarsha/models"
	"context"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginLogic *logic.LoginLogic
}

func NewAuthHandler(group *gin.RouterGroup) {
	h := &AuthHandler{
		loginLogic: logic.NewLoginLogic(),
	}

	group.POST("/login", h.Login)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest

	if err := ctx.BindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	} else if vErr := loginReq.Validate(); vErr != nil {
		ctx.JSON(400, gin.H{
			"error": vErr.Error(),
		})
		return
	}

	res, rErr := h.loginLogic.Login(context.Background(), &loginReq)
	if rErr != nil {
		ctx.JSON(403, gin.H{
			"error": rErr.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}
