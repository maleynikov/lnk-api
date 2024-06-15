package main

import (
	"fmt"
	"os"
)

func main() {
	srv := &Server{}
	if err := srv.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
