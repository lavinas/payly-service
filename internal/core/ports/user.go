package ports

type User interface {
	GetActive(email string) (string, error)
}
