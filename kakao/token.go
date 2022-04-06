package kakao

import (
	"context"
	"net/http"
	"time"

	"github.com/LUXROBO/kakao-api/helper"
	"golang.org/x/oauth2"
)

// GetToken 토큰처리
// code: 로그인후 받은 인가코드값
// redirectURL: 인가코드요청시 필수로 보낸 redirectURL값
func (c Client) GetToken(ctx context.Context, code string, redirectURL string) (*oauth2.Token, error) {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	var (
		conf = &oauth2.Config{
			ClientID: c.ClientID,
			Endpoint: oauth2.Endpoint{
				AuthURL:  c.AuthURL,
				TokenURL: c.TokenURL,
			},
			RedirectURL: redirectURL,
		}
	)
	token, err := conf.Exchange(ctx, code)
	if helper.WorkCheck("Exchange", err) != nil {
		return nil, err
	}

	_, err = c.GetProfile(token.AccessToken)
	if helper.WorkCheck("GetProfile", err) != nil {
		return nil, err
	}
	return token, nil
}
