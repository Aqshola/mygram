package helpers

import (
	"net/url"
)

func ParseAndValidateUrl(rawurl string) (string, error) {
	u, err := url.ParseRequestURI(rawurl)
	if err != nil || u.Host == "" {
		u, repErr := url.ParseRequestURI("https://" + rawurl)
		if repErr != nil {
			return "", repErr
		}
		return u.String(), nil
	}
	return u.String(), nil
}
