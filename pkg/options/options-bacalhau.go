package options

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/pkg/executor/bacalhau"
)

func GetDefaultBacalhauOptions() bacalhau.BacalhauExecutorOptions {
	return bacalhau.BacalhauExecutorOptions{
		ApiHost: config.Conf.GetString("BACALHAU_API_HOST"),
	}
}

func AddBacalhauCliFlags(cmd *cobra.Command, bacalhauOptions *bacalhau.BacalhauExecutorOptions) {
	cmd.PersistentFlags().StringVar(
		&bacalhauOptions.ApiHost, "bacalhau-api-host", bacalhauOptions.ApiHost,
		`The api hostname for the bacalhau cluster to run jobs`,
	)
}

func CheckBacalhauOptions(options bacalhau.BacalhauExecutorOptions) error {
	if options.ApiHost == "" {
		return fmt.Errorf("No bacalhau service specified - please use BACALHAU_API_HOST or --bacalhau-api-host")
	}
	return nil
}
