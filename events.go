package streamdeck

import (
	"encoding/json"
)

// An ApplicationDidLaunchEvent is emitted when an application specified in the plugin manifest's
// "applicationsToMonitor" configuration has launched.
type ApplicationDidLaunchEvent struct {
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

func (e *ApplicationDidLaunchEvent) unmarshal(data []byte) (*ApplicationDidLaunchEvent, error) {
	if err := json.Unmarshal(data, e); err != nil {
		return nil, err
	}
	return e, nil
}

// An ApplicationDidTerminateEvent is emitted when an application specified in the plugin manifest's
// "applicationsToMonitor" configuration has terminated.
type ApplicationDidTerminateEvent struct {
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

func (e *ApplicationDidTerminateEvent) unmarshal(data []byte) (*ApplicationDidTerminateEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A DeviceDidConnectEvent is emitted when a Stream Deck device is plugged in to the computer.
type DeviceDidConnectEvent struct {
	Device     string `json:"device"`
	DeviceInfo struct {
		Type int `json:"type"`
		Size struct {
			Rows    int `json:"rows"`
			Columns int `json:"columns"`
		} `json:"size"`
	} `json:"deviceInfo"`
}

func (e *DeviceDidConnectEvent) unmarshal(data []byte) (*DeviceDidConnectEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A DeviceDidDisconnectEvent is emitted when a Stream Deck is unplugged from the computer.
type DeviceDidDisconnectEvent struct {
	Device string `json:"device"`
}

func (e *DeviceDidDisconnectEvent) unmarshal(data []byte) (*DeviceDidDisconnectEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A KeyDownEvent is emitted when a button on the Stream Deck is pressed that is associated with a
// context belonging to this plugin.
type KeyDownEvent struct {
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		IsInMultiAction  bool            `json:"isInMultiAction"`
		Settings         json.RawMessage `json:"settings"`
		State            int             `json:"state"`
		UserDesiredState int             `json:"userDesiredState"`
	} `json:"payload"`
}

func (e *KeyDownEvent) unmarshal(data []byte) (*KeyDownEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A KeyUpEvent is emitted when a previously pressed button on the Stream Deck is released that is
// associated with a context belonging to this plugin.
type KeyUpEvent struct {
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		IsInMultiAction  bool            `json:"isInMultiAction"`
		Settings         json.RawMessage `json:"settings"`
		State            int             `json:"state"`
		UserDesiredState int             `json:"userDesiredState"`
	} `json:"payload"`
}

func (e *KeyUpEvent) unmarshal(data []byte) (*KeyUpEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A TitleParametersDidChangeEvent is emitted when the user changes the title parameters of a
// context in the Stream Deck application.
type TitleParametersDidChangeEvent struct {
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		Settings        json.RawMessage `json:"settings"`
		State           int             `json:"state"`
		Title           string          `json:"title"`
		TitleParameters struct {
			FontFamily     string `json:"fontFamily"`
			FontSize       int    `json:"fontSize"`
			FontStyle      string `json:"fontStyle"`
			FontUnderline  bool   `json:"fontUnderline"`
			ShowTitle      bool   `json:"showTitle"`
			TitleAlignment string `json:"titleAlignment"`
			TitleColor     string `json:"titleColor"`
		} `json:"titleParameters"`
	} `json:"payload"`
}

func (e *TitleParametersDidChangeEvent) unmarshal(data []byte) (*TitleParametersDidChangeEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A WillAppearEvent is emitted when a context is about to be displayed, either when the Stream Deck
// application is started or when the user navigates to a page or profile containing the context.
type WillAppearEvent struct {
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		IsInMultiAction bool            `json:"isInMultiAction"`
		Settings        json.RawMessage `json:"settings"`
		State           int             `json:"state"`
	} `json:"payload"`
}

func (e *WillAppearEvent) unmarshal(data []byte) (*WillAppearEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}

// A WillDisappearEvent is emitted when a context is about to be hidden, usually when the user
// navigates to a another page or profile.
type WillDisappearEvent struct {
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		IsInMultiAction bool            `json:"isInMultiAction"`
		Settings        json.RawMessage `json:"settings"`
		State           int             `json:"state"`
	} `json:"payload"`
}

func (e *WillDisappearEvent) unmarshal(data []byte) (*WillDisappearEvent, error) {
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return e, nil
}
