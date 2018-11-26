package models

// BodyResp is a model to process data from the JSON
type BodyResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiration  int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
