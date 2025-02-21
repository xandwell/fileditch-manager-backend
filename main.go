package main

import (
	"fmt"

	"github.com/xandwell/fileditch-manager-backend/dump"
)

func main() {
	/* TODO: Handle command execution and arguments */
	err := dump.Upload()
	if err != nil {
		fmt.Printf("Congrats")
	}
}
