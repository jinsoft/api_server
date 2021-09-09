package config

import "time"

type System struct {
	HttpPort int `yaml:"httpPort"`
}

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

type Redis struct {
	Addr     string
	DB       int
	Password string
}

type Zap struct {
	Level       string `yaml:"level"`
	Director    string `yaml:"director"`
	LinkName    string `yaml:"linkName"`
	Format      string `yaml:"format"`
	Prefix      string `yaml:"prefix"`
	EncodeLevel string `yaml:"encodeLevel"`
}

type JWTSettings struct {
	Secret string        `yaml:"secret"`
	Issure string        `yaml:"issure"`
	Expire time.Duration `yaml:"expire"`
}
