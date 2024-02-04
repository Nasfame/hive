package root

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
)

var Fatal = FatalErrorHandler

func newRootCmd(conf *viper.Viper, subCommands ...*cobra.Command) *cobra.Command {

	// RunCmd, SolverCmd, ResourceProviderCmd, mediatorCmd, JobCreatorCmd, VersionCmd *cobra.JobCreatorCmd

	version := conf.GetString(enums.VERSION)

	commit := conf.GetString(enums.COMMIT_SHA)

	cmd := &cobra.Command{
		Use:   getCommandLineExecutable(),
		Short: "CoopHive",
		Long:  fmt.Sprintf("CoopHive: %s \nCommit: %s \n", version, commit),
	}
	//
	// cmd.AddCommand(SolverCmd)
	// cmd.AddCommand(ResourceProviderCmd)
	// cmd.AddCommand(RunCmd)
	// cmd.AddCommand(mediatorCmd)
	// cmd.AddCommand(JobCreatorCmd)
	// cmd.AddCommand(VersionCmd)

	for _, subCmd := range subCommands {
		cmd.AddCommand(subCmd)
	}

	return cmd
}

/*
command line processing
*/
func getCommandLineExecutable() string {
	return os.Args[0]
}

func getDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func getDefaultServeOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

/*
useful tools
*/
func FatalErrorHandler(cmd *cobra.Command, msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		cmd.Print(msg)
	}
	os.Exit(code)
}