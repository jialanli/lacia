package lacia

import (
	"testing"
	"time"
)

func TestDoHttpRequest(t *testing.T) {
	_, err := DoHttpRequest("GET", "https://baidu.com", time.Duration(10*time.Second), nil)
	t.Log(err)

	curlHeader := make(map[string]string)
	curlHeader["Content-Type"] = "application/json"
	_, err = DoHttpRequestWithHeader("GET", "https://baidu.com", time.Duration(10*time.Second), nil, curlHeader)
	t.Log(err)
}
