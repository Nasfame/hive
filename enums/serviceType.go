//go:generate stringer -type=ServiceType -linecomment -output=serviceType_string.go
package enums

import (
	"path"
	"strings"

	"github.com/CoopHive/hive/utils"
)

type ServiceType int

const (
	JC ServiceType = iota
	RP

	SOLVER
	MEDIATOR
)

// GetPrivateKeyEnum returns the enum associated with retrieving the private key for the service
func (s ServiceType) GetPrivateKeyEnum() string {
	switch s {
	case JC:
		return JC_PRIVATE_KEY
	case RP:
		return RP_PRIVATE_KEY
	case SOLVER:
		return SOLVER_PRIVATE_KEY
	case MEDIATOR:
		return MEDIATOR_PRIVATE_KEY
	default:
		return WEB3_PRIVATE_KEY
		// log.Fatal("unknown service type")
	}
}

// TODO:similar implementation in system/Service

func (s ServiceType) ProcessSyncServiceDirectory(appDir string, syncFunc func(serviceDir string)) (string, error) {
	serviceName := s.String()

	serviceName = strings.ToLower(serviceName)
	serviceDir := path.Join(appDir, serviceName)

	defer syncFunc(serviceDir)

	return utils.EnsureDir(serviceDir)
}

// ProcessDirectoryForService returns the directory for the service.
// Deprecated on b0c1d4f29508f6718e72fca14f2bb4302774dbff
func (s ServiceType) ProcessDirectoryForService(appDataDir string) (string, error) {
	serviceName := s.String()
	serviceDir := path.Join(appDataDir, serviceName)
	return utils.EnsureDir(serviceDir)
}
