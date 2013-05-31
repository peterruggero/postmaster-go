package postmaster

type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	ZipCode    string `json:"zip_code"`
	County     string
	Latitude   string
	Longitude  string
	Notes      string
	Active     bool
	Commercial bool
}

type AddressResponse struct {
	Status    string
	Addresses []Address
}

func (p *Postmaster) Validate(addr Address) (*AddressResponse, error) {
	params := MapStruct(addr)
	res := AddressResponse{}
	_, err := p.Post("v1", "validate", params, &res)
	return &res, err
}
