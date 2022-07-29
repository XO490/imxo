.PHONY: all build clean
.SILENT:

filename = "imxo"

build-server:
#	 Linux
	GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./bin/$(filename)-server-l64 cmd/server/main.go
	GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./bin/$(filename)-server-l32 cmd/server/main.go

#	Windows
	GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H windows" -o ./bin/$(filename)-server-w64.exe cmd/server/main.go

build-client:
#	 Linux
	GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./bin/$(filename)-client-l64 cmd/client/main.go
	GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./bin/$(filename)-client-l32 cmd/client/main.go

#	Windows
	GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H windows" -o ./bin/$(filename)-client-w64.exe cmd/client/main.go

build: build-server build-client

#run: build
#	./bin/$(filename)__linux-x86_64
