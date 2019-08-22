package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/log"
	"github.com/KWRI/demo-service/core/test/helpers"
)

const (
	EnvVariable          = "ENV_VARIABLE"
	UsageMessage         = "usage message"
	DefaultIntValue      = 2
	DefaultStringValue   = "default string"
	DefaultBoolValue     = false
	DefaultDurationValue = 1 * time.Second
	DefaultLogLevelValue = log.LevelDebug
	IntValue             = 13
	StringValue          = "string"
	BoolValue            = true
	DurationValue        = "2s"
	LogValue             = log.LevelDebug
)

//
// GetIntParameter : initialize config with int parameter : expected well initialized int parameter object.
//
func TestConfigGetIntParameterShouldReturnIntParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, strconv.Itoa(IntValue))
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewInt(EnvVariable, UsageMessage, DefaultIntValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, IntValue, config.GetInt(EnvVariable))
}

//
// GetIntParameter : initialize config with int parameter. don't set the ENV value.
// 				   : expected well initialized int parameter object. int value
// 					 should goes from the default value.
//
func TestConfigGetIntParameterShouldReturnDefaultIntParameter(t *testing.T) {

	helpers.ResetEnvVariables()

	config := New()
	config.RegisterParameters(NewInt(EnvVariable, UsageMessage, DefaultIntValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, DefaultIntValue, config.GetInt(EnvVariable))
}

//
// GetStringParameter : initialize config with string parameter : expected well initialized string parameter object.
//
func TestConfigGetStringParameterShouldReturnStringParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, StringValue)
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewString(EnvVariable, UsageMessage, DefaultStringValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, StringValue, config.GetString(EnvVariable))
}

//
// GetStringParameter : initialize config with string parameter. don't set the ENV value.
// 				   	  : expected well initialized string parameter object. string value
// 					 	should goes from the default value.
//
func TestConfigGetStringParameterShouldReturnDefaultStringParameter(t *testing.T) {

	helpers.ResetEnvVariables()

	config := New()
	config.RegisterParameters(NewString(EnvVariable, UsageMessage, DefaultStringValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, DefaultStringValue, config.GetString(EnvVariable))
}

//
// GetBase64StringParameter : initialize config with base64 string parameter
// 						    : expected well initialized base64 string parameter object.
//
func TestConfigGetBase64StringParameterShouldReturnStringParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, base64.StdEncoding.EncodeToString([]byte(StringValue)))
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewBase64String(
		EnvVariable,
		UsageMessage,
		base64.StdEncoding.EncodeToString([]byte(DefaultStringValue)),
	))

	assert.Nil(t, config.Parse())
	assert.Equal(t, []byte(StringValue), config.GetBase64String(EnvVariable))
}

//
// GetBase64StringParameter : initialize config with string parameter. don't set the ENV value.
// 				   	  		: expected well initialized base64 string parameter object. base64 string value
// 					 		  should goes from the default value.
//
func TestConfigGetBase64StringParameterShouldReturnDefaultStringParameter(t *testing.T) {

	helpers.ResetEnvVariables()

	config := New()
	config.RegisterParameters(NewBase64String(
		EnvVariable,
		UsageMessage,
		base64.StdEncoding.EncodeToString([]byte(DefaultStringValue)),
	))

	assert.Nil(t, config.Parse())
	assert.Equal(t, []byte(DefaultStringValue), config.GetBase64String(EnvVariable))
}

//
// GetBoolParameter : initialize config with bool parameter : expected well initialized bool parameter object.
//
func TestConfigGetBoolParameterShouldReturnBoolParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, strconv.FormatBool(BoolValue))
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewBool(EnvVariable, UsageMessage, DefaultBoolValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, BoolValue, config.GetBool(EnvVariable))
}

//
// GetBoolParameter : initialize config with bool parameter. don't set the ENV value.
// 				    : expected well initialized bool parameter object. bool value
// 					  should goes from the default value.
//
func TestConfigGetBoolParameterShouldReturnDefaultBoolParameter(t *testing.T) {

	helpers.ResetEnvVariables()

	config := New()
	config.RegisterParameters(NewBool(EnvVariable, UsageMessage, DefaultBoolValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, DefaultBoolValue, config.GetBool(EnvVariable))
}

//
// GetDurationParameter : initialize config with duration parameter : expected well initialized duration parameter object.
//
func TestConfigGetDurationParameterShouldReturnDurationParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, DurationValue)
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewDuration(EnvVariable, UsageMessage, DefaultDurationValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, 2*time.Second, config.GetDuration(EnvVariable))
}

//
// GetDuration : initialize config with duration parameter. don't set the ENV value.
// 			   : expected well initialized duration parameter object. duration value
// 			     should goes from the default value.
//
func TestConfigGetDurationShouldReturnDefaultDurationParameter(t *testing.T) {

	helpers.ResetEnvVariables()

	config := New()
	config.RegisterParameters(NewDuration(EnvVariable, UsageMessage, DefaultDurationValue))

	assert.Nil(t, config.Parse())
	assert.Equal(t, DefaultDurationValue, config.GetDuration(EnvVariable))
}

//
// GetLogLevel : initialize config with logger parameter : expected well initialized logger parameter object.
//
func TestConfigGetLogLevelShouldReturnLogLevelParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, LogValue)
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewLoggerLevel(EnvVariable, UsageMessage))

	assert.Nil(t, config.Parse())
	assert.Equal(t, LogValue, config.GetLogLevel(EnvVariable))
}

//
// GetLogLevel : initialize config with logger parameter. don't set the ENV value.
// 			   : expected well initialized logger parameter object. logger value
// 				 should goes from the default value.
//
func TestConfigGetLogLevelShouldReturnDefaultLogLevelParameter(t *testing.T) {

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, LogValue)
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewLoggerLevel(EnvVariable, UsageMessage))

	assert.Nil(t, config.Parse())
	assert.Equal(t, DefaultLogLevelValue, config.GetLogLevel(EnvVariable))
}

//
// GetCassandraInfo : initialize config with cassandra parameter
// 					: expected well initialized cassandra parameter object.
//
func TestConfigGetCassandraInfoShouldReturnCassandraInfoParameter(t *testing.T) {

	const (
		CassandraScheme   = "cassandra"
		CassandraUser     = "user"
		CassandraPassword = "password"
		CassandraHost     = "host:9160"
		CassandraKeyspace = "keyspace"
		CassandraDC       = "DC1"
	)

	var (
		ConnectionString = fmt.Sprintf(
			"%s://%s:%s@%s/%s?dc=%s",
			CassandraScheme,
			CassandraUser,
			CassandraPassword,
			CassandraHost,
			CassandraKeyspace,
			CassandraDC,
		)
	)

	helpers.ResetEnvVariables()
	os.Setenv(EnvVariable, ConnectionString)
	defer func() {
		os.Unsetenv(EnvVariable)
	}()

	config := New()
	config.RegisterParameters(NewCassandraInfo(EnvVariable, UsageMessage))

	assert.Nil(t, config.Parse())

	cassandraInfoParameter := config.GetCassandraConnectionInfo(EnvVariable)

	assert.Equal(t, ConnectionString, cassandraInfoParameter.GetConnectionString())
	assert.Equal(t, CassandraUser, cassandraInfoParameter.GetUser())
	assert.Equal(t, CassandraPassword, cassandraInfoParameter.GetPassword())
	assert.Equal(t, CassandraKeyspace, cassandraInfoParameter.GetKeyspace())
	assert.Equal(t, CassandraDC, cassandraInfoParameter.GetDataCenter())
	assert.Equal(t, []string{CassandraHost}, cassandraInfoParameter.GetHosts())
}
