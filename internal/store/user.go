package store

import (
	"awesomeProject/internal/model"
	"errors"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) error {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Updates(u).Error
}

func (us *UserStore) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := us.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var user model.User

	if err := us.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) IsFollower(userID uint, followerID uint) (bool, error) {
	var follow model.Follow

	if err := us.db.Where("following_id = ? AND follower_id = ?", userID, followerID).First(&follow).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (us *UserStore) AddFollower(u *model.User, followerID uint) error {
	return us.db.Model(&u).Association("Followers").Append(&model.Follow{FollowerID: followerID, FollowingID: u.ID})
}

func (us *UserStore) RemoveFollower(u *model.User, followerID uint) error {
	err := us.db.Exec("delete FROM 'follows' WHERE 'follower_id'=? and 'following_id'=?", followerID, u.ID).Error
	if err != nil {
		return err
	}
	return nil
}
