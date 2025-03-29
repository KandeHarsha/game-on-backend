package models

import "time"

type Organization struct {
	Id           string         `json:"Id"`
	Name         string         `json:"Name"`
	CreatedDate  time.Time      `json:"CreatedDate"`
	IsActive     bool           `json:"IsActive"`
	ModifiedDate time.Time      `json:"ModifiedDate"`
	Display      *Display       `json:"Display"`
	Domains      []Domain       `json:"Domains"`
	Metadata     map[string]any `json:"Metadata"`
	Connections  []Connection   `json:"Connections"`
	Policies     *Policy        `json:"Policies"`
}

type Display struct {
	LogoURL string `json:"LogoURL"`
	Name    string `json:"Name"`
}

type Domain struct {
	Id                   string `json:"Id"`
	DomainName           string `json:"DomainName"`
	IsVerified           bool   `json:"IsVerified"`
	VerificationStrategy string `json:"VerificationStrategy"`
	VerificationToken    string `json:"VerificationToken"`
}

type Connection struct {
	// IDPEntityId string `json:"IDPEntityId"`
	// IDPMetadataUrl string `json:"IDPMetadataUrl"`
	// IsIDPInitiated bool `json:"IsIDPInitiated"`
	// IDPCertificate Certificate `json:"IDPCertificate"`
	// IsActive bool `json:"IsActive"`
	// CreatedDate time.Time `json:"CreatedDate"`
	// ModifiedDate time.Time `json:"ModifiedDate"`
	// Metadata map[string]string `json:"Metadata"`
}

type Policy struct {
	// Id string `json:"Id"`
	// Name string `json:"Name"`
	// Description string `json:"Description"`
	// IsActive bool `json:"IsActive"`
	// CreatedDate time.Time `json:"CreatedDate"`
	// ModifiedDate time.Time `json:"ModifiedDate"`
	// Metadata map[string]string `json:"Metadata"`
}

type OrganizationData struct {
	Data []Organization
}

type CreateOrgRequest struct {
	Name     string         `json:"Name"`
	Metadata map[string]any `json:"Metadata,omitempty"`
}

type CreateOrgResponse struct {
	Organization
}

type GetOrganizationResponse struct {
	Data Organization `json:"Data"`
}

type AddUserToOrganizationRequest struct {
	RoleIds []string `json:"roleIds"`
}

type AddUserToOrganizationResponse struct {
	Data []OrganizationUser `json:"Data"`
}

type OrganizationUser struct {
	Id          string    `json:"Id"`
	Uid         string    `json:"Uid"`
	RoleId      string    `json:"RoleId"`
	OrgId       string    `json:"OrgId"`
	CreatedDate time.Time `json:"CreatedDate"`
}

type UpdateOrgRequest struct {
	Metadata map[string]any `json:"Metadata,omitempty"`
}

type UpdateOrgRespnse struct {
	Organization
}
