package resourceprovider

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
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
	serviceType := enums.RP

	resourceProviderCmd := &cobra.Command{
		Use:     "rp",
		Aliases: []string{"resource-provider", "resourceprovider"},
		Short:   "Start the CoopHive resource provider service.",
		Long:    "Start the CoopHive resource provider service.",
		Example: "hive rp --offer-cpu 3000 --offer-gpu 1 --offer-ram 1024 --offer-count 1 (rp is offering 1 offer of {3000mCPUs, 1024MB RAM,1 GPU }",
		RunE: func(cmd *cobra.Command, _ []string) error {
			serviceType.ProcessSyncServiceDirectory(s.Conf.GetString(enums.APP_DIR), func(appDir string) {
				config.SetAppDir(s.Conf, appDir)
			})
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
