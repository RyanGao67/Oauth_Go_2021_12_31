package model

type RealmAccess struct {
	Roles []string `json:"roles"`
}
type Account struct {
	Roles []string `json:"roles"`
}
type ResourceAccess struct {
	Account Account `json:"Account"`
}

type Tokenclaim struct {
	Jti               string         `json:"jti"`
	Exp               int            `json:"exp"`
	Nbf               int            `json:"nbf"`
	Iat               int            `json:"iat"`
	Iss               string         `json:"iss"`
	Aud               string         `json:"aud"`
	Sub               string         `json:"sub"`
	Typ               string         `json:"typ"`
	Azp               string         `json:"azp"`
	AuthTime          int            `json:"auth_time"`
	SessionState      string         `json:"session_state"`
	Acr               string         `json:"acr"`
	AllowOrigins      []string       `json:"allow_origins"`
	RealmAccess       RealmAccess    `json:"realm_access"`
	ResourceAccess    ResourceAccess `json:"resource_access"`
	Scope             string         `json:"scope"`
	EmailVerified     bool           `json:"email-verified"`
	Name              string         `json:"name"`
	PreferredUsername string         `json:"preferred_username"`
	GivenName         string         `json:"given_name"`
	FamilyName        string         `json:"family_name"`
	Email             string         `json:"email"`
}
