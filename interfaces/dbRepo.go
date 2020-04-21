package interfaces

type DBRepo interface {
	QueryOne(model interface{}, query string, params ...interface{}) (err error)
	Query(model interface{}, query string, params ...interface{}) (err error)
	Exec(query string, params ...interface{}) (err error)
	ExecOne(query string, params ...interface{}) (err error)
}
