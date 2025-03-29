package logic

import (
	"KandeHarsha/config"
	"KandeHarsha/models"
	"KandeHarsha/service/loginradius"
	"context"
	"errors"
	"fmt"
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
	fmt.Println("resp", resp)
	if err != nil {
		return nil, err
	}
	return models.LoginResponse{
		AccessToken:  token,
		Profile:      resp.Profile,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (l *LoginLogic) GetUserByUid(ctx context.Context, uid string) (interface{}, error) {
	resp, vErr := loginLogicInstance.loginRadius.GetUserByUid(ctx, uid)
	if vErr != nil {
		return nil, errors.New((vErr.Description))
	}
	return resp, nil
}

func (l *LoginLogic) GetProfile(ctx context.Context, token string) (interface{}, error) {
	accesstokenJwt, lErr := l.loginRadius.GetProfileByToken(ctx, token)
	if lErr != nil {
		return nil, errors.New((lErr.Description))
	}
	return accesstokenJwt, nil
}
