package leveler

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var Connection *grpc.ClientConn

var opts []grpc.DialOption

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