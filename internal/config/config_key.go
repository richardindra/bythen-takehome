package config

type (
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
	}

	ServerConfig struct {
		Port string `yaml:"port"`
	}

	DatabaseConfig struct {
		Master string `yaml:"master"`
	}

)
