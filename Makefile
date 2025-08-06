APP_NAME=ultahost-agent
BUILD_DIR=dist
ENTRY_POINT=cmd/agent/main.go

build:
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) \
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY_POINT)
	strip $(BUILD_DIR)/$(APP_NAME) || true
	upx --best --lzma $(BUILD_DIR)/$(APP_NAME) || true

size:
	du -h $(BUILD_DIR)/$(APP_NAME)

clean:
	rm -f $(BUILD_DIR)/$(APP_NAME)
