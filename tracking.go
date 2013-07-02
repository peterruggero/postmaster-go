package postmaster

// TrackingHistory is a part of TrackingResponse.
type TrackingHistory struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Timestamp   int      `json:"timestamp"`
	Street      []string `json:"street"`
	PostalCode  string   `json:"postal_code"`
	CountryCode string   `json:"country_code"`
	City        string   `json:"city"`
	Code        string   `json:"code"`
	State       string   `json:"state"`
	Text        string   `json:"text"`
}

// TrackingResponse is being sent back from API when tracking shipment and
// tracking shipment by its reference number.
type TrackingResponse struct {
	Status     string            `json:"status"`
	LastUpdate int               `json:"last_update"`
	SignedBy   string            `json:"signed_by"`
	History    []TrackingHistory `json:"history"`
}

// TrackRef method allows to track shipment by its reference number.
func (p *Postmaster) TrackRef(trackingNumber string) (*TrackingResponse, error) {
	params := make(map[string]string)
	params["tracking"] = trackingNumber
	res := TrackingResponse{}
	_, err := get(p, "v1", "track", params, &res)
	return &res, err
}
