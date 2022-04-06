package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/LUXROBO/kakao-api/logger"
)

// AssertEqual ...
func AssertEqual(t *testing.T, a interface{}, b interface{}, errorMsg string) {
	if a != b {
		msg := fmt.Sprintf("%v != %v", a, b)
		if len(errorMsg) == 0 {
			t.Fatal(msg)
		} else {
			t.Fatalf("%s => %s", errorMsg, msg)
		}
	}
}

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

// RequestGetAPI GET요청
func RequestGetAPI(tokenStr string, apiURL string) ([]byte, error) {
	token := "Bearer " + tokenStr
	client := &http.Client{}

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Host", "kapi.kakao.com")
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
