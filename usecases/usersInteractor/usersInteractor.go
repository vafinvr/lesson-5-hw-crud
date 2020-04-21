package usersInteractor

import (
	"fmt"
	"lesson-5-hw-crud/domain"
)

type usersRepo interface {
	Create(user *domain.User) error
	Read(id int64) (user *domain.User, err error)
	Update(user *domain.User) error
	Delete(userID int64) error
}

type usersInteractor struct {
	usersRepo usersRepo
}

func New(repo usersRepo) *usersInteractor {
	return &usersInteractor{
		usersRepo: repo,
	}
}

func (u *usersInteractor) Create(user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, fmt.Errorf("empty data")
	}

	if err := u.usersRepo.Create(user); err != nil {
		return nil, err
	}

	return u.usersRepo.Read(user.ID)
}

func (u *usersInteractor) Read(id int64) (*domain.User, error) {
	return u.usersRepo.Read(id)
}

func (u *usersInteractor) Update(user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, fmt.Errorf("empty data")
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user id can't be zero")
	}

	if err := u.usersRepo.Update(user); err != nil {
		return nil, err
	}

	return u.usersRepo.Read(user.ID)
}

func (u *usersInteractor) Delete(id int64) error {
	return u.usersRepo.Delete(id)
}
