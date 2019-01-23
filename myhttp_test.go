package myhttp

import "testing"

func TestHttpClient_Get(t *testing.T) {
	client := NewHttpClient(nil)
	resp, err := client.Get("https://www.baidu.com", nil)
	if err != nil {
		t.Error(err)
	}
	if resp.Resp.Request.URL.String() == "https://www.baidu.com" {
		return
	}
	t.Error(resp.Resp.Request.URL.String())
}

func TestHttpClient_MustGet(t *testing.T) {
	client := NewHttpClient(nil)
	resp:= client.MustGet("https://www.baidu.com", nil)
	if resp.Resp.Request.URL.String() == "https://www.baidu.com" {
		return
	}
	t.Error(resp.Resp.Request.URL.String())
}

func TestHttpClient_PostForm(t *testing.T) {

}

func TestHttpClient_Post(t *testing.T) {
}

func TestHttpClient_PostJson(t *testing.T) {
}
