package repositories

import (
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lavinas/payly-service/internal/core/ports"
)

type userMySQL struct {
	conf ports.Config
	db   *sql.DB
}

func NewUserMySQL(conf ports.Config) *userMySQL {
	return &userMySQL{conf: conf, db: connect(conf)}
}

func (u *userMySQL) GetActive(email string) (int, string, error) {
	q := "select id, password from " + "user where email = '" + email + "' and status = 1;"
	r, err := u.db.Query(q)
	if err != nil {
		panic(err)
	}
	if !r.Next() {
		return 0, "", errors.New("user not found")
	}
	var id int
	var pass string
	if err := r.Scan(&id, &pass); err != nil {
		if err != nil {
			panic(err)
		}
	}
	return id, pass, nil
}

func connect(conf ports.Config) *sql.DB {
	u, p, h, o, m := connectParams(conf)
	s := u + ":" + p + "@tcp(" + h + ")/" + m
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic("db connection error: " + err.Error())
	}
	db.SetMaxIdleConns(o)
	return db
}

func connectParams(conf ports.Config) (string, string, string, int, string) {
	dbConf, err := conf.GetGroup("db")
	if err != nil {
		panic("DB configuration error: db group")
	}
	u := getParam(dbConf, "user")
	p := getParam(dbConf, "pass")
	h := getParam(dbConf, "host")
	o := getParam(dbConf, "pool_size")
	m := getParam(dbConf, "main_db")
	getParam(dbConf, "shellbox_db")
	oi, err := strconv.Atoi(o)
	if err != nil {
		panic("db configuration error: pool_size is not int")
	}
	return u, p, h, oi, m
}

func getParam(g map[string]interface{}, p string) string {
	r, ok := g[p]
	if !ok {
		panic("db configuration error:" + p)
	}
	return r.(string)
}
