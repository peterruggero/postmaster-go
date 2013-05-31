package postmaster

type Address struct {
	Company    string
	Contact    string
	Line1      string
	Line2      string
	Line3      string
	City       string
	State      string
	ZipCode    string `json:"zip_code"`
	County     string
	Latitude   string
	Longitude  string
	Notes      string
	PhoneNo    string `json:"phone_no"`
	Active     bool   `dontMap:"true"`
	Commercial bool   `dontMap:"true"`
	Residental bool   `dontMap:"true"`
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
