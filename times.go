package postmaster

type TimeResponseItem struct {
	Service           string
	DeliveryTimestamp int    `json:"delivery_timestamp"`
	DeliveryDesc      string `json:"delivery_desc"`
}

type TimeResponse struct {
	Services []TimeResponseItem
}

type TimeMessage struct {
	FromZip    string `json:"from_zip"`
	ToZip      string `json:"to_zip"`
	Weight     float32
	Carrier    string
	Commercial bool
}

func (p *Postmaster) Time(t TimeMessage) (*TimeResponse, error) {
	params := MapStruct(t)
	res := TimeResponse{}
	_, err := p.Post("v1", "times", params, &res)
	return &res, err
}
