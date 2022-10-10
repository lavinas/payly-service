package ports

type User interface {
	GetActive(email string) (int, string, error)
}
