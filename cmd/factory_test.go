package cmd

import (
	"testing"
)


func (t *testing.T) TestBuildResourceList() {
	// PROBLEM: currently no support for passing in an in-memory configuration to drive the building of the resource client list -- this could be reimplemented, but might not provide any additional value in the long run.  I may need to keep this as an integration test only.

}

func (t *testing.T) TestAddCommands() {
	// This function can be spoofed since it only adds commands to the parent command and doesn't use the resource client object to make calls to the server
}