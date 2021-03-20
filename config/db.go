package config

type Mysql struct {
	Path         string `yaml:"path"`
	Database     string `yaml:"database"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Debug        bool   `yaml:"logMode"`
	LogZap       string `yaml:"logZap"`
}
