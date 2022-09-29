package server

import (
	"net/http"

	"github.com/halivor/cncamp/module2/internal/server"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.NewRouter()
			http.ListenAndServe(":8080", nil)
		},
	}
}
