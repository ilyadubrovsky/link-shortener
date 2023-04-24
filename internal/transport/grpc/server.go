package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"link-shortener/internal/apperror"
	"link-shortener/internal/entity/link"
	"link-shortener/internal/transport/grpc/server/link-shortener/protos"
	"net"
)

type shortenLinkService interface {
	ShortenURL(dto link.ShortenURLDTO) (string, error)
	GetRawURL(dto link.GetRawURLDTO) (string, error)
}

type Server struct {
	protos.UnimplementedLinkShortenerServer
	shortenLinkService
}

func NewServer(shortenLinkService shortenLinkService) *Server {
	return &Server{shortenLinkService: shortenLinkService}
}

func (s *Server) GetRawURL(ctx context.Context, req *protos.GetRawURLRequest) (*protos.GetRawURLResponse, error) {
	dto := link.GetRawURLDTO{Token: req.GetToken()}

	if err := dto.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rawURL, err := s.shortenLinkService.GetRawURL(link.GetRawURLDTO{Token: dto.Token})
	if err != nil {
		switch err {
		case apperror.ErrNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case apperror.ErrInternalServer:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &protos.GetRawURLResponse{RawURL: rawURL}, nil
}

func (s *Server) ShortenURL(ctx context.Context, req *protos.ShortenURLRequest) (*protos.ShortenURLResponse, error) {
	dto := link.ShortenURLDTO{RawURL: req.RawURL}

	if err := dto.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.shortenLinkService.ShortenURL(dto)
	if err != nil {
		switch err {
		case apperror.ErrBadRequest:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case apperror.ErrInternalServer:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &protos.ShortenURLResponse{Token: token}, nil
}

func (s *Server) Run(addr string) error {
	srv := grpc.NewServer()
	protos.RegisterLinkShortenerServer(srv, s)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if err = srv.Serve(l); err != nil {
		return err
	}

	return nil
}
