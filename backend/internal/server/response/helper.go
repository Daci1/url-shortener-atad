package response

import (
	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/helper"
)

func New[T any](resourceType string, attributes T) *APIResponse[T] {
	return &APIResponse[T]{
		Data: Data[T]{
			Type:       resourceType,
			Attributes: attributes,
		},
	}
}

func UrlAttributesFromEntity(entity db.UrlEntity) UrlAttributes {
	return UrlAttributes{
		Id:          entity.Id,
		ShortUrl:    entity.ShortUrl,
		OriginalUrl: entity.OriginalUrl,
		CreatedAt:   entity.CreatedAt.UTC().String(),
		DeletedAt:   helper.If(entity.DeletedAt.Valid, entity.DeletedAt.Time.String(), ""),
	}
}

func ToCredentialsResponseAttributes(username, token, refreshToken string) CredentialsResponseAttributes {
	return CredentialsResponseAttributes{
		Username:     username,
		Token:        token,
		RefreshToken: refreshToken,
	}
}
