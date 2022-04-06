package test

import (
	"testing"

	"github.com/LUXROBO/kakao-api/helper"
	"github.com/LUXROBO/kakao-api/logger"
)

// authCode는 한번사용한 코드는 재요청시 에러코드 발생
// 내부적으로 캐쉬처리 필요
func TestGetToken(t *testing.T) {
	const authCode = "VnT7jW5GV13QfCMUC7zRidstK4wtkl09_y0CcDYt4lYtzdlwzWfjCLHdSXPkiTq13u_V0QopcBQAAAGAAOdODw"

	svc := New()
	token, err := svc.client.GetToken(svc.ctx, authCode, KAKAO_REDIRECT_URL)
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug("token", token)
}

// TestKakaoProfile 카카오 프로필 테스트
// 입력값 accessToken
func TestKakaoProfile(t *testing.T) {
	const accessToken = "YpyZ8nM8RRcqnXrqFnEhXeZC5MpoahNQTsiySwo9dNkAAAF__w613A"

	svc := New()
	profileData, err := svc.client.GetProfile(accessToken)
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug("profileData", profileData)
	helper.AssertEqual(t, profileData.KakaoAccount.Email, "walter.jung@luxrobo.com", "")
}
