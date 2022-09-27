package queue

import (
	"github.com/halivor/cncamp/module1/internal/queue"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "queue",
		Run: func(cmd *cobra.Command, args []string) {
			queue.Run()
		},
	}
}
