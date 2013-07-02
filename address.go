package postmaster

// Address is used in Shipment requests (as From or To fields), or in validating
// addresses.
type Address struct {
	Company    string `json:"company,omitempty"`
	Contact    string `json:"contact,omitempty"`
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	Line3      string `json:"line3,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	ZipCode    string `json:"zip_code,omitempty"`
	Country    string `json:"country,omitempty"`
	Latitude   string `json:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty"`
	Notes      string `json:"notes,omitempty"`
	PhoneNo    string `json:"phone_no,omitempty"`
	Active     bool   `json:"active,omitempty"`
	Commercial bool   `json:"commercial,omitempty"`
	Residental bool   `json:"residental,omitempty"`
}

// AddressResponse is being sent back from API when asking to validate an address.
type AddressResponse struct {
	Status    string
	Addresses []Address
}

// Validate tries to validate given address.
func (p *Postmaster) Validate(addr *Address) (*AddressResponse, error) {
	res := new(AddressResponse)
	_, err := post(p, "v1", "validate", addr, &res)
	return res, err
}
