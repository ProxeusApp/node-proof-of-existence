package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigurationTestSuite struct {
	suite.Suite

	originalEnvVariable string // Used to store previous environment variable, if set. We're going to reset that at the end
}

func (me *ConfigurationTestSuite) SetupTest() {
	me.originalEnvVariable = os.Getenv("PROXEUS_INSTANCE_URL")
}

func (me *ConfigurationTestSuite) Test_ConfigurationEnvironmentVariablesAreRead() {
	err := os.Setenv("PROXEUS_INSTANCE_URL", "test url")
	assert.Nil(me.T(), err)

	configuration, _ := ReadConfiguration()
	assert.Equal(me.T(), "test url", configuration.Get("PROXEUS_INSTANCE_URL"))
}

func (me *ConfigurationTestSuite) Test_ConfigurationDefaultValuesAreReturnedWhenEnvironmentVariableMissing() {
	err := os.Unsetenv("PROXEUS_INSTANCE_URL")
	assert.Nil(me.T(), err)

	configuration, _ := ReadConfiguration()
	assert.Equal(me.T(), "http://127.0.0.1:1323", configuration.Get("PROXEUS_INSTANCE_URL"))
}

func (me *ConfigurationTestSuite) TearDownTest() {
	err := os.Setenv("PROXEUS_INSTANCE_URL", me.originalEnvVariable)
	if err != nil {
		panic(err)
	}
}

func TestConfigurationTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigurationTestSuite))
}
