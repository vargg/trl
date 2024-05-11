BINARY_NAME=trl

.PHONY: install build uninstall help

.DEFAULT_GOAL := help

install: build
	mkdir -p ~/.trl
	mkdir -p ~/.local/bin
	cp ./src/conf/settings.yaml ~/.trl/
	cp ./bin/$(BINARY_NAME) ~/.local/bin/

build:
	go build -C ./src -o ../bin/$(BINARY_NAME)

clean:
	rm -rf ./bin

uninstall:
	rm -f ~/.local/bin/$(BINARY_NAME)
	rm -f ~/.trl/settings.yaml

help:
	@echo "Available commands:"
	@echo "  build     - Compile the project and generate binary."
	@echo "  install   - Install the binary and configuration files."
	@echo "  clean     - Remove any build artifacts."
	@echo "  uninstall - Remove the installed binary and configuration files."
	@echo "  help      - Display this help message."
