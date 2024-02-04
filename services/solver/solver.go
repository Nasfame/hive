package solver

import (
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/internal/genesis"
	optionsfactory "github.com/CoopHive/hive/pkg/options"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/solver/solver"
	memorystore "github.com/CoopHive/hive/services/solver/solver/store/memory"
)

type service struct {
	*genesis.Service
}

func newSolverCmd(s0 *genesis.Service) *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	s := &service{s0}

	solverCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the CoopHive solver service.",
		Long:    "Start the CoopHive solver service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessSolverOptions(options)
			if err != nil {
				return err
			}
			return s.runSolver(cmd, options)
		},
	}

	optionsfactory.AddSolverCliFlags(solverCmd, &options)

	return solverCmd
}

func (s *service) runSolver(cmd *cobra.Command, options solver.SolverOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	solverStore, err := memorystore.NewSolverStoreMemory(s.Conf)
	if err != nil {
		return err
	}

	solverService, err := solver.NewSolver(options, solverStore, web3SDK)
	if err != nil {
		return err
	}

	solverErrors := solverService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-solverErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
