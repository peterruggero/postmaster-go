package postmaster

import (
	"github.com/jmcvetta/restclient"
)

// TrackingHistory is a part of TrackingResponse.
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

// TrackingResponse is being sent back from API when tracking shipment and
// tracking shipment by its reference number. 
type TrackingResponse struct {
	Status     string
	LastUpdate int `json:"last_update"`
	SignedBy   string `json:"signed_by"`
	History    []TrackingHistory
}

// TrackRef method allows to track shipment by its reference number.
func (p *Postmaster) TrackRef(trackingNumber string) (*TrackingResponse, error) {
	params := restclient.Params{
		"tracking": trackingNumber,
	}
	res := TrackingResponse{}
	_, err := p.get("v1", "track", params, &res)
	return &res, err
}
