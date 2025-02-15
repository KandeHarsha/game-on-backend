package logic

import (
	"KandeHarsha/config"
	"KandeHarsha/models"
	"KandeHarsha/service/loginradius"
	"context"
	"errors"
	"sync"
)

type LoginLogic struct {
	loginRadius      *loginradius.Config
	configInstance   *config.Config
	accessTokenLogic *AccessTokenLogic
}

var (
	loginLogicInstance *LoginLogic
	loginLogicOnce     sync.Once
)

func NewLoginLogic() *LoginLogic {
	loginLogicOnce.Do(func() {
		loginLogicInstance = &LoginLogic{
			loginRadius:      loginradius.GetInstance(),
			configInstance:   config.GetInstance(),
			accessTokenLogic: NewAccessTokenLogic(),
		}
	})
	return loginLogicInstance
}

func (l *LoginLogic) Login(ctx context.Context, loginRequestModel *models.LoginRequest) (interface{}, error) {
	resp, vErr := loginLogicInstance.loginRadius.Login(ctx, loginRequestModel)
	if vErr != nil {
		return nil, errors.New((vErr.Description))
	}
	token, err := l.accessTokenLogic.GenerateAccessToken(resp)
	if err != nil {
		return nil, err
	}
	return models.LoginResponse{
		AccessToken: token,
	}, nil
}
