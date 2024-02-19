package solver

import (
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/solver/solver"
	memorystore "github.com/CoopHive/hive/services/solver/solver/store/memory"
)

type service struct {
	*genesis.Service
}

func newSolverCmd(s0 *genesis.Service) *cobra.Command {
	options := NewSolverOptions()

	s := &service{s0}

	solverCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the CoopHive solver service.",
		Long:    "Start the CoopHive solver service.",
		Example: "SERVER_URL=0.0.0.0:8080 hive solver",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := ProcessSolverOptions(options)
			if err != nil {
				return err
			}
			return s.runSolver(cmd, options)
		},
	}

	AddSolverCliFlags(solverCmd, &options)

	return solverCmd
}

func (s *service) runSolver(cmd *cobra.Command, options solver.SolverOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		s.Log.Errorf("failed to start due to %v", err)
		return err
	}

	s.Log.Info("solver address ", web3SDK.GetAddress())
	solverAddresses, err := web3SDK.GetSolverAddresses()

	if err != nil {
		s.Log.Errorf("couldn't find solver addresses due to %v", solverAddresses)
	}

	s.Log.Info("default solver addresses ", solverAddresses)

	solverStore, err := memorystore.NewSolverStoreMemory(s.Conf)
	if err != nil {
		return err
	}

	solverService, err := solver.NewSolver(options, solverStore, web3SDK)
	if err != nil {
		return err
	}

	solverErrors := solverService.Start(commandCtx.Ctx, commandCtx.Cm)

	signal.NotifyContext(commandCtx.Ctx, os.Interrupt)

	for {
		/*if commandCtx.Ctx.Done() == nil {
			s.Log.Info("solver service stopped")
			return nil
		}*/

		select {
		case <-commandCtx.Ctx.Done():
			return nil
		case err := <-solverErrors:
			commandCtx.Cleanup()
			return err
		}
	}
}
