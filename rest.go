package postmaster

import (
	"github.com/jmcvetta/restclient"
	"net/http"
)

// get makes a HTTP GET request. Parameters must be provided in params.
var get = func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.userinfo,
		Method:   "GET",
		Params:   params,
		Result:   result,
		Error:    &err,
		Header:   p.headers,
	}
	status, e = p.client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

// put makes a HTTP PUT request. Parameters must be provided in params, and will
// be translated into query string.
var put = func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.userinfo,
		Method:   "PUT",
		Data:     urlencode(params),
		Result:   result,
		Error:    &err,
		Header:   p.headers,
	}
	status, e = p.client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

// post makes a HTTP POST request. Parameters must be provided in params, and will
// be translated into query string.
var post = func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.userinfo,
		Method:   "POST",
		Data:     urlencode(params),
		Result:   result,
		Error:    &err,
		Header:   p.headers,
	}
	status, e = p.client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

// postJson makes a HTTP POST request, with parameters as struct and being encoded to JSON.
// Currently the only function that utilizes this is *postmaster.Fit(), but it may change in future.
// Remember that every field of params structure must have a "json" comment, or json.Marshal will
// use its tentacles to make bad things to your data!
var postJson = func(p *Postmaster, version string, endpoint string, params interface{}, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	headers := make(http.Header)
	for k, v := range *p.headers {
		headers[k] = v
	}
	headers["Content-Type"] = []string{"application/json"}
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.userinfo,
		Method:   "POST",
		Data:     params,
		Result:   result,
		Error:    &err,
		Header:   &headers,
	}
	status, e = p.client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

// delete makes a HTTP DELETE request. Parameters must be provided in params, and will
// be translated into query string.
var del = func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.userinfo,
		Method:   "DELETE",
		Data:     urlencode(params),
		Result:   result,
		Error:    &err,
		Header:   p.headers,
	}
	status, e = p.client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}
