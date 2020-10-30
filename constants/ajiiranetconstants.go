package constants

import "os"

// AjiiraNetCorrelationID : To store correlation id in request context and response header and for logging
const AjiiraNetCorrelationID = "X-Correlation-Id"

// AjiiraCorrelationID : To store correlation id in request context and response header and for logging
const AjiiraCorrelationID = "Correlation-ID"

// CRUD :
const (
	CREATE = "CREATE"
	MODIFY = "MODIFY"
	FETCH  = "FETCH"
)

// URL :
var (
	AjiraNetServiceProtocol = os.Getenv("AJIIRA_NET_SERVICE_PROTOCOL")
	AjiraNetServiceHostName = os.Getenv("AJIIRA_NET_SERVICE_HOST_NAME")
	AjiraNetServicePortNo   = os.Getenv("AJIIRA_NET_SERVICE_PORT_NUMBER")
)

// DEVICES :
const (
	COMPUTER = "COMPUTER"
	REPEATER = "REPEATER"
)

// Messages And Error :
const (
	MsgDeviceAddSuccess             = "msg.device.add.success"
	MsgDeviceConnectionSuccess      = "msg.device.connection.success"
	MsgStrengthUpdateSuccess        = "msg.strength.update.success"
	MsgRoutes                       = "msg.routes"
	ErrorInvalidMessage             = "error.invalid.message"
	ErrorInvalidCommandSyntax       = "error.invalid.command.syntax"
	ErrorInvalidRequest             = "error.invalid.request"
	ErrorInvalidRequestBody         = "error.invalid.request.body"
	ErrorNotValidDeviceType         = "error.not.valid.device.type"
	ErrorNotValidNode               = "error.not.valid.node"
	ErrorDeviceAlreadyExist         = "error.device.already.exist"
	ErrorValueMustBeInteger         = "error.value.must.integer"
	ErrorDeviceNotFound             = "error.device.not.found"
	ErrorDeviceConnectItself        = "error.device.connection.itself"
	ErrorDeviceConnectAlready       = "error.device.connected.already"
	ErrorNodeNotFound               = "error.node.not.found"
	ErrorRouteNotFound              = "error.route.not.found"
	ErrorRouteCannotCalculateRouter = "error.route.cannot.calculate.router"
)

// AcceptedLanguages :
const (
	EnIN = "en-IN"
)
