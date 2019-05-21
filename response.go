package myhttp

import (
	"github.com/axgle/mahonia"
	"github.com/wylyeak/simplejson"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Resp *http.Response

	bodyFlag bool

	strBody string
	bodyErr error
}

func (r *Response) toggleBodyFlag() {
	r.bodyFlag = true
}

func (r *Response) StringBody() (string, error) {
	if r.bodyFlag {
		return r.strBody, r.bodyErr
	} else {
		defer r.Resp.Body.Close()
		defer r.toggleBodyFlag()
		charset := "UTF-8"
		contentType, ok := r.Resp.Header["Content-Type"]
		if ok && contentType != nil && len(contentType) >= 1 {
			charset = strings.Split(contentType[0], "=")[1]
		}
		body, err := ioutil.ReadAll(mahonia.NewDecoder(charset).NewReader(r.Resp.Body))
		if err != nil {
			r.strBody = "ERROR"
			r.bodyErr = err
			return r.strBody, r.bodyErr
		}
		r.strBody = string(body)
		return r.strBody, err
	}
}

func (r *Response) MustStringBody() (string) {
	body, err := r.StringBody()
	if err != nil {
		panic(err)
	}
	return body
}

func (r *Response) JsonObjectBody() (*simplejson.JSONObject, error) {
	str, err := r.StringBody()
	if err != nil {
		return nil, err
	}
	json, err := simplejson.NewJSONObjectFromString(str)
	return json, err
}

func (r *Response) MustJsonObjectBody() (*simplejson.JSONObject) {
	jsonObject, err := r.JsonObjectBody()
	if err != nil {
		panic(err)
	}
	return jsonObject
}

func (r *Response) JsonArrayBody() (*simplejson.JSONArray, error) {
	str, err := r.StringBody()
	if err != nil {
		return nil, err
	}
	json, err := simplejson.NewJSONArrayFromString(str)
	return json, err
}
