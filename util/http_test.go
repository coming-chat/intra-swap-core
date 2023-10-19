package util

import (
	"net/http"
	"testing"
)

func TestGetReq(t *testing.T) {
	resp, err := GetReq[map[string]any]("https://api.geckoterminal.com/api/v2/search/pools?query=ETH&network=base", map[string]string{
		"Accept":     "application/json;version=20230302",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestRequest(t *testing.T) {
	res, err := Request[map[string]interface{}, map[string]interface{}](http.MethodGet, "https://api.geckoterminal.com/api/v2/search/pools?query=ETH&network=base", map[string]string{
		"Accept":     "application/json;version=20230302",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
