/*
	WARNING!!!
	Because of many, many inconsistencies in our API (duh) this file is a STUB
	until someone fixes our API calls and docs.
*/
package postmaster

import (
	"errors"
	"fmt"
)

type Package struct {
	p              *Postmaster `dontMap:"true"`
	Id             int         `dontMap:"true"`
	Name           string
	Width          float32
	Height         float32
	Length         float32
	Weight         float32
	WeightUnits    string `dontMap:"true" json:"weight_units"`
	Type           string `dontMap:"true"`
	LabelUrl       string `dontMap:"true" json:"label_url"`
	DimensionUnits string `dontMap:"true" json:"dimension_units"`
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
	Count       int
}

func (p *Postmaster) Package() (pack *Package) {
	pack = new(Package)
	pack.p = p
	pack.Id = -1
	return
}

func (p *Package) Create() (int, error) {
	if p.Id != -1 {
		return -1, errors.New("You can't create an existing package.")
	}
	params := MapStruct(p)
	res := map[string]int{}
	_, err := p.p.Post("v1", "packages", params, &res)
	if err == nil {
		p.Id = res["id"]
	}
	return p.Id, err
}

func (p *Package) Get() (*Package, error) {
	if p.Id == -1 {
		return nil, errors.New("You must provide a package ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", p.Id)
	_, err := p.p.Get("v1", endpoint, nil, p)
	return p, err
}

func (p *Package) Delete() (*Package, error) {
	if p.Id == -1 {
		return nil, errors.New("You must provide a package ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", p.Id)
	res := map[string]string{}
	_, err := p.p.Delete("v1", endpoint, nil, &res)
	p = p.p.Package()
	return p, err
}

func (p *Package) Update() (*Package, error) {
	if p.Id == -1 {
		return nil, errors.New("You must provide a package ID.")
	}
	endpoint := fmt.Sprintf("packages/%d", p.Id)
	params := MapStruct(p)
	res := map[string]string{}
	_, err := p.p.Put("v1", endpoint, params, &res)
	return p, err
}
