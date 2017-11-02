package leveler

import (
	"fmt"
	//"log"
	"github.com/spf13/cobra"
	endpoints "leveler/endpoints"
)

type Requirement struct {
	Client endpoints.RequirementEndpointClient
}

// CLIENT FUNCTIONS

func (requirement Requirement) Usage() string {
	return "requirement"
}

func (requirement Requirement) ShortDescription() string {
	return "TODO"
}

func (requirement Requirement) LongDescription() string {
	return `TODO`
}

func (requirement Requirement) AddFlags(operation string, cmd *cobra.Command) {

}

func (requirement Requirement) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement Requirement) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement Requirement) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}	

func (requirement Requirement) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement Requirement) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
} 
