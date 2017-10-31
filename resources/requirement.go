package leveler

import (
	"fmt"
	"github.com/spf13/cobra"
	// endpoints "leveler/endpoints"
)

type Requirement struct {}

// CLIENT FUNCTIONS

func (requirement *Requirement) Usage() string {
	return "requirement"
}

func (requirement *Requirement) ShortDescription() string {
	return "TODO"
}

func (requirement *Requirement) LongDescription() string {
	return `TODO`
}

func (requirement *Requirement) AddFlags(operation string, cmd *cobra.Command) {

}

func (requirement *Requirement) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement *Requirement) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement *Requirement) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}	

func (requirement *Requirement) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (requirement *Requirement) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
} 


// SERVER FUNCTIONS

func (requirement *Requirement) Create() {

}

func (requirement *Requirement) Get() {
	
}

func (requirement *Requirement) List() {
	
}

func (requirement *Requirement) Update() {
	
}

func (requirement *Requirement) Delete() {
	
}