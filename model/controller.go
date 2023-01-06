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
	//Password        string `local:"true"`
	AtcRating       int
	ProtocolVersion int
	//TODO: Add remaining data for VATSIM Connection with "local" struct tag
	//PositionUpdate ControllerPositionUpdate
	//ClientRequests []ClientRequest
}
