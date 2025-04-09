package redirect

// Needs for redirect - Status code profile, percentages of traffic by status code
// New Architecture
// Client configures thresholds, in percentages of traffic, for each status code.
// Application continues to accumulate profile and compare periodically against threshold
// to determine success or failure.
// So we have a threshold vs actual metrics comparison.
// No cloud database interaction
// The client will be responsible for creating the appropriate thresholds based on peak vs off-peak
// traffic

// Q. How to handle if 1 of 3 instances fail??
// A. There needs to be information in the collective to consolidate this, and also configuration
//    for the redirect. Also need to allow client overrides for success or failure.

// StatusCodeProfile - counts of status codes, can also be used for configuring client thresholds.
// Same struct used for add, get, configured percentage threshold
// TODO : locally convert a raw count into a rate/percentage
type StatusCodeProfile struct {
	Count     int `json:"count"`
	Status2xx int `json:"status2xx"`
	Status4xx int `json:"status4xx"`
	Status5xx int `json:"status5xx"`
}

// TODO : Do we need to differentiate peak vs off-peak No
// Analysis is done by comparing local metrics vs global service metrics
// Need a configuration to set the threshold for comparison
// If no distinction between peak and off peak, can the client change the thresholds?
