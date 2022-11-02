package main

import (
	"github.com/halivor/cncamp/module3/cmd/busyloop"
	"github.com/halivor/cncamp/module3/cmd/memalloc"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "exec",
	}
	rootCmd.AddCommand(busyloop.NewCommand(), memalloc.NewCommand())
	rootCmd.Execute()
}
