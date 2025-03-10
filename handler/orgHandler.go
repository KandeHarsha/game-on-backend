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
	group.GET("/org/:orgId", h.GetOrg)
	group.POST("/org", h.CreateOrg)
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
