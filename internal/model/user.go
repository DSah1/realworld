package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not nul"`
	Email     string `gorm:"uniqueIndex;not nul"`
	Password  string `gorm:"uniqueIndex;not nul"`
	Bio       *string
	Image     *string
	Followers []Follow  `gorm:"foreignKey:FollowingID"`
	Following []Follow  `gorm:"foreignKey:FollowerID"`
	Favorites []Article `gorm:"many2many:favorites"`
}

type Follow struct {
	Follower    User
	FollowerID  uint `gorm:"primaryKey" sql:"type:int not null"`
	Following   User
	FollowingID uint `gorm:"primaryKey" sql:"type:int not null"`
}

func (Follow) TableName() string {
	return "follows"
}

func (u *User) HashPassword(plainPassword string) error {
	if len(plainPassword) == 0 {
		u.Password = ""
		return errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(h)
	return nil
}

func (u *User) CheckPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}

func (u *User) followedBy(id uint) bool {
	if u.Followers == nil {
		return false
	}
	for _, f := range u.Followers {
		if f.FollowerID == id {
			return true
		}
	}
	return false
}
