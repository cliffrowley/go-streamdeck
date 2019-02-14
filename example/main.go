package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cliffrowley/go-streamdeck/example/example"
	"github.com/cliffrowley/go-streamdeck/streamdeck"
	"github.com/mitchellh/go-homedir"
)

var client *streamdeck.Client
var contexts = make(map[string]*example.Context, 0)

func main() {
	// Get user home directory
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	// Open up a log file
	fh, err := os.Create(filepath.Join(home, "example-log.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	// Set logging to use our log file
	log.SetOutput(fh)

	// Create the client
	c, err := streamdeck.New()
	if err != nil {
		log.Fatalln(err)
	}
	client = c

	// Register willApepar handler
	client.OnWillAppear(func(e *streamdeck.WillAppearEvent) {
		if _, ok := contexts[e.Context]; !ok {
			contexts[e.Context] = &example.Context{ID: e.Context, Client: client}
		}
		contexts[e.Context].WillAppear(e)
	})

	// Register keyUp handler
	client.OnKeyUp(func(e *streamdeck.KeyUpEvent) {
		if ctx, ok := contexts[e.Context]; ok {
			ctx.KeyUp(e)
		}
	})

	// Register keyDown handler
	client.OnKeyDown(func(e *streamdeck.KeyDownEvent) {
		if ctx, ok := contexts[e.Context]; ok {
			ctx.KeyDown(e)
		}
	})

	// Run the client
	err = client.Run()
	if err != nil {
		log.Fatalf("Error running client: %v", err)
	}
}
