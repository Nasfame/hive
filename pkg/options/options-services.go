package options

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/dto"
)

func GetDefaultServicesOptions() dto.ServiceConfig {
	return dto.ServiceConfig{
		// TODO: refactor to config
		Solver:   GetDefaultServeOptionString("SERVICE_SOLVER", "0xd4646ef9f7336b06841db3019b617ceadf435316"),
		Mediator: GetDefaultServeOptionStringArray("SERVICE_MEDIATORS", []string{"0x2d83ced7562e406151bd49c749654429907543b4"}),
	}
}

func AddServicesCliFlags(cmd *cobra.Command, servicesConfig *dto.ServiceConfig) {
	cmd.PersistentFlags().StringVar(
		&servicesConfig.Solver, "service-solver", servicesConfig.Solver,
		`The solver to connect to (SERVICE_SOLVER)`,
	)
	cmd.PersistentFlags().StringSliceVar(
		&servicesConfig.Mediator, "service-mediators", servicesConfig.Mediator,
		`The mediators we trust (SERVICE_MEDIATORS)`,
	)
}

func ProcessServicesOptions(options dto.ServiceConfig) (dto.ServiceConfig, error) {
	return options, nil
}

func CheckServicesOptions(options dto.ServiceConfig) error {
	if options.Solver == "" {
		return fmt.Errorf("No solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	if len(options.Mediator) == 0 {
		return fmt.Errorf("No mediators services specified - please use SERVICE_MEDIATORS or --service-mediators")
	}
	return nil
}
