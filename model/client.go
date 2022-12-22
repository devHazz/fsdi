package model

type ClientIdentify struct {
	callsign        string
	clientId        int
	versionExpanded string
	networkId       int
	uniqueId        int
}

type ClientRequest struct {
	callsign    string
	requestType string
}
