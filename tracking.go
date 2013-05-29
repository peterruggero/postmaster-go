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
	params := restclient.Params{
		"tracking": number,
	}
	// Response
	res := TrackingResponse{}
	// Error
	e := struct {
		Message string
	}{}
	// Do!
	status, err := p.Get("v1", "track", params, &res, &e)
	if err != nil {
		log.Fatal(status)
		log.Fatal(err)
	}
	return res
}
