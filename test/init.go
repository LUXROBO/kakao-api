package test

import (
	"context"

	kakaoapi "github.com/LUXROBO/kakao-api/kakao"
)

// KAKAO_CLIENT_ID 카카오 앱ID
const KAKAO_CLIENT_ID = "a07075fd6083b99006d0f14f7180ac55"

// KAKAO_REDIRECT_URL 카카오 인가코드 redirectURL정보
const KAKAO_REDIRECT_URL = "http://localhost:8081/user/kakao/callback"

// Service ...
type Service struct {
	ctx    context.Context
	client *kakaoapi.Client
}

// New 생성자
func New() *Service {
	ctx := context.Background()
	client := kakaoapi.NewClient(KAKAO_CLIENT_ID, KAKAO_REDIRECT_URL)
	return &Service{
		ctx:    ctx,
		client: client,
	}
}
