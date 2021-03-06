package lacia

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// send Request
func DoHttpRequest(method, url string, timeout time.Duration, body io.Reader) ([]byte, error) {
	httpClient := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("do http request failed:%s", resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, err
}

// send Request with selfHeader
func DoHttpRequestWithHeader(method, url string, timeout time.Duration, body io.Reader, header map[string]string) (res []byte, err error) {
	//var timeout = time.Duration(30 * time.Second)
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
