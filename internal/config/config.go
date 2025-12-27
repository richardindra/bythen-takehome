package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	config      *Config
	watchConfig WatchCfg
)

type (
	WatchCfg struct {
		Path string
		Name string
	}

	option struct {
		configFile string
	}
)

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := os.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	configSplit := strings.Split(opt.configFile, "/")
	lengthSplit := len(configSplit)
	watchConfig.Path = strings.Join(configSplit[:lengthSplit-1], "/")
	watchConfig.Name = configSplit[lengthSplit-1]

	return yaml.Unmarshal(out, &config)
}

func PrepareWatchPath() {
	viper.SetConfigName(watchConfig.Name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(watchConfig.Path)
}

// Option ...
type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {

	configPath := "./files/etc/blog/blog.development.yaml"

	if os.Getenv("GOPATH") == "" {
		configPath = "./files/etc/blog/blog.development.yaml"
	}

	return configPath
}

// Get ...
func Get() *Config {
	return config
}
