APP_NAME=pdfdecrypt
BIN_DIR=/usr/local/bin
GO=go

PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

.PHONY: all build install clean cross

all: build

build:
	$(GO) build -o $(APP_NAME) main.go

install: build
	@echo "🔧 Installation of $(APP_NAME) in $(BIN_DIR)"
	@mv $(APP_NAME) $(BIN_DIR)/$(APP_NAME)
	@echo "✅ Install end : run '$(APP_NAME)' "

clean:
	@rm -f $(APP_NAME)
	@rm -rf dist
	@echo "🧹 Clean end"

cross:
	@mkdir -p dist
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*}; \
		GOARCH=$${platform#*/}; \
		out=dist/$(APP_NAME)-$$GOOS-$$GOARCH; \
		if [ $$GOOS = "windows" ]; then out=$$out.exe; fi; \
		echo "🚀 Build for $$platform → $$out"; \
		GOOS=$$GOOS GOARCH=$$GOARCH $(GO) build -o $$out main.go || exit 1; \
	done
	@echo "✅ Cross-compile end (binaires dans ./dist)"
