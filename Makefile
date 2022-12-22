NAME=hget
COMMIT = $$(git describe --always)

build:
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/${NAME}

clean:
	rm -f bin/*

arch:
	GOOS=darwin  GOARCH=amd64 go build -o bin/${NAME}-darwin-amd64
	GOOS=darwin  GOARCH=arm64 go build -o bin/${NAME}-darwin-arm64
	GOOS=freebsd GOARCH=amd64 go build -o bin/${NAME}-freebsd-amd64
	GOOS=linux   GOARCH=amd64 go build -o bin/${NAME}linux-amd64
	GOOS=netbsd  GOARCH=amd64 go build -o bin/${NAME}netbsd-amd64
	GOOS=openbsd GOARCH=amd64 go build -o bin/${NAME}openbsd-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/${NAME}windows-amd64.exe