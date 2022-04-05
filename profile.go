package kakaoapi

import (
	"encoding/json"
	"kakao-api/logger"
)

// GetProfileEmail 카카오 프로필
func (c Client) GetProfileEmail(tokenStr string) (string, error) {
	body, err := c.requestGetAPI(tokenStr, KakaoAPIPAth+"/v2/user/me")
	if WorkCheck("requestAPI", err) != nil {
		return "", err
	}

	profile := &KakaoData{}
	err = json.Unmarshal(body, &profile)
	logger.Debug("profile", profile)
	return profile.KakaoAccount.Email, err
}
