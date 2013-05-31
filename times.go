package postmaster

// TimeResponseItem is a part of TimeResponse.
type TimeResponseItem struct {
	Service           string
	DeliveryTimestamp int    `json:"delivery_timestamp"`
	DeliveryDesc      string `json:"delivery_desc"`
}

// TimeResponse is being returned by Postmaster.Time().
type TimeResponse struct {
	Services []TimeResponseItem
}

// TimeMessage is being sent to API when calling Postmaster.Time().
type TimeMessage struct {
	FromZip    string `json:"from_zip"`
	ToZip      string `json:"to_zip"`
	Weight     float32
	Carrier    string
	Commercial bool
}

// Time asks API for time to transport a shipment between two ZIP codes.
func (p *Postmaster) Time(t TimeMessage) (*TimeResponse, error) {
	params := MapStruct(t)
	res := TimeResponse{}
	_, err := p.Post("v1", "times", params, &res)
	return &res, err
}
