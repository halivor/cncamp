package memalloc

import (
	"time"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "memalloc",
		Run: func(cmd *cobra.Command, args []string) {
			x := [][]byte{}
			for i := 0; ; i++ {
				xx := make([]byte, 1024*1024)
				xx[1024] = 'x'
				x = append(x, xx)
				time.Sleep(time.Second * 1)
			}
		},
	}
}
