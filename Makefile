NAME=hget
COMMIT = $$(git describe --always)
LDFLAGS = -ldflags "-X main.GitCommit=\"$(COMMIT)\""

build:
	go build ${LDFLAGS} -o bin/${NAME}

clean:
	rm -f bin/*

arch:
	GOOS=darwin  GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-darwin-amd64
	GOOS=darwin  GOARCH=arm64 go build ${LDFLAGS} -o bin/${NAME}-darwin-arm64
	GOOS=freebsd GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-freebsd-amd64
	GOOS=linux   GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-linux-amd64
	GOOS=netbsd  GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-netbsd-amd64
	GOOS=openbsd GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-openbsd-amd64
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${NAME}-windows-amd64.exe

pack: arch
	@for f in $$(ls bin --ignore='*.exe'); do tar cvJf "bin/$$f.tar.xz" "bin/$$f" && rm -f "bin/$$f"; done
	@for f in $$(ls bin/*.exe); do zip -9 "bin/$$(basename "$$f" ".exe").zip" "$$f" && rm -f "$$f"; done
