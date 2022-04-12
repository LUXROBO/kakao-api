package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LUXROBO/kakao-api/types"
	"github.com/stretchr/testify/assert"
)

func Test_MustPanic(t *testing.T) {
	defer func() { recover() }()

	err := errors.New("must panic")
	Must(err)

	t.Errorf("did not panic")
}

func Test_RequestGetAPI(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", types.KakaoAPI+"/v2/user/me", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/v2/user/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode("ok")
	})

	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, http.StatusOK)
}
