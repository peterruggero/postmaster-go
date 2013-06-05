package postmaster

import (
	"testing"
)

func TestShipmentNew(t *testing.T) {
	pm := New("apikey")
	s := pm.Shipment()
	if s.Id != -1 {
		t.Error("new shipment should have ID = -1")
	}
	if s.p != pm {
		t.Error("new shipment should have Postmaster instance initialized")
	}
}

func TestShipmentCreate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	s := pm.Shipment()
	s.Create()
	ret := <-c
	if ret.endpoint != "shipments" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	s.Id = 1
	_, err := s.Create()
	if err == nil {
		t.Error("it shouldn't be possible to create an existing shipment")
	}
}

func TestShipmentGet(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	s := pm.Shipment()
	_, err := s.Get()
	if err == nil {
		t.Error("it shouldn't be possible to get a non-existing shipment")
	}

	s.Id = 1234
	_, err = s.Get()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "shipments/1234" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestShipmentVoid(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	del = restMock(c, nil, 100, nil)

	pm := New("apikey")
	s := pm.Shipment()
	_, err := s.Void()
	if err == nil {
		t.Error("it shouldn't be possible to void a non-existing shipment")
	}

	s.Id = 1234
	_, err = s.Void()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "shipments/1234/void" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestShipmentTrack(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	s := pm.Shipment()
	_, err := s.Track()
	if err == nil {
		t.Error("it shouldn't be possible to track a non-existing shipment")
	}

	s.Id = 1234
	_, err = s.Track()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "shipments/1234/track" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestShipmentList(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	pm.ListShipments(10, "cursor", "Delivered")
	ret := <-c
	if ret.params["limit"] != "10" {
		t.Error("wrong limit in params")
	}
	if ret.params["cursor"] != "cursor" {
		t.Error("wrong cursor in params")
	}
	if ret.params["status"] != "Delivered" {
		t.Error("wrong status in params")
	}
	if ret.endpoint != "shipments" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestShipmentFind(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	_, err := pm.FindShipments("", 10, "cursor")
	if err == nil {
		t.Error("you shouldn't be able to give empty search query")
	}

	pm.FindShipments("query", 10, "cursor")
	ret := <-c
	if ret.params["limit"] != "10" {
		t.Error("wrong limit in params")
	}
	if ret.params["cursor"] != "cursor" {
		t.Error("wrong cursor in params")
	}
	if ret.params["q"] != "query" {
		t.Error("wrong query in params")
	}
	if ret.endpoint != "shipments/search" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

