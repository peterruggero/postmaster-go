package postmaster

import (
	"fmt"
	"github.com/jmcvetta/restclient"
	"net/url"
	"strings"
)

func (p *Postmaster) MakeUrl(version string, endpoint string) string {
	var url string
	if p.BaseUrl != "" {
		url = p.BaseUrl
	} else {
		url = "http://api.postmaster.io"
	}
	return fmt.Sprintf("%s/%s/%s", url, version, endpoint)
}

func (p *Postmaster) Urlencode(params map[string]string) string {
	arr := make([]string, len(params))
	for k, v := range params {
		if fmt.Sprintf("%s", v) != "" {
			arr = append(arr, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		}
	}
	return strings.Join(arr, "&")
}

func (p *Postmaster) Get(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "GET",
		Params:   params,
		Result:   result,
		Error:    &err,
		Header:   p.Headers,
	}
	status, e = p.Client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

func (p *Postmaster) Put(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "PUT",
		Data:     p.Urlencode(params),
		Result:   result,
		Error:    &err,
		Header:   p.Headers,
	}
	status, e = p.Client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

func (p *Postmaster) Post(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	fmt.Println(p.Urlencode(params))
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "POST",
		Data:     p.Urlencode(params),
		Result:   result,
		Error:    &err,
		Header:   p.Headers,
	}
	status, e = p.Client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

func (p *Postmaster) Delete(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "DELETE",
		Result:   result,
		Error:    &err,
		Header:   p.Headers,
	}
	status, e = p.Client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}
