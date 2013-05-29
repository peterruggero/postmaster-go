package postmaster

import (
	"github.com/jmcvetta/restclient"
)

func (p *Postmaster) MakeUrl(version string, endpoint string) string {
	if p.BaseUrl != "" {
		return p.BaseUrl + "/" + version + "/" + endpoint
	} else {
		return "http://api.postmaster.io" + "/" + version + "/" + endpoint
	}
}

func (p *Postmaster) Get(version string, endpoint string, params restclient.Params, result interface{}, err interface{}) (status int, e error) {
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "GET",
		Params:   params,
		Result:   result,
		Error:    err,
	}
	return p.Client.Do(&rr)
}

func (p *Postmaster) Put(version string, endpoint string, data interface{}, result interface{}, err interface{}) (status int, e error) {
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "PUT",
		Data:     data,
		Result:   result,
		Error:    err,
	}
	return p.Client.Do(&rr)
}

func (p *Postmaster) Post(version string, endpoint string, data interface{}, result interface{}, err interface{}) (status int, e error) {
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "POST",
		Data:     data,
		Result:   result,
		Error:    err,
	}
	return p.Client.Do(&rr)
}

func (p *Postmaster) Delete(version string, endpoint string, result interface{}, err interface{}) (status int, e error) {
	rr := restclient.RequestResponse{
		Url:      p.MakeUrl(version, endpoint),
		Userinfo: p.Userinfo,
		Method:   "DELETE",
		Result:   result,
		Error:    err,
	}
	return p.Client.Do(&rr)
}
