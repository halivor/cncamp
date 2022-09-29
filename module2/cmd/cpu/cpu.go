package cpu

import (
	"github.com/halivor/cncamp/module2/internal/cpu"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "cpu",
		Run: func(cmd *cobra.Command, args []string) {
			go cpu.Run(1)
			cpu.Run(2)
		},
	}
}
