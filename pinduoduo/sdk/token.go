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

type PopAuthTokenResponse struct {
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
// 生成授权码,商家APP扫码授权

func (g *TokenAPI) PopAuthTokenCreate(mustParams ...Params) (resp *PopAuthTokenResponse, err error) {
	if err != nil {
		return nil, err
	}
	params := NewParamsWithType("pdd.pop.auth.token.create", mustParams...)
	params.Set("grant_type", "authorization_code")
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "pop_auth_token_create_response")
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// 刷新token
func (g *TokenAPI) PopAuthTokenRefresh(mustParams ...Params) (resp *PopAuthTokenResponse, err error) {
	params := NewParamsWithType("pdd.pop.auth.token.refresh", mustParams...)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "pop_auth_token_refresh_response")
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return
}
