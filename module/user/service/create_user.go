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
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	pswd := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Name: req.Name,
		Password: pswd,
	}})
}

