package test

import (
	"testing"

	"github.com/LUXROBO/kakao-api/helper"
	"github.com/LUXROBO/kakao-api/logger"
)

// authCode는 한번사용한 코드는 재요청시 에러코드 발생
// 내부적으로 캐쉬처리 필요
func TestGetToken(t *testing.T) {
	const authCode = "XOdsoOS370U47RGkQONm9nHbBTv7SJJoYR6J2o4Zx6doj1kUn0Yvz7BJ7_hFDqakKvNMpAo9cxgAAAGAA2-ulw"

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
	const accessToken = "oUJdzykMGCuyRpj3CWVRUkv4HYaa_NLMZkfPsgo9dRkAAAGAA3BVTg"

	svc := New()
	profileData, err := svc.client.GetProfile(accessToken)
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug("profileData", profileData)
	helper.AssertEqual(t, profileData.KakaoAccount.Email, "walter.jung@luxrobo.com", "")
}
