package domains

type Auth struct {
	Id       string `json:"client_id"`
	Secret   string `json:"client_secret"`
	Type     string `json:"grant_type"`
	Username string `json:"username"`
	Password string `json:"password"`
}
