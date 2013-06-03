/*
	WARNING!!!
	Because of many, many inconsistencies in our API (duh) this file is a STUB
	until someone fixes our API calls and docs.
*/
package postmaster

import (
	"errors"
	"fmt"
	"strconv"
)

// Box (other name: Package; used Box here in order not to confuse with Shipment's Package)
// is a container that we try to fit Items into 
type Box struct {
	p           *Postmaster `dontMap:"true"`
	Id          int         `dontMap:"true"`
	Name        string
	Width       float32
	Height      float32
	Length      float32
	Weight      float32
	SizeUnits   string `dontMap:"true" json:"size_units"`
	WeightUnits string `dontMap:"true" json:"weight_units"`
}

// Item is an object we try to fit into Boxes
type Item struct {
	SKU         string
	Name        string
	Width       float32
	Height      float32
	Length      float32
	Weight      float32
	SizeUnits   string `json:"size_units"`
	WeightUnits string `json:"weight_units"`
}

// BoxList is API response for List() function
type BoxList struct {
	Results        []Box
	Cursor         string
	PreviousCursor string `json:"previous_cursor"`
}

// FitMessage is API response for trying to fit items into boxes
type FitMessage struct {
	Packages     []Box
	Items        []Item
	PackageLimit int `json:"package_limit"`
}

// Box() creates new Box and assigns all important variables. Use this instead
// of new(postmaster.Box).
func (p *Postmaster) Box() (b *Box) {
	b = new(Box)
	b.p = p
	b.Id = -1
	return
}

// Create creates new Box. Existing *Box receiver's fields will be overwritten.
// You musn't invoke this function from an existing Box (i.e. Box with ID > -1).
func (b *Box) Create() (*Box, error) {
	if b.Id != -1 {
		return nil, errors.New("You can't create an existing box.")
	}
	params := mapStruct(b)
	res := map[string]int{}
	_, err := b.p.post("v1", "packages", params, &res)
	if err == nil {
		b.Id = res["id"]
	}
	return b, err
}

// Get fetches Box from API and stores it in *Box receiver.
// You musn't invoke this function from an "empty" box (i.e. Box with ID == -1).
func (b *Box) Get() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	_, err := b.p.get("v1", endpoint, nil, b)
	return b, err
}

// Delete deletes Box, and replaces *Box receiver with an empty one.
// You musn't invoke this function from an "empty" box (i.e. Box with ID == -1).
func (b *Box) Delete() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	res := map[string]string{}
	_, err := b.p.del("v1", endpoint, nil, &res)
	b = b.p.Box()
	return b, err
}

// Update updates Box.
// You musn't invoke this function from an "empty" box (i.e. Box with ID == -1).
func (b *Box) Update() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	params := mapStruct(b)
	res := map[string]string{}
	_, err := b.p.put("v1", endpoint, params, &res)
	return b, err
}

// List returns a list of boxes, with limit and cursor (e.g. for pagination).
func (b *Box) List(limit int, cursor string) (*BoxList, error) {
	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	res := new(BoxList)
	_, err := b.p.get("v1", "packages", params, &res)
	// Set Postmaster "base" object for each package, so we can use API with them
	for k, _ := range res.Results {
		res.Results[k].p = b.p
	}
	return res, err
}
