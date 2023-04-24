package app

import (
	"context"
	"fmt"
	"link-shortener/internal/config"
	"link-shortener/internal/entity/link"
	"link-shortener/internal/service"
	"link-shortener/internal/storage/inmemory"
	"link-shortener/internal/storage/postgres"
	grpc2 "link-shortener/internal/transport/grpc"
	"link-shortener/internal/transport/http"
	"link-shortener/pkg/client/postgresql"
	"link-shortener/pkg/logging"
	"os"
)

type Service interface {
	ShortenURL(dto link.ShortenURLDTO) (string, error)
	GetRawURL(dto link.GetRawURLDTO) (string, error)
}

type Repository interface {
	CreateToken(ctx context.Context, link *link.Link) error
	GetToken(ctx context.Context, rawURL string) (*string, error)
	GetRawURL(ctx context.Context, token string) (*string, error)
}

type Server interface {
	Run(addr string) error
}

type App struct {
	logger *logging.Logger
	cfg    *config.Config
	server Server
}

func New(cfg *config.Config) (*App, error) {
	var (
		a                  App
		linksStorage       Repository
		shortenLinkService Service
	)

	a.cfg = cfg

	if a.cfg.StorageConnect == "" {
		a.cfg.StorageConnect = os.Getenv("STORAGE_CONNECT")
	}

	a.logger = logging.GetLogger()

	if cfg.StorageConnect == "postgresql" {
		a.logger.Info("postgresql client initializing")
		pgConfig := postgresql.NewPgConfig(cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host,
			cfg.Postgres.Port, cfg.Postgres.Database)

		pgClient, err := postgresql.NewClient(context.Background(), pgConfig)
		if err != nil {
			return nil, fmt.Errorf("postgres: %v", err)
		}

		a.logger.Info("postgresql links storage initializing")
		linksStorage = postgres.NewLinksStorage(pgClient, a.logger)
	} else {
		a.logger.Info("inmemory storage initializing")
		linksStorage = inmemory.NewLinksStorage()
	}

	a.logger.Info("shorten link service initializing")
	shortenLinkService = service.NewShortenLinkService(linksStorage)

	if a.cfg.TransportType == "grpc" {
		a.logger.Info("grpc server initializing")
		a.server = grpc2.NewServer(shortenLinkService)
	} else {
		a.logger.Info("http server initializing")
		handler := http.NewHandler(shortenLinkService)
		a.server = http.NewServer(handler)
	}

	return &a, nil
}

func (a *App) Run() error {
	a.logger.Info("server is starting")
	return a.server.Run(fmt.Sprintf(":%s", a.cfg.ServerPort))
}
