package streamdeck

import "encoding/json"

type (
	// A SetTitleCommand corresponds to a "setTitle" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#settitle
	setTitleCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`

		Payload *setTitlePayload `json:"payload"`
	}

	setTitlePayload struct {
		Title  string `json:"title"`
		Target string `json:"target"`
	}

	// A SetImageCommand corresponds to a "setImage" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#setimage
	setImageCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`

		Payload *setImagePayload `json:"payload"`
	}

	setImagePayload struct {
		Image  string `json:"image"`
		Target string `json:"target"`
	}

	// A ShowAlertCommand corresponds to a "showAlert" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#showalert
	showAlertCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`
	}

	// A ShowOKCommand corresponds to a "showAlert" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#showok
	showOKCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`
	}

	// A SetSettingsCommand corresponds to a "setSettings" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#setsettings
	setSettingsCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`

		Payload json.RawMessage `json:"payload"`
	}

	// A SetStateCommand corresponds to a "setState" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#setstate
	setStateCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`

		Payload *setStatePayload `json:"payload"`
	}

	setStatePayload struct {
		State int `json:"state"`
	}

	// A SendToPropertyInspectorCommand corresponds to a "sendToPropertyInspector" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#sendtopropertyinspector
	sendToPropertyInspectorCommand struct {
		Name    string `json:"event"`
		Action  string `json:"action"`
		Context string `json:"context"`

		Payload json.RawMessage `json:"payload"`
	}

	// A SwitchToProfileCommand corresponds to a "switchToProfile" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#switchtoprofile
	switchToProfileCommand struct {
		Name    string `json:"event"`
		Context string `json:"context"`
		Device  string `json:"device"`

		Payload *switchToProfilePayload `json:"payload"`
	}

	switchToProfilePayload struct {
		Profile string `json:"profile"`
	}

	// A OpenURLCommand corresponds to a "openUrl" event.
	//
	// https://developer.elgato.com/documentation/stream-deck/sdk/events-sent/#switchtoprofile
	openURLCommand struct {
		Name string `json:"event"`

		Payload *openURLPayload `json:"payload"`
	}

	openURLPayload struct {
		URL string `json:"url"`
	}
)

// SendToPropertyInspector sends the given JSON data to property inspector for a context.
func (c *Client) SendToPropertyInspector(context string, action string, data json.RawMessage) error {
	return c.write(&sendToPropertyInspectorCommand{
		Name:    "sendToPropertyInspector",
		Context: context,
		Action:  action,
		Payload: data,
	})
}

// OpenURL instructs the Stream Deck software to open the specified URL in the default browser.
func (c *Client) OpenURL(url string) error {
	return c.write(&openURLCommand{
		Name: "openUrl",
		Payload: &openURLPayload{
			URL: url,
		},
	})
}

// SetImage sets the image for a context.
func (c *Client) SetImage(context string, image string, target string) error {
	return c.write(&setImageCommand{
		Name:    "setImage",
		Context: context,
		Payload: &setImagePayload{
			Image:  image,
			Target: target,
		},
	})
}

// SetSettings sets the settings for a context.
func (c *Client) SetSettings(context string, settings json.RawMessage) error {
	return c.write(&setSettingsCommand{
		Name:    "setSettings",
		Context: context,
		Payload: settings,
	})
}

// SetState sets the state for a context.
func (c *Client) SetState(context string, state int) error {
	return c.write(&setStateCommand{
		Name:    "setState",
		Context: context,
		Payload: &setStatePayload{
			State: state,
		},
	})
}

// SetTitle sets the title for a context.
func (c *Client) SetTitle(context string, title string, target string) error {
	return c.write(&setTitleCommand{
		Name:    "setTitle",
		Context: context,
		Payload: &setTitlePayload{
			Title:  title,
			Target: target,
		},
	})
}

// ShowAlert shows a temporary alert for a context.
func (c *Client) ShowAlert(context string) error {
	return c.write(&showAlertCommand{
		Name:    "showAlert",
		Context: context,
	})
}

// ShowOK shows a temporary ok checkmark for a context.
func (c *Client) ShowOK(context string) error {
	return c.write(&showOKCommand{
		Name:    "showOk",
		Context: context,
	})
}

// SwitchToProfile switches the Stream Deck to the specified read-only profile.
func (c *Client) SwitchToProfile(context string, device string, profile string) error {
	return c.write(&switchToProfileCommand{
		Name:    "switchToProfile",
		Context: context,
		Device:  device,
		Payload: &switchToProfilePayload{
			Profile: profile,
		},
	})
}
