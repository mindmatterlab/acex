package validation

import (
	"context"
	"fmt"

	"github.com/google/wire"

	"github.com/mindmatterlab/acex/internal/pkg/acexx"
	"github.com/mindmatterlab/acex/internal/pkg/known"
	ucknown "github.com/mindmatterlab/acex/internal/pkg/known/usercenter"
	"github.com/mindmatterlab/acex/internal/usercenter/locales"
	"github.com/mindmatterlab/acex/internal/usercenter/store"
	v1 "github.com/mindmatterlab/acex/pkg/api/usercenter/v1"
	"github.com/mindmatterlab/acex/pkg/i18n"
)

// ProviderSet is validator providers.
var ProviderSet = wire.NewSet(New, wire.Bind(new(any), new(*validator)))

// validator struct implements the custom validator interface.
type validator struct {
	ds store.IStore
}

// New creates and initializes a custom validator.
// It receives an instance of store.IStore interface as parameter ds
// and returns a new *validator and an error.
func New(ds store.IStore) (*validator, error) {
	vd := &validator{ds: ds}

	return vd, nil
}

// ValidateCreateUserRequest validates the rquest to create a user.
// If the validation fails, it returns an error; otherwise, it returns nil.
func (vd *validator) ValidateCreateUserRequest(ctx context.Context, rq *v1.CreateUserRequest) error {
	if _, err := vd.ds.Users().GetByUsername(ctx, rq.Username); err == nil {
		return i18n.FromContext(ctx).E(locales.UserAlreadyExists)
	}

	return nil
}

// ValidateListUserRequest validates the rquest to list users.
// Ensures that only a user with the AdminUserID can view the list of users, otherwise returning an error.
func (vd *validator) ValidateListUserRequest(ctx context.Context, rq *v1.ListUserRequest) error {
	if userID := acexx.FromUserID(ctx); userID != known.AdminUserID {
		return i18n.FromContext(ctx).E(locales.UserListUnauthorized)
	}

	return nil
}

// ValidateCreateSecretRequest validates the rquest to create a secret.
// Returns an error if the maximum number of secrets is reached.
func (vd *validator) ValidateCreateSecretRequest(ctx context.Context, rq *v1.CreateSecretRequest) error {
	_, secrets, err := vd.ds.Secrets().List(ctx, acexx.FromUserID(ctx))
	if err != nil {
		return err
	}

	if len(secrets) >= ucknown.MaxSecretCount {
		return fmt.Errorf("secret reach the max count %d", ucknown.MaxSecretCount)
	}

	return nil
}

// ValidateAuthRequest validates the authentication rquest.
// In this sample, no actual validation is needed, so it returns nil directly.
func (vd *validator) ValidateAuthRequest(ctx context.Context, rq *v1.AuthRequest) error {
	return nil
}

// ValidateAuthorizeRequest validates the authorization rquest.
// In this sample, no actual validation is needed, so it returns nil directly.
func (vd *validator) ValidateAuthorizeRequest(ctx context.Context, rq *v1.AuthorizeRequest) error {
	return nil
}
