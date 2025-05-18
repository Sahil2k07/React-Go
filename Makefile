FRONTEND_DIR=client
BACKEND_DIR=server
BUILD_DIR=build
PUBLIC_DIR=$(BUILD_DIR)/public
GO_BIN=$(BUILD_DIR)/React-Go

.PHONY: all frontend backend run clean

all: frontend backend

frontend:
	cd $(FRONTEND_DIR) && bun install && bun run build
	rm -rf $(PUBLIC_DIR)
	mkdir -p $(PUBLIC_DIR)
	cp -r $(FRONTEND_DIR)/dist/* $(PUBLIC_DIR)/

backend:
	cd $(BACKEND_DIR) && go build -o ../$(GO_BIN)

run:
	cd $(BUILD_DIR) && ./React-Go

clean:
	rm -rf $(BUILD_DIR)
