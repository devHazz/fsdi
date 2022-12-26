package model

type ControllerPositionUpdate struct {
	Callsign     string
	Rating       int
	Lat          string
	Long         string
	Altitude     string
	Frequency    string
	FacilityType int
	VisualRange  int
}

type AddControllerRequest struct {
	Callsign  string
	FullName  string
	NetworkId int
	//Password       string
	AtcRating       int
	ProtocolVersion int
	//PositionUpdate ControllerPositionUpdate
	//ClientRequests []ClientRequest
}
