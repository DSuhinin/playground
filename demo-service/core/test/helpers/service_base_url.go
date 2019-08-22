package helpers

import (
	"fmt"
	"os"
	"strings"
)

//
// ServiceBaseURLEnvVariable service URL under test.
//
const ServiceBaseURLEnvVariable = "SERVICE_BASE_URL"

//
// GetServiceBaseURL returns the HTTP server base URL for functional tests.
//
func GetServiceBaseURL() string {

	url := os.Getenv(ServiceBaseURLEnvVariable)
	if url == "" {
		panic(fmt.Sprintf("functional tests variable (%s) is not set", ServiceBaseURLEnvVariable))
	}

	return strings.Replace(url, "http://", "", -1)
}
