package domains

type AuthError struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}
