package lacia

import (
	"fmt"
	"testing"
	"time"
)

func TestDoHttpRequest(t *testing.T) {
	//var buff bytes.Buffer
	//buff.Write(reqDataJson)
	//body := bytes.NewReader(postByte)
	_, err := DoHttpRequest("GET", "https://console.cloud.tencent.com/im", time.Duration(30*time.Second), nil)
	fmt.Println(err)

	curlHeader := make(map[string]string)
	curlHeader["Content-Type"] = "application/json"
	_, err = DoHttpRequestWithHeader("GET", "https://console.cloud.tencent.com/trtc", time.Duration(30*time.Second), nil, curlHeader)
	fmt.Println(err)
}
