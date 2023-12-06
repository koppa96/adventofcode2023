package main

import (
	"fmt"
	"github.com/koppa96/adventofcode2023/cmd"
	"os"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}
