package response

import (
	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/helper"
	"github.com/Daci1/url-shortener-atad/internal/types"
)

func New[T any](resourceType string, attributes T) *types.APIResponse[T] {
	return &types.APIResponse[T]{
		Data: types.Data[T]{
			Type:       resourceType,
			Attributes: attributes,
		},
	}
}

func UrlAttributesFromEntity(entity db.UrlEntity) types.UrlAttributes {
	return types.UrlAttributes{
		Id:          entity.Id,
		ShortUrl:    entity.ShortUrl,
		OriginalUrl: entity.OriginalUrl,
		UserId:      entity.UserId,
		CreatedAt:   entity.CreatedAt.UTC().String(),
		DeletedAt:   helper.If(entity.DeletedAt.Valid, entity.DeletedAt.Time.String(), ""),
	}
}

func ToCredentialsResponseAttributes(username, token, refreshToken string) types.CredentialsResponseAttributes {
	return types.CredentialsResponseAttributes{
		Username:     username,
		Token:        token,
		RefreshToken: refreshToken,
	}
}
