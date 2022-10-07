package domains

type Token struct {
	Code   string `json:"access_token"`
	Type   string `json:"token_type"`
	Expire int    `json:"expires_in"`
}
