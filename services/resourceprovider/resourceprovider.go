package resourceprovider

import (
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/dealmaker"
)

func (s *service) runResourceProvider(cmd *cobra.Command, options ResourceProviderOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if options.Dealer != config.DEFAULT_DEALER {
		if err := s.dealMakerService.LoadPlugin(options.Dealer); err != nil {
			s.Log.Errorf("Dealer %s is not supported on this platform", options.Dealer)
		}
	}

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	executor, err := bacalhau.NewBacalhauExecutor(options.Bacalhau)
	if err != nil {
		return err
	}

	resourceProviderService, err := NewResourceProvider(options, web3SDK, executor, s.dealMakerService)
	if err != nil {
		return err
	}

	resourecProviderErrors := resourceProviderService.Start(commandCtx.Ctx, commandCtx.Cm)

	signal.NotifyContext(commandCtx.Ctx, os.Interrupt)

	for {
		select {
		case err := <-resourecProviderErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}

type service struct {
	dealMakerService *dealmaker.Service
	*genesis.Service
}
