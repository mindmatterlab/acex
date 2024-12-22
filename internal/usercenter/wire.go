//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package usercenter

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/mindmatterlab/acex/internal/pkg/bootstrap"
	"github.com/mindmatterlab/acex/internal/pkg/validation"
	"github.com/mindmatterlab/acex/internal/usercenter/auth"
	"github.com/mindmatterlab/acex/internal/usercenter/biz"
	"github.com/mindmatterlab/acex/internal/usercenter/server"
	"github.com/mindmatterlab/acex/internal/usercenter/service"
	"github.com/mindmatterlab/acex/internal/usercenter/store"
	customvalidation "github.com/mindmatterlab/acex/internal/usercenter/validation"
	"github.com/mindmatterlab/acex/pkg/db"
	genericoptions "github.com/mindmatterlab/acex/pkg/options"
)

// wireApp builds and returns a Kratos app with the given options.
// It uses the Wire library to automatically generate the dependency injection code.
func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
	*genericoptions.JWTOptions,
	*genericoptions.RedisOptions,
	*genericoptions.EtcdOptions,
	*genericoptions.KafkaOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		bootstrap.NewEtcdRegistrar,
		server.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		auth.ProviderSet,
		store.SetterProviderSet,
		NewAuthenticator,
		validation.ProviderSet,
		customvalidation.ProviderSet,
	)

	return nil, nil, nil
}
