package models

import (
	"errors"
	"time"
)

type RegisterRequest struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	EmailVerified bool   `json:"emailVerified"`
}

type CreateAccountRequest struct {
	UserName      string      `json:"username"`
	Email         []EmailType `json:"email"`
	Password      string      `json:"password"`
	EmailVerified bool        `json:"emailVerified"`
}

func (r *RegisterRequest) Validate() error {
	if r.Username == "" || r.Email == "" || r.Password == "" {
		return errors.New("username, email, and password are required")
	} else if len(r.Username) == 0 || len(r.Email) == 0 || len(r.Password) == 0 {
		return errors.New("username, email, and password are required")
	}
	return nil
}

type RegisterResponse struct {
	IsPosted    bool     `json:"isPosted"`
	Data        UserData `json:"data"`
	EmailExists bool     `json:"emailExists"`
}

type UserData struct {
	Uid         string      `json:"uid"`
	UserName    string      `json:"username"`
	Email       []EmailType `json:"email"`
	CreatedDate time.Time   `json:"createdDate"`
}

type EmailType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RegisterAPIResponse struct {
	Message string           `json:"message"`
	Status  bool             `json:"status"`
	Data    RegisterResponse `json:"data"`
}
