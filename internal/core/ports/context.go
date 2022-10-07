package ports

type Context interface {
	BindJSON(obj any) error
	IndentedJSON(code int, obj any)
}
