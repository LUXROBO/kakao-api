package kakaoapi

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// GetToken 토큰처리
func (c Client) GetToken(ctx context.Context, code string) (*oauth2.Token, error) {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	var (
		conf = &oauth2.Config{
			ClientID: c.ClientID,
			Endpoint: oauth2.Endpoint{
				AuthURL:  c.AuthURL,
				TokenURL: c.TokenURL,
			},
			RedirectURL: c.RedirectURL,
		}
	)
	token, err := conf.Exchange(ctx, code)
	if WorkCheck("Exchange", err) != nil {
		return nil, err
	}

	_, err = c.GetProfileEmail(token.AccessToken)
	if WorkCheck("GetProfileEmail", err) != nil {
		return nil, err
	}
	return token, nil
}