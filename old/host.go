package leveler

import (
	"fmt"
	"github.com/spf13/cobra"
	service "leveler/grpc"
)

type Host struct {
	Client service.HostEndpointClient
}

// CLIENT FUNCTIONS

func (host Host) Usage() string {
	return "host"
}

func (host Host) ShortDescription() string {
	return "TODO"
}

func (host Host) LongDescription() string {
	return `TODO`
}

func (host Host) AddFlags(operation string, cmd *cobra.Command) {

}

func (host Host) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (host Host) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (host Host) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}	

func (host Host) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
}

func (host Host) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it!")
} 
