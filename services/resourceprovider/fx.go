package resourceprovider

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/services/dealmaker"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service

	DealMakerService *dealmaker.Service `name:"dealmaker"`
}

type out struct {
	fx.Out

	ResourceProviderCmd *cobra.Command `name:"rp"`
}

func newServices(i in) (o out) {

	s := &service{
		i.DealMakerService,
		i.Service,
	}

	cmd := s.newResourceProviderCmd()

	o = out{
		ResourceProviderCmd: cmd,
	}
	return
}

func (s *service) newResourceProviderCmd() *cobra.Command {
	options := NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resourceprovider",
		Aliases: []string{"resource-provider", "rp"},
		Short:   "Start the CoopHive resource provider service.",
		Long:    "Start the CoopHive resource provider service.",
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
