# postmaster-go

**postmaster-go** is Postmaster.io API library for Go language.


## Postmaster - simple API for shipping packages
Postmaster takes the pain out of sending shipments via UPS, Fedex, and USPS. Save money before you ship, while you ship, and after you ship.

For more information, see [postmaster.io](https://www.postmaster.io/).


## Requirements
- Go 1.1 or above
- [restclient](https://github.com/jmcvetta/restclient)


## Installation

    go get github.com/postmaster/postmaster-go


## Basic usage

### Getting started

First and foremost, you must import the library in your `import` instruction:

	import (
		...
		"github.com/postmaster/postmaster-go"
	)

Then, you must initialize new instance.

	key := "<YOUR_API_KEY>"
	var pm *postmaster.Postmaster = postmaster.New(key)

We assume that `pm` is your initialized Postmaster object.

If case you'd want to change API's base URL:

	pm.SetBaseUrl("http://some.url.com")


### Errors

Every function returns base object (which usually is some structure) and an error variable (of type `error`). If everything is OK, error will be `nil`. If something goes wrong, API's error message will be stored in error variable.

	res, err := pm.Rate(rateMsg)
	if err != nil {
		// Something is wrong!
	} else {
		// Everything is OK
	}


### Testing

You can run tests by executing `go test` command.


## Usage

Refer to Postmaster.io API documentation to see which fields are required.  
For more detailed information about structures being used, see [autogenerated documentation](http://godoc.org/github.com/postmaster/postmaster-go).


### Shipments


#### Base usage

Don't use `new(postmaster.Shipment)`, use `ship := pm.Shipment()` instead. This creates new `Shipment` structure and sets all necessary fields.

`Create()` and `Get()` return the `Shipment` object itself. `Create()`, `Get()` and `Void()` modify object itself, so you may ommit the first returned variable, e.g.:

	_, err = ship.Create()

**Note**: `Track()` returns `TrackingResponse`, so be sure to assign it to a variable!


#### Create ([documentation](https://www.postmaster.io/docs#create))

	ship := pm.Shipment()
	// Fill ship
	ship, err := ship.Create()

**Note**: you can't create an existing shipment (i.e. the one with ID > -1).  
**Note 2**: in case of successful creation, shipment's ID field will be modified.


#### Get

	ship := pm.Shipment()
	// Set ship ID
	ship, err := ship.Get()

**Note**: you can't get the shipment unless it has ID > -1.  


#### List shipments

	ships, err := pm.ListShipments(10, "", "Delivered")


#### Find shipments

	ships, err := pm.FindShipments("texas", 10, "")

**Note**: you must provide search query.


#### Void ([documentation](https://www.postmaster.io/docs#cancel))

	ship := pm.Shipment()
	// Set ship ID
	success, err := ship.Void()

**Note**: you can't void the shipment unless it has ID > -1.  
**Note 2**: `success` variable is of type `bool`.


#### Track ([documentation](https://www.postmaster.io/docs#track))

	ship := pm.Shipment()
	// Set ship ID, get shipment etc.
	res, err := ship.Track()

Response object: `TrackingResponse`.

**Note**: you can't void the shipment unless it has ID > -1.  
**Note 2**: in case you need to track a shipment that was not created by Postmaster, check Tracking by Reference below.


### Tracking by Reference ([documentation](https://www.postmaster.io/docs#track_ref))

Request object: string containing tracking number.

Usage:

	res, err := pm.TrackRef("1Z1896X70305267337")

Response object: `TrackingResponse`.

**Note**: in case you need to track a shipment that was created by Postmaster, check Shipment Track above.


### Monitoring external shipments ([documentation](https://www.postmaster.io/docs#track_mon))

Don't use `new(postmaster.TrackingExternal)`, use `ship := pm.TrackingExternal()` instead. This creates new `TrackingExternal` structure and sets all necessary fields.

	tr := pm.TrackingExternal()
	tr.Tracking = "1Z1896X70305267337"
	tr.Url = "http://your-website.com/webhook"
	tr.Events = []string{"Delivered", "Exception"}
	tr.Put()

Response object: `boolean` indicating whether operation succeeded.

### Boxes ([documentation](https://www.postmaster.io/docs#createbox))

#### Basic usage

**Note**: API documentation uses "package" and "box" alternately. In our case, "fitting" package is different than "shipment" package, so it's named "Box" instead.

Don't use `new(postmaster.Box)`, use `ship := pm.Box()` instead. This creates new `Box` structure and sets all necessary fields.

`Create()`, `Get()`, `Delete()` and `Update()` must be used with \*Box receiver, and they return (and modify) the `Box` object itself. `ListBoxes()` and `Fit()` must be used with \*Postmaster receiver.


#### Create

	b := pm.Box()
	// Fill box
	b, err := b.Create()

**Note**: you can't create an existing box (i.e. the one with ID > -1).  
**Note 2**: in case of successful creation, box's ID field will be modified.


#### Get

	b := pm.Box()
	// Set box's ID
	b, err := b.Get()

**Note**: you can't get the box unless it has ID > -1.  


#### Update

	b := pm.Box()
	// Set box's ID
	b, err := b.Update()

**Note**: you can't update the box unless it has ID > -1.  


#### Delete

	b := pm.Box()
	// Set box's ID
	b, err := b.Delete()

**Note**: you can't update the box unless it has ID > -1.  
**Note 2**: `Delete()` replaces existing `Box` object with a new one.


#### ListBoxes ([documentation](https://www.postmaster.io/docs#listbox))

Request objects:

- `limit int`: how many boxes to fetch
- `cursor string`: in case you need some pagination

Usage:

	res, err := pm.ListBoxes(10, "some_cursor")

Response: `BoxList` object, which contains `Cursor`, `PreviousCursor` and an array of `Box` (`[]Boxes`).


#### Fit ([documentation](https://www.postmaster.io/docs#rates))

Request objects:

- `boxes []Box`
- `items []Item`
- `limit int`

Usage:

	res, err := pm.Fit(boxes, items, limit)

Response object: `FitResponse` with array of boxes (each entry containing box and items that were put inside), leftovers (items that couldn't be fit) and a boolean `AllFit` field informing whether fitting operation has been successful.

### Shipment Rates ([documentation](https://www.postmaster.io/docs#fitbox))

Request object: `RateMessage`.

Usage:

	rateMsg := new(postmaster.RateMessage)
	// Fill rateMsg
	res, err := pm.Rate(rateMsg)

There might be two possible responses, depending on whether `Carrier` was provided in `RateMessage`:

- `Carrier` was provided: `RateResponse`,
- `Carrier` was not provided: `RateResponseBest`, containing `map[string]RateResponse` for each carrier.


### Shipment Times ([documentation](https://www.postmaster.io/docs#get_time))

Request object: `TimeMessage`.

Usage:

	timeMsg := new(postmaster.TimeMessage)
	// Fill timeMsg
	res, err := pm.Time(timeMsg)

Response object: `TimeResponse` containing an array of `TimeResponseItem`.


### Validating Addresses ([documentation](https://www.postmaster.io/docs#validate))

Request object: `Address`.

Usage:

	addr := new(postmaster.Address)
	// Fill addr
	res, err := pm.Validate(addr)

Response object: `AddressResponse` (containing `Status` string [which should be "OK" in case everything is, well, OK] and `Addresses` array of `Address` objects).