package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	util "leveler/util"
	pb "leveler/cmdconfig"
)

var opts []grpc.DialOption
var resourceList = buildResourceList()

func CreateCommand(resource Resource) *cobra.Command {

	var create = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.CreateRequest(cmd)
		},
	}

	fmt.Printf("%v", create)

	resource.AddFlags("create", create)

	return create	
}

func GetCommand(resource Resource) *cobra.Command {

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

func ListCommand(resource Resource) *cobra.Command {

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

func UpdateCommand(resource Resource) *cobra.Command {

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

func DeleteCommand(resource Resource) *cobra.Command {

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

func buildResourceList() []pb.Resource {
	opts = append(opts, grpc.WithInsecure())
	clientConn, err := grpc.Dial("127.0.0.1:8080", opts...) // TODO: move server and port to config file
	
	if err != nil {
		fmt.Printf("Couldn't connect to server: %s", err)
		os.Exit(1)
	}

	// read the resources.yml file to get a list of resources
	dir, err := filepath.Abs(filepath.Dir("resources.json"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, "resources.json"))
	if err != nil {
		fmt.Printf("Error reading resource configuration file: %v", err)
		os.Exit(1)
	}

	var m = pb.ResourceCmdConfig{}

	err = util.ConvertJsonToProto(contents, &m)
	if err != nil {
		fmt.Printf("Error processing resource configuration file: %v", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		resourceList = append(resourceList, ResourceClient{Client: service.NewResourceEndpointClient(clientConn), pb.CmdConfig: *res})
	}
}