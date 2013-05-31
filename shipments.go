package postmaster

import ( 
	"fmt"
	"errors"
)

// Shipment "object"
type Shipment struct {
	p       *Postmaster
	Id      int
	To      Address
	From    Address
	Package Package
	Carrier string
	Service string
}

func (p *Postmaster) Shipment() (s *Shipment) {
	s = new(Shipment)
	s.p = p
	s.Id = -1 // default for "null" Shipment
	return
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
