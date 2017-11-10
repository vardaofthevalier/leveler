package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	util "leveler/util"
	endpoints "leveler/grpc"
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

func buildResourceList() []Resource {
	var r []Resource

	opts = append(opts, grpc.WithInsecure())
	clientConn, err := grpc.Dial("127.0.0.1:8080", opts...) // TODO: move server and port to config file
	
	if err != nil {
		fmt.Printf("Couldn't connect to server: %s", err)
		os.Exit(1)
	}

	// read the resources.yml file to get a list of resources
	var resourcesYml map[string]interface{}

	contents, err = ioutil.ReadFile("resources.yml")

	if err != nil {
		return err
	}

	err = util.ConvertYamlToMap(contents, m)
	if err != nil {
		fmt.Printf("Couldn't create resource map: %v", err)
		os.Exit(1)
	}

	for r := range m["Resources"] {
		var operations []*service.CmdOperation

		for o := range r["Operations"] {
			var options []*service.Option
			for opt := range o["Options"] {
				options = append(options, &service.Option{Name: opt["Name"], ShortName: opt["ShortName"], Type: opt["Type"], Default: opt["Default"], Description: opt["Description"],})
			}
			operations = append(operations, &service.CmdOperation{Name: o["Name"], Options: options,})
		}
		cmdConfig := &service.CmdConfig{
			Name: r["Name"],
			Usage: r["Usage"],
			ShortDescription: r["ShortDescription"],
			LongDescription: r["LongDescription"],
			Operations: operations,
		}

		r = append(r, resource{Client: endpoints.ResourceEndpointClient(clientConn), CmdConfig: cmdConfig})
	}

	return r
}