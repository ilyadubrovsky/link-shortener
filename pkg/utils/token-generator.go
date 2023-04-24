package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateToken(alphabet string, lenOfToken int) (string, error) {
	lenOfAlphabet := int64(len(alphabet))
	token := make([]byte, lenOfToken)

	for i := 0; i < lenOfToken; i++ {
		symbolIndex, err := rand.Int(rand.Reader, big.NewInt(lenOfAlphabet))
		if err != nil {
			return "", fmt.Errorf("generate token: %v", err)
		}

		token[i] = alphabet[symbolIndex.Int64()]
	}

	return string(token), nil
}
