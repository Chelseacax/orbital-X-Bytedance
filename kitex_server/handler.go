package main

import (
	"context"
	management "github.com/Chelseacax/orbital-X-Bytedance/hertz_server/kitex_gen/user/management"
)

// UserManagementImpl implements the last service interface defined in the IDL.
type UserManagementImpl struct{}

// QueryUser implements the UserManagementImpl interface.
func (s *UserManagementImpl) QueryUser(ctx context.Context, req *management.QueryUserRequest) (resp *management.QueryUserResponse, err error) {
	// TODO: Your code here...
	return
}

// AddUser implements the UserManagementImpl interface.
func (s *UserManagementImpl) AddUser(ctx context.Context, req *management.AddUserRequest) (resp *management.AddUserResponse, err error) {
	// TODO: Your code here...
	return
}
