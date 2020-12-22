package lacia

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// send Request
func DoHttpRequest(method, url string, body io.Reader) ([]byte, string, error) {
	httpClient := &http.Client{}

	// generate request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, "", errors.New(fmt.Sprintf("do http request failed:%s", resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, resp.Header.Get("Content-Type"), err
}
