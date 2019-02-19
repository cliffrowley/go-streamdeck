package streamdeck

import "encoding/json"

// Valid targets for some commands that require it, specifically "setTItle" and "setImage".
const (
	TargetBoth     = 0
	TargetHardware = 1
	TargetSoftware = 2
)

type openURLPayload struct {
	URL string `json:"url"`
}

type openURLCommand struct {
	Name    string          `json:"event"`
	Payload *openURLPayload `json:"payload"`
}

type sendToPropertyInspectorCommand struct {
	Name    string          `json:"event"`
	Action  string          `json:"action"`
	Context string          `json:"context"`
	Payload json.RawMessage `json:"payload"`
}

type setImagePayload struct {
	Image  string `json:"image"`
	Target string `json:"target"`
}

type setImageCommand struct {
	Name    string           `json:"event"`
	Context string           `json:"context"`
	Payload *setImagePayload `json:"payload"`
}

type setTitlePayload struct {
	Title  string `json:"title"`
	Target int    `json:"target"`
}

type setTitleCommand struct {
	Name    string           `json:"event"`
	Context string           `json:"context"`
	Payload *setTitlePayload `json:"payload"`
}

type showAlertCommand struct {
	Name    string `json:"event"`
	Context string `json:"context"`
}

type showOKCommand struct {
	Name    string `json:"event"`
	Context string `json:"context"`
}

type setSettingsCommand struct {
	Name    string          `json:"event"`
	Context string          `json:"context"`
	Payload json.RawMessage `json:"payload"`
}

// SetStatePayload contains the payload for a SetStateCommand.
type setStatePayload struct {
	State int `json:"state"`
}

type setStateCommand struct {
	Name    string           `json:"event"`
	Context string           `json:"context"`
	Payload *setStatePayload `json:"payload"`
}

type switchToProfilePayload struct {
	Profile string `json:"profile"`
}

type switchToProfileCommand struct {
	Name    string                  `json:"event"`
	Context string                  `json:"context"`
	Device  string                  `json:"device"`
	Payload *switchToProfilePayload `json:"payload"`
}
