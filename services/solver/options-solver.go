package solver

import (
	"github.com/sirupsen/logrus"

	options2 "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/services/solver/solver"

	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	options := solver.SolverOptions{
		Server: GetDefaultServerOptions(),
		Web3:   options2.GetDefaultWeb3Options(),
	}
	options.Web3.Service = system.SolverService
	return options
}

func AddSolverCliFlags(cmd *cobra.Command, options *solver.SolverOptions) {
	options2.AddWeb3CliFlags(cmd, &options.Web3)
	AddServerCliFlags(cmd, &options.Server)
}

func CheckSolverOptions(options solver.SolverOptions) error {
	err := options2.CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckServerOptions(options.Server)
	if err != nil {
		return err
	}
	return nil
}

func ProcessSolverOptions(options solver.SolverOptions) (solver.SolverOptions, error) {
	newWeb3Options, err := options2.ProcessWeb3Options(options.Web3)
	if err != nil {
		logrus.Debugf("failed to process web3 options %v", err)
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckSolverOptions(options)
}
