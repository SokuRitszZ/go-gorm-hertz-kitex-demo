package db

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, name string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("name = ?", name).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
