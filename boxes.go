package postmaster

import (
	"errors"
	"fmt"
	"strconv"
)

// Box (other name: Package; used Box here in order not to confuse with Shipment's Package)
// is a container that we try to fit Items into.
// Reason why this structure is wholly JSON-mapped (every field has a "json" comment)
// is that fitting package request is the only request that sends data using JSON
// and not query params. Thus, we can't use our mapStruct, but must use json package
// and its Marshal function.
type Box struct {
	p           *Postmaster `json:"-" dontMap:"true"`
	Id          int         `json:"-" dontMap:"true"`
	Name        string      `json:"name"`
	Width       float32     `json:"width"`
	Height      float32     `json:"height"`
	Length      float32     `json:"length"`
	Weight      float32     `json:"weight"`
	SizeUnits   string      `dontMap:"true" json:"size_units,omitempty"`
	WeightUnits string      `dontMap:"true" json:"weight_units,omitempty"`
}

// Item is an object we try to fit into Boxes.
// Reason why this structure is wholly JSON-mapped (every field has a "json" comment)
// is that fitting package request is the only request that sends data using JSON
// and not query params. Thus, we can't use our mapStruct, but must use json package
// and its Marshal function.
type Item struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name,omitempty"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Length      float32 `json:"length"`
	Weight      float32 `json:"weight"`
	Count       int     `json:"count"`
	SizeUnits   string  `json:"size_units,omitempty"`
	WeightUnits string  `json:"weight_units,omitempty"`
}

// BoxList is API response for List() function.
type BoxList struct {
	Results        []Box
	Cursor         string
	PreviousCursor string `json:"previous_cursor"`
}

// FitMessage is being sent to API in order to check whether given items fit given boxes.
// Reason why this structure is wholly JSON-mapped (every field has a "json" comment)
// is that fitting package request is the only request that sends data using JSON
// and not query params. Thus, we can't use our mapStruct, but must use json package
// and its Marshal function.
type FitMessage struct {
	Boxes        []Box  `json:"boxes"`
	Items        []Item `json:"items"`
	PackageLimit int    `json:"package_limit"`
}

// FitResponse is API response for trying to fit items into boxes.
type FitResponse struct {
	Boxes []struct {
		Box   Box
		Items []Item
	}
	Leftovers []Item
	AllFit    bool `json:"all_fit"`
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

// Fit checks if given items can be packed into given boxes.
// Currently this is the only function that utilizes postJson function.
func (p *Postmaster) Fit(boxes []Box, items []Item, limit int) (*FitResponse, error) {
	params := FitMessage{
		Boxes:        boxes,
		Items:        items,
		PackageLimit: limit,
	}
	res := new(FitResponse)
	_, err := p.postJson("v1", "packages/fit", params, &res)
	return res, err
}
