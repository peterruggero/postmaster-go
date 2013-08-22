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

// TrackingExternal is used in requests for monitoring external packages.
type TrackingExternal struct {
	p          *Postmaster `json:"-"`
	TrackingNo string      `json:"tracking_no,omitempty"`
	Url        string      `json:"url,omitempty"`
	Sms        string      `json:"sms,omitempty"`
	Events     []string    `json:"events,omitempty"`
}

func (p *Postmaster) TrackingExternal() (t *TrackingExternal) {
	t = new(TrackingExternal)
	t.p = p
	return
}

// Put sends TrackingExternal object to the server.
func (t *TrackingExternal) Put() (success bool, err error) {
	res := new(interface{})
	var status int
	status, err = post(t.p, "v1", "track", t, &res)
	success = status == 200
	return
}

// TrackRef method allows to track shipment by its reference number.
func (p *Postmaster) TrackRef(trackingNumber string) (*TrackingResponse, error) {
	params := make(map[string]string)
	params["tracking"] = trackingNumber
	res := TrackingResponse{}
	_, err := get(p, "v1", "track", params, &res)
	return &res, err
}
