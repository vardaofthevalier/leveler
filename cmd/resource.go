package cmd

import (
	"os"
	"fmt"
	"reflect"
	"context"
	"io/ioutil"
	"leveler/server"
	yaml "gopkg.in/yaml.v2"
	"github.com/spf13/cobra"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
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

type ResourceCommander struct {
	CmdConfig CmdConfig
	Client server.ResourcesClient
}

func (r ResourceCommander) Usage() string {
	return *r.CmdConfig.Usage
}

func (r ResourceCommander) ShortDescription() string {
	return *r.CmdConfig.ShortDescription
}

func (r ResourceCommander) LongDescription() string {
	return *r.CmdConfig.LongDescription
}

func (r ResourceCommander) getId(cmd *cobra.Command) string {
	return cmd.Flags().Arg(0)
}

func (r ResourceCommander) processFlags(cmd *cobra.Command) (reflect.Type, string, error) {
	// create protobuf type
	pb := proto.MessageType(*r.CmdConfig.ProtobufType)

	// process options
	if r.CmdConfig.GetFromOptions() != nil {
		for _, opt := range r.CmdConfig.GetFromOptions().Options {
			var required bool 
			for _, r := range opt.Required {
				if r == cmd.Name() {
					required = true
					break
				}
			}

			if *opt.Type == "string" {
				k, err := cmd.Flags().GetString(*opt.Name)
				if err != nil && required {
					fmt.Printf("'%s' is a required parameter!", *opt.Name)
			 		os.Exit(1)
				}

				reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetString(k)

			} else if *opt.Type == "bool" {
				k, err := cmd.Flags().GetBool(*opt.Name)
				if err != nil && required {
					fmt.Printf("'%s' is a required parameter!", *opt.Name)
			 		os.Exit(1)
				}

				reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetBool(k)

			} else if *opt.Type == "int64" {
				k, err := cmd.Flags().GetInt64(*opt.Name)
				if err != nil && required {
					fmt.Printf("'%s' is a required parameter!", *opt.Name)
			 		os.Exit(1)
				}

				reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetInt(k)

			} else {
				// TODO: error and implement additional types
			}
		}

		for _, sc := range r.CmdConfig.GetFromOptions().Subcommands {
			child := proto.MessageType(*sc.ProtobufType)

			for _, opt := range sc.Options {
				var required bool 
				for _, r := range opt.Required {
					if r == cmd.Name() {
						required = true
						break
					}
				}
				if *opt.Type == "string" {
					k, err := cmd.Flags().GetString(*opt.Name)
					if err != nil && required {
						fmt.Printf("'%s' is a required parameter!", *opt.Name)
				 		os.Exit(1)
					}

					reflect.ValueOf(child).Elem().FieldByName(*opt.Name).SetString(k)

				} else if *opt.Type == "bool" {
					k, err := cmd.Flags().GetBool(*opt.Name)
					if err != nil && required {
						fmt.Printf("'%s' is a required parameter!", opt.Name)
				 		os.Exit(1)
					}

					reflect.ValueOf(child).Elem().FieldByName(*opt.Name).SetBool(k)

				} else if *opt.Type == "int64" {
					k, err := cmd.Flags().GetInt64(*opt.Name)
					if err != nil &&required {
						fmt.Printf("'%s' is a required parameter!", *opt.Name)
				 		os.Exit(1)
					}

					reflect.ValueOf(child).Elem().FieldByName(*opt.Name).SetInt(k)

				} else {
					// TODO: error and implement additional types
				}
			}

			reflect.ValueOf(pb).Elem().FieldByName(*sc.ParentField).Set(reflect.ValueOf(child))
		}
	} else if r.CmdConfig.GetFromFile() != nil {
		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("'file' is a required parameter!")
		 	os.Exit(1)
		}

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(1)
		}

		err = yaml.Unmarshal(contents, pb)
		if err != nil {
			fmt.Printf("Error unmarshaling file contents: %v", err)
			os.Exit(1)
		}

		for _, opt := range r.CmdConfig.GetFromFile().MergeOptions {
			for _, r := range opt.Required {
				if r == cmd.Name() {
					if *opt.Type == "string" {
						k, err := cmd.Flags().GetString(*opt.Name)
						if err != nil {
							fmt.Printf("'%s' is a required parameter!", *opt.Name)
					 		os.Exit(1)
						}

						reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetString(k)

					} else if *opt.Type == "bool" {
						k, err := cmd.Flags().GetBool(*opt.Name)
						if err != nil {
							fmt.Printf("'%s' is a required parameter!", *opt.Name)
					 		os.Exit(1)
						}

						reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetBool(k)

					} else if *opt.Type == "int64" {
						k, err := cmd.Flags().GetInt64(*opt.Name)
						if err != nil {
							fmt.Printf("'%s' is a required parameter!", *opt.Name)
					 		os.Exit(1)
						}

						reflect.ValueOf(pb).Elem().FieldByName(*opt.Name).SetInt(k)

					} else {
						// TODO: error and implement additional types
					}
				}
			}
		}
	}

	return pb, *r.CmdConfig.ProtobufType, nil
}

func (r ResourceCommander) AddRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Add(context.Background(), resource)

	if err != nil {
		fmt.Printf("Error adding resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully added resource!")
}

func (r ResourceCommander) CreateRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	resp, err := r.Client.Create(context.Background(), resource)

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCommander) GetRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	resp, err := r.Client.Get(context.Background(), resource)

	if err != nil {
		fmt.Printf("Error getting resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCommander) ListRequest(cmd *cobra.Command) {
	query, _ := cmd.Flags().GetString("query") 

	_, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	q := &server.Query{
		Type: pbType,
		Query: query,
	}

	resp, err := r.Client.List(context.Background(), q)

	if err != nil {
		fmt.Println("Error listing resources: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}	

func (r ResourceCommander) UpdateRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Update(context.Background(), resource)
	if err != nil {
		fmt.Println("Error updating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully updated resource!")
}

func (r ResourceCommander) PatchRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Patch(context.Background(), resource)
	if err != nil {
		fmt.Println("Error patching resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully patched resource!")
}

func (r ResourceCommander) RemoveRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Remove(context.Background(), resource)
	if err != nil {
		fmt.Println("Error removing resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully removed resource!")
}

func (r ResourceCommander) DeleteRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Delete(context.Background(), resource)
	if err != nil {
		fmt.Println("Error deleting resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully deleted resource!")
} 

func (r ResourceCommander) ApplyRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Apply(context.Background(), resource)
	if err != nil {
		fmt.Println("Error applying resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully applied resource!")
}

func (r ResourceCommander) RunRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	resp, err := r.Client.Run(context.Background(), resource)

	if err != nil {
		fmt.Printf("Error running resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%+v\n", resp)
}

func (r ResourceCommander) CancelRequest(cmd *cobra.Command) {
	pb, pbType, err := r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	m, err := ptypes.MarshalAny(pb.(proto.Message))
	if err != nil {
		fmt.Printf("Couldn't marshal protobuf message: %v", err)
		os.Exit(1)
	}

	resource := &server.Resource{
		Type: pbType,
		Message: m,
	}

	_, err = r.Client.Cancel(context.Background(), resource)
	if err != nil {
		fmt.Println("Error cancelling resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println("Successfully cancelled resource!")
}