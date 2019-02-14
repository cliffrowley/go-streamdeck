package streamdeck

import (
	"encoding/json"
	"fmt"
)

// An ApplicationDidLaunchEvent corresponds to an "applicationDidLaunch" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#applicationdidlaunch
type ApplicationDidLaunchEvent struct {
	Name string `json:"event"`

	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

// OnApplicationDidLaunch registers a handler for the applicationDidLaunch event.
func (c *Client) OnApplicationDidLaunch(fn func(*ApplicationDidLaunchEvent)) {
	c.handlers["applicationDidLaunch"] = func(data []byte) error {
		e := ApplicationDidLaunchEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// An ApplicationDidTerminateEvent corresponds to an "applicationDidTerminate" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#applicationdidterminate
type ApplicationDidTerminateEvent struct {
	Name string `json:"event"`

	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

// OnApplicationDidTerminate registers a handler for the applicationDidTerminate event.
func (c *Client) OnApplicationDidTerminate(fn func(*ApplicationDidTerminateEvent)) {
	c.handlers["applicationDidTerminate"] = func(data []byte) error {
		e := ApplicationDidTerminateEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A DeviceDidConnectEvent corresponds to a "deviceDidConnect" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#devicedidconnect
type DeviceDidConnectEvent struct {
	Name   string `json:"event"`
	Device string `json:"device"`

	Payload struct {
		Device     string `json:"device"`
		DeviceInfo struct {
			Type int `json:"type"`
			Size struct {
				Rows    int `json:"rows"`
				Columns int `json:"columns"`
			} `json:"size"`
		} `json:"deviceInfo"`
	} `json:"payload"`
}

// OnDeviceDidConnect registers a handler for the deviceDidConnect event.
func (c *Client) OnDeviceDidConnect(fn func(*DeviceDidConnectEvent)) {
	c.handlers["deviceDidConnect"] = func(data []byte) error {
		e := DeviceDidConnectEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A DeviceDidDisconnectEvent corresponds to a "deviceDidDisconnect" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#devicediddisconnect
type DeviceDidDisconnectEvent struct {
	Name string `json:"event"`

	Payload struct {
		Settings json.RawMessage `json:"settings"`
		Device   string          `json:"device"`
	} `json:"payload"`
}

// OnDeviceDidDisconnect registers a handler for the deviceDidDisconnect event.
func (c *Client) OnDeviceDidDisconnect(fn func(*DeviceDidDisconnectEvent)) {
	c.handlers["deviceDidDisconnect"] = func(data []byte) error {
		e := DeviceDidDisconnectEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A KeyDownEvent corresponds to "keyDown" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#keydown
type KeyDownEvent struct {
	Name    string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`

	Payload struct {
		Settings    json.RawMessage `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State            int  `json:"state"`
		UserDesiredState int  `json:"userDesiredState"`
		IsInMultiAction  bool `json:"isInMultiAction"`
	} `json:"payload"`
}

// OnKeyDown registers a handler for the keyDown event.
func (c *Client) OnKeyDown(fn func(*KeyDownEvent)) {
	c.handlers["keyDown"] = func(data []byte) error {
		e := KeyDownEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A KeyUpEvent corresponds to a "keyUp" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#keyup
type KeyUpEvent struct {
	Name    string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`

	Payload struct {
		Settings    json.RawMessage `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State            int  `json:"state"`
		UserDesiredState int  `json:"userDesiredState"`
		IsInMultiAction  bool `json:"isInMultiAction"`
	} `json:"payload"`
}

// OnKeyUp registers a handler for the keyUp event.
func (c *Client) OnKeyUp(fn func(*KeyUpEvent)) {
	c.handlers["keyUp"] = func(data []byte) error {
		e := KeyUpEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A TitleParametersDidChangeEvent corresponds to a "titleParametersDidChange" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#titleparametersdidchange
type TitleParametersDidChangeEvent struct {
	Name    string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`

	Payload struct {
		Settings    json.RawMessage `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State           int  `json:"state"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

// OnTitleParametersDidChange registers a handler for the titleParametersDidChange event.
func (c *Client) OnTitleParametersDidChange(fn func(*TitleParametersDidChangeEvent)) {
	c.handlers["titleParametersDidChange"] = func(data []byte) error {
		e := TitleParametersDidChangeEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A WillAppearEvent corresponds to a "willAppear" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#willappear
type WillAppearEvent struct {
	Name    string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`

	Payload struct {
		Settings    json.RawMessage `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State           int  `json:"state"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

// OnWillAppear registers a handler for the willAppear event.
func (c *Client) OnWillAppear(fn func(*WillAppearEvent)) {
	c.handlers["willAppear"] = func(data []byte) error {
		e := WillAppearEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}

// A WillDisappearEvent corresponds to a "willDisappear" event.
//
// https://developer.elgato.com/documentation/stream-deck/sdk/events-received/#willdisappear
type WillDisappearEvent struct {
	Name    string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device"`

	Payload struct {
		Settings    json.RawMessage `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State           int  `json:"state"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

// OnWillDisappear registers a handler for the willDisappear event.
func (c *Client) OnWillDisappear(fn func(*WillDisappearEvent)) {
	c.handlers["willDisappear"] = func(data []byte) error {
		e := WillDisappearEvent{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return fmt.Errorf("unmarshalling event: %v", err)
		}
		fn(&e)
		return nil
	}
}
