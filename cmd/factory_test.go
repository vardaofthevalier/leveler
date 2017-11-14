package cmd

import (
	"testing"
)


func TestBuildResourceList(t *testing.T) {
	// PROBLEM: currently no support for passing in an in-memory configuration to drive the building of the resource client list -- this could be reimplemented, but might not provide any additional value in the long run.  I may need to keep this as an integration test only (i.e., if mock == true, skip)

}

func TestAddCommands(t *testing.T) {
	// This function can be spoofed since it only adds commands to the parent command and doesn't use the resource client object to make calls to the server
	// things to check:
	// - are all supported operations (and no others) added to the parent command?
	// - are all the options added to the supported child commands?
	// PROBLEM: currently the command-line parsing logic needs some work to correct some minor issues.  See todo/features_and_improvements.txt for more info.
}