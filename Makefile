MAC_BINARY_NAME=duupe-client
LINUX_BINARY_NAME=duupe-client-linux
WINDOWS_BINARY_NAME=duupe-client.exe
LDFLAGS=-ldflags="-s -w"
#WINDOWS_LDFLAGS=-ldflags="-s -w -H=windowsgui"
WINDOWS_LDFLAGS=-ldflags="-s -w"

full-build: quick-build
	upx -9 ${MAC_BINARY_NAME}
	upx -9 ${LINUX_BINARY_NAME}
	upx -9 ${WINDOWS_BINARY_NAME}
	rm -f ${WINDOWS_BINARY_NAME}.7z
	7z a ${WINDOWS_BINARY_NAME}.7z ${WINDOWS_BINARY_NAME}
	cp ${WINDOWS_BINARY_NAME} ~/Desktop

quick-build:
	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o ${MAC_BINARY_NAME} .
	GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o ${LINUX_BINARY_NAME} .
	GOARCH=amd64 GOOS=windows go build ${WINDOWS_LDFLAGS} -o ${WINDOWS_BINARY_NAME} .

debug-build:
	GOARCH=amd64 GOOS=darwin go build -o ${MAC_BINARY_NAME} .

run: debug-build
	./${MAC_BINARY_NAME}

clean:
	go clean
	rm -f ${MAC_BINARY_NAME}
	rm -f ${LINUX_BINARY_NAME}
	rm -f ${WINDOWS_BINARY_NAME}
	rm -f ${WINDOWS_BINARY_NAME}.7z
	rm -rf ~/Desktop/${WINDOWS_BINARY_NAME}
