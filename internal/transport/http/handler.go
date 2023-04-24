package http

import (
	"github.com/gin-gonic/gin"
	"link-shortener/internal/apperror"
	"link-shortener/internal/entity/link"
	"net/http"
)

const (
	apiURL     = "/api"
	shortenURL = ""
	getRawURL  = "/:token"
)

//go:generate mockgen -source=handler.go -destination=mocks/mock.go
type shortenLinkService interface {
	ShortenURL(dto link.ShortenURLDTO) (string, error)
	GetRawURL(dto link.GetRawURLDTO) (string, error)
}

type Handler struct {
	shortenLinkService
}

func NewHandler(shortenLinkService shortenLinkService) *Handler {
	return &Handler{shortenLinkService: shortenLinkService}
}

func (h *Handler) Register(r *gin.Engine) {
	apiGroup := r.Group(apiURL)
	{
		apiGroup.POST(shortenURL, h.shortenURL)
		apiGroup.GET(getRawURL, h.getRawURL)
	}
}

func (h *Handler) shortenURL(c *gin.Context) {
	var dto link.ShortenURLDTO

	if err := c.BindJSON(&dto); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := dto.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.ShortenURL(dto)
	if err != nil {
		switch err {
		case apperror.ErrBadRequest:
			c.String(http.StatusBadRequest, err.Error())
		case apperror.ErrInternalServer:
			c.String(http.StatusInternalServerError, err.Error())
		default:
			c.String(http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *Handler) getRawURL(c *gin.Context) {
	token := c.Params.ByName("token")

	dto := link.GetRawURLDTO{Token: token}

	if err := dto.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	rawURL, err := h.GetRawURL(dto)
	if err != nil {
		switch err {
		case apperror.ErrNotFound:
			c.String(http.StatusNotFound, err.Error())
		case apperror.ErrInternalServer:
			c.String(http.StatusInternalServerError, err.Error())
		default:
			c.String(http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"raw_url": rawURL})
}
