package ports

import "os"

type Config interface {
	GetGroup(group string) (map[string]interface{}, error)
	GetField(string, string) (string, error)
}

type AuthJWT interface {
	Get(string, string, int) (string, int64, error)
}

type Log interface {
	GetFile() *os.File
	Info(string)
	Error(string)
}
