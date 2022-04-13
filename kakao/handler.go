package kakao

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LUXROBO/kakao-api/logger"
)

// Auth 인증요청 handler
func (c Client) Auth(w http.ResponseWriter, r *http.Request) {
	authLoginURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code",
		c.AuthURL,
		c.ClientID,
		c.RedirectURL,
	)
	logger.Debug("authLoginURL", authLoginURL)
	http.Redirect(w, r, authLoginURL, http.StatusMovedPermanently)
}

// Callback 인증요청후 Callback handler
func (c Client) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	logger.Debug("code", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(code)
}
