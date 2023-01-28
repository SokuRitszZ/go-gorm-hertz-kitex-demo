package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"ghkd/kitex_gen/user"
	"ghkd/module/user/model/db"
	"ghkd/pkg/errno"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx}
}

func (s *CreateUserService) Do(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Name)
	if err != nil {
		return err
	}
	fmt.Println(119911)
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	fmt.Println(220022)
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	fmt.Println(334433)
	pswd := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Name: req.Name,
		Password: pswd,
	}})
}

