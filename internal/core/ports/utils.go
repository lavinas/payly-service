package ports

type Config interface {
	GetGroup(group string) (map[string]interface{}, error)
	GetField(string, string) (string, error)
}

type AuthJWT interface {
	Get(string, string, int) (string, int64, error)
}
