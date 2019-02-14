PLUGIN_NAME=com.github.cliffrowley.example.sdPlugin
SD_PLUGIN_DIR="$(HOME)/Library/Application Support/com.elgato.StreamDeck/Plugins"
OUTPUT_DIR=example/$(PLUGIN_NAME)
DIST_DIR=dist
BINARY_NAME=example

all: build-mac build-windows

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME) example/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME).exe example/main.go

sim:
	go run example/main.go -port 9876 -pluginUUID ABCD1234 -registerEvent foo -info "{}"

dist: build-mac build-windows
	rm -rf $(DIST_DIR)
	mkdir -p $(DIST_DIR)
	sd-distro-tool $(OUTPUT_DIR) $(DIST_DIR)

install: all
	rm -rf $(SD_PLUGIN_DIR)/$(PLUGIN_NAME)
	cp -r $(OUTPUT_DIR) $(SD_PLUGIN_DIR)
	osascript -e 'quit app "Stream Deck"'
	sleep 1
	open -a "Stream Deck"

clean:
	go clean
	rm -rf $(OUTPUT_DIR)/$(BINARY_NAME)
	rm -rf $(OUTPUT_DIR)/$(BINARY_NAME).exe
	rm -rf $(DIST_DIR)
