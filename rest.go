package postmaster

import (
	"github.com/jmcvetta/restclient"
	"net/http"
)

// Get makes a HTTP GET request. Parameters must be provided in params.
func (p *Postmaster) get(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
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

// Put makes a HTTP PUT request. Parameters must be provided in params, and will
// be translated into query string.
func (p *Postmaster) put(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "PUT",
		Data:     urlencode(params),
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

// Post makes a HTTP POST request. Parameters must be provided in params, and will
// be translated into query string.
func (p *Postmaster) post(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "POST",
		Data:     urlencode(params),
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

// PostJson makes a HTTP POST request, with parameters as struct and being encoded to JSON.
// Currently the only function that utilizes this is *postmaster.Fit(), but it may change in future.
// Remember that every field of params structure must have a "json" comment, or json.Marshal will
// use its tentacles to make bad things to your data!
func (p *Postmaster) postJson(version string, endpoint string, params interface{}, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	headers := make(http.Header)
	for k, v := range *p.Headers {
		headers[k] = v
	}
	headers["Content-Type"] = []string{"application/json"}
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "POST",
		Data:     params,
		Result:   result,
		Error:    &err,
		Header:   &headers,
	}
	status, e = p.Client.Do(&rr)
	if status >= 300 {
		e = err
	}
	return
}

// Delete makes a HTTP DELETE request. Parameters must be provided in params, and will
// be translated into query string.
func (p *Postmaster) del(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "DELETE",
		Data:     urlencode(params),
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
