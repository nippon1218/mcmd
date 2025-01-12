package main

import (
	"fmt"
	"mcmd/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	cmd.RegisterCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
