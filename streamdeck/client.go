package streamdeck

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/tidwall/gjson"

	"github.com/gorilla/websocket"
)

type handler func([]byte) error

// A Client encapsulates communication with the Stream Deck software.
type Client struct {
	Host          string
	Port          int
	UUID          string
	RegisterEvent string
	Info          string

	handlers map[string]handler
	ws       *websocket.Conn
	stop     chan bool
}

func (c *Client) write(command interface{}) error {
	err := c.ws.WriteJSON(command)
	if err != nil {
		return fmt.Errorf("marshalling command: %v", err)
	}
	return nil
}

// New returns a new client configured via the command line.
func New() (*Client, error) {
	c := &Client{}

	flag.IntVar(&c.Port, "port", 0, "the port to connect to")
	flag.StringVar(&c.UUID, "pluginUUID", "", "the plugin UUID")
	flag.StringVar(&c.RegisterEvent, "registerEvent", "", "the plugin register event")
	flag.StringVar(&c.Info, "info", "", "the plugin info")

	flag.Parse()

	if c.Port < 1 || c.Port > 65535 {
		return nil, errors.New("invalid port specified")
	}

	if c.UUID == "" {
		return nil, errors.New("invalid pluginUUID specified")
	}

	if c.RegisterEvent == "" {
		return nil, errors.New("invalid registerEvent specified")
	}

	if c.Info == "" {
		return nil, errors.New("invalid info specified")
	}

	c.handlers = make(map[string]handler, 0)

	return c, nil
}

// Run begins the event loop and does not return unless stopped or an error occurs.
func (c *Client) Run() error {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%v", c.Port)}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("connecting: %v", err)
	}
	c.ws = ws
	defer c.ws.Close()

	msg := fmt.Sprintf(`{"event":"%v", "uuid":"%v"}`, c.RegisterEvent, c.UUID)
	c.ws.WriteMessage(websocket.TextMessage, []byte(msg))

	c.stop = make(chan (bool), 0)

	go func() {
		for {
			_, data, err := c.ws.ReadMessage()
			if err != nil {
				log.Printf("error reading message: %v\n", err)
				c.stop <- true
				return
			}

			res := gjson.ParseBytes(data).Get("event")
			if res.Exists() {
				if h, ok := c.handlers[res.String()]; ok {
					h(data)
				}
			} else {
				log.Printf("unknown data received by client: %v\n", string(data))
			}
		}
	}()

	for {
		select {
		case <-c.stop:
			log.Println("stopping gracefully")
			c.ws.Close()
			return nil
		}
	}
}

// Stop terminates the event loop.
func (c *Client) Stop() {
	c.stop <- true
}
