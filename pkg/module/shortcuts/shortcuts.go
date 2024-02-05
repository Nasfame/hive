package shortcuts

import (
	"fmt"
	"strings"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/pkg/dto"
)

// TODO: enforce sha1 for tags on the server side (like a pin file)

// parse something with no slashes in it as
// github.com/CoopHive/hive-module-<shortcode>

func GetModule(name string) (dto.ModuleConfig, error) {
	// parse name per following valid formats
	// github.com/user/repo:tag --> Repo: https://github.com/user/repo; Hash = tag
	// bar:tag --> Repo = https://github.com/CoopHive/hive-module-<bar>, Hash = tag
	if name == "" {
		return dto.ModuleConfig{}, fmt.Errorf("module name is empty")
	}
	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return dto.ModuleConfig{}, fmt.Errorf("invalid module name format: %s", name)
	}
	repo, hash := parts[0], parts[1]
	if strings.Contains(name, "/") {
		// 3rd party module
		repo = fmt.Sprintf("https://%s", repo)
	} else {
		// CoopHive std module
		repo = fmt.Sprintf(config.STD_MODULE_FORMAT, repo)
	}

	// TODO: docs for authoring a module
	module := dto.ModuleConfig{
		Name: "", // TODO:
		Repo: repo,
		Hash: hash,
		Path: config.MODULE_PATH,
	}

	return module, nil
}
