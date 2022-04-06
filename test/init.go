package test

import (
	"context"

	kakaoapi "github.com/LUXROBO/kakao-api/kakao"
)

const KAKAO_CLIENT_ID = "a07075fd6083b99006d0f14f7180ac55"
const KAKAO_REDIRECT_URL = "http://localhost:8081/user/kakao/callback"

type service struct {
	ctx    context.Context
	client *kakaoapi.Client
}

func New() *service {
	ctx := context.Background()
	client := kakaoapi.NewClient(KAKAO_CLIENT_ID, KAKAO_REDIRECT_URL)
	return &service{
		ctx:    ctx,
		client: client,
	}
}
