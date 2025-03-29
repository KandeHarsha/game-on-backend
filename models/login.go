package models

import (
	"KandeHarsha/service/loginradius/schema"
	"errors"
)

type LoginRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	InvitationToken string `json:"invitation_token,omitempty"`
}

func (l *LoginRequest) Validate() error {
	if l.Email == "" || l.Password == "" {
		return errors.New("email and password are required")
	} else if len(l.Email) == 0 || len(l.Password) == 0 {
		return errors.New("email and password are required")
	}
	return nil
}

type LoginResponse struct {
	AccessToken  string                  `json:"access_token"`
	Profile      schema.IdentityResponse `json:"profile"`
	RefreshToken string                  `json:"refresh_token"`
}

type User struct {
	UserName     string
	Uid          string
	Roles        []string
	FullName     string
	Email        []EmailType
	CustomFields map[string]string
}
