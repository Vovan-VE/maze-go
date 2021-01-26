GO ?= go

BIN_DIR = bin
BIN = maze

.PHONY: all
all: $(BIN_DIR)/$(BIN)

$(BIN_DIR)/$(BIN): $(BIN_DIR)/%: $(BIN_DIR)/
	go build -o $(BIN_DIR)/$* cmd/$*/main.go

$(BIN_DIR)/:
	mkdir -p $(BIN_DIR)

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	rm -f $(addprefix $(BIN_DIR)/,$(BIN))
	rm -fd $(BIN_DIR)
