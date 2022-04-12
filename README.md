[![Maintainability](https://api.codeclimate.com/v1/badges/12c7a9e048efae2a0615/maintainability)](https://codeclimate.com/github/LUXROBO/kakao-api/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/12c7a9e048efae2a0615/test_coverage)](https://codeclimate.com/github/LUXROBO/kakao-api/test_coverage)

# KAKAO
카카오 API 모든 서비스 담당 ...


### Test & coverage output

```

make testing

```

### 사용방법

```go
package main

import (
	kakaoapi "github.com/LUXROBO/kakao-api/kakao"
)

func main() {
  // 환경설정 
  // KAKAO_CLIENT_ID: 카카오 REST API키 
  // KAKAO_REDIRECT_URL: 카카오 Redreict API주소 
  // 사용하기전 위와같이 환경설정 해줘야함. 
	kakaoClient := kakaoapi.NewClient(os.Getenv("KAKAO_CLIENT_ID"), os.Getenv("KAKAO_REDIRECT_URL"))
  ....

}

```
