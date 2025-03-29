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
	group.GET("/user/:uid", h.getUser)
	group.GET("/profile", h.profileByAccesstoken)
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

func (h *AuthHandler) getUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "uid is required",
		})
		return
	}
	user, err := h.loginLogic.GetUserByUid(context.Background(), uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, user)
}

func (h *AuthHandler) profileByAccesstoken(ctx *gin.Context) {
	token := ctx.Query("access_token")
	if len(token) == 0 {
		ctx.JSON(403, gin.H{
			"error": "access_token is required",
		})
	}

	resp, err := h.loginLogic.GetProfile(context.Background(), token)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, resp)
}
