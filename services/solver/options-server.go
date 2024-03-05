package solver

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/utils"
)

func GetDefaultServerOptions() http.ServerOptions {
	o := http.ServerOptions{
		URL:  config.Conf.GetString(enums.SERVER_URL),
		Host: config.Conf.GetString(enums.SERVER_HOST),
		Port: config.Conf.GetInt(enums.SERVER_PORT),
	}
	o.URL = strings.TrimSpace(o.URL)
	o.Host = strings.TrimSpace(o.Host)
	return o
}

func AddServerCliFlags(cmd *cobra.Command, serverOptions *http.ServerOptions) {
	// TODO: change server-port to port
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
	if options.Host == "" {
		return fmt.Errorf("SERVER_HOST is required")
	}

	u := options.URL

	if u == "" || strings.Contains(u, "curl") {
		ip := utils.GetPublicIP()
		if ip != "" {
			u = fmt.Sprintf("http://%s:%d", ip, options.Port)
			log.Info().Msgf("setting %s to public ip:%s", enums.SERVER_URL, u)
			options.URL = u
		}
	}

	if options.URL == "" {
		return fmt.Errorf("SERVER_URL is required")
	}
	return nil
}
