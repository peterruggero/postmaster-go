package postmaster

import "fmt"

func (p *Postmaster) Void(shipmentId int) (bool, error) {
	endpoint := fmt.Sprintf("shipments/%d/void", shipmentId)
	var res map[string]string
	_, err := p.Delete("v1", endpoint, nil, &res)
	return res["message"] == "OK", err
}
