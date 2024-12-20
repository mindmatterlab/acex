package secret

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mindmatterlab/go-pro/internal/usercenter/model"
	v1 "github.com/mindmatterlab/go-pro/pkg/api/usercenter/v1"
)

// ModelToReply converts a model.SecretM to a v1.SecretReply. It copies the data from
// secretM to secret and sets the CreatedAt and UpdatedAt fields to their respective timestamps.
func ModelToReply(secretM *model.SecretM) *v1.SecretReply {
	var secret v1.SecretReply
	_ = copier.Copy(&secret, secretM)
	secret.CreatedAt = timestamppb.New(secretM.CreatedAt)
	secret.UpdatedAt = timestamppb.New(secretM.UpdatedAt)
	return &secret
}
