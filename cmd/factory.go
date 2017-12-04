package cmd

import (
	"os"
	"os/user"
	"fmt"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"leveler/resources"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var opts []grpc.DialOption
var resourceList = buildResourceClientList()


func AddOptions(options []*resources.Option, cmd *cobra.Command) {
	// process string options
	for _, f := range options {
		if f.Type == "string" {
			cmd.PersistentFlags().StringVarP(new(string), f.Name, string(f.Name[0]), string(f.Default), f.Description)
		} else if f.Type == "bool" {
			cmd.PersistentFlags().BoolVarP(new(bool), f.Name, string(f.Name[0]), bool(f.Default), f.Description)
		} else if f.Type == "int64" {
			cmd.PersistentFlags().Int64VarP(new(int64), f.Name, string(f.Name[0]), int64(f.Default), f.Description)
		} else {
			// TODO: handle error situation (unimplemented type)
			// also... implement more types!
		}
	}
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
							Run: func(cmd *cobra.Command, args []string) {
								resource.CreateRequest(cmd)  // TODO: when subcommands are present, need to move this to the innermost command
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(create, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(create, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(create)

					case "add":
						var add = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.AddRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(add, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(add, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(add)

					case "get": 
						var get = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.GetRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(get, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(get, resource.CmdConfig.Subcommands)
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

						// SPECIAL CASE: the list operation is implemented to take in a query for all types
						list.PersistentFlags().StringVarP(new(string), "query", "q", "", "A query for filtering list results")

						if resource.CmdConfig.Options != nil {
							AddOptions(list, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(list, resource.CmdConfig.Subcommands)
						}

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

						if resource.CmdConfig.Options != nil {
							AddOptions(update, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(update, resource.CmdConfig.Subcommands)
						}

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

						if resource.CmdConfig.Options != nil {
							AddOptions(patch, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(patch, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(patch)

					case "remove": 
						var remove = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.RemoveRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(remove, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(remove, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(remove)

					case "delete":
						var delete = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.DeleteRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(delete, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(delete, resource.CmdConfig.Subcommands)
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

						if resource.CmdConfig.Options != nil {
							AddOptions(apply, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(apply, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(apply)

					case "run": 
						var run = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.RunRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(run, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(run, resource.CmdConfig.Subcommands)
						}

						parent.AddCommand(run)

					case "cancel":
						var cancel = &cobra.Command{
							Use:   resource.Usage(),
							Short: resource.ShortDescription(),
							Long: resource.LongDescription(),
							Run: func(cmd *cobra.Command, args []string) {
								resource.CancelRequest(cmd)
							},
						}

						if resource.CmdConfig.Options != nil {
							AddOptions(cancel, resource.cmdConfig.Options)
						}

						if resource.CmdConfig.Subcommands != nil {
							resource.AddSubcommands(cancel, resource.CmdConfig.Subcommands)
						}

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

func AddSubcommands(parent *cobra.Command, subcommands []*cobra.Command) {
	// TODO: implement this
}

func buildResourceClientList() []ResourceClient {
	var r []*ResourceCmd
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
		pb := proto.MessageType(res.CmdConfig.ProtobufType)
		r = append(r, &ResourceCmd{Client: pb.GetClient(clientConn), CmdConfig: *res})
	}

	return r
}