# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: install help

GOBIN = ./build/bin
GO ?= latest
GORUN = go build

#? install: Build install
install:
	$(GORUN) -o ./build/bin/blockchain
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."
help: Makefile
	@echo " Choose a command run in blockchain_go:"
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'
