package mps

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewTunnelHandler(t *testing.T) {
	tunnel := NewTunnelHandler()
	//tunnel.Transport().Proxy = func(r *http.Request) (*url.URL, error) {
	//	//return url.Parse("http://59.58.58.92:4235")
	//	return url.Parse("http://127.0.0.1:7890")
	//}
	tunnel.Transport().Dial = nil
	tunnelSrv := httptest.NewServer(tunnel)
	defer tunnelSrv.Close()

	req, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	http.DefaultClient.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse(tunnelSrv.URL)
		},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(err)
	log.Println(resp.Status)
	log.Println(string(body))
}
