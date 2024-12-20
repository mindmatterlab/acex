package user

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/mindmatterlab/go-pro/internal/pkg/gcontext"
	validationutil "github.com/mindmatterlab/go-pro/internal/pkg/util/validation"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// Get retrieves a single user from the database.
func (b *userBiz) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.UserReply, error) {
	filters := map[string]any{"username": rq.Username}
	if !validationutil.IsAdminUser(gcontext.FromUserID(ctx)) {
		filters["user_id"] = gcontext.FromUserID(ctx)
	}

	userM, err := b.ds.Users().Fetch(ctx, filters)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorUserNotFound(err.Error())
		}

		return nil, err
	}

	return ModelToReply(userM), nil
}
