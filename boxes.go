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

type BoxList struct {
	Results        []Box
	Cursor         string
	PreviousCursor string `json:"previous_cursor"`
}

type FitMessage struct {
	Packages     []Box
	Items        []Item
	PackageLimit int `json:"package_limit"`
}

func (p *Postmaster) Box() (b *Box) {
	b = new(Box)
	b.p = p
	b.Id = -1
	return
}

func (b *Box) Create() (*Box, error) {
	if b.Id != -1 {
		return nil, errors.New("You can't create an existing box.")
	}
	params := MapStruct(b)
	res := map[string]int{}
	_, err := b.p.Post("v1", "packages", params, &res)
	if err == nil {
		b.Id = res["id"]
	}
	return b, err
}

func (b *Box) Get() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	_, err := b.p.Get("v1", endpoint, nil, b)
	return b, err
}

func (b *Box) Delete() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	res := map[string]string{}
	_, err := b.p.Delete("v1", endpoint, nil, &res)
	b = b.p.Box()
	return b, err
}

func (b *Box) Update() (*Box, error) {
	if b.Id == -1 {
		return nil, errors.New("You must provide a box ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", b.Id)
	params := MapStruct(b)
	res := map[string]string{}
	_, err := b.p.Put("v1", endpoint, params, &res)
	return b, err
}

func (b *Box) List(limit int, cursor string) (*BoxList, error) {
	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	res := new(BoxList)
	_, err := b.p.Get("v1", "packages", params, &res)
	// Set Postmaster "base" object for each package, so we can use API with them
	for k, _ := range res.Results {
		res.Results[k].p = b.p
	}
	return res, err
}
