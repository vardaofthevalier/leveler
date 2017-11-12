package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	util "leveler/util"
	cmdconfig "leveler/cmdconfig"
	service "leveler/grpc"
)

var opts []grpc.DialOption
var resourceList = buildResourceClientList()

//type cmdRunFn func(*cobra.Command, []string)

// func CreateCommand(resource Resource, runFn cmdRunFn) *cobra.Command {

// 	var create = &cobra.Command{
// 		Use:   resource.Usage(),
// 		Short: resource.ShortDescription(),
// 		Long: resource.LongDescription(),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			resource.CreateRequest(cmd)
// 		},
// 	}

// 	resource.AddOptions("create", create)

// 	return create	
// }

// func GetCommand(resource Resource) *cobra.Command {

// 	var get = &cobra.Command{
// 		Use:   resource.Usage(),
// 		Short: resource.ShortDescription(),
// 		Long: resource.LongDescription(),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			resource.GetRequest(cmd)
// 		},
// 	}

// 	resource.AddOptions("get", get)

// 	return get
// }

// func ListCommand(resource Resource) *cobra.Command {

// 	var list = &cobra.Command{
// 		Use:   resource.Usage(),
// 		Short: resource.ShortDescription(),
// 		Long: resource.LongDescription(),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			resource.ListRequest(cmd)
// 		},
// 	}

// 	resource.AddOptions("list", list)

// 	return list
// }

// func UpdateCommand(resource Resource) *cobra.Command {

// 	var update = &cobra.Command{
// 		Use:   resource.Usage(),
// 		Short: resource.ShortDescription(),
// 		Long: resource.LongDescription(),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			resource.UpdateRequest(cmd)
// 		},
// 	}

// 	resource.AddOptions("update", update)

// 	return update
// }

// func DeleteCommand(resource Resource) *cobra.Command {

// 	var delete = &cobra.Command{
// 		Use:   resource.Usage(),
// 		Short: resource.ShortDescription(),
// 		Long: resource.LongDescription(),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			resource.DeleteRequest(cmd)
// 		},
// 	}

// 	resource.AddOptions("delete", delete)

// 	return delete
// }

func AddCommands(parent *cobra.Command) {
	var runFn cmdRunFn
	var supported bool
	for _, resource := range resourceList {
		for _, o := range resource.CmdConfig.SupportedOperations {
			if o == parent.Name() {
				supported = true
			}
		}

		if supported {
			switch parent.Name() {
			case "create":
				var create = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.CreateRequest(cmd)
					},
				}

				resource.AddOptions(create)
				parent.AddCommand(create)

			case "get": 
				var get = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.GetRequest(cmd)
					},
				}

				resource.AddOptions("get", get)

			case "list":
				var list = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.ListRequest(cmd)
					},
				}

				resource.AddOptions(list)

			case "update":
				var update = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.UpdateRequest(cmd)
					},
				}

				resource.AddOptions(update)

				return update

			case "patch":
				fmt.Println("Operation 'patch' not yet implemented")
				os.Exit(1)

			case "delete":
				var delete = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.DeleteRequest(cmd)
					},
				}

				resource.AddOptions(delete)

			case "apply":
				fmt.Println("Operation 'patch' not yet implemented")
				os.Exit(1)

			default:
				fmt.Printf("Unknown operation '%s'", operation)
				os.Exit(1)
			}
			

		} else {
			runFn = func(cmd *cobra.Command, []string) { 
				fmt.Printf("Unsupported operation '%s' for '%s'", operation, cmd.Name()) 
				os.Exit(1)
			}
		}

		
	}
}

// func AddListCommands(parent *cobra.Command) {
// 	for _, r := range resourceList {
// 		parent.AddCommand(ListCommand(r))
// 	}

// 	var runFn cmdRunFn
// 	var supported bool
// 	for _, r := range resourceList {
// 		for _, o := range r.CmdConfig.SupportedOperations {
// 			if o == "list" {
// 				supported = true
// 			}
// 		}

// 		if supported {
// 			runFn = func(cmd *cobra.Command, []string) { resource.ListRequest(cmd) }

// 		} else {
// 			runFn = func(cmd *cobra.Command, []string) { 
// 				fmt.Printf() 
// 				os.Exit(1)
// 			}
// 		}

// 		parent.AddCommand(CreateCommand(r, runFn))
// 	}
// }

// func AddGetCommands(parent *cobra.Command) {
// 	for _, r := range resourceList {
// 		parent.AddCommand(GetCommand(r))
// 	}
// }

// func AddUpdateCommands(parent *cobra.Command) {
// 	for _, r := range resourceList {
// 		parent.AddCommand(UpdateCommand(r))
// 	}
// }

// func AddDeleteCommands(parent *cobra.Command) {
// 	for _, r := range resourceList {
// 		parent.AddCommand(DeleteCommand(r))
// 	}
// }

func buildResourceClientList() []ResourceClient {
	var r []ResourceClient
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

	var m = cmdconfig.ResourceCmdConfig{}

	err = util.ConvertJsonToProto(contents, &m)
	if err != nil {
		fmt.Printf("Error processing resource configuration file: %v", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		r = append(r, ResourceClient{Client: service.NewResourceEndpointClient(clientConn), CmdConfig: *res})
	}

	return r
}