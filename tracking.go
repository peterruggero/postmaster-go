package postmaster

import (
	"fmt"
	"github.com/jmcvetta/restclient"
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

func (p *Postmaster) TrackRef(trackingNumber string) (*TrackingResponse, error) {
	params := restclient.Params{
		"tracking": trackingNumber,
	}
	res := TrackingResponse{}
	_, err := p.Get("v1", "track", params, &res)
	return &res, err
}

func (p *Postmaster) TrackShipment(shipmentId int) (*TrackingResponse, error) {
	endpoint := fmt.Sprintf("shipments/%d/track", shipmentId)
	res := TrackingResponse{}
	_, err := p.Get("v1", endpoint, nil, &res)
	return &res, err
}
