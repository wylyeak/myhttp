package myhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
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

type IClient interface {
	PostForm(url string, data url.Values) (*Response, error)
	PostJson(url string, obj interface{}) (*Response, error)
	MustPostJson(url string, obj interface{}) (*Response)
	Get(url string, data url.Values) (*Response, error)
	MustGet(url string, data url.Values) (*Response)
}

func (client *HttpClient) PostForm(url string, data url.Values) (*Response, error) {
	return client.post(url, FormContentType, strings.NewReader(data.Encode()))
}

func (client *HttpClient) MustPostForm(url string, data url.Values) (*Response) {
	resp, err := client.PostForm(url, data)
	if err != nil {
		panic(err)
	}
	return resp
}


func (client *HttpClient) PostJson(url string, obj interface{}) (*Response, error) {
	by, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return client.post(url, JsonContentType, strings.NewReader(string(by)))
}

func (client *HttpClient) MustPostJson(url string, obj interface{}) (*Response) {
	resp, err := client.PostJson(url, obj)
	if err != nil {
		panic(err)
	}
	return resp
}

func (client *HttpClient) post(url string, contentType string, body io.Reader) (*Response, error) {
	req, err := client.newRequest("POST", url, contentType, body)
	if err != nil {
		return nil, err
	}
	return client.do(req)
}

func (client *HttpClient) mustPost(url string, contentType string, body io.Reader) (*Response) {
	resp, err := client.post(url, contentType, body)
	if err != nil {
		panic(err)
	}
	return resp
}

func (client *HttpClient) Get(url string, data url.Values) (*Response, error) {
	urlGet := &urlGet{url: url, data: data}
	req, err := client.newRequest("GET", urlGet.parseUrl(), JsonContentType, nil)
	if err != nil {
		return nil, err
	}
	return client.do(req)
}

func (client *HttpClient) MustGet(url string, data url.Values) (*Response) {
	resp, err := client.Get(url, data)
	if err != nil {
		panic(err)
	}
	return resp
}

func (client *HttpClient) do(req *http.Request) (*Response, error) {
	oResp, err := client.client.Do(req)
	return &Response{Resp: oResp}, err
}

func (client *HttpClient) newRequest(method, urlStr string, contentType string, body io.Reader) (*http.Request, error) {
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

func NewHttpClient(header map[string]string) *HttpClient {
	jar, _ := cookiejar.New(nil)
	return &HttpClient{client: &http.Client{Jar: jar, Timeout: time.Second * 10}, header: header}
}
