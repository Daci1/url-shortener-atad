package service

import (
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

func (s *UrlService) CreateUrl(originalUrl string) (*response.APIResponse[response.UrlAttributes], error) {
	currentCounter, err := s.urlRepository.GetCounterAndIncrement()
	if err != nil {
		return nil, err
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
		return nil, err
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
