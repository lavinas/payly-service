package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lavinas/payly-service/internal/core/ports"
	"strconv"
	"time"
)

type cache struct {
	id      int
	pass    string
	expires int64
}

type userMySQL struct {
	db    *sql.DB
	cache map[string]cache
	time  int64
}

func NewUserMySQL(conf ports.Config) *userMySQL {
	return &userMySQL{
		db:    connect(conf),
		cache: make(map[string]cache),
		time:  cachetime(conf)}
}

func (u *userMySQL) GetActive(email string) (int, string, error) {
	c, ok := u.cache[email]
	if ok && c.expires >= time.Now().Unix() {
		return c.id, c.pass, nil
	}
	id, pass, err := getDBUser(email, u.db)
	if err != nil {
		return id, pass, err
	}
	deletecache(u.cache)
	u.cache[email] = cache{id: id, pass: pass, expires: time.Now().Unix() + u.time}
	return id, pass, nil
}

func cachetime(conf ports.Config) int64 {
	t, err := conf.GetField("db", "cache_time")
	if err != nil {
		panic("db configuration error: cachetime")
	}
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		panic("db configuration error: cachetime")
	}
	return i
}

func deletecache(c map[string]cache) {
	for key, element := range c {
		if element.expires < time.Now().Unix(){
			delete(c, key)
		}
	}
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

func getDBUser(email string, db *sql.DB) (int, string, error) {
	q := "select id, password from " + "user where email = '" + email + "' and status = 1;"
	r, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer r.Close()
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