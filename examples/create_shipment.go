package main

import (
	"github.com/postmaster/postmaster-go"
	"fmt"
)

func main() {
	pm := postmaster.New("tt_MTpxWHoyaG43MncyRGw2UmpNcTFvTnEwZHB6Q0k")

	s := pm.Shipment()
	s.To = &postmaster.Address{
		Company: "ASLS",
		Contact: "Joe Smith",
		Line1: "1110 Someplace Ave.",
		City: "Austin",
		State: "TX",
		ZipCode: "78704",
		PhoneNo: "5551234444",
	}
	s.Carrier = "ups"
	s.Service = "ground"
	s.Package = &postmaster.Package{
		Width: 5,
		Height: 5,
		Length: 5,
		Weight: 20,
	}
	_, err := s.Create()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Shipment created. Its ID: ", s.Id)
	}
}
