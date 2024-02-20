package options

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/dto"
)

func GetDefaultServicesOptions() dto.ServiceConfig {
	return dto.ServiceConfig{
		Solver:   config.Conf.GetString(enums.HIVE_SOLVER),
		Mediator: strings.Split(config.Conf.GetString(enums.HIVE_MEDIATION), ","),
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
		return fmt.Errorf("no solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	if len(options.Mediator) == 0 {
		return fmt.Errorf("no mediators services specified - please use SERVICE_MEDIATORS or --service-mediators")
	}
	return nil
}
