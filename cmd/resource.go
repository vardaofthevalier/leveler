package cmd

import (
	"os"
	"fmt"
	"context"
	"reflect"
	"leveler/resources"
	"github.com/spf13/cobra"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// type ResourceCommander interface {
// 	Usage() string
// 	ShortDescription() string
// 	LongDescription() string
// 	AddOptions(cmd *cobra.Command)
	
// 	CreateRequest(cmd *cobra.Command) 
// 	GetRequest(cmd *cobra.Command)
// 	ListRequest(cmd *cobra.Command) 
// 	UpdateRequest(cmd *cobra.Command) 
// 	DeleteRequest(cmd *cobra.Command)
// }

type ResourceCmd struct {
	CmdConfig resources.CmdConfig
	Client interface{}
}

func (r ResourceCmd) Usage() string {
	return *r.CmdConfig.Usage
}

func (r ResourceCmd) ShortDescription() string {
	return *r.CmdConfig.ShortDescription
}

func (r ResourceCmd) LongDescription() string {
	return *r.CmdConfig.LongDescription
}

func (r ResourceCmd) getId(cmd *cobra.Command) string {
	return cmd.Flags().Arg(0)
}

func (r ResourceCmd) processFlags(cmd *cobra.Command) (*proto.Message, error) {
	// create protobuf type
	pb := proto.MessageType(r.CmdConfig.ProtobufType)

	// process options
	for _, opt := range r.CmdConfig.Options {
		if opt.Type == "string" {
			k, err := cmd.Flags().GetString(opt.Name)
			if err != nil && opt.Required {
				fmt.Printf("'%s' is a required parameter!", opt.Name)
		 		os.Exit(1)
			}

			reflect.ValueOf(pb).Elem().Field(opt.Name).SetString(k)

		} elif opt.Type == "bool" {
			k, err := cmd.Flags().GetBool(opt.Name)
			if err != nil && opt.Required {
				fmt.Printf("'%s' is a required parameter!", opt.Name)
		 		os.Exit(1)
			}

			reflect.ValueOf(pb).Elem().Field(opt.Name).SetBool(k)

		} elif opt.Type == "int64" {
			k, err := cmd.Flags().GetInt64(opt.Name)
			if err != nil && opt.Required {
				fmt.Printf("'%s' is a required parameter!", opt.Name)
		 		os.Exit(1)
			}

			reflect.ValueOf(pb).Elem().Field(opt.Name).SetInt(k)

		} else {
			// TODO: error and implement additional types
		}
	}

	// process file options

	for _, opt := range r.CmdConfig.FileOptions {
		/*
			TODO: 
			- get file by opt name
			- try to open file (fail if error)
			- try to read contents of (yaml) file into pb message
		*/
	}

	// process subcommands, if any
	for _, sc := range r.CmdConfig.SubCommands {

	}

	return pb, nil
}

func (r ResourceCmd) AddRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Add(r.Client)

	if err != nil {
		fmt.Printf("Error adding resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) CreateRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Create(r.Client)

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) GetRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Get(r.Client)

	if err != nil {
		fmt.Println("Error retrieving resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) ListRequest(cmd *cobra.Command) {
	q, _ := cmd.Flags().GetString("query") 

	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.List(query, r.Client)

	if err != nil {
		fmt.Println("Error listing resources: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}	

func (r ResourceCmd) UpdateRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Update()

	if err != nil {
		fmt.Printf("Error updating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) PatchRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Patch()

	if err != nil {
		fmt.Printf("Error patchingresource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) RemoveRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Remove()

	if err != nil {
		fmt.Printf("Error removing resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) DeleteRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Delete()

	if err != nil {
		fmt.Printf("Error deleting resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
} 

func (r ResourceCmd) ApplyRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Apply()

	if err != nil {
		fmt.Printf("Error applying resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) RunRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Run()

	if err != nil {
		fmt.Printf("Error runnin resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCmd) CancelRequest(cmd *cobra.Command) {
	pb, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resp, err := pb.Cancel()

	if err != nil {
		fmt.Printf("Error cancelling resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}