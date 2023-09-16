package util

import (
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
