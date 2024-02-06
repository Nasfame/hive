package resourceprovider

import (
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/dealmaker"
)

func (s *service) newResourceProviderCmd() *cobra.Command {
	options := NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the CoopHive resource-provider service.",
		Long:    "Start the CoopHive resource-provider service.",
		Aliases: []string{"resourceprovider"},
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := ProcessResourceProviderOptions(options)
			if err != nil {
				return err
			}
			return s.runResourceProvider(cmd, options)
		},
	}

	AddResourceProviderCliFlags(resourceProviderCmd, &options)

	return resourceProviderCmd
}

func (s *service) runResourceProvider(cmd *cobra.Command, options ResourceProviderOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

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
