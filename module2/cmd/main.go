package main

import (
	"github.com/halivor/cncamp/module2/cmd/cpu"
	"github.com/halivor/cncamp/module2/cmd/mutex"
	"github.com/halivor/cncamp/module2/cmd/queue"
	"github.com/halivor/cncamp/module2/cmd/server"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "exec",
	}
	rootCmd.AddCommand(
		mutex.NewCommand(),
		cpu.NewCommand(),
		queue.NewCommand(),
		server.NewCommand(),
	)
	rootCmd.Execute()
}
