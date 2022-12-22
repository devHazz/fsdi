package parser

// FSD Commands
const (
	SERVER_IDENTIFY = iota + 1
	CLIENT_IDENTIFY
	ADD_ATC
	ADD_PILOT
	MOTD
	DELETE_ATC
	DELETE_PILOT
	CLIENT_REQUEST
	CLIENT_RESPONSE
	SERVER_ERROR = 100
)

// FSD Errors
const (
	ERR_OK         = iota /* No error */
	ERR_CSINUSE           /* Callsign in use */
	ERR_CSINVALID         /* Callsign invalid */
	ERR_REGISTERED        /* Already registered */
	ERR_SYNTAX            /* Syntax error */
	ERR_SRCINVALID        /* Invalid source in packet */
	ERR_CIDINVALID        /* Invalid CID/password */
	ERR_NOSUCHCS          /* No such callsign */
	ERR_NOFP              /* No flightplan */
	ERR_NOWEATHER         /* No such weather profile*/
	ERR_REVISION          /* Invalid protocol revision */
	ERR_LEVEL             /* Requested level too high */
	ERR_SERVFULL          /* No more clients */
	ERR_CSSUSPEND         /* CID/PID suspended */
	ERR_UNAUTH     = 16   /* Custom implementation for VATSIM, for software whitelisting. */
)
