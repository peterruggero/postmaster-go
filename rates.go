package postmaster

type RateResponse struct {
	Service        string
	Charge         string
	Currency       string
	GuarrantedDays string `json:"guarranted_days"`
}

type RateResponseBest struct {
	Fedex RateResponse
	UPS   RateResponse
	USPS  RateResponse
	Best  string
}

type RateMessage struct {
	FromZip    string `json:"from_zip"`
	ToZip      string `json:"to_zip"`
	Weight     float32
	Carrier    string
	Packaging  string
	Commercial bool
	Service    string
}

// if carrier is empty, returns RateResponseBest
// if not, single RateResponse
func (p *Postmaster) Rate(r RateMessage) (interface{}, error) {
	params := MapStruct(r)
	if r.Carrier != "" {
		res := RateResponse{}
		_, err := p.Post("v1", "rates", params, &res)
		return res, err
	} else {
		res := RateResponseBest{}
		_, err := p.Post("v1", "rates", params, &res)
		return res, err
	}
}
