package model

type ClientIdentify struct {
	Callsign        string
	ClientId        int
	VersionExpanded string
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
