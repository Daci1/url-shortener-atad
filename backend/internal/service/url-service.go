package service

import (
	"fmt"
	"github.com/Daci1/url-shortener-atad/internal/errs"
	"time"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server/response"
	"github.com/Daci1/url-shortener-atad/internal/shortener"
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
	currentCounter, err := s.urlRepository.GetCounterAndIncrement()
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while retrieving counter: %s", err.Error()))
	}
	shortUrl := shortener.ToBase62(currentCounter)

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
	currentCounter, err := s.urlRepository.GetCounterAndIncrement()
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while retrieving counter: %s", err.Error()))
	}
	shortUrl := shortener.ToBase62(currentCounter)

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

func (s *UrlService) GetUrl(shortUrl string) (string, error) {
	urlEntity, err := s.urlRepository.GetByShortUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return urlEntity.OriginalUrl, nil
}
