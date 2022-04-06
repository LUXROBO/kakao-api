package kakaoapi

import (
	"encoding/json"

	"github.com/LUXROBO/kakao-api/logger"
	"github.com/LUXROBO/kakao-api/schema"
)

// GetProfileEmail 카카오 프로필
func (c Client) GetProfileEmail(tokenStr string) (*schema.ProfileData, error) {
	body, err := c.requestGetAPI(tokenStr, KakaoAPIPAth+"/v2/user/me")
	if WorkCheck("requestAPI", err) != nil {
		return nil, err
	}

	profile := &schema.ProfileData{}
	err = json.Unmarshal(body, &profile)
	logger.Debug("profile", profile)
	return profile, err
}
