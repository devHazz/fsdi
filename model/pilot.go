package model

type AddPilotRequest struct {
	Callsign        string
	NetworkId       int
	Unk1            int
	ProtocolVersion int
	Rating          int
	FullName        string
}
