package solver

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/options"
)

func GetDefaultServerOptions() http.ServerOptions {
	o := http.ServerOptions{
		URL:  options.GetDefaultServeOptionString("SERVER_URL", config.Conf.GetString(enums.SERVER_URL)),
		Host: options.GetDefaultServeOptionString("SERVER_HOST", config.Conf.GetString(enums.SERVER_HOST)),
		Port: options.GetDefaultServeOptionInt("SERVER_PORT", config.Conf.GetInt(enums.SERVER_PORT)),
	}
	o.URL = strings.TrimSpace(o.URL)
	o.Host = strings.TrimSpace(o.Host)
	return o
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
