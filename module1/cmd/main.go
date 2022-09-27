package main

import (
	"github.com/halivor/cncamp/module1/cmd/queue"
	"github.com/halivor/cncamp/module1/cmd/slice"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "command",
	}
	rootCmd.AddCommand(slice.NewCommand(), queue.NewCommand())
	rootCmd.Execute()
}
