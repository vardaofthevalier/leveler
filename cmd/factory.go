package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	util "leveler/util"
	service "leveler/grpc"
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
	contents, err := ioutil.ReadFile("resources.yml")
	if err != nil {
		fmt.Printf("Error reading resource configuration file: %v", err)
		os.Exit(1)
	}

	var m ResourceCmdConfig
	err = util.ConvertFromYaml(contents, m)
	if err != nil {
		fmt.Printf("Couldn't create resource map: %v", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		r = append(r, ResourceClient{Client: service.NewResourceEndpointClient(clientConn), CmdConfig: *res})
	}

	// resources, ok := m["Resources"].([]interface{})
	// if !ok {
	// 	// TODO
	// }

	// var name, shortname, typ, def, usage, shortdescription, description *string


	// for _, r := range resources {
	// 	var operations []*CmdOperation

	// 	r2, ok := r.(map[string]interface{})
	// 	if !ok {
	// 		// TODO
	// 	}
	// 	ops, ok := r2["Operations"].([]interface{})
	// 	if !ok {
	// 		// TODO
	// 	}

	// 	for _, o := range ops {
	// 		var options []*Option

	// 		o2, ok := o.(map[string]interface{})
	// 		if !ok {
	// 			// TODO
	// 		}
	// 		opts, ok := o2["Options"].([]interface{})
	// 		if !ok {
	// 			// TODO
	// 		}

	// 		for _, opt := range opts {
	// 			opt2, ok := opt.(map[string]string)
	// 			if !ok {
	// 				// TODO
	// 			}
	// 			*name = opt2["Name"]
	// 			*shortname = opt2["ShortName"]
	// 			*typ = opt2["Type"]
	// 			*def = opt2["Default"]
	// 			*description = opt2["Description"]
	// 			options = append(options, &Option{Name: name, ShortName: shortname, Type: typ, Default: def, Description: description,})
	// 		}
	// 		// TODO: start here 
	// 		o3, ok := o2["Name"].(string)
	// 		if !ok {
	// 			// TODO
	// 		}
	// 		*name = o3
	// 		operations = append(operations, &CmdOperation{Name: name, Options: options,}) // TODO: name should be an Operation
	// 	}

	// 	n, ok := r2["Name"].(string)
	// 	if !ok {
	// 		// TODO
	// 	}

	// 	u, ok := r2["Usage"].(string)
	// 	if !ok {
	// 		// TODO
	// 	}

	// 	s, ok := r2["ShortDescription"].(string)
	// 	if !ok {
	// 		// TODO
	// 	}

	// 	l, ok := r2["LongDescription"].(string)
	// 	if !ok {
	// 		// TODO
	// 	}

	// 	// TODO: pointers
	// 	cmdConfig := &CmdConfig{
	// 		Name: n,
	// 		Usage: r,
	// 		ShortDescription: s,
	// 		LongDescription: l,
	// 		Operations: operations,
	// 	}
	//}


	return r
}