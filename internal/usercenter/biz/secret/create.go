package secret

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/mindmatterlab/go-pro/internal/pkg/gcontext"
	"github.com/mindmatterlab/go-pro/internal/usercenter/model"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// Create creates a new secret.
func (b *secretBiz) Create(ctx context.Context, rq *v1.CreateSecretRequest) (*v1.SecretReply, error) {
	var secretM model.SecretM
	_ = copier.Copy(&secretM, rq)
	secretM.UserID = gcontext.FromUserID(ctx)

	if err := b.ds.Secrets().Create(ctx, &secretM); err != nil {
		return nil, v1.ErrorSecretCreateFailed("create secret failed: %s", err.Error())
	}

	return ModelToReply(&secretM), nil
}
