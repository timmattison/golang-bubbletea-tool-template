TOOL_NAME=bubbletea-tool-template
MAC_BINARY_NAME=${TOOL_NAME}
LINUX_AMD64_BINARY_NAME=${TOOL_NAME}-linux-amd64
LINUX_ARM32_BINARY_NAME=${TOOL_NAME}-linux-arm32
LINUX_ARM64_BINARY_NAME=${TOOL_NAME}-linux-arm64
WINDOWS_BINARY_NAME=${TOOL_NAME}.exe
LDFLAGS=-ldflags="-s -w"
#WINDOWS_LDFLAGS=-ldflags="-s -w -H=windowsgui"
WINDOWS_LDFLAGS=-ldflags="-s -w"

full-build: format-and-lint quick-build
	upx -9 ${MAC_BINARY_NAME}
	upx -9 ${LINUX_AMD64_BINARY_NAME}
	upx -9 ${LINUX_ARM64_BINARY_NAME}
	upx -9 ${LINUX_ARM32_BINARY_NAME}
	upx -9 ${WINDOWS_BINARY_NAME}

format-and-lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .

quick-build:
	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o ${MAC_BINARY_NAME} .
	GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o ${LINUX_AMD64_BINARY_NAME} .
	GOARCH=arm64 GOOS=linux go build ${LDFLAGS} -o ${LINUX_ARM64_BINARY_NAME} .
	GOARCH=arm GOOS=linux go build ${LDFLAGS} -o ${LINUX_ARM32_BINARY_NAME} .
	GOARCH=amd64 GOOS=windows go build ${WINDOWS_LDFLAGS} -o ${WINDOWS_BINARY_NAME} .

debug-build:
	GOARCH=amd64 GOOS=darwin go build -o ${MAC_BINARY_NAME} .

run: debug-build
	./${MAC_BINARY_NAME}

clean:
	go clean
	rm -f ${MAC_BINARY_NAME}
	rm -f ${LINUX_AMD64_BINARY_NAME}
	rm -f ${LINUX_ARM64_BINARY_NAME}
	rm -f ${LINUX_ARM32_BINARY_NAME}
	rm -f ${WINDOWS_BINARY_NAME}
