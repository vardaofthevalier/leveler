package leveler

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	resources "leveler/resources"
	endpoints "leveler/endpoints"
)

var opts []grpc.DialOption
var resourceList = buildResourceList()

func CreateCommand(resource resources.Resource) *cobra.Command {

	var create = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.CreateRequest(cmd)
		},
	}

	resource.AddFlags("create", create)

	return create	
}

func GetCommand(resource resources.Resource) *cobra.Command {

	var get = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.GetRequest(cmd)
		},
	}

	resource.AddFlags("get", get)

	return get
}

func ListCommand(resource resources.Resource) *cobra.Command {

	var list = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.ListRequest(cmd)
		},
	}

	resource.AddFlags("list", list)

	return list
}

func UpdateCommand(resource resources.Resource) *cobra.Command {

	var update = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.UpdateRequest(cmd)
		},
	}

	resource.AddFlags("update", update)

	return update
}

func DeleteCommand(resource resources.Resource) *cobra.Command {

	var delete = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.DeleteRequest(cmd)
		},
	}

	resource.AddFlags("delete", delete)

	return delete
}

func AddCreateCommands(parent *cobra.Command) {
	fmt.Println(resourceList)
	for _, r := range resourceList {
		parent.AddCommand(CreateCommand(r))
	}
}

func AddListCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(ListCommand(r))
	}
}

func AddGetCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(GetCommand(r))
	}
}

func AddUpdateCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(UpdateCommand(r))
	}
}

func AddDeleteCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(DeleteCommand(r))
	}
}

func buildResourceList() []resources.Resource {
	var r []resources.Resource

	opts = append(opts, grpc.WithInsecure())
	clientConn, err := grpc.Dial("127.0.0.1:8080", opts...) // TODO: move server and port to config file
	if err != nil {
		fmt.Printf("Couldn't connect to server: %s", err)
		os.Exit(1)
	}

	r = append(r, resources.Action{endpoints.NewActionEndpointClient(clientConn)})
	r = append(r, resources.Requirement{endpoints.NewRequirementEndpointClient(clientConn)})
	r = append(r, resources.Role{endpoints.NewRoleEndpointClient(clientConn)})
	r = append(r, resources.Host{endpoints.NewHostEndpointClient(clientConn)})

	return r
}