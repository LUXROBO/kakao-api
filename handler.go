package kakaoapi

import (
	"context"
	"encoding/json"
	"kakao-api/logger"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// Auth 인증요청 handler
func (c Client) Auth(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, c.AuthLoginURL, http.StatusMovedPermanently)
	})
}

// Callback 인증요청후 Callback handler
func (c Client) Callback(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")

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
			return
		}

		email, err := c.GetProfileEmail(token.AccessToken)
		if WorkCheck("GetProfileEmail", err) != nil {
			return
		}

		logger.Debug("profile", email)
		logger.Debug("token", token)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(token)
	})
}
