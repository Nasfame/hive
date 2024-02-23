//go:generate stringer -type=ServiceType -linecomment -output=serviceType_string.go
package enums

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
