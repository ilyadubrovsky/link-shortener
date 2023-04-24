package utils

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const (
	alphabetTest   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	lenOfTokenTest = 10
)

func TestGenerateToken(t *testing.T) {
	t.Run("token consists of provided alphabet", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			token, err := GenerateToken(alphabetTest, lenOfTokenTest)
			assert.Nil(t, err)
			assert.Regexp(t, regexp.MustCompile(`^[a-zA-Z0-9_]+$`), token)
		}
	})

	t.Run("token has provided len", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			token, err := GenerateToken(alphabetTest, lenOfTokenTest)
			assert.Nil(t, err)
			assert.Len(t, token, lenOfTokenTest)
		}
	})
}
