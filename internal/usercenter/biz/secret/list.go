package secret

import (
	"context"

	"github.com/mindmatterlab/go-pro/internal/pkg/gcontext"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// List returns a list of secrets.
func (b *secretBiz) List(ctx context.Context, rq *v1.ListSecretRequest) (*v1.ListSecretResponse, error) {
	count, list, err := b.ds.Secrets().List(ctx, gcontext.FromUserID(ctx))
	if err != nil {
		return nil, err
	}

	secrets := make([]*v1.SecretReply, 0)
	for _, item := range list {
		secrets = append(secrets, ModelToReply(item))
	}

	return &v1.ListSecretResponse{TotalCount: count, Secrets: secrets}, nil
}
