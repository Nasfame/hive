package config

type configMap[T ~string] map[T]*argvMeta

type argvMeta struct {
	desc       string
	defaultVal string
}

// //go:embed version.txt
// var version string TODO: another way to embed
