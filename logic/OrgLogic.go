package logic

import (
	"KandeHarsha/models"
	"KandeHarsha/service/loginradius"
	"context"
	"errors"
	"sync"
)

type OrgLogic struct {
	loginRadius *loginradius.Config
}

var (
	orgLogicInstance *OrgLogic
	orgLogicOnce     sync.Once
)

func NewOrgLogic() *OrgLogic {
	orgLogicOnce.Do(func() {
		orgLogicInstance = &OrgLogic{
			loginRadius: loginradius.GetInstance(),
		}
	})
	return orgLogicInstance
}

func (o *OrgLogic) GetAllOrgs(ctx context.Context) (interface{}, error) {
	resp, vErr := orgLogicInstance.loginRadius.GetAllOrgs(ctx)
	if vErr != nil {
		return nil, errors.New((vErr.Description))
	}
	return models.OrganizationData{
		Data: resp.Data,
	}, nil
}
