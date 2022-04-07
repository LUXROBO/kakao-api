package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	kakaoapi "github.com/LUXROBO/kakao-api/kakao"
	"github.com/LUXROBO/kakao-api/logger"
	"github.com/rs/cors"
)

// KAKAO_CLIENT_ID 카카오 앱ID
const KAKAO_CLIENT_ID = "a07075fd6083b99006d0f14f7180ac55"

// KAKAO_REDIRECT_URL 카카오 인가코드 redirectURL정보
const KAKAO_REDIRECT_URL = "http://localhost:8081/user/kakao/callback"

func main() {
	errc := make(chan error)
	ctx := context.Background()
	kakaoClient := kakaoapi.NewClient(KAKAO_CLIENT_ID, KAKAO_REDIRECT_URL)

	mux := http.NewServeMux()
	mux.Handle("/user/kakao/auth", kakaoClient.Auth(ctx))
	mux.Handle("/user/kakao/callback", kakaoClient.Callback(ctx))
	serverStart("8081", mux, errc)
	logger.Debug("exit", <-errc)
}

func serverStart(
	port string,
	mux *http.ServeMux,
	errc chan error,
) {
	var (
		httpAddr          = flag.String("http", ":"+port, "HTTP listen address")
		readHeaderTimeout = 60 * time.Second
		writeTimeout      = 60 * time.Second
		idleTimeout       = 60 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	op := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{
			"Origin",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"DeviceInfo",
			"Authorization",
			"X-Requested-With",
		},
		AllowCredentials: false,
	}

	handler := cors.New(op).Handler(mux)
	http := &http.Server{
		Addr:              *httpAddr,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	go func() {
		logger.Infof(
			"connect to http://localhost%s",
			*httpAddr,
		)
		errc <- http.ListenAndServe()
	}()
}
