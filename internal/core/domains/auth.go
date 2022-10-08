package domains

type AuthIn struct {
	Id       string `json:"client_id"`
	Secret   string `json:"client_secret"`
	Type     string `json:"grant_type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthError struct {
	Code        int    `json:"-"`
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

type AuthToken struct {
	Code   string `json:"access_token"`
	Type   string `json:"token_type"`
	Expire int    `json:"expires_in"`
}
