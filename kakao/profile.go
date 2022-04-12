package kakao

import (
	"encoding/json"

	"github.com/LUXROBO/kakao-api/helper"
	"github.com/LUXROBO/kakao-api/logger"
	"github.com/LUXROBO/kakao-api/types"
)

// GetProfile 카카오 프로필
func (c Client) GetProfile(accessToken string) (*types.ProfileData, error) {
	body, err := helper.RequestGetAPI(accessToken, types.KakaoAPI+"/v2/user/me")
	if helper.WorkCheck("RequestGetAPI", err) != nil {
		return nil, err
	}

	profile := &types.ProfileData{}
	logger.Debug("body", string(body))
	err = json.Unmarshal(body, &profile)

	logger.Debug("profile", profile)
	return profile, err
}
