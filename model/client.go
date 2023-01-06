package model

type ClientIdentify struct {
	Callsign        string
	Unk1            string
	ClientId        int
	VersionExpanded string
	Unk2            int
	Unk3            int
	NetworkId       int
	UniqueId        int
}

type ClientRequest struct {
	Requester   string
	Requestee   string
	RequestType string
}

type ClientResponse struct {
}

type DeleteClientRequest struct {
	Callsign  string
	NetworkId int
}
