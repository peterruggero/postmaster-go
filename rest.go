package postmaster

import (
	"github.com/jmcvetta/restclient"
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
