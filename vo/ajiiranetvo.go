package vo

// AJIIRADevice :
type AJIIRADevice struct {
	Name     string                `json:"name,omitempty"`
	Type     string                `json:"type,omitempty"`
	Strength *AJIIRADeviceStrength `json:"value,omitempty"`
}

// AJIIRADeviceVO :
type AJIIRADeviceVO struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// AJIRADevices :
type AJIRADevices struct {
	Devices []*AJIIRADeviceVO `json:"devices,omitempty"`
}

// AJIIRADeviceStrength :
type AJIIRADeviceStrength struct {
	Value int `json:"value,omitempty"`
}

// AJIIRADeviceConnection :
type AJIIRADeviceConnection struct {
	Source string   `json:"source,omitempty"`
	Target []string `json:"targets,omitempty"`
}

// DeviceStrength :
type DeviceStrength struct {
	Value string `json:"value,omitempty"`
}

// AJRequestBodyVO : Text to JSON
type AJRequestBodyVO struct {
	RequsetType      string `json:"requsetType,omitempty"`
	Path             string `json:"path,omitempty"`
	ContentType      string `json:"contentType,omitempty"`
	RequestBody      []byte `json:"requestBody,omitempty"`
	AcceptedLanguage string `json:"acceptedLanguage,omitempty"`
}

// AJIRASuccessMsg :
type AJIRASuccessMsg struct {
	Message    string `json:"msg,omitempty"`
	HTTPStatus int    `json:"-"`
}

// AJIRAErrorMsg :
type AJIRAErrorMsg struct {
	Error      string `json:"error,omitempty"`
	HTTPStatus int    `json:"-"`
}

// ErrorVO :
type ErrorVO struct {
	Criticality  string           `json:",omitempty"`
	Message      string           `json:",omitempty"`
	Description  string           `json:",omitempty"`
	ErrorDetails []*ErrorDetailVO `json:",omitempty"`
	DebugID      string           `json:",omitempty"`
}

// ErrorDetailVO :
type ErrorDetailVO struct {
	Field    string
	Value    interface{}
	Location string
	Issue    string
}

// ErrorJSON : ErrorJSON revamped struct
type ErrorJSON struct {
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}
