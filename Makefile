NAME=osc-bridge
OUTDIR=build
MAIN=cmd/cli/main.go
VERSION=$(shell git describe --tags --always)
FLAGS+=-trimpath
FLAGS+=-tags timetzdata
FLAGS+=-ldflags "-s -w -X main.AppVersion=$(VERSION)"
export CGO_ENABLED=0

all: windows-amd64

generate:
	go generate ./...

windows-amd64: generate
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o $(OUTDIR)/$(NAME)-$@.exe $(MAIN)

sha256sum:
	rm -r $(OUTDIR)/*.sha256
	cd $(OUTDIR); for file in *; do sha256sum $$file > $$file.sha256; done