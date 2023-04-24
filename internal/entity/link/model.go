package link

import (
	"fmt"
	"net/url"
	"regexp"
)

type Link struct {
	RawURL string `json:"raw_url"`
	Token  string `json:"token"`
}

type ShortenURLDTO struct {
	RawURL string `json:"raw_url"`
}

// Validate TODO maybe make it less strict
func (dto *ShortenURLDTO) Validate() error {
	if dto.RawURL == "" {
		return fmt.Errorf("provided url is empty")
	}

	parsedURL, err := url.Parse(dto.RawURL)
	if err != nil {
		return err
	}

	if parsedURL.Scheme == "" {
		return fmt.Errorf("scheme of provided url is empty")
	}

	if parsedURL.Host == "" {
		return fmt.Errorf("host of provided url is empty")
	}

	return nil
}

func (dto *ShortenURLDTO) ConvertToLink() *Link {
	return &Link{
		RawURL: dto.RawURL,
		Token:  "",
	}
}

type GetRawURLDTO struct {
	Token string `json:"token"`
}

func (dto *GetRawURLDTO) Validate() error {
	isStringAlphabet := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	lenOfToken := 10

	if !isStringAlphabet(dto.Token) || len(dto.Token) != lenOfToken {
		return fmt.Errorf("provided token not valid")
	}

	return nil
}
