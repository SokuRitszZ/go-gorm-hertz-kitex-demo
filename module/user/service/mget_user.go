package service

import (
	"context"
	"ghkd/kitex_gen/user"
	"ghkd/module/user/model/db"
	"ghkd/module/user/pack"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx}
}

func (s *MGetUserService) Do(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
