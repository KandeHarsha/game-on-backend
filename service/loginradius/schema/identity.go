package schema

import "time"

type Identity struct {
	Provider  string  `bsonn:"provider" json:"provider"`
	FullName  string  `bsonn:"fullName" json:"fullName"`
	FirstName string  `bsonn:"firstName" json:"firstName"`
	LastName  string  `bsonn:"lastName" json:"lastName"`
	ID        string  `bsonn:"id" json:"id"`
	Email     []Email `bsonn:"email" json:"email"`
	Uid       string  `bsonn:"uid" json:"uid"`
}

type Email struct {
	Type  string `bsonn:"type" json:"type"`
	Value string `bsonn:"value" json:"value"`
}

type IdentityResponse struct {
	Identities []Identity `json:"Identities"`
	Identity
}

type IdentityResponseWithToken struct {
	AccessToken string           `json:"access_token"`
	Profile     IdentityResponse `json:"profile"`
	ExpiresIn   time.Time        `json:"expires_in"`
}
