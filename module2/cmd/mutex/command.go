package mutex

import (
	"github.com/halivor/cncamp/module2/internal/mutex"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "mutex",
		Run: func(cmd *cobra.Command, args []string) {
			mutex.Run()
		},
	}
}
