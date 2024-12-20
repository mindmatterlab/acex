package secret

import (
	"context"

	"github.com/mindmatterlab/go-pro/internal/pkg/gcontext"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// Delete deletes a secret.
func (b *secretBiz) Delete(ctx context.Context, rq *v1.DeleteSecretRequest) error {
	return b.ds.Secrets().Delete(ctx, gcontext.FromUserID(ctx), rq.Name)
}
