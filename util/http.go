package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetReq[ResultT any](url string, headers, params map[string]string) (ResultT, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	if len(params) != 0 {
		url += "?"
	}
	for k, v := range params {
		url = url + k + "=" + v + "&"
	}
	var result ResultT
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	jsonDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}
		return result, fmt.Errorf("code: %d, Resp: %s", resp.StatusCode, string(body))
	}
	err = jsonDecoder.Decode(&result)
	return result, err
}

func Request[TBody any, TResult any](method, url string, headers, params map[string]string, reqBody TBody) (TResult, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	if len(params) != 0 {
		url += "?"
	}
	for k, v := range params {
		url = url + k + "=" + v + "&"
	}
	var (
		result    TResult
		bodyBytes = new(bytes.Buffer)
	)
	jsonEn := json.NewEncoder(bodyBytes)
	err := jsonEn.Encode(reqBody)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest(method, url, bodyBytes)
	if err != nil {
		return result, err
	}
	req.Header.Set("accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	jsonDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}
		return result, fmt.Errorf("code: %d, Resp: %s", resp.StatusCode, string(body))
	}
	err = jsonDecoder.Decode(&result)
	return result, err
}
