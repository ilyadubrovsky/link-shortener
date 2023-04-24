package service

import (
	"context"
	"link-shortener/internal/apperror"
	"link-shortener/internal/entity/link"
	"link-shortener/pkg/logging"
	"link-shortener/pkg/utils"
	"time"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	lenOfToken = 10
)

//go:generate mockgen -source=link-shortener.go -destination=mocks/mock.go
type linksRepository interface {
	CreateToken(ctx context.Context, link *link.Link) error
	GetToken(ctx context.Context, rawURL string) (*string, error)
	GetRawURL(ctx context.Context, token string) (*string, error)
}

type Service struct {
	repository linksRepository
	logger     *logging.Logger
}

func (s *Service) ShortenURL(dto link.ShortenURLDTO) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	l := dto.ConvertToLink()

	token, err := s.repository.GetToken(ctx, l.RawURL)
	if err != nil {
		s.logger.Errorf("failed to ShortenURL due to error: %v", err)
		s.logger.Debugf("RawURL: %s", l.RawURL)
		return "", apperror.ErrInternalServer
	}

	if token != nil {
		return *token, nil
	}

	l.Token, err = utils.GenerateToken(alphabet, lenOfToken)
	if err != nil {
		s.logger.Errorf("failed to ShortenURL due to error: %v", err)
		s.logger.Debugf("alphabet: %s, lenOfToken: %d", alphabet, lenOfToken)
		return "", apperror.ErrBadRequest
	}

	// TODO maybe attempt to regenerate if token exists
	if err = s.repository.CreateToken(ctx, l); err != nil {
		s.logger.Errorf("failed to ShortenURL due to error: %v", err)
		s.logger.Debugf("link: %s", l)
		return "", apperror.ErrInternalServer
	}

	return l.Token, nil
}

func (s *Service) GetRawURL(dto link.GetRawURLDTO) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rawURL, err := s.repository.GetRawURL(ctx, dto.Token)
	if err != nil {
		s.logger.Errorf("failed to GetRawURL due to error: %v", err)
		s.logger.Debugf("token: %s", dto.Token)
		return "", apperror.ErrInternalServer
	}

	if rawURL != nil {
		return *rawURL, nil
	}

	return "", apperror.ErrNotFound
}

func NewShortenLinkService(repository linksRepository) *Service {
	return &Service{repository: repository}
}
