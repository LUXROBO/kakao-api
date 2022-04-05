package kakaoapi

import "kakao-api/logger"

// Must 에러일경우 panic 처리
func Must(err error) {
	if err != nil {
		logger.Error("Must", err)
		panic(err)
	}
}

// WorkCheck 에러일경우 처리
func WorkCheck(message string, err error) error {
	if err != nil {
		logger.Errorf("%s - %s", message, err)
		return err
	}
	return nil
}
