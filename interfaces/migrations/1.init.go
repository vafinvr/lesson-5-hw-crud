package migrations

func (m *migrations) init1() error {
	query := `CREATE TABLE "users" (
  "id" serial NOT NULL,
  "username" character varying(256) NOT NULL,
  "first_name" character varying(256) NOT NULL,
  "last_name" character varying(256) NOT NULL,
  "email" character varying(256) NOT NULL,
  "phone" character varying(256) NOT NULL
);`

	return m.db.ExecOne(query)
}
