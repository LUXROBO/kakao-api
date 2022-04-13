package kakao

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LUXROBO/kakao-api/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// KakaoSuite ....
type KakaoSuite struct {
	ctx         context.Context
	clientID    string
	redirectURL string
	suite.Suite
	client *Client
}

// TestKakaoSuite ...
func TestKakaoSuite(t *testing.T) {
	suite.Run(t, new(KakaoSuite))
}

// SetupTest 세팅
func (s *KakaoSuite) SetupTest() {
	s.ctx = context.Background()
	s.clientID = "a07075fd6083b99006d0f14f7180ac55"
	s.redirectURL = "http://localhost:8081/user/kakao/callback"
	s.client = NewClient(s.clientID, s.redirectURL)
}

// Test_Auth 인증요청 페이지이동
func (s *KakaoSuite) Test_Auth() {
	req := httptest.NewRequest(http.MethodGet, "/user/kakao/auth", nil)
	w := httptest.NewRecorder()

	s.client.Auth(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.T().Errorf("expected error to be nil got %v", err)
	}

	assert.Equal(s.T(), true, strings.Contains(string(data), "Moved Permanently"), "equal")
}

// Test_Callback ....
func (s *KakaoSuite) Test_Callback() {
	req := httptest.NewRequest(http.MethodGet, "/user/kakao/callback", nil)
	w := httptest.NewRecorder()

	s.client.Callback(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.T().Errorf("expected error to be nil got %v", err)
	}
}

// Test_Profile 토큰을 이용한 프로필정보
func (s *KakaoSuite) Test_Profile() {
	_, err := s.client.GetProfile("token")
	if err != nil {
		s.T().Errorf("expected error to be nil got %v", err)
	}
}

// Test_Token 토큰발행 테스트가 애매해서 작성한 상태...
func (s *KakaoSuite) Test_Token() {
	var (
		authCode    = "code"
		redirectURL = "http://localhost:8081/user/kakao/callback"
	)
	_, err := s.client.GetToken(s.ctx, authCode, redirectURL)

	// 에러발생
	assert.Equal(s.T(), true, helper.ErrorCheck(err), err)
}
