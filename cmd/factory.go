package cmd

import (
	"os"
	"os/user"
	"fmt"
	"bytes"
	"strconv"
	"io/ioutil"
	"path/filepath"
	"leveler/resources"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/jsonpb"
)

var opts []grpc.DialOption
var resourceList = buildResourceCommanderList()

type runFn func(*cobra.Command) 

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

func AddOptions(cmd *cobra.Command, options []*resources.Option) {
	// process string options
	for _, f := range options {
		for _, n := range f.Required {
			if cmd.Name() == n {
				if *f.Type == "string" {
					cmd.PersistentFlags().StringVarP(new(string), *f.Name, string((*f.Name)[0]), *f.Default, *f.Description)
				} else if *f.Type == "bool" {
					d, err := strconv.ParseBool(*f.Default)
					if err != nil {
						// TODO: handle error
					}
					cmd.PersistentFlags().BoolVarP(new(bool), *f.Name, string((*f.Name)[0]), d, *f.Description)
				} else if *f.Type == "int64" {
					d, err := strconv.ParseInt(*f.Default, 10, 64)
					if err != nil {
						// TODO: handle error
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

func AddFileOption(cmd *cobra.Command, options []*resources.Option) {
	AddOptions(cmd, options)
	cmd.PersistentFlags().StringVarP(new(string), "file", "f", "", "Resource configuration file")
}

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
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(create, resource, resource.CreateRequest)
						parent.AddCommand(create)

					case "add":
						var add = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(add, resource, resource.AddRequest)
						parent.AddCommand(add)

					case "get": 
						var get = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(get, resource, resource.GetRequest)
						parent.AddCommand(get)

					case "list":
						var list = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						// SPECIAL CASE: the list operation is implemented to take in a query for all types
						list.PersistentFlags().StringVarP(new(string), "query", "q", "", "A query for filtering list results")
						PrepareCmd(list, resource, resource.ListRequest)
						parent.AddCommand(list)

					case "update":
						var update = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(update, resource, resource.UpdateRequest)
						parent.AddCommand(update)

					case "patch": 
						var patch = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(patch, resource, resource.PatchRequest)
						parent.AddCommand(patch)

					case "remove": 
						var remove = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(remove, resource, resource.RemoveRequest)
						parent.AddCommand(remove)

					case "delete":
						var delete = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(delete, resource, resource.DeleteRequest)
						parent.AddCommand(delete)

					case "apply":
						var apply = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(apply, resource, resource.ApplyRequest)
						parent.AddCommand(apply)

					case "run": 
						var run = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(run, resource, resource.RunRequest)
						parent.AddCommand(run)

					case "cancel":
						var cancel = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {},
						}

						PrepareCmd(cancel, resource, resource.CancelRequest)
						parent.AddCommand(cancel)

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

func AddSubcommands(parent *cobra.Command, subcommands []*resources.SubCmdConfig, run runFn) {
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

func buildResourceCommanderList() []ResourceCommander {
	var r []ResourceCommander
	opts = append(opts, grpc.WithInsecure())  // TODO: set appropriate options
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

	contents, err := ioutil.ReadFile(filepath.Join(u.HomeDir, ".leveler", "resources.json"))  // TODO: read from centralized schema stored on the server
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
		client := resources.NewResourcesClient(clientConn)
		r = append(r, ResourceCommander{Client: &client, CmdConfig: *res})
	}

	return r
}