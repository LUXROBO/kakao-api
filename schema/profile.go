package schema

// Property ...
type Property struct {
	Nickname string `json:"nickname"`
}

// Profile ...
type Profile struct {
	Nickname string `json:"nickname"`
}

// Account ...
type Account struct {
	ProfileNeedsAgreement  bool    `json:"profile_needs_agreement"`
	Profile                Profile `json:"profile"`
	HasEmail               bool    `json:"has_email"`
	EmailNeedsAgreement    bool    `json:"email_needs_agreement"`
	IsEmailValid           bool    `json:"is_email_valid"`
	IsEmailValified        bool    `json:"is_email_verified"`
	Email                  string  `json:"email"`
	HasAgeRange            bool    `json:"has_age_range"`
	AgeRangeNeedsAgreement bool    `json:"age_range_needs_agreement"`
	AgeRange               string  `json:"age_range"`
	HasBirthday            bool    `json:"has_birthday"`
	BirthdayNeedsAgreement bool    `json:"birthday_needs_agreement"`
	Birthday               string  `json:"birthday"`
	Hasgendar              bool    `json:"has_gender"`
	GenderNeedsAgreement   bool    `json:"gender_needs_agreement"`
	Gender                 string  `json:"gender"`
}

// ProfileData ...
type ProfileData struct {
	ID           int64    `json:"id"`
	ConnectedAt  string   `json:"connected_at"`
	Properties   Property `json:"properties"`
	KakaoAccount Account  `json:"kakao_account"`
}
