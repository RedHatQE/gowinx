# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test

# https://golang.org/cmd/link/
LDFLAGS := $(VERSION_VARIABLES) -extldflags='-static' ${GO_EXTRA_LDFLAGS}

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/windows-amd64/main.exe

$(BUILD_DIR)/windows-amd64/main.exe: $(SOURCES)
	GOARCH=amd64 GOOS=windows go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/windows-amd64/main.exe $(GO_EXTRA_BUILDFLAGS) .