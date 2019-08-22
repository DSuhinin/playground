package helpers

import (
	"fmt"
	"strings"
)

//
// GenerateTestEndpoint makes generation of the test endpoint and replace
// all given placeholders with theirs values.
//
func GenerateTestEndpoint(serverAddress, endpoint string, placeholderList map[string]string) string {

	if placeholderList != nil {
		for placeholder, value := range placeholderList {
			endpoint = strings.Replace(endpoint, placeholder, value, -1)
		}
	}

	return fmt.Sprintf(
		"http://%s%s", serverAddress, endpoint,
	)
}
