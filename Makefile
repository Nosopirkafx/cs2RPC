BINARY := bin/faceit-rpc.exe
PORT ?= 42157

.PHONY: build pack run clean tidy

build:
	cd backend && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ../$(BINARY) .

pack: build
	upx --best --lzma $(BINARY)

run:
	cd backend && go run .

tidy:
	cd backend && go mod tidy

clean:
	rm -rf bin
