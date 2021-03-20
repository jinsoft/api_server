package config

type Server struct {
	Redis  Redis  `yaml:"redis"`
	Mysql  Mysql  `yaml:"mysql"`
	Zap    Zap    `yaml:"zap"`
	System System `yaml:"system"`
}
