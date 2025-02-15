package logic

import (
	"KandeHarsha/config"
	"KandeHarsha/entities"
	"KandeHarsha/service/loginradius/schema"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenLogic struct {
	configInstance *config.Config
}

var (
	accessTokenLogicInstance *AccessTokenLogic
	accessTokenLogicOnce     sync.Once
)

func NewAccessTokenLogic() *AccessTokenLogic {
	accessTokenLogicOnce.Do(func() {
		accessTokenLogicInstance = &AccessTokenLogic{
			configInstance: config.GetInstance(),
		}
	})
	return accessTokenLogicInstance
}

func (a *AccessTokenLogic) GenerateAccessToken(prop *schema.IdentityResponseWithToken) (string, error) {
	t := time.Now().UTC()
	accessTokenJwt := entities.AccessTokenJWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GameOn",
			Subject:   prop.Profile.ID,
			Audience:  jwt.ClaimStrings{"api://gameon"},
			ExpiresAt: jwt.NewNumericDate(prop.ExpiresIn),
			NotBefore: jwt.NewNumericDate(t),
			IssuedAt:  jwt.NewNumericDate(t),
			ID:        uuid.NewString(),
		},
		SessionID: prop.AccessToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenJwt)
	ss, err := token.SignedString([]byte(a.configInstance.TokenSignKey))

	return ss, err
}

func (a *AccessTokenLogic) ParseToken(tokenStr string) (*entities.AccessTokenJWT, error) {
	var accessTokenJwt entities.AccessTokenJWT
	token, err := jwt.ParseWithClaims(tokenStr, &accessTokenJwt, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.configInstance.TokenSignKey), nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return &accessTokenJwt, nil
}
