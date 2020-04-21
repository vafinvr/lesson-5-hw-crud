package pgdb

import (
	"github.com/go-pg/pg/v9"
	"lesson-5-hw-crud/common"
)

type pgdb struct {
	conn *pg.DB
	log  common.Logger
}

func New(addr, db, user, password, dbNetwork string, log common.Logger) *pgdb {
	conn := connect(addr, db, user, password, dbNetwork)

	return &pgdb{
		conn: conn,
		log:  log,
	}
}

// Connect - connect to database
func connect(addr, db, user, password, dbNetwork string) *pg.DB {
	return pg.Connect(&pg.Options{
		Network:  dbNetwork,
		User:     user,
		Addr:     addr,
		Database: db,
		Password: password,
	})
}

// Disconnect - close db
func disconnect(db *pg.DB) error {
	return db.Close()
}

func (db *pgdb) QueryOne(model interface{}, query string, params ...interface{}) (err error) {
	s, err := db.conn.Prepare(query)
	if err != nil {
		return
	}
	defer s.Close()
	_, err = s.QueryOne(model, params...)
	return
}

func (db *pgdb) Query(model interface{}, query string, params ...interface{}) (err error) {
	s, err := db.conn.Prepare(query)
	if err != nil {
		return
	}
	defer s.Close()
	_, err = s.Query(model, params...)
	return
}

func (db *pgdb) Exec(query string, params ...interface{}) (err error) {
	s, err := db.conn.Prepare(query)
	if err != nil {
		return
	}
	defer s.Close()
	_, err = s.Exec(params...)
	return
}

func (db *pgdb) ExecOne(query string, params ...interface{}) (err error) {
	s, err := db.conn.Prepare(query)
	if err != nil {
		return
	}
	defer s.Close()
	_, err = s.ExecOne(params...)
	return
}
