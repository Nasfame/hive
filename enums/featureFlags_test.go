package enums_test

import (
	"strings"
	"testing"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/utils"
)

func TestFeatureFlagStringMethod(t *testing.T) {
	featureFlag := enums.PanicIfResultNotFound
	str := featureFlag.String()

	t.Logf("%v.String()=%v", featureFlag, str)

	typeName := utils.GetTypeString(featureFlag)
	t.Logf("typeName=%s", typeName)

	if strings.Contains(str, typeName) {
		t.Errorf("String method should not contain 'FeatureFlag', got: %s", str)
	}
}
