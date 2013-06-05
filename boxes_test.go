package postmaster

import (
	"testing"
)

func TestBoxNew(t *testing.T) {
	pm := New("apikey")
	b := pm.Box()
	if b.Id != -1 {
		t.Error("new box should have ID = -1")
	}
	if b.p != pm {
		t.Error("new box should have Postmaster instance initialized")
	}
}

func TestBoxCreate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	b := pm.Box()
	b.Create()
	ret := <-c
	if ret.endpoint != "packages" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	b.Id = 1
	_, err := b.Create()
	if err == nil {
		t.Error("it shouldn't be possible to create an existing box")
	}
}

func TestBoxGet(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	b := pm.Box()
	_, err := b.Get()
	if err == nil {
		t.Error("it shouldn't be possible to get a non-existing box")
	}

	b.Id = 1234
	_, err = b.Get()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "packages/1234" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestBoxDelete(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	del = restMock(c, nil, 100, nil)

	pm := New("apikey")
	b := pm.Box()
	_, err := b.Delete()
	if err == nil {
		t.Error("it shouldn't be possible to delete a non-existing box")
	}

	b.Id = 1234
	_, err = b.Delete()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "packages/1234" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestBoxUpdate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	put = restMock(c, nil, 100, nil)

	pm := New("apikey")
	b := pm.Box()
	_, err := b.Update()
	if err == nil {
		t.Error("it shouldn't be possible to update a non-existing box")
	}

	b.Id = 1234
	_, err = b.Update()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "packages/1234" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestBoxList(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	get = restMock(c, nil, 100, nil)

	pm := New("apikey")
	pm.ListBoxes(10, "cursor")
	ret := <-c
	if ret.params["limit"] != "10" {
		t.Error("wrong limit in params")
	}
	if ret.params["cursor"] != "cursor" {
		t.Error("wrong cursor in params")
	}
	if ret.endpoint != "packages" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestFit(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	postJson = restMockJson(c, nil, 100, nil)

	pm := New("apikey")
	boxes := []Box{Box{}, Box{}}
	items := []Item{Item{}}
	pm.Fit(boxes, items, 10)
	ret := <-c
	params := ret.paramsJson.(FitMessage)
	if len(params.Boxes) != 2 {
		t.Error("wrong boxes count")
	}
	if len(params.Items) != 1 {
		t.Error("wrong items count")
	}
	if params.PackageLimit != 10 {
		t.Error("wrong package limit")
	}
	if ret.endpoint != "packages/fit" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}
