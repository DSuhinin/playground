package helpers

import (
	"os"

	"github.com/namsral/flag"
)

//
// ResetEnvVariables makes reset of parsed ENV variables.
//
func ResetEnvVariables() {

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
