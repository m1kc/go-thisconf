// Package thisconf is a simple package to do things that Viper doesn't.
package thisconf

import (
	"github.com/spf13/viper"
	"github.com/utrack/goroadie"
)

// Load loads app config from files & env variables. Config structure
// is whatever you like (and whatever the libs are able to unmarshal into).
func Load(conf interface{}, envPrefix string) (err error) {
	// First, read stuff from config file using Viper.
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		return
	}

	// Then, overwrite them with env variables using goroadie.
	goroadie.Process(envPrefix, conf)

	return
}
