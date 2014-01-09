package postmaster

import (
	"errors"
	"fmt"
	"strconv"
)

// Shipment is a base object used in Shipment API requests.
// Options will always be a nested map of strings, meaning that you need to type
// every nested interface as map[string]interface{} or map[string]string.
type Shipment struct {
	p  *Postmaster `json:"-"`
	Id int         `json:"id,omitempty"`
	// These fields are filled by User
	To         *Address               `json:"to,omitempty"`
	From       *Address               `json:"from,omitempty"`
	Package    *Package               `json:"package,omitempty"`
	Packages   []Package              `json:"packages,omitempty"`
	Carrier    string                 `json:"carrier"`
	Service    string                 `json:"service"`
	PONumber   string                 `json:"po_number,omitempty"`
	References []string               `json:"references,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty"`
	Signature  string                 `json:"signature,omitempty"`
	Label      *Label                 `json:"label,omitempty"`
	// These fields are returned by server
	Status       string   `json:"status,omitempty"`
	Tracking     []string `json:"tracking,omitempty"`
	PackageCount int      `json:"package_count,omitempty"`
	CreatedAt    int      `json:"created_at,omitempty"`
	Cost         int      `json:"cost,omitempty"`
	Prepaid      bool     `json:"prepaid,omitempty"`
}

// ShipmentList is returned when asking for list of shipments.
type ShipmentList struct {
	Results        []Shipment `json:"results"`
	Cursor         string     `json:"cursor,omitempty"`
	PreviousCursor string     `json:"previous_cursor,omitempty"`
}

// Package (not to be confused with packages in fitting API, which are called "Boxes")
// is being used in Shipment request.
type Package struct {
	Id             int     `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Width          float32 `json:"width,omitempty"`
	Height         float32 `json:"height,omitempty"`
	Length         float32 `json:"length,omitempty"`
	Weight         float32 `json:"weight,omitempty"`
	Customs        *Custom `json:"customs,omitempty"`
	DimensionUnits string  `json:"dimension_units,omitempty"`
	WeightUnits    string  `json:"weight_units,omitempty"`
	Type           string  `json:"type,omitempty"`
	LabelUrl       string  `json:"label_url,omitempty"`
}

// CustomContent is being used as a single item in Custom object.
type CustomContent struct {
	Description     string  `json:"description,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	Value           string  `json:"value,omitempty"`
	Weight          float32 `json:"weight,omitempty"`
	WeightUnits     string  `json:"weight_units,omitempty"`
	HSTariffNumber  string  `json:"hs_tariff_number,omitempty"`
	CountryOfOrigin string  `json:"country_of_origin,omitempty"`
}

// Custom is being used per Package. It is necessary only in international
// packages.
type Custom struct {
	Type          string          `json:"type,omitempty"`
	Comments      string          `json:"comments,omitempty"`
	InvoiceNumber string          `json:"invoice_number,omitempty"`
	Contents      []CustomContent `json:"contents,omitempty"`
}

// Label is used per Shipment
type Label struct {
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Size   string `json:"size,omitempty"`
}

// Shipment creates a brand new Shipment structure. Don't use new(postmaster.Shipment),
// use this function instead.
func (p *Postmaster) Shipment() (s *Shipment) {
	s = new(Shipment)
	s.p = p
	s.Id = -1 // default for "null" Shipment
	return
}

// Create creates new Shipment in API.
// You musn't invoke this function from an existing Shipment (i.e. shipment.Id > -1).
func (s *Shipment) Create() (*Shipment, error) {
	if s.Id != -1 {
		return nil, errors.New("You can't create an existing shipment.")
	}
	_, err := post(s.p, "v1", "shipments", s, s)
	return s, err
}

// Get fetches single Shipment from API, and replaces existing Shipment structure.
// You musn't invoke this function from an "empty" Shipment (i.e. shipment.Id == -1).
func (s *Shipment) Get() (*Shipment, error) {
	if s.Id == -1 {
		return nil, errors.New("You must provide a shipment ID.")
	}
	endpoint := fmt.Sprintf("shipments/%d", s.Id)
	_, err := get(s.p, "v1", endpoint, nil, s)
	return s, err
}

// Void sets Shipment's status to "voided".
// You musn't invoke this function from an "empty" Shipment (i.e. shipment.Id == -1).
func (s *Shipment) Void() (bool, error) {
	if s.Id == -1 {
		return false, errors.New("You must provide a shipment ID.")
	}
	endpoint := fmt.Sprintf("shipments/%d/void", s.Id)
	var res map[string]string
	_, err := del(s.p, "v1", endpoint, nil, &res)
	if res["message"] == "OK" {
		s.Status = "Voided"
	}
	return res["message"] == "OK", err
}

// Track returns TrackingResponse for Shipment.
// You musn't invoke this function from an "empty" Shipment (i.e. shipment.Id == -1).
// In order to track shipment just by its tracking number, use Postmaster.TrackRef()
// function.
func (s *Shipment) Track() (*TrackingResponse, error) {
	if s.Id == -1 {
		return nil, errors.New("You must provide a shipment ID.")
	}
	endpoint := fmt.Sprintf("shipments/%d/track", s.Id)
	res := TrackingResponse{}
	_, err := get(s.p, "v1", endpoint, nil, &res)
	return &res, err
}

// ListShipments returns a list of shipments, with limit, status and cursor (e.g. for pagination).
func (p *Postmaster) ListShipments(limit int, cursor string, status string) (*ShipmentList, error) {
	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	if status != "" {
		params["status"] = status
	}
	res := new(ShipmentList)
	_, err := get(p, "v1", "shipments", params, &res)
	// Set Postmaster "base" object for each shipment, so we can use API with them
	for k, _ := range res.Results {
		res.Results[k].p = p
	}
	return res, err
}

// FindShipments returns a list of shipments matching given search query, with limit,
// status and cursor (e.g. for pagination).
func (p *Postmaster) FindShipments(q string, limit int, cursor string) (*ShipmentList, error) {
	params := make(map[string]string)
	if q == "" {
		return nil, errors.New("You must provide search query.")
	}
	params["q"] = q
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	res := new(ShipmentList)
	_, err := get(p, "v1", "shipments/search", params, &res)
	// Set Postmaster "base" object for each shipment, so we can use API with them
	for k, _ := range res.Results {
		res.Results[k].p = p
	}
	return res, err
}
