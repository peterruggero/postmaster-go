package postmaster

import (
	"github.com/jmcvetta/restclient"
	"fmt"
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

func (p *Postmaster) TrackRef(trackingNumber string) TrackingResponse {
	params := restclient.Params{
		"tracking": trackingNumber,
	}
	res := TrackingResponse{}
	e := struct {
		Message string
	}{}
	status, err := p.Get("v1", "track", params, &res, &e)
	if err != nil {
		log.Fatal(status)
		log.Fatal(err)
	}
	return res
}

func (p *Postmaster) TrackShipment(shipmentId int) TrackingResponse {
	endpoint := fmt.Sprintf("shipments/%d/track", shipmentId)
	res := TrackingResponse{}
	e := struct {
		Message string
	}{}
	p.Get("v1", endpoint, nil, &res, &e)
	return res
}