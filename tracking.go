package postmaster

import (
	"github.com/jmcvetta/restclient"
)

type TrackingHistory struct {
	Status      string
	Description string
	Timestamp   int
	Street      []string
	PostalCode  string `json:"postal_code"`
	CountryCode string `json:"country_code"`
	City        string
	Code        string
	State       string
	Text        string
}

type TrackingResponse struct {
	Status     string
	LastUpdate int `json:"last_update"`
	SignedBy   string `json:"signed_by"`
	History    []TrackingHistory
}

func (p *Postmaster) TrackRef(trackingNumber string) (*TrackingResponse, error) {
	params := restclient.Params{
		"tracking": trackingNumber,
	}
	res := TrackingResponse{}
	_, err := p.Get("v1", "track", params, &res)
	return &res, err
}
