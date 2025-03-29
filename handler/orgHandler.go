package handler

import (
	"KandeHarsha/logic"
	"KandeHarsha/models"

	"github.com/gin-gonic/gin"
)

type OrgHandler struct {
	orgLogic *logic.OrgLogic
}

func NewOrgHandler(group *gin.RouterGroup) {
	h := &OrgHandler{
		orgLogic: logic.NewOrgLogic(),
	}

	group.GET("/orgs", h.GetAllOrgs)
	group.PUT("org/:orgId", h.updateOrg)
	group.GET("/org/:orgId", h.GetOrg)
	group.POST("/org", h.CreateOrg)
	group.PUT("user/:userId/org/:orgId", h.AddUserToOrg)
}

func (h *OrgHandler) GetAllOrgs(ctx *gin.Context) {
	res, err := h.orgLogic.GetAllOrgs(ctx)
	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, res)
}

func (h *OrgHandler) GetOrg(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	res, err := h.orgLogic.GetOrg(ctx, orgId)
	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, res)
}

func (h *OrgHandler) CreateOrg(ctx *gin.Context) {
	var orgRequest models.CreateOrgRequest
	if err := ctx.BindJSON(&orgRequest); err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	res, err := h.orgLogic.CreateOrg(ctx, &orgRequest)
	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, res)
}

func (h *OrgHandler) AddUserToOrg(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	userID := ctx.Param("userId")
	var assignRole models.AddUserToOrganizationRequest
	if err := ctx.BindJSON(&assignRole); err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	if orgId == "" || userID == "" {
		ctx.JSON(403, gin.H{
			"error": "orgId and userId are required",
		})
	}

	res, err := h.orgLogic.AddUserToOrg(ctx, &assignRole, orgId, userID)

	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, res)
}

func (h *OrgHandler) updateOrg(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	var updateOrg models.UpdateOrgRequest
	if err := ctx.BindJSON(&updateOrg); err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	if orgId == "" {
		ctx.JSON(403, gin.H{
			"error": "orgId is required",
		})
	}

	res, err := h.orgLogic.UpdateOrg(ctx, &updateOrg, orgId)

	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, res)
}
