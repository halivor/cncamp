package slice

import (
	"github.com/halivor/cncamp/module1/internal/slice"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "slice",
		Run: func(cmd *cobra.Command, args []string) {
			slice.Run()
		},
	}
}
