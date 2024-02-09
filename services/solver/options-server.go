package solver

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/options"
)

func GetDefaultServerOptions() http.ServerOptions {
	return http.ServerOptions{
		// TODO: move to appConfig
		URL:  options.GetDefaultServeOptionString("SERVER_URL", ""),
		Host: options.GetDefaultServeOptionString("SERVER_HOST", "0.0.0.0"),
		Port: options.GetDefaultServeOptionInt("SERVER_PORT", 8080), //nolint:gomnd
	}
}

func AddServerCliFlags(cmd *cobra.Command, serverOptions *http.ServerOptions) {
	cmd.PersistentFlags().StringVar(
		&serverOptions.URL, "server-url", serverOptions.URL,
		`The URL the api server is listening on (SERVER_URL).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.Host, "server-host", serverOptions.Host,
		`The host to bind the api server to (SERVER_HOST).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.Port, "server-port", serverOptions.Port,
		`The port to bind the api server to (SERVER_PORT).`,
	)
}

func CheckServerOptions(options http.ServerOptions) error {
	if options.URL == "" {
		return fmt.Errorf("SERVER_URL is required")
	}
	return nil
}
