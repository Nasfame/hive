package enums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrivateKeyEnumForJCServiceType(t *testing.T) {
	serviceType := JC
	privateKeyEnum := serviceType.GetPrivateKeyEnum()
	assert.Equal(t, JC_PRIVATE_KEY, privateKeyEnum)
}

func TestGetPrivateKeyEnumForNegativeServiceTypeValues(t *testing.T) {
	serviceType := ServiceType(-1)
	privateKeyEnum := serviceType.GetPrivateKeyEnum()
	assert.Equal(t, WEB3_PRIVATE_KEY, privateKeyEnum)
}

func TestGetPrivateKeyEnumForRPServiceType(t *testing.T) {
	serviceType := RP
	privateKeyEnum := serviceType.GetPrivateKeyEnum()
	assert.Equal(t, RP_PRIVATE_KEY, privateKeyEnum)
}

func TestGetPrivateKeyEnumForSolverServiceType(t *testing.T) {
	serviceType := SOLVER
	privateKeyEnum := serviceType.GetPrivateKeyEnum()
	assert.Equal(t, SOLVER_PRIVATE_KEY, privateKeyEnum)
}

func TestGetPrivateKeyEnumForUnknownServiceType(t *testing.T) {
	serviceType := ServiceType(10) // Unknown service type
	privateKeyEnum := serviceType.GetPrivateKeyEnum()
	assert.Equal(t, WEB3_PRIVATE_KEY, privateKeyEnum)
}
