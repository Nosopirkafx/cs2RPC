BINARY := bin/faceit-rpc.exe

.PHONY: build xpi dist run clean tidy

build:
	cd backend && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ../$(BINARY) .

xpi:
	cmd /c install\pack_firefox.bat

dist:
	cmd /c install\pack_dist.bat

run:
	cd backend && go run .

tidy:
	cd backend && go mod tidy

clean:
	rm -rf bin dist
