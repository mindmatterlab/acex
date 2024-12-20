package user

import (
	"context"

	"github.com/mindmatterlab/go-pro/internal/pkg/gcontext"
	validationutil "github.com/mindmatterlab/go-pro/internal/pkg/util/validation"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// Delete deletes a user from the database.
func (b *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	filters := map[string]any{"username": rq.Username}
	if !validationutil.IsAdminUser(gcontext.FromUserID(ctx)) {
		filters["user_id"] = gcontext.FromUserID(ctx)
	}

	return b.ds.Users().Delete(ctx, filters)
}
