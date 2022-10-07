package handlers

type userMysql struct {
}

func NewUserMysql() *userMysql {
	return &userMysql{}
}

func (u *userMysql) GetActive(email string) (password string, err error) {
	return "8dasd8sd8asd8eyyy", nil
}
