package example

import (
	"encoding/json"
	"fmt"

	"github.com/cliffrowley/go-streamdeck/streamdeck"
)

// Settings contains the settings for a context.
type Settings struct {
	Count int `json:"count"`
}

// Context encapsulates a single context.
type Context struct {
	ID       string
	Client   *streamdeck.Client
	Settings *Settings
}

func (c *Context) importSettings(data json.RawMessage) {
	s := &Settings{}
	json.Unmarshal(data, &s)
	c.Settings = s
}

func (c *Context) increment() {
	c.Settings.Count++
	j, _ := json.Marshal(c.Settings)
	c.Client.SetSettings(c.ID, json.RawMessage(j))
}

func (c *Context) updateTitle() {
	c.Client.SetTitle(c.ID, fmt.Sprintf("%v", c.Settings.Count), "both")
}

// WillAppear handles the willAppear event.
func (c *Context) WillAppear(e *streamdeck.WillAppearEvent) {
	c.importSettings(e.Payload.Settings)
	c.updateTitle()
}

// KeyUp handles the keyUp event.
func (c *Context) KeyUp(e *streamdeck.KeyUpEvent) {
	c.importSettings(e.Payload.Settings)
}

// KeyDown handles the keyDown event.
func (c *Context) KeyDown(e *streamdeck.KeyDownEvent) {
	c.importSettings(e.Payload.Settings)
	c.increment()
	c.updateTitle()
}
