package postmaster

import (
	"errors"
	"fmt"
)

// Shipment "object"
type Shipment struct {
	p       *Postmaster `dontMap:"true"`
	Id      int         `dontMap:"true"`
	To      Address
	From    Address
	Package Package
	Carrier string
	Service string
}

type ShipmentResponse struct {
	Id           int
	To           Address
	From         Address
	Package      Package
	Carrier      string
	Service      string
	Status       string
	Tracking     []string
	PackageCount int `json:"package_count"`
	CreatedAt    int `json:"created_at"`
	Cost         int
	Packages     []Package
	Prepaid      bool
}

func (p *Postmaster) Shipment() (s *Shipment) {
	s = new(Shipment)
	s.p = p
	s.Id = -1 // default for "null" Shipment
	return
}

func (s *Shipment) Create() (*ShipmentResponse, error) {
	if s.Id != -1 {
		return nil, errors.New("You can't create an existing shipment.")
	}
	params := MapStruct(s)
	res := ShipmentResponse{}
	_, err := s.p.Post("v1", "shipments", params, &res)
	return &res, err
}

func (s *Shipment) Void() (bool, error) {
	if s.Id == -1 {
		return false, errors.New("You must provide an shipment ID.")
	}
	endpoint := fmt.Sprintf("shipments/%d/void", s.Id)
	var res map[string]string
	_, err := s.p.Delete("v1", endpoint, nil, &res)
	return res["message"] == "OK", err
}

func (s *Shipment) Track() (*TrackingResponse, error) {
	if s.Id == -1 {
		return nil, errors.New("You must provide an shipment ID.")
	}
	endpoint := fmt.Sprintf("shipments/%d/track", s.Id)
	res := TrackingResponse{}
	_, err := s.p.Get("v1", endpoint, nil, &res)
	return &res, err
}
