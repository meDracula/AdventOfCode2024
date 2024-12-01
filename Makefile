.DEFAULT_GOAL := all

########
# Global
########
.PHONY: all requirements deps test lint build clean
all: requirements deps test lint build
	@echo "INFO: All steps completed 🚀"

requirements: go golangci-lint
	@echo "INFO: all required tools are installed"

deps: requirements
	go mod download
	go mod verify
	@echo "INFO: Dependencies are installed 📦"

test: requirements
	go mod tidy
	go test ./...
	@echo "INFO: Test are green ✔"

lint: requirements
	golangci-lint run --config .golangci.yml ./...
	@echo "INFO: Linted, well done 🦾"

clean: requirements
	go clean
	rm -rf dist/
	@echo "INFO: Clean 🧹"

##############
# Requirements
##############
.PHONY: go golangci-lint
# Install https://go.dev/doc/install
go: ; @which go > /dev/null

# Install https://golangci-lint.run/welcome/install/
golangci-lint: ; @which golangci-lint > /dev/null
