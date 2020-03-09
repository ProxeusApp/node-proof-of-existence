package main

import (
	"github.com/spf13/viper"
)

// The configuration only supports the following environment variables:
// PROXEUS_INSTANCE_URL, SERVICE_NAME, SERVICE_URL, SERVICE_PORT, SERVICE_SECRET, AUTH_KEY
// Since it's an example only, we will provide default values.
// Please set these for production use
type Configuration struct{}

func ReadConfiguration() (Configuration, error) {
	viper.SetDefault("PROXEUS_INSTANCE_URL", "http://127.0.0.1:1323")
	viper.SetDefault("SERVICE_NAME", "proof-of-existence")
	viper.SetDefault("SERVICE_URL", "127.0.0.1")
	viper.SetDefault("SERVICE_PORT", "8012")
	viper.SetDefault("SERVICE_SECRET", "my secret")
	viper.SetDefault("AUTH_KEY", "auth")
	viper.AutomaticEnv()

	return Configuration{}, nil
}

func (me Configuration) Get(key string) string {
	return viper.GetString(key)
}
