package config

type configMap[T ~string | ~int] map[T]*argvMeta

type argvMeta struct {
	desc       string
	defaultVal string
}
