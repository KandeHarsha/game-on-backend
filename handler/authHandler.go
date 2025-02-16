package handler

import (
	"KandeHarsha/logic"
	"KandeHarsha/models"
	"KandeHarsha/service/loginradius"
	"context"
	"net/http"

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
	group.POST("/register", h.Register)
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

func (h *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.RegisterAPIResponse{
			Message: "Invalid request body: " + err.Error(),
			Status:  false,
		})
		return
	} else if vErr := request.Validate(); vErr != nil {
		c.JSON(403, gin.H{
			"error": vErr.Error(),
		})
		return
	}

	response, err := loginradius.GetInstance().Register(c, &request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegisterAPIResponse{
			Message: "Error registering user: " + err.Message,
			Status:  false,
		})
		return
	}

	if response.Data.EmailExists {
		c.JSON(http.StatusConflict, models.RegisterAPIResponse{
			Message: "Email already exists",
			Status:  false,
		})
		return
	}

	c.JSON(http.StatusOK, models.RegisterAPIResponse{
		Message: "User registered successfully",
		Status:  true,
		Data:    response.Data,
	})

}
