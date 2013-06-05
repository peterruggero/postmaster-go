package postmaster

import (
	"testing"
)

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
