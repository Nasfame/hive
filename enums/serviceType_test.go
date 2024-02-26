package enums

import (
	"os"
	"path"
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

func TestProcessDirectoryForService(t *testing.T) {
	// Set up the test

	services := []ServiceType{JC, RP, SOLVER, MEDIATOR}

	for _, serviceName := range services {
		testProcessDirectoryForService(serviceName, t)
	}

}

func testProcessDirectoryForService(serviceName ServiceType, t *testing.T) {
	appDataDir := t.TempDir()

	serviceDir := path.Join(appDataDir, serviceName.String())

	// Call the method
	serviceDirGen, err := serviceName.ProcessDirectoryForService(appDataDir)
	if err != nil {
		t.Fatalf("Error processing directory for service: %v", err)
	}
	assert.Equal(t, serviceDir, serviceDirGen)

	// Check if the directory was created with the correct name
	_, err = os.Stat(serviceDir)
	if os.IsNotExist(err) {
		t.Errorf("Expected directory %s to be created, but it does not exist", serviceDir)
	}

	// Check if the directory was created with the correct name
	_, err = os.Stat(serviceDir)
	if os.IsNotExist(err) {
		t.Errorf("Expected directory %s to be created, but it does not exist", serviceDir)
	} else {
		t.Logf("service dir : %v", serviceDir)
	}
}
