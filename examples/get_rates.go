package main

import (
	"github.com/postmaster/postmaster-go"
	"fmt"
)

func main() {
	pm := postmaster.New("tt_MTpxWHoyaG43MncyRGw2UmpNcTFvTnEwZHB6Q0k")

	r := new(postmaster.RateMessage)
	r.FromZip = "28771"
	r.ToZip = "78704"
	r.Weight = 2
	resTemp, err := pm.Rate(r)
	res := resTemp.(postmaster.RateResponseBest)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Best offer is from %s\n", res.Best)
		fmt.Printf("Their price is %s\n", res.Rates[res.Best].Charge)
	}
}
