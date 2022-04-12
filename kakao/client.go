package kakao

import (
	"fmt"

	"github.com/LUXROBO/kakao-api/types"
)

// Client ....
type Client struct {
	// ClientID 카카오 REST API키
	ClientID string
	// 카카오 인가코드 요청 URL
	AuthURL string
	// 인가코드를 통한 토큰요청 URL
	TokenURL string
	// 카카오 Callback Redirect URL
	RedirectURL string
}

// NewClient 카카오 서비스
func NewClient(
	clientID string,
	redirectURL string,
) *Client {
	client := &Client{
		ClientID:    clientID,
		RedirectURL: redirectURL,
		AuthURL:     fmt.Sprintf("%s/oauth/authorize", types.KakaoAuthAPI),
		TokenURL:    fmt.Sprintf("%s/oauth/token", types.KakaoAuthAPI),
	}
	return client
}
