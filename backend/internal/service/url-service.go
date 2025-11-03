package service

import (
	"fmt"
	"github.com/Daci1/url-shortener-atad/internal/errs"
	"time"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server/response"
	"github.com/google/uuid"
)

type UrlService struct {
	urlRepository *db.UrlRepository
}

func NewUrlService(urlRepo *db.UrlRepository) *UrlService {
	return &UrlService{
		urlRepository: urlRepo,
	}
}

func (s *UrlService) CreateUrl(originalUrl string) (*response.APIResponse[response.UrlAttributes], errs.CustomError) {
	shortUrl, err := s.urlRepository.GetNextAvailableShortUrl()
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while retrieving next available url: %s", err.Error()))
	}

	urlEntity := &db.UrlEntity{
		Id:          uuid.NewString(),
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
		CreatedAt:   time.Now(),
	}

	err = s.urlRepository.CreateUrl(*urlEntity)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while creating url: %s", err.Error()))
	}

	attributes := response.UrlAttributesFromEntity(*urlEntity)
	return response.New("urls", attributes), nil
}

func (s *UrlService) CreateUrlWithUser(originalUrl, userId string) (*response.APIResponse[response.UrlAttributes], errs.CustomError) {
	shortUrl, err := s.urlRepository.GetNextAvailableShortUrl()
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while retrieving counter: %s", err.Error()))
	}

	urlEntity := &db.UrlEntity{
		Id:          uuid.NewString(),
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
		UserId:      userId,
		CreatedAt:   time.Now(),
	}

	err = s.urlRepository.CreateUrlWithUser(*urlEntity)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while creating url: %s", err.Error()))
	}

	attributes := response.UrlAttributesFromEntity(*urlEntity)
	return response.New("urls", attributes), nil
}

func (s *UrlService) CreateCustomUrl(originalUrl, desiredUrl, userId string) (*response.APIResponse[response.UrlAttributes], errs.CustomError) {
	shortUrlExists, err := s.urlRepository.ShortUrlExists(desiredUrl)
	if err != nil {
		return nil, err
	}

	if shortUrlExists {
		return nil, errs.Conflict("Short url already exists")
	}

	urlEntity := &db.UrlEntity{
		Id:          uuid.NewString(),
		ShortUrl:    desiredUrl,
		OriginalUrl: originalUrl,
		UserId:      userId,
		CreatedAt:   time.Now(),
	}

	err = s.urlRepository.CreateUrlWithUser(*urlEntity)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while creating url: %s", err.Error()))
	}

	attributes := response.UrlAttributesFromEntity(*urlEntity)
	return response.New("urls", attributes), nil
}

func (s *UrlService) GetUrl(shortUrl string) (string, errs.CustomError) {
	urlEntity, err := s.urlRepository.GetByShortUrlAndIncrementAnalytics(shortUrl)
	if err != nil {
		return "", err
	}

	return urlEntity.OriginalUrl, nil
}
