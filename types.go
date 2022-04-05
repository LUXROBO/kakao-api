package kakaoapi

// KakaoProperty ...
type KakaoProperty struct {
	Nickname string `json:"nickname"`
}

// KakaoProfile ...
type KakaoProfile struct {
	Nickname string `json:"nickname"`
}

// KakaoAccount ...
type KakaoAccount struct {
	ProfileNeedsAgreement  bool         `json:"profile_needs_agreement"`
	Profile                KakaoProfile `json:"profile"`
	HasEmail               bool         `json:"has_email"`
	EmailNeedsAgreement    bool         `json:"email_needs_agreement"`
	IsEmailValid           bool         `json:"is_email_valid"`
	IsEmailValified        bool         `json:"is_email_verified"`
	Email                  string       `json:"email"`
	HasAgeRange            bool         `json:"has_age_range"`
	AgeRangeNeedsAgreement bool         `json:"age_range_needs_agreement"`
	AgeRange               string       `json:"age_range"`
	HasBirthday            bool         `json:"has_birthday"`
	BirthdayNeedsAgreement bool         `json:"birthday_needs_agreement"`
	Birthday               string       `json:"birthday"`
	Hasgendar              bool         `json:"has_gender"`
	GenderNeedsAgreement   bool         `json:"gender_needs_agreement"`
	Gender                 string       `json:"gender"`
}

// KakaoData ...
type KakaoData struct {
	ID           int64         `json:"id"`
	ConnectedAt  string        `json:"connected_at"`
	Properties   KakaoProperty `json:"properties"`
	KakaoAccount KakaoAccount  `json:"kakao_account"`
}
