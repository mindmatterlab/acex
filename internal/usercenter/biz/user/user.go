package user

//go:generate mockgen -self_package github.com/mindmatterlab/go-pro/internal/usercenter/biz/user -destination mock_user.go -package user github.com/mindmatterlab/go-pro/internal/usercenter/biz/user UserBiz

import (
	"context"

	"github.com/mindmatterlab/go-pro/internal/usercenter/store"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// UserBiz defines methods used to handle user request.
type UserBiz interface {
	Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.UserReply, error)
	List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error)
	Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.UserReply, error)
	Update(ctx context.Context, rq *v1.UpdateUserRequest) error
	Delete(ctx context.Context, rq *v1.DeleteUserRequest) error

	// extensions apis
	UpdatePassword(ctx context.Context, rq *v1.UpdatePasswordRequest) error
}

// userBiz struct implements the UserBiz interface and contains a store.IStore instance.
type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

// New returns a new instance of userBiz.
func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
