package postmaster

// Address is used in Shipment requests (as From or To fields), or in validating
// addresses.
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

// AddressResponse is being sent back from API when asking to validate an address.
type AddressResponse struct {
	Status    string
	Addresses []Address
}

// Validate tries to validate given address.
func (p *Postmaster) Validate(addr *Address) (*AddressResponse, error) {
	params := mapStruct(addr)
	res := new(AddressResponse)
	_, err := p.post("v1", "validate", params, &res)
	return res, err
}
