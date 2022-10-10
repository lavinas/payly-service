package domains

type AuthIn struct {
	Id       string `json:"client_id" form:"client_id"`
	Secret   string `json:"client_secret" form:"client_secret"`
	Type     string `json:"grant_type" form:"grant_type"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AuthError struct {
	Code        int    `json:"-"`
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

type AuthToken struct {
	Code   string `json:"access_token"`
	Type   string `json:"token_type"`
	Expire int64  `json:"expires_in"`
}
