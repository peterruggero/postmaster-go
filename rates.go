package postmaster

// RateResponse contains response for single Carrier.
type RateResponse struct {
	Service  string // Type of service
	Charge   string // Cost of sending the shipment
	Currency string // Currency
}

// rateResponseBestTemp is temporary, as name indicates.
type rateResponseBestTemp struct {
	Fedex RateResponse // Rate for Fedex
	UPS   RateResponse // Rate for UPS
	USPS  RateResponse // Rate for USPS
	Best  string       // Lowercase carrier name that offers the best deal
}

// RateResponseBest is being returned if Carrier is empty.
type RateResponseBest struct {
	Rates map[string]RateResponse
	Best  string // Lowercase carrier name that offers the best deal
}

// RateMessage is being used in query to find delivery rates for single package.
type RateMessage struct {
	FromZip    string  `json:"from_zip"` // The source zip code
	ToZip      string  `json:"to_zip"`   // The destination zip code
	Weight     float32 // The weight of the package in pounds
	Carrier    string  // Which carrier to query
	Packaging  string  // What type of packaging this shipment will use (optional, default: CUSTOM)
	Commercial bool    // Is the package going to a commercial address?
	Service    string  // Which service level to quote (optional, default: GROUND)
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
		return &res, err
	} else {
		resTemp := rateResponseBestTemp{}
		_, err := p.post("v1", "rates", params, &resTemp)
		res := RateResponseBest{
			Rates: make(map[string]RateResponse),
			Best:  resTemp.Best,
		}
		res.Rates["fedex"] = resTemp.Fedex
		res.Rates["ups"] = resTemp.UPS
		res.Rates["usps"] = resTemp.USPS
		return &res, err
	}
}
