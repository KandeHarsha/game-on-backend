package loginradius

import (
	"KandeHarsha/config"
	"KandeHarsha/models"
	"KandeHarsha/service/loginradius/lib"
	"KandeHarsha/service/loginradius/schema"
	"context"
	"net/http"
	"sync"
)

type Config struct {
	ApiBaseURL string
	ApiKey     string
	ApiSecret  string
}

var loginradiusInstance *Config
var loginradiusOnce sync.Once

func GetInstance() *Config {
	loginradiusOnce.Do(func() {
		configInstance := config.GetInstance()
		loginradiusInstance = &Config{
			ApiBaseURL: configInstance.TenantAPIEndPoint,
			ApiKey:     configInstance.TenantKey,
			ApiSecret:  configInstance.TenantSecret,
		}
	})
	return loginradiusInstance
}

func (c *Config) getPath(path string) string {
	return c.ApiBaseURL + path
}

func (c *Config) Login(ctx context.Context, loginRequestModel *models.LoginRequest) (*schema.IdentityResponseWithToken, *schema.ErrorResponse) {
	r := lib.Request{
		Method: http.MethodPost,
		Path:   c.getPath("/identity/v2/auth/login"),
		Query: map[string]string{
			"apikey":           c.ApiKey,
			"invitation_token": loginRequestModel.InvitationToken,
		},
		Payload:  loginRequestModel,
		Response: &schema.IdentityResponseWithToken{},
	}
	if vErr := r.Do(); vErr != nil {
		return nil, vErr
	}
	response := r.Response.(*schema.IdentityResponseWithToken)
	return response, nil
}

func (c *Config) Register(ctx context.Context, registerRequestModel *models.RegisterRequest) (*models.RegisterAPIResponse, *schema.ErrorResponse) {
	lrRegisterRequest := models.CreateAccountRequest{
		Email: []models.EmailType{
			{
				Type:  "Primary",
				Value: registerRequestModel.Email,
			},
		},
		Password: registerRequestModel.Password,
		UserName: registerRequestModel.Username,
	}
	r := lib.Request{
		Method: http.MethodPost,
		Path:   c.getPath("/identity/v2/manage/account"),
		Query: map[string]string{
			"apikey":    c.ApiKey,
			"apisecret": c.ApiSecret,
		},
		Payload:  lrRegisterRequest,
		Response: &models.RegisterAPIResponse{},
	}
	if vErr := r.Do(); vErr != nil {
		return nil, vErr
	}
	response := r.Response.(*models.RegisterAPIResponse)
	return response, nil
}
