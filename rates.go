package postmaster

// RateResponse contains response for single Carrier.
type RateResponse struct {
	Service        string
	Charge         string
	Currency       string
	GuarrantedDays string `json:"guarranted_days"`
}

// RateResponseBest is being returned if Carrier is empty.
type RateResponseBest struct {
	Fedex RateResponse
	UPS   RateResponse
	USPS  RateResponse
	Best  string
}

// RateMessage is being used in query to find delivery rates for single package.
type RateMessage struct {
	FromZip    string `json:"from_zip"`
	ToZip      string `json:"to_zip"`
	Weight     float32
	Carrier    string
	Packaging  string
	Commercial bool
	Service    string
}

// Rate asks API for delivery cost between two ZIP codes. If you provide a Carrier
// in your RateMessage, single RateResponse for given Carrier will be returned.
// If Carrier is left empty, a RateResponseBest structure is returned, with one
// RateResponse per carrier.
func (p *Postmaster) Rate(r *RateMessage) (interface{}, error) {
	params := mapStruct(r)
	if r.Carrier != "" {
		res := RateResponse{}
		_, err := p.post("v1", "rates", params, &res)
		return res, err
	} else {
		res := RateResponseBest{}
		_, err := p.post("v1", "rates", params, &res)
		return res, err
	}
}
