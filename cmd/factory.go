package cmd

import (
	"os"
	"os/user"
	"fmt"
	"bytes"
	"strconv"
	"io/ioutil"
	"path/filepath"
	"leveler/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/jsonpb"
)

var opts []grpc.DialOption
var resourceList = BuildResourceCommanderList()

type runFn func(*cobra.Command) 

func GetResourceList() []ResourceCommander {
	return resourceList
}

func PrepareCmd(c *cobra.Command, resource ResourceCommander, run runFn) {
	if resource.CmdConfig.GetFromOptions() != nil {
		AddOptions(c, resource.CmdConfig.GetFromOptions().Options)

		if resource.CmdConfig.GetFromOptions().Subcommands != nil {
			AddSubcommands(c, resource.CmdConfig.GetFromOptions().Subcommands, run)

		} else {
			c.Run = func(cmd *cobra.Command, args []string) {
				run(cmd)  
			}
		}
	} else if resource.CmdConfig.GetFromFile() != nil {
		AddFileOption(c, resource.CmdConfig.GetFromFile().MergeOptions)
		c.Run = func(cmd *cobra.Command, args []string) {
			run(cmd)  
		}
	} else {
		// TODO: error
	}
}

func AddOptions(cmd *cobra.Command, options []*server.Option) {
	// process string options
	for _, f := range options {
		for _, n := range f.Required {
			if cmd.Name() == n.String() {
				if *f.Type == "string" {
					cmd.PersistentFlags().StringVarP(new(string), *f.Name, string((*f.Name)[0]), *f.Default, *f.Description)
				} else if *f.Type == "bool" {
					d, err := strconv.ParseBool(*f.Default)
					if err != nil {
						fmt.Printf("Error parsing boolean value: %v\n", err)
						os.Exit(1)
					}
					cmd.PersistentFlags().BoolVarP(new(bool), *f.Name, string((*f.Name)[0]), d, *f.Description)
				} else if *f.Type == "int64" {
					d, err := strconv.ParseInt(*f.Default, 10, 64)
					if err != nil {
						fmt.Printf("Error parsing int64 value: %v\n", err)
						os.Exit(1)
					}
					cmd.PersistentFlags().Int64VarP(new(int64), *f.Name, string((*f.Name)[0]), d, *f.Description)
				} else {
					// TODO: handle error situation (unimplemented type)
					// also... implement more types!
				}
			}
		}
	}
}

func AddFileOption(cmd *cobra.Command, options []*server.Option) {
	AddOptions(cmd, options)
	cmd.PersistentFlags().StringVarP(new(string), "file", "f", "", "Resource configuration file")
}

func AddCommands(resource ResourceCommander, parent *cobra.Command) {
	// var onInit []func()
	var unsupportedFn = func(cmd *cobra.Command) {
		fmt.Printf("Unsupported operation '%s' for the specified resource!\n", cmd.Use)
		os.Exit(1)
	}

	supportMap := make(map[string]bool)
	for o, _ := range server.Operation_value {
		supportMap[o] = false
	}

	for _, s := range resource.CmdConfig.SupportedOperations {
		supportMap[s.String()] = true
	}

	for p, supported := range supportMap {
		if p == parent.Name() {
			switch parent.Name() {
			case "create":
				var create = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(create, resource, resource.CreateRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'create' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(create, ResourceCommander{}, unsupportedFn)
					create.Hidden = true
					create.SetHelpFunc(help)
				}

				parent.AddCommand(create)

			case "add":
				var add = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(add, resource, resource.AddRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'add' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(add, ResourceCommander{}, unsupportedFn)
					add.Hidden = true
					add.SetHelpFunc(help)
				}

				parent.AddCommand(add)

			case "get": 
				var get = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(get, resource, resource.GetRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'get' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(get, ResourceCommander{}, unsupportedFn)
					get.Hidden = true
					get.SetHelpFunc(help)
				}

				parent.AddCommand(get)

			case "list":
				var list = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
					TraverseChildren: true,
				}

				// SPECIAL CASE: the list operation is implemented to take in a query for all types
				list.PersistentFlags().StringVarP(new(string), "query", "q", "", "A query for filtering list results")

				if supported {
					PrepareCmd(list, resource, resource.ListRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'list' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(list, ResourceCommander{}, unsupportedFn)
					list.Hidden = true
					list.SetHelpFunc(help)
				}

				parent.AddCommand(list)

			case "update":
				var update = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(update, resource, resource.UpdateRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'update' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(update, ResourceCommander{}, unsupportedFn)
					update.Hidden = true
					update.SetHelpFunc(help)
				}

				parent.AddCommand(update)

			case "patch": 
				var patch = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(patch, resource, resource.PatchRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'patch' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(patch, ResourceCommander{}, unsupportedFn)
					patch.Hidden = true
					patch.SetHelpFunc(help)
				}

				parent.AddCommand(patch)

			case "remove": 
				var remove = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(remove, resource, resource.RemoveRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'remove' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(remove, ResourceCommander{}, unsupportedFn)
					remove.Hidden = true
					remove.SetHelpFunc(help)
				}	

				parent.AddCommand(remove)

			case "delete":
				var delete = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(delete, resource, resource.DeleteRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'delete' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(delete, ResourceCommander{}, unsupportedFn)
					delete.Hidden = true
					delete.SetHelpFunc(help)
				}

				parent.AddCommand(delete)

			case "apply":
				var apply = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(apply, resource, resource.ApplyRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'apply' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(apply, ResourceCommander{}, unsupportedFn)
					apply.Hidden = true
					apply.SetHelpFunc(help)
				}

				parent.AddCommand(apply)

			case "run": 
				var run = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(run, resource, resource.RunRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'run' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(run, ResourceCommander{}, unsupportedFn)
					run.Hidden = true
					run.SetHelpFunc(help)
				}

				parent.AddCommand(run)

			case "cancel":
				var cancel = &cobra.Command{
					Use:   resource.Usage(),
					Short: resource.ShortDescription(),
					Long: resource.LongDescription(),
					Run: func(cmd *cobra.Command, args []string) {},
				}

				if supported {
					PrepareCmd(cancel, resource, resource.CancelRequest)
				} else {
					help := func(cmd *cobra.Command, args []string) {
						fmt.Printf("Unsupported operation 'cancel' for resource '%s'!\n\n", resource.Usage())
						cmd.Parent().Help()
					}

					PrepareCmd(cancel, ResourceCommander{}, unsupportedFn)
					cancel.Hidden = true
					cancel.SetHelpFunc(help)
				}

				parent.AddCommand(cancel)

			default:
				fmt.Printf("Unknown operation '%s' in resource configuration", parent.Name())
				os.Exit(1)
			}
		} 
	}
}

func AddSubcommands(parent *cobra.Command, subcommands []*server.SubCmdConfig, run runFn) {
	for _, s := range subcommands {
		sub := &cobra.Command{
			Use:   *s.Usage,
			Short: *s.ShortDescription,
			Long: *s.LongDescription,
			Run: func(cmd *cobra.Command, args []string) {
				run(cmd)
			},
		}

		if s.Options != nil {
			AddOptions(sub, s.Options)
		}
		
		parent.AddCommand(sub)
	}
}

func BuildResourceCommanderList() []ResourceCommander {
	var r []ResourceCommander
	opts = append(opts, grpc.WithInsecure())  // TODO: set appropriate options
	clientConn, err := grpc.Dial("127.0.0.1:8080", opts...) // TODO: move server and port to config file
	
	if err != nil {
		fmt.Printf("Couldn't connect to server: %s\n", err)
		os.Exit(1)
	}

	// read the resources definition file to get a list of resources -- TODO: move this filename to the config file
	u, err := user.Current() 
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(filepath.Join(u.HomeDir, ".leveler", "resources.json"))  // TODO: read from centralized schema stored on the server -- maybe implement a GetSchema endpoint that just sends the message directly
	if err != nil {
		fmt.Printf("Error reading resource configuration file: %v\n", err)
		os.Exit(1)
	}

	var m = server.ResourceCmdConfig{}
	var jsonUnmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}

	reader := bytes.NewReader(contents)
	err = jsonUnmarshaler.Unmarshal(reader, &m)
	if err != nil {
		fmt.Printf("Error processing resource configuration file: %v\n", err)
		os.Exit(1)
	}

	for _, res := range m.Resources {
		client := server.NewResourcesClient(clientConn)
		r = append(r, ResourceCommander{Client: client, CmdConfig: *res})
	}

	return r
}