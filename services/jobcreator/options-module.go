package jobcreator

import (
	"github.com/spf13/cobra"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/module"
	"github.com/CoopHive/hive/pkg/options"
)

func GetDefaultModuleOptions() dto.ModuleConfig {
	return dto.ModuleConfig{
		// the shortcut name
		Name: options.GetDefaultServeOptionString("MODULE_NAME", ""),
		// the repo we can clone from
		Repo: options.GetDefaultServeOptionString("MODULE_REPO", ""),
		// the hash to checkout the repo
		Hash: options.GetDefaultServeOptionString("MODULE_HASH", ""),
		// the path to the go template file
		Path: options.GetDefaultServeOptionString("MODULE_PATH", ""),
	}
}

func AddModuleCliFlags(cmd *cobra.Command, moduleConfig *dto.ModuleConfig) {
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Name, "module-name", moduleConfig.Name,
		`The name of the shortcut module (MODULE_NAME)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Repo, "module-repo", moduleConfig.Repo,
		`The (http) git repo we can close (MODULE_REPO)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Hash, "module-hash", moduleConfig.Hash,
		`The hash of the repo we can checkout (MODULE_HASH)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Path, "module-path", moduleConfig.Path,
		`The path in the repo to the go template (MODULE_PATH)`,
	)
}

// see if we have a shortcut and fill in the other values if we do
func ProcessModuleOptions(options dto.ModuleConfig) (dto.ModuleConfig, error) {
	return module.ProcessModule(options)
}

func CheckModuleOptions(options dto.ModuleConfig) error {
	return module.CheckModuleOptions(options)
}
