package busyloop

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "busyloop",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				for i := 0; ; {
					i++
				}
			}()
			for i := 0; ; {
				i++
			}
		},
	}
}
