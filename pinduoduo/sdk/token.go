package sdk

import (
	"encoding/json"
)

type TokenAPI struct {
	Context *Context
}

//func NewTokenAPI(cfg *Config) *TokenAPI {
//	return &TokenAPI{Context: NewContext(cfg)}
//}

func newTokenAPIWithContext(c *Context) *TokenAPI {
	return &TokenAPI{Context: c}
}

type PopAuthTokenCreateResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresAt             string `json:"expires_at"`
	ExpiresIn             string `json:"expires_in"`
	OwnerId               string `json:"owner_id"`
	OwnerName             string `json:"owner_name"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

//type TokenResponse struct {
//	PopAuthTokenCreateResponse PopAuthTokenCreateResponse `json:"pop_auth_token_create_response"`
//}

func (g *TokenAPI) TokenGet(code string) (res *PopAuthTokenCreateResponse, err error) {
	params := NewParamsWithType("pdd.pop.auth.token.create")
	params.Set("code", code)

	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "pop_auth_token_create_response")
	json.Unmarshal(bytes, &res)
	return
}
