package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	util "leveler/util"
	resources "leveler/resources"
)

var opts []grpc.DialOption
var resourceList = buildResourceClientList()

func AddCommands(parent *cobra.Command) {
	var supported bool
	for _, resource := range resourceList {
		for _, o := range resource.CmdConfig.SupportedOperations {
			if o.String() == parent.Name() {
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

				//resource.AddOptions(get)  
				parent.AddCommand(get)

			case "list":
				var list = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.ListRequest(cmd)
					},
				}

				//resource.AddOptions(list)  // TODO: add list options function?
				list.PersistentFlags().StringVarP(new(string), "query", "q", "", "A query for filtering list results")
				parent.AddCommand(list)

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
				parent.AddCommand(update)

			case "patch":
				continue
				// fmt.Println("Operation 'patch' not yet implemented")
				// os.Exit(1)

			case "delete":
				var delete = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {
						resource.DeleteRequest(cmd)
					},
				}

				//resource.AddOptions(delete)
				parent.AddCommand(delete)

			case "apply":
				continue
				// fmt.Println("Operation 'apply' not yet implemented")
				// os.Exit(1)

			default:
				continue
				// fmt.Printf("Unknown operation '%s'", parent.Name())
				// os.Exit(1)
			}
			

		} else {
			fmt.Printf("Unsupported operation '%s' for '%s'", parent.Name(), resource.CmdConfig.Name) 
			os.Exit(1)
		}

		
	}
}

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

	var m = resources.ResourceCmdConfig{}

	err = util.ConvertJsonToProto(contents, &m)
	if err != nil {
		fmt.Printf("Error processing resource configuration file: %v", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		r = append(r, ResourceClient{Client: resources.NewResourceEndpointClient(clientConn), CmdConfig: *res})
	}

	return r
}