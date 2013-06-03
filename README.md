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

### Initializing library

	key := "<YOUR_API_KEY>"
	var pm *Postmaster = postmaster.New(key)

If case you'd want to change API's base URL:

	pm.BaseUrl = "http://some.url.com"

### Errors

Every function returns base object (which usually is some structure) and an error variable (of type `error`). If everything is OK, error will be `nil`. If something goes wrong, API's error message will be stored in error variable.

	res, err := pm.Rate(rateMsg)
	if err != nil {
		// Something is wrong!
	} else {
		// Everything is OK
	}

## Usage

### Rates ([documentation](https://www.postmaster.io/docs#rates))

Request object:

	type RateMessage struct {
		FromZip    string  // The source zip code (required)
		ToZip      string  // The destination zip code
		Weight     float32 // The weight of the package in pounds
		Carrier    string  // Which carrier to query
		Packaging  string  // What type of packaging this shipment will use (optional, default: CUSTOM)
		Commercial bool    // Is the package going to a commercial address?
		Service    string  // Which service level to quote (default: GROUND)
	}

There might be two possible responses, depending on whether Carrier was provided:

	type RateResponse struct { // Carrier was specified
		Service        string // Type of service
		Charge         string // Cost of sending the shipment
		Currency       string // Currency
	}

	type RateResponseBest struct { // Carrier was not specified
		Fedex RateResponse // Rate for Fedex
		UPS   RateResponse // Rate for UPS
		USPS  RateResponse // Rate for USPS
		Best  string       // Lowercase carrier name that offers the best deal
	}

### Times ([documentation](https://www.postmaster.io/docs#get_time))

Request object:

	type TimeMessage struct {
		FromZip    string  // The source zip code
		ToZip      string  // The destination zip code
		Weight     float32 // The weight of the package in pounds
		Carrier    string  // Which carrier to query
		Commercial bool    // Is the package going to a commercial address?
	}

Response object:

	type TimeResponse struct {
		Services []TimeResponseItem // Delivery time for each service
	}

`TimeResponseItem` structure:

	type TimeResponseItem struct {
		Service           string // Service type
		DeliveryTimestamp int    // Presumed delivery date timestamp
		DeliveryDesc      string // Additional description
	}
