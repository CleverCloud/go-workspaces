BINARY=bin/myApp

build:
	echo "Build the application as ./${BINARY}"
	go build -o ${BINARY} cmd/main.go