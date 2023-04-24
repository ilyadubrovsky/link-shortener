package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"link-shortener/internal/apperror"
	"link-shortener/internal/entity/link"
	mock_service "link-shortener/internal/service/mocks"
	"link-shortener/pkg/logging"
	"testing"
)

func TestService_ShortenURL(t *testing.T) {
	type mockBehaviour func(s *mock_service.MocklinksRepository, rawURL string)

	testTable := []struct {
		name          string
		dto           link.ShortenURLDTO
		mockBehaviour mockBehaviour
		expectedToken string
		expectedError error
	}{
		{
			name: "Ok",
			dto:  link.ShortenURLDTO{RawURL: "https://www.google.com/"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, rawURL string) {
				s.EXPECT().GetToken(gomock.Any(), rawURL).Return(ref("01ToKeN_89"), nil)
			},
			expectedToken: "01ToKeN_89",
			expectedError: nil,
		},
		{
			name: "RepositoryError1",
			dto:  link.ShortenURLDTO{RawURL: "www.ozon.ru/"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, rawURL string) {
				s.EXPECT().GetToken(gomock.Any(), rawURL).Return(nil, assert.AnError)
			},
			expectedToken: "",
			expectedError: apperror.ErrInternalServer,
		},
		{
			name: "RepositoryError2",
			dto:  link.ShortenURLDTO{RawURL: "ozon.ru"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, rawURL string) {
				s.EXPECT().GetToken(gomock.Any(), rawURL).Return(nil, nil)
				s.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			expectedToken: "",
			expectedError: apperror.ErrInternalServer,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_service.NewMocklinksRepository(c)
			testCase.mockBehaviour(repository, testCase.dto.RawURL)

			service := Service{
				repository: repository,
				logger:     logging.GetLogger(),
			}

			token, err := service.ShortenURL(testCase.dto)

			assert.Equal(t, testCase.expectedToken, token)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestService_GetRawURL(t *testing.T) {
	type mockBehaviour func(s *mock_service.MocklinksRepository, token string)

	testTable := []struct {
		name           string
		dto            link.GetRawURLDTO
		mockBehaviour  mockBehaviour
		expectedRawURL string
		expectedError  error
	}{
		{
			name: "RawURLFound",
			dto:  link.GetRawURLDTO{Token: "ozon__2023"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, token string) {
				s.EXPECT().GetRawURL(gomock.Any(), token).Return(ref("https://www.ozon.ru/"), nil)
			},
			expectedRawURL: "https://www.ozon.ru/",
			expectedError:  nil,
		},
		{
			name: "RawURLNotFound",
			dto:  link.GetRawURLDTO{Token: "OZON__2023"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, token string) {
				s.EXPECT().GetRawURL(gomock.Any(), token).Return(nil, nil)
			},
			expectedRawURL: "",
			expectedError:  apperror.ErrNotFound,
		},
		{
			name: "RepositoryError",
			dto:  link.GetRawURLDTO{Token: "ozonozon04"},
			mockBehaviour: func(s *mock_service.MocklinksRepository, token string) {
				s.EXPECT().GetRawURL(gomock.Any(), token).Return(nil, assert.AnError)
			},
			expectedRawURL: "",
			expectedError:  apperror.ErrInternalServer,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_service.NewMocklinksRepository(c)
			testCase.mockBehaviour(repository, testCase.dto.Token)

			service := Service{
				repository: repository,
				logger:     logging.GetLogger(),
			}

			rawURL, err := service.GetRawURL(testCase.dto)

			assert.Equal(t, testCase.expectedRawURL, rawURL)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func ref(s string) *string {
	return &s
}
