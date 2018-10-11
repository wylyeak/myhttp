package myhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	FormContentType = "application/x-www-form-urlencoded"
	UserAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
	JsonContentType = "application/json;charset=UTF-8"
)

type HttpClient struct {
	client *http.Client
	header map[string]string
}

func (client *HttpClient) NewRequest(method, urlStr string, contentType string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return req, err
	} else {
		req.Header.Set("User-Agent", UserAgent)
		req.Header.Set("Content-Type", contentType)
		if client.header != nil {
			for k, v := range client.header {
				req.Header.Set(k, v)
			}
		}
		return req, err
	}
}

func (client *HttpClient) PostForm(url string, data url.Values) (*Response, error) {
	return client.Post(url, FormContentType, strings.NewReader(data.Encode()))
}

func (client *HttpClient) PostJson(url string, obj map[string]interface{}) (*Response, error) {
	by, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return client.Post(url, JsonContentType, strings.NewReader(string(by)))
}

func (client *HttpClient) MustPostJson(url string, obj interface{}) (*Response) {
	by, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	response, err := client.Post(url, JsonContentType, strings.NewReader(string(by)))
	if err != nil {
		panic(err)
	}
	return response
}

func (client *HttpClient) Post(url string, contentType string, body io.Reader) (*Response, error) {
	req, err := client.NewRequest("POST", url, contentType, body)
	if err != nil {
		return nil, err
	}
	return client.do(req)
}

func (client *HttpClient) Get(url string, data url.Values) (*Response, error) {
	urlGet := &urlGet{url: url, data: data}
	req, err := client.NewRequest("GET", urlGet.parseUrl(), JsonContentType, nil)
	if err != nil {
		return nil, err
	}
	return client.do(req)
}

func (client *HttpClient) MustGet(url string, data url.Values) (*Response) {
	urlGet := &urlGet{url: url, data: data}
	req, err := client.NewRequest("GET", urlGet.parseUrl(), JsonContentType, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.do(req)
	if err != nil {
		panic(err)
	}
	return response
}

func (client *HttpClient) do(req *http.Request) (*Response, error) {
	oResp, err := client.client.Do(req)
	return &Response{Resp: oResp}, err
}

func NewHttpClient(header map[string]string) *HttpClient {
	return &HttpClient{client: &http.Client{Timeout: time.Second * 10}, header: header}
}
