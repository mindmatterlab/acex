package user

import (
	"context"

	"github.com/mindmatterlab/acex/internal/pkg/acexx"
	validationutil "github.com/mindmatterlab/acex/internal/pkg/util/validation"
	v1 "github.com/mindmatterlab/acex/pkg/api/usercenter/v1"
)

// Delete deletes a user from the database.
func (b *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	filters := map[string]any{"username": rq.Username}
	if !validationutil.IsAdminUser(acexx.FromUserID(ctx)) {
		filters["user_id"] = acexx.FromUserID(ctx)
	}

	return b.ds.Users().Delete(ctx, filters)
}
