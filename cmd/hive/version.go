package hive

import (
	"fmt"

	"github.com/spf13/cobra"

	optionsfactory "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
)

var VERSION string

var COMMIT_SHA string

const GO_BINARY_URL = "https://github.com/CoopHive/hive/releases/"

func newVersionCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Get the CoopHive version",
		Long:    "Get the CoopHive version",
		Example: "CoopHive version",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runVersion(cmd)
		},
	}

	optionsfactory.AddSolverCliFlags(versionCmd, &options)

	return versionCmd
}

func runVersion(cmd *cobra.Command) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if VERSION == "" {
		fmt.Printf("version not found: download the latest binary from %s", GO_BINARY_URL)
		// unnecessary help shows up when returned as error, so shortciruting here
		return nil
	}

	fmt.Printf("CoopHive: %s\n", VERSION)
	fmt.Printf("Commit: %s\n", COMMIT_SHA)

	// TODO: suggest auto updating to the latest version if the current version is not the latest version

	return nil
}
