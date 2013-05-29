package postmaster

import (
	"github.com/jmcvetta/restclient"
	"log"
)

type TrackingHistory struct {
	Status      string
	City        string
	Timestamp   int
	Code        string
	Description string
}

type TrackingResponse struct {
	Status     string
	LastUpdate int `json:"last_update"`
	History    []TrackingHistory
}

func (p *Postmaster) Track(number string) TrackingResponse {
	// Request
	dd := restclient.Params{
		"tracking": number,
	}
	// Response
	res := TrackingResponse{}
	// Error
	e := struct {
		Message string
	}{}
	// Do!
	var url string
	if p.BaseUrl != "" {
		url = p.BaseUrl + "/v1/track"
	} else {
		url = "http://api.postmaster.io/v1/track"
	}
	rr := restclient.RequestResponse{
		Url:      url,
		Userinfo: p.Userinfo,
		Method:   "GET",
		Params:   dd,
		Result:   &res,
		Error:    &e,
	}
	status, err := p.Client.Do(&rr)
	if err != nil {
		log.Fatal(status)
		log.Fatal(err)
	}
	return res
}
