package secret

import (
	"context"

	"github.com/mindmatterlab/acex/internal/pkg/acexx"
	v1 "github.com/mindmatterlab/acex/pkg/api/usercenter/v1"
)

// Delete deletes a secret.
func (b *secretBiz) Delete(ctx context.Context, rq *v1.DeleteSecretRequest) error {
	return b.ds.Secrets().Delete(ctx, acexx.FromUserID(ctx), rq.Name)
}
