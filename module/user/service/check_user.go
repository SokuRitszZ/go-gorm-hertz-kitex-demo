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

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx}
}

func (s *CheckUserService) Do(req *user.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	pswd := fmt.Sprintf("%x", h.Sum(nil))
	
	name := req.Name
	users, err := db.QueryUser(s.ctx, name)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationErr
	}
	u := users[0]
	if u.Password != pswd {
		return 0, errno.AuthorizationErr
	}
	return int64(u.ID), nil
}
