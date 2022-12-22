package model

type ControllerPositionUpdate struct {
	callsign     string
	rating       int
	lat          string
	long         string
	altitude     string
	frequency    string
	facilityType int
	visualRange  int
}

type AddControllerRequest struct {
	callsign       string
	fullName       string
	networkId      int
	password       string
	atcRating      int
	positionUpdate ControllerPositionUpdate
	clientRequests []ClientRequest
}
