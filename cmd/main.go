package main

import (
	compiledCheckerCC "github.com/CleverCloud/go-workspaces/checker"
	httpServerCC "github.com/CleverCloud/go-workspaces/httpServer"
)

func main() {
	compiledCheckerCC.CheckAndShow()
	httpServerCC.Launch("8080", "Hello, World... from Clever Cloud!")
}
