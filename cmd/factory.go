package cmd

import (
	"os"
	"os/user"
	"fmt"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	resources "leveler/resources"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

var opts []grpc.DialOption
var resourceList = buildResourceClientList()

func AddCommands(parent *cobra.Command) {
	var supported bool
	var unsupportedFn = func(cmd *cobra.Command, args []string) {
		fmt.Printf("Unsupported operation '%s' for the specified resource!\n", cmd.Use)
		fmt.Println(cmd.UsageString())
		os.Exit(1)
	}

	for _, resource := range resourceList {  
		for o, _ := range resources.Operation_value {
			for _, s := range resource.CmdConfig.SupportedOperations {
				supported = false

				if o == s.String() {
					supported = true
				}

				if supported && s.String() == parent.Name() {
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

					case "patch":  // TODO: fully implement the patch operation
						var patch = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								fmt.Println("'patch' operation not yet implemented!")
								os.Exit(1)
							},
						}

						resource.AddOptions(patch)
						parent.AddCommand(patch)

					case "delete":
						var delete = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.DeleteRequest(cmd)
							},
						}

						parent.AddCommand(delete)

					case "apply":
						var apply = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								fmt.Println("'apply' operation not yet implemented!")
								os.Exit(1)
							},
						}

						parent.AddCommand(apply)

					default:
						fmt.Printf("Unknown operation '%s' in resource configuration", parent.Name())
						os.Exit(1)
					}
				} else if !supported {
					var unsupported = &cobra.Command{
						Use: parent.Use,
						Short: parent.Short,
						Long: parent.Long,
						Run: unsupportedFn,
					}

					*parent = *unsupported
				}
			} 
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

	// read the resources definition file to get a list of resources -- TODO: move this filename to the config file
	u, err := user.Current() 
	if err != nil {
		fmt.Printf("Error getting user home directory: %v", err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(filepath.Join(u.HomeDir, ".leveler", "resources.json"))
	if err != nil {
		fmt.Printf("Error reading resource configuration file: %v", err)
		os.Exit(1)
	}

	var m = resources.ResourceCmdConfig{}
	var jsonUnmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}

	reader := bytes.NewReader(contents)
	err = jsonUnmarshaler.Unmarshal(reader, &m)
	if err != nil {
		fmt.Printf("Error processing resource configuration file: %v", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		r = append(r, ResourceClient{Client: resources.NewResourceEndpointClient(clientConn), CmdConfig: *res})
	}

	return r
}