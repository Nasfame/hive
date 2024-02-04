package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
	optionsfactory "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
)

func newVersionCmd(conf *viper.Viper) *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Get the CoopHive version",
		Long:    "Get the CoopHive version",
		Example: "CoopHive version",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runVersion(conf, cmd)
		},
	}

	optionsfactory.AddSolverCliFlags(versionCmd, &options)

	return versionCmd
}

func runVersion(conf *viper.Viper, cmd *cobra.Command) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	VERSION := conf.GetString(enums.VERSION)
	GO_BINARY_URL := conf.GetString(enums.RELEASE_URL)
	COMMIT := conf.GetString(enums.COMMIT_SHA)

	if VERSION == "" {
		fmt.Printf("version not found: download the latest binary from %s", GO_BINARY_URL)
		// unnecessary help shows up when returned as error, so shortciruting here
		return nil
	}

	fmt.Printf("CoopHive: %s\n", VERSION)
	fmt.Printf("Commit: %s\n", COMMIT)

	// TODO: suggest auto updating to the latest version if the current version is not the latest version

	return nil
}
