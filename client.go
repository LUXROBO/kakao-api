package kakaoapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	// KakaoAuthPath 인증 API
	KakaoAuthPath = "https://kauth.kakao.com"
	// KakaoAPIPAth API
	KakaoAPIPAth = "https://kapi.kakao.com"
)

type Client struct {
	ClientID     string
	AuthURL      string
	AuthLoginURL string
	TokenURL     string
	RedirectURL  string
}

// NewClient 카카오 서비스
func NewClient() *Client {
	client := &Client{
		ClientID:    os.Getenv("KAKAO_CLIENT_ID"),
		AuthURL:     fmt.Sprintf("%s/oauth/authorize", KakaoAuthPath),
		TokenURL:    fmt.Sprintf("%s/oauth/token", KakaoAuthPath),
		RedirectURL: os.Getenv("KAKAO_REDIRECT_URL"),
	}
	client.AuthLoginURL = fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code",
		client.AuthURL,
		client.ClientID,
		client.RedirectURL,
	)
	return client
}

func (c Client) requestGetAPI(tokenStr string, apiURL string) ([]byte, error) {
	token := "Bearer " + tokenStr
	client := &http.Client{}

	profileURL := fmt.Sprintf("%s/%s", KakaoAPIPAth, apiURL)
	req, _ := http.NewRequest("GET", profileURL, nil)
	req.Header.Set("Host", "kapi.kakao.com")
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
