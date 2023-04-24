package http

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"link-shortener/internal/apperror"
	"link-shortener/internal/entity/link"
	mock_http "link-shortener/internal/transport/http/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_shortenURL(t *testing.T) {
	type mockBehaviour func(s *mock_http.MockshortenLinkService, dto link.ShortenURLDTO)

	testTable := []struct {
		name                 string
		inputBody            string
		inputDTO             link.ShortenURLDTO
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"raw_url":"https://www.ozon.ru/"}`,
			inputDTO:  link.ShortenURLDTO{RawURL: "https://www.ozon.ru/"},
			mockBehaviour: func(s *mock_http.MockshortenLinkService, dto link.ShortenURLDTO) {
				s.EXPECT().ShortenURL(dto).Return("T0kenT0ken", nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"token":"T0kenT0ken"}`,
		},
		{
			name:                 "EmptyRawURL",
			inputBody:            `{"raw_url":""}`,
			inputDTO:             link.ShortenURLDTO{},
			mockBehaviour:        func(s *mock_http.MockshortenLinkService, dto link.ShortenURLDTO) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "provided url is empty",
		},
		{
			name:      "ServiceError",
			inputBody: `{"raw_url":"https://www.google.com/"}`,
			inputDTO:  link.ShortenURLDTO{RawURL: "https://www.google.com/"},
			mockBehaviour: func(s *mock_http.MockshortenLinkService, dto link.ShortenURLDTO) {
				s.EXPECT().ShortenURL(dto).Return("", assert.AnError)
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: "internal server",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockshortenLinkService(c)
			testCase.mockBehaviour(service, testCase.inputDTO)

			h := NewHandler(service)

			target := fmt.Sprintf("%s%s", apiURL, shortenURL)

			r := gin.Default()
			r.POST(target, h.shortenURL)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, target, bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_getRawURL(t *testing.T) {
	type mockBehaviour func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO)

	testTable := []struct {
		name                 string
		inputParam           string
		inputDTO             link.GetRawURLDTO
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			inputParam: "ozon00ozon",
			inputDTO:   link.GetRawURLDTO{Token: "ozon00ozon"},
			mockBehaviour: func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO) {
				s.EXPECT().GetRawURL(dto).Return("https://www.ozon.ru/", nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"raw_url":"https://www.ozon.ru/"}`,
		},
		{
			name:       "ServiceError",
			inputParam: "ozon_20000",
			inputDTO:   link.GetRawURLDTO{Token: "ozon_20000"},
			mockBehaviour: func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO) {
				s.EXPECT().GetRawURL(dto).Return("", assert.AnError)
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: "internal server",
		},
		{
			name:                 "TokenNotValid1",
			inputParam:           "*#12010*_Z",
			inputDTO:             link.GetRawURLDTO{},
			mockBehaviour:        func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "provided token not valid",
		},
		{
			name:                 "TokenNotValid2",
			inputParam:           "SDVlvdsls_VlSVdl0",
			inputDTO:             link.GetRawURLDTO{},
			mockBehaviour:        func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "provided token not valid",
		},
		{
			name:       "TokenNotFound",
			inputParam: "0Z0N______",
			inputDTO:   link.GetRawURLDTO{Token: "0Z0N______"},
			mockBehaviour: func(s *mock_http.MockshortenLinkService, dto link.GetRawURLDTO) {
				s.EXPECT().GetRawURL(dto).Return("", apperror.ErrNotFound)
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: "not found",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockshortenLinkService(c)
			testCase.mockBehaviour(service, testCase.inputDTO)

			h := NewHandler(service)

			r := gin.Default()
			r.GET(fmt.Sprintf("%s%s", apiURL, getRawURL), h.getRawURL)

			target := fmt.Sprintf("%s/%s", apiURL, testCase.inputParam)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, target, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
