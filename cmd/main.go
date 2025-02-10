package main

import "github.com/jaleoncordero/FileTransfer/internal/transfer"

// TODO: create settings struct & use that instead of env vars
func main() {

	// TODO: implement cmd type handling to work with multiple plugins
	transfer.Run()
}
