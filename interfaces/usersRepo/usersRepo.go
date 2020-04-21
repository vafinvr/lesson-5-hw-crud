package usersRepo

import (
	"fmt"
	"lesson-5-hw-crud/domain"
	"lesson-5-hw-crud/interfaces"
)

type usersRepo struct {
	db interfaces.DBRepo
}

func New(db interfaces.DBRepo) *usersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) Create(user *domain.User) error {
	if user == nil {
		return fmt.Errorf("failed create: no data")
	}
	query := `INSERT INTO users
		(username, first_name, last_name, email, phone)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id`

	return r.db.QueryOne(&user.ID, query, user.Username, user.FirstName, user.LastName, user.Email, user.Phone)
}

func (r *usersRepo) Read(id int64) (user *domain.User, err error) {
	query := `SELECT
		username, first_name, last_name, email, phone
	FROM users
	WHERE id=$1`

	user = new(domain.User)
	err = r.db.QueryOne(user, query, id)

	return
}

func (r *usersRepo) Update(user *domain.User) error {
	if user == nil {
		return fmt.Errorf("failed update: no data")
	}
	query := `UPDATE users
	SET username=$2, first_name=$3, last_name=$4, email=$5, phone=$6
	WHERE id=$1`

	return r.db.ExecOne(query, user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Phone)
}

func (r *usersRepo) Delete(userID int64) error {
	query := `DELETE FROM users
	WHERE id=$1`

	return r.db.Exec(query, userID)
}
