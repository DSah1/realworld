package user

import "awesomeProject/model"

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string2 string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error

	IsFollower(uint, uint) (bool, error)
	AddFollower(*model.User, uint) error
	RemoveFollower(*model.User, uint) error
}
