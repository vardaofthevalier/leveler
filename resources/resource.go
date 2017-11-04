package leveler

import (
	"github.com/spf13/cobra"
)

type Resource interface {
	Usage() string
	ShortDescription() string
	LongDescription() string
	AddFlags(operation string, cmd *cobra.Command)
	
	CreateRequest(cmd *cobra.Command)
	GetRequest(cmd *cobra.Command)
	ListRequest(cmd *cobra.Command)
	UpdateRequest(cmd *cobra.Command)
	DeleteRequest(cmd *cobra.Command)
}