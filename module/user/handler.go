package main

import (
	"context"
	"ghkd/kitex_gen/user"
	"ghkd/module/user/pack"
	"ghkd/module/user/service"
	"ghkd/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}
	
	err = service.NewCreateUserService(ctx).Do(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MGetUserResponse)
	
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}

	users, err := service.NewMGetUserService(ctx).Do(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return 
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}

	uID, err := service.NewCheckUserService(ctx).Do(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return 
	}
	
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uID
	return
}
