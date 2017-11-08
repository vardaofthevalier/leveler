package leveler

import (
	"fmt"
	"github.com/spf13/cobra"
	service "leveler/grpc"
)

type Role struct {
	Client service.RoleEndpointClient
}

// CLIENT FUNCTIONS

func (role Role) Usage() string {
	return "role"
}

func (role Role) ShortDescription() string {
	return "TODO"
}

func (role Role) LongDescription() string {
	return `TODO`
}

func (role Role) AddFlags(operation string, cmd *cobra.Command) {
}

func (role Role) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (role Role) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (role Role) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}	

func (role Role) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (role Role) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
} 