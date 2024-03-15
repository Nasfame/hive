package config

import (
	"github.com/CoopHive/hive/enums"
)

type argvMeta1[T ~string | ~bool | ~int | any] struct {
	desc       string
	defaultVal T
}

var featureFlags = map[enums.FeatureFlag]argvMeta1[bool]{
	enums.PanicIfResultNotFound: {
		desc:       "Panic if result not found",
		defaultVal: true,
	},
}
