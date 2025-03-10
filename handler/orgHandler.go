package handler

import (
	"KandeHarsha/logic"

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
