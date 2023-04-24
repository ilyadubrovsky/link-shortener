package inmemory

import (
	"context"
	"link-shortener/internal/entity/link"
	"sync"
)

type linksStorage struct {
	mu            sync.RWMutex
	TokenToRawURL map[string]string
	RawURLToToken map[string]string
}

func (s *linksStorage) CreateToken(_ context.Context, link *link.Link) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TokenToRawURL[link.Token] = link.RawURL
	s.RawURLToToken[link.RawURL] = link.Token

	return nil
}

func (s *linksStorage) GetToken(_ context.Context, rawURL string) (*string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	token, ok := s.RawURLToToken[rawURL]
	if !ok {
		return nil, nil
	}

	return &token, nil
}

func (s *linksStorage) GetRawURL(_ context.Context, token string) (*string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rawURL, ok := s.TokenToRawURL[token]
	if !ok {
		return nil, nil
	}

	return &rawURL, nil
}

func NewLinksStorage() *linksStorage {
	return &linksStorage{TokenToRawURL: make(map[string]string), RawURLToToken: make(map[string]string)}
}
