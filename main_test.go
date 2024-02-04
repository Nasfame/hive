package main_test

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
)

var ctx = context.Background()

// TestAppConfig tests the config module
func TestAppConfig(t *testing.T) {
	a := assert.New(t)
	testConf := func(conf *viper.Viper) {
		appName := conf.GetString(enums.APP_NAME)

		// a.Equal(appName, "CoopHive", "check AppName")

		if appName != "CoopHive" {
			t.Error("AppName is not CoopHive")
		} else {
			t.Logf("AppName is %s", appName)
		}

		modulePath := conf.GetString(enums.MODULE_PATH)

		a.Contains(modulePath, "module.coophive", "check ModulePath")

	}

	app := fxtest.New(t, config.Module, fx.Invoke(testConf))

	// Start the application.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()
	// if err := app.Start(ctx); err != nil {
	// 	t.Fatal(err)
	// }

	app.RequireStart()

	// Stop the application.
	if err := app.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}
