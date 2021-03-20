package config

type Zap struct {
	Level       string `yaml:"level"`
	Director    string `yaml:"director"`
	LinkName    string `yaml:"linkName"`
	Format      string `yaml:"format"`
	Prefix      string `yaml:"prefix"`
	EncodeLevel string `yaml:"encodeLevel"`
}
