package postmaster

import (
	"github.com/jmcvetta/restclient"
)

func (p *Postmaster) Get(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
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

func (p *Postmaster) Put(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
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

func (p *Postmaster) Post(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
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

func (p *Postmaster) Delete(version string, endpoint string, params restclient.Params, result interface{}) (status int, e error) {
	err := new(PostmasterError)
	rr := restclient.RequestResponse{
		Url:      p.makeUrl(version, endpoint),
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
