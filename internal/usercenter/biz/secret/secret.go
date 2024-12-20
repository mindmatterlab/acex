package secret

//go:generate mockgen -self_package github.com/mindmatterlab/go-pro/internal/usercenter/biz/secret -destination mock_secret.go -package secret github.com/mindmatterlab/go-pro/internal/usercenter/biz/secret SecretBiz

import (
	"context"

	"github.com/mindmatterlab/go-pro/internal/usercenter/store"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// SecretBiz defines functions used to handle secret rquest.
type SecretBiz interface {
	Create(ctx context.Context, rq *v1.CreateSecretRequest) (*v1.SecretReply, error)
	List(ctx context.Context, rq *v1.ListSecretRequest) (*v1.ListSecretResponse, error)
	Get(ctx context.Context, rq *v1.GetSecretRequest) (*v1.SecretReply, error)
	Update(ctx context.Context, rq *v1.UpdateSecretRequest) error
	Delete(ctx context.Context, rq *v1.DeleteSecretRequest) error
}

type secretBiz struct {
	ds store.IStore
}

var _ SecretBiz = (*secretBiz)(nil)

// New creates a new instance of the secretBiz struct.
func New(ds store.IStore) *secretBiz {
	return &secretBiz{ds: ds}
}
