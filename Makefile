GOCMD=go
GOBUILD=$(GOCMD) build
PATH := "${CURDIR}/bin:$(PATH)"

.PHONY: bindown

bin/golangci-lint: bindown
	script/bindown install $(notdir $@)

bin/shellcheck: bindown
	script/bindown install $(notdir $@)

bin/octo: bindown
	script/bindown install $(notdir $@)
