package postmaster

// TimeResponseItem is a part of TimeResponse.
type TimeResponseItem struct {
	Service           string // Service type
	DeliveryTimestamp int    `json:"delivery_timestamp"` // Presumed delivery date timestamp
	DeliveryDesc      string `json:"delivery_desc"`      // Additional description
}

// TimeResponse is being returned by Postmaster.Time().
type TimeResponse struct {
	Services []TimeResponseItem // Delivery time for each service
}

// TimeMessage is being sent to API when calling Postmaster.Time().
type TimeMessage struct {
	FromZip    string  `json:"from_zip"` // The source zip code
	ToZip      string  `json:"to_zip"`   // The destination zip code
	Weight     float32 // The weight of the package in pounds
	Carrier    string  // Which carrier to query
	Commercial bool    // Is the package going to a commercial address?
}

// Time asks API for time to transport a shipment between two ZIP codes.
func (p *Postmaster) Time(t *TimeMessage) (*TimeResponse, error) {
	params := mapStruct(t)
	res := TimeResponse{}
	_, err := post(p, "v1", "times", params, &res)
	return &res, err
}
