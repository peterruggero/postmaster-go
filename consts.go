package postmaster

var VERSION float32 = 1.0

// SERVICE_LEVELS unifies different carriers' service levels, for example:
// Fedex's "Standard Overnight", UPS' "Next Day Air Saver" and USPS' "Express"
// all become our's "1DAY".
// For reference, see http://postmaster.io/docs#services
var SERVICE_LEVELS []string = []string{
	"GROUND",
	"3DAY",
	"2DAY",
	"2DAY_EARLY",
	"1DAY",
	"1DAY_EARLY",
	"1DAY_MORNING",
	"INTL_SURFACE",
	"INTL_PRIORITY",
	"INTL_EXPRESS",
}

// PACKAGE_TYPES unifies different carriers' packages terminology, for example:
// "Fedex Envelope", "UPS Letter" and "Legal Envelopes" all become "LETTER".
// For reference, see http://postmaster.io/docs#packages
var PACKAGE_TYPES []string = []string{
	"TUBE",
	"LETTER",
	"PAK",
	"CARRIER_BOX_SMALL",
	"CARRIER_BOX_MEDIUM",
	"CARRIER_BOX_LARGE",
	"CUSTOM",
}
