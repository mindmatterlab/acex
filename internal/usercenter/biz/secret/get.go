package secret

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/mindmatterlab/acex/internal/pkg/acexx"
	v1 "github.com/mindmatterlab/acex/pkg/api/usercenter/v1"
)

// Get returns a single secret.
func (b *secretBiz) Get(ctx context.Context, rq *v1.GetSecretRequest) (*v1.SecretReply, error) {
	secretM, err := b.ds.Secrets().Get(ctx, acexx.FromUserID(ctx), rq.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorSecretNotFound(err.Error())
		}

		return nil, err
	}

	return ModelToReply(secretM), nil
}
