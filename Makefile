# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test


# https://golang.org/cmd/link/
LDFLAGS := $(VERSION_VARIABLES) -extldflags='-static' ${GO_EXTRA_LDFLAGS}

.PHONY: clean ## Remove all build artifacts
clean: 
	rm -rf $(BUILD_DIR)

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/gowinx.exe

$(BUILD_DIR)/gowinx.exe: $(SOURCES)
	OARCH=amd64 GOOS=windows go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/gowinx.exe $(GO_EXTRA_BUILDFLAGS) ./cmd