package streamdeck

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// Stream Deck device types.
const (
	StreamDeck     = 0
	StreamDeckMini = 1
)

// A Device represents a Stream Deck device.
type Device struct {
	ID   string
	Type int
	Size *Size
}

// Size specifies the size of a device in rows and columns.
type Size struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

type clientInfo struct {
	Application struct {
		Language string `json:"language"`
		Platform string `json:"platform"`
		Version  string `json:"version"`
	} `json:"application"`

	Devices []struct {
		ID   string `json:"id"`
		Size struct {
			Rows    int `json:"rows"`
			Columns int `json:"columns"`
		} `json:"size"`
		Type int `json:"type"`
	} `json:"devices"`
}

type conn interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}

type config struct {
	Info          string
	Port          int
	PluginUUID    string
	RegisterEvent string
}

// A Client encapsulates communication with the Stream Deck software.
type Client struct {
	language string
	platform string
	version  string

	sendLock sync.Mutex

	devices     map[string]*Device
	devicesLock sync.Mutex

	applicationDidLaunchHandler     ApplicationDidLaunchHandler
	applicationDidTerminateHandler  ApplicationDidTerminateHandler
	deviceDidConnectHandler         DeviceDidConnectHandler
	deviceDidDisconnectHandler      DeviceDidDisconnectHandler
	keyDownHandler                  KeyDownHandler
	keyUpHandler                    KeyUpHandler
	titleParametersDidChangeHandler TitleParametersDidChangeHandler
	willAppearHandler               WillAppearHandler
	willDisappearHandler            WillDisappearHandler

	conn conn
}

// Connect returns a new Client configured via the command line.
func Connect() (*Client, error) {
	cfg := config{}
	flag.IntVar(&cfg.Port, "port", 0, "the port to connect to")
	flag.StringVar(&cfg.PluginUUID, "pluginUUID", "", "the plugin UUID")
	flag.StringVar(&cfg.RegisterEvent, "registerEvent", "", "the plugin register event")
	flag.StringVar(&cfg.Info, "info", "", "the plugin info")
	flag.Parse()

	info := &clientInfo{}
	err := json.Unmarshal([]byte(cfg.Info), &info)
	if err != nil {
		return nil, err
	}

	c := &Client{
		devices:  make(map[string]*Device, 0),
		language: info.Application.Language,
		platform: info.Application.Platform,
		version:  info.Application.Version,
	}

	for _, d := range info.Devices {
		c.addDevice(d.ID, d.Type, d.Size.Columns, d.Size.Rows)
	}

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%v", cfg.Port)}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("connecting: %v", err)
	}
	c.conn = ws

	c.sendRegisterEvent(cfg.RegisterEvent, cfg.PluginUUID)

	return c, nil
}

func (c *Client) addDevice(ID string, deviceType int, columns int, rows int) {
	c.devicesLock.Lock()
	defer c.devicesLock.Unlock()
	c.devices[ID] = &Device{ID: ID, Size: &Size{Columns: columns, Rows: rows}}
}

func (c *Client) removeDevice(ID string) {
	c.devicesLock.Lock()
	defer c.devicesLock.Unlock()
	delete(c.devices, ID)
}

func (c *Client) sendCommand(cmd interface{}) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	return c.send(data)
}

func (c *Client) send(data []byte) error {
	c.sendLock.Lock()
	defer c.sendLock.Unlock()
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) sendRegisterEvent(registerEvent string, pluginUUID string) error {
	msg := fmt.Sprintf(`{"event":"%v", "uuid":"%v"}`, registerEvent, pluginUUID)
	return c.send([]byte(msg))
}

func (c *Client) dispatch(data []byte) error {
	switch gjson.ParseBytes(data).Get("event").String() {
	case "applicationDidLaunch":
		if c.applicationDidLaunchHandler != nil {
			evt, err := (&ApplicationDidLaunchEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.applicationDidLaunchHandler.ApplicationDidLaunch(evt)
		}
	case "applicationDidTerminate":
		if c.applicationDidTerminateHandler != nil {
			evt, err := (&ApplicationDidTerminateEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.applicationDidTerminateHandler.ApplicationDidTerminate(evt)
		}
	case "deviceDidConnect":
		if c.deviceDidConnectHandler != nil {
			evt, err := (&DeviceDidConnectEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.addDevice(evt.Device, evt.DeviceInfo.Type, evt.DeviceInfo.Size.Columns, evt.DeviceInfo.Size.Rows)
			c.deviceDidConnectHandler.DeviceDidConnect(evt)
		}
	case "deviceDidDisconnect":
		if c.deviceDidDisconnectHandler != nil {
			evt, err := (&DeviceDidDisconnectEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.removeDevice(evt.Device)
			c.deviceDidDisconnectHandler.DeviceDidDisconnect(evt)
		}
	case "keyDown":
		if c.keyDownHandler != nil {
			evt, err := (&KeyDownEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.keyDownHandler.KeyDown(evt)
		}
	case "keyUp":
		if c.keyUpHandler != nil {
			evt, err := (&KeyUpEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.keyUpHandler.KeyUp(evt)
		}
	case "titleParametersDidChange":
		if c.titleParametersDidChangeHandler != nil {
			evt, err := (&TitleParametersDidChangeEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.titleParametersDidChangeHandler.TitleParametersDidChange(evt)
		}
	case "willAppear":
		if c.willAppearHandler != nil {
			evt, err := (&WillAppearEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.willAppearHandler.WillAppear(evt)
		}
	case "willDisappear":
		if c.willDisappearHandler != nil {
			evt, err := (&WillDisappearEvent{}).unmarshal(data)
			if err != nil {
				return err
			}
			c.willDisappearHandler.WillDisappear(evt)
		}
	default:
		log.Println("Unknown event received: ", string(data))
	}
	return nil
}

// GetDevice returns the device with the given id.
func (c *Client) GetDevice(id string) *Device {
	if d, ok := c.devices[id]; ok {
		return d
	}
	return nil
}

// GetLanguage returns the language as specified by the Stream Deck software.
func (c *Client) GetLanguage() string {
	return c.language
}

// GetPlatform returns the platform as specified by the Stream Deck software.
func (c *Client) GetPlatform() string {
	return c.platform
}

// GetVersion returns the version as specified by the Stream Deck software.
func (c *Client) GetVersion() string {
	return c.version
}

// HandleApplicationDidLaunch registers a handler for ApplicationDidLaunchEvents.
func (c *Client) HandleApplicationDidLaunch(h ApplicationDidLaunchHandler) {
	c.applicationDidLaunchHandler = h
}

// HandleApplicationDidLaunchFunc registers a handler func for ApplicationDidLaunchEvents.
func (c *Client) HandleApplicationDidLaunchFunc(f ApplicationDidLaunchHandlerFunc) {
	c.HandleApplicationDidLaunch(f)
}

// HandleApplicationDidTerminate registers a handler for ApplicationDidTerminateEvents.
func (c *Client) HandleApplicationDidTerminate(h ApplicationDidTerminateHandler) {
	c.applicationDidTerminateHandler = h
}

// HandleApplicationDidTerminateFunc registers a handler func for ApplicationDidTerminateEvents.
func (c *Client) HandleApplicationDidTerminateFunc(f ApplicationDidTerminateHandlerFunc) {
	c.HandleApplicationDidTerminate(f)
}

// HandleDeviceDidConnect registers a handler for DeviceDidConnectEvents.
func (c *Client) HandleDeviceDidConnect(h DeviceDidConnectHandler) {
	c.deviceDidConnectHandler = h
}

// HandleDeviceDidConnectFunc registers a handler func for DeviceDidConnectEvents.
func (c *Client) HandleDeviceDidConnectFunc(f DeviceDidConnectHandlerFunc) {
	c.HandleDeviceDidConnect(f)
}

// HandleDeviceDidDisconnect registers a handler for DeviceDidDisconnectEvents.
func (c *Client) HandleDeviceDidDisconnect(h DeviceDidDisconnectHandler) {
	c.deviceDidDisconnectHandler = h
}

// HandleDeviceDidDisconnectFunc registers a handler func for DeviceDidDisconnectEvents.
func (c *Client) HandleDeviceDidDisconnectFunc(f DeviceDidDisconnectHandlerFunc) {
	c.HandleDeviceDidDisconnectFunc(f)
}

// HandleKeyDown registers a handler for KeyDownEvents.
func (c *Client) HandleKeyDown(h KeyDownHandler) {
	c.keyDownHandler = h
}

// HandleKeyDownFunc registers a handler func for KeyDownEvents.
func (c *Client) HandleKeyDownFunc(f KeyDownHandlerFunc) {
	c.HandleKeyDown(f)
}

// HandleKeyUp registers a handler for KeyUpEvents.
func (c *Client) HandleKeyUp(h KeyUpHandler) {
	c.keyUpHandler = h
}

// HandleKeyUpFunc registers a handler func for KeyUpEvents.
func (c *Client) HandleKeyUpFunc(f KeyUpHandlerFunc) {
	c.HandleKeyUp(f)
}

// HandleTitleParametersDidChange registers a handler for TitleParametersDidChangeEvents.
func (c *Client) HandleTitleParametersDidChange(h TitleParametersDidChangeHandler) {
	c.titleParametersDidChangeHandler = h
}

// HandleTitleParametersDidChangeFunc registers a handler func for TitleParametersDidChangeEvents.
func (c *Client) HandleTitleParametersDidChangeFunc(f TitleParametersDidChangeHandlerFunc) {
	c.HandleTitleParametersDidChange(f)
}

// HandleWillAppear registers a handler for the "willAppear" event.
func (c *Client) HandleWillAppear(h WillAppearHandler) {
	c.willAppearHandler = h
}

// HandleWillAppearFunc registers a handler func for the "willAppear" event.
func (c *Client) HandleWillAppearFunc(f WillAppearHandlerFunc) {
	c.HandleWillAppear(f)
}

// HandleWillDisappear registers a handler for WillDisappearEvents.
func (c *Client) HandleWillDisappear(h WillDisappearHandler) {
	c.willDisappearHandler = h
}

// HandleWillDisappearFunc registers a handler func for WillDisappearEvents.
func (c *Client) HandleWillDisappearFunc(f WillDisappearHandlerFunc) {
	c.HandleWillDisappear(f)
}

// OpenURL instructs the Stream Deck software to open the specified URL in the default browser.
func (c *Client) OpenURL(url string) error {
	return c.sendCommand(openURLCommand{
		Name:    "openUrl",
		Payload: &openURLPayload{URL: url},
	})
}

// SendToPropertyInspector sends JSON data to the property inspector.
func (c *Client) SendToPropertyInspector(context string, action string, data json.RawMessage) error {
	return c.sendCommand(sendToPropertyInspectorCommand{
		Name:    "sendToPropertyInspector",
		Context: context,
		Action:  action,
		Payload: data,
	})
}

// SetImage sets the image for a context.
func (c *Client) SetImage(context string, image string, target string) error {
	return c.sendCommand(setImageCommand{
		Name:    "setImage",
		Context: context,
		Payload: &setImagePayload{Image: image, Target: target},
	})
}

// SetTitle sets the title for a context.
func (c *Client) SetTitle(context string, title string, target int) error {
	return c.sendCommand(setTitleCommand{
		Name:    "setTitle",
		Context: context,
		Payload: &setTitlePayload{Title: title, Target: target},
	})
}

// ShowAlert shows a temporary alert for a context.
func (c *Client) ShowAlert(context string) error {
	return c.sendCommand(showAlertCommand{
		Name:    "showAlert",
		Context: context,
	})
}

// ShowOK shows a temporary OK checkmark for a context.
func (c *Client) ShowOK(context string) error {
	return c.sendCommand(showOKCommand{
		Name:    "showOk",
		Context: context,
	})
}

// SetSettings sets the settings for a context.
func (c *Client) SetSettings(context string, settings json.RawMessage) error {
	return c.sendCommand(setSettingsCommand{
		Name:    "setSettings",
		Context: context,
		Payload: settings,
	})
}

// SetState sets the state for a context.
func (c *Client) SetState(context string, state int) error {
	return c.sendCommand(setStateCommand{
		Name:    "setState",
		Context: context,
		Payload: &setStatePayload{State: state},
	})
}

// SwitchToProfile switches the Stream Deck to the given read-only profile, specified in the
// plugin's manifest.
func (c *Client) SwitchToProfile(context string, device string, profile string) error {
	return c.sendCommand(switchToProfileCommand{
		Name:    "switchToProfile",
		Context: context,
		Device:  device,
		Payload: &switchToProfilePayload{Profile: profile},
	})
}

// Run begins the event loop and does not return unless stopped or an error occurs.
func (c *Client) Run() error {
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				return err
			}
			break
		}

		c.dispatch(data)
	}

	return nil
}

// Stop terminates the event loop.
func (c *Client) Stop() {
	c.conn.Close()
}
