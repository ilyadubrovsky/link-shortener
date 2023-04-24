package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"link-shortener/internal/entity/link"
	"link-shortener/pkg/logging"
)

const linksTable = "links"

type postgresqlClient interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type linksStorage struct {
	client postgresqlClient
	logger *logging.Logger
}

func (s *linksStorage) CreateToken(ctx context.Context, link *link.Link) error {
	q := fmt.Sprintf("INSERT INTO %s (token, raw_url) VALUES ($1, $2)", linksTable)
	s.logger.Tracef("SQL: %s", q)

	_, err := s.client.Exec(ctx, q, link.Token, link.RawURL)
	if err != nil {
		return fmt.Errorf("repository create token: %v", err)
	}

	return nil
}

func (s *linksStorage) GetToken(ctx context.Context, rawURL string) (*string, error) {
	q := fmt.Sprintf("SELECT token FROM %s WHERE raw_url = $1", linksTable)
	s.logger.Tracef("SQL: %s", q)

	row := s.client.QueryRow(ctx, q, rawURL)

	var token string

	err := row.Scan(&token)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("repository get token: %v", err)
	}

	return &token, nil
}

func (s *linksStorage) GetRawURL(ctx context.Context, token string) (*string, error) {
	q := fmt.Sprintf("SELECT raw_url FROM %s WHERE token = $1", linksTable)
	s.logger.Tracef("SQL: %s", q)

	row := s.client.QueryRow(ctx, q, token)

	var rawURL string

	err := row.Scan(&rawURL)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("repository get raw url: %v", err)
	}

	return &rawURL, nil
}

func NewLinksStorage(client postgresqlClient, logger *logging.Logger) *linksStorage {
	return &linksStorage{client: client, logger: logger}
}
