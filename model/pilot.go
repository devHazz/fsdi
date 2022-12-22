package model

type AddPilotRequest struct {
	callsign        string
	networkId       int
	protocolVersion int
	rating          int
	fullName        string
}
