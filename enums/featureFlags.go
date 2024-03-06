//go:generate stringer -type=FeatureFlag --trimprefix=FeatureFlag -linecomment -output=featureFlags_string.go

package enums

type FeatureFlag int

const (
	PanicIfResultNotFound FeatureFlag = iota
)
