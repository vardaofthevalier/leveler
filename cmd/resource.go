package cmd

import (
	"os"
	"fmt"
	"context"
	"github.com/spf13/cobra"
	util "leveler/util"
	cmdconfig "leveler/cmdconfig"
	service "leveler/grpc"
)

type Resource interface {
	Usage() string
	ShortDescription() string
	LongDescription() string
	AddFlags(operation string, cmd *cobra.Command)
	
	CreateRequest(cmd *cobra.Command) 
	GetRequest(cmd *cobra.Command)
	ListRequest(cmd *cobra.Command) 
	UpdateRequest(cmd *cobra.Command) 
	DeleteRequest(cmd *cobra.Command)
}

type ResourceClient struct {
	Client service.ResourceEndpointClient
	CmdConfig cmdconfig.CmdConfig
}

// CLIENT FUNCTIONS

func (r ResourceClient) Usage() string {
	return *r.CmdConfig.Usage
}

func (r ResourceClient) ShortDescription() string {
	return *r.CmdConfig.ShortDescription
}

func (r ResourceClient) LongDescription() string {
	return *r.CmdConfig.LongDescription
}

func (r ResourceClient) AddFlags(operation string, cmd *cobra.Command) {
	// TODO: find the correct operation within the resource
	for _, o := range r.CmdConfig.Operations {
		if o.Name.String() == operation {
			for _, f := range o.Options {
				var s string
				switch *f.Type {
				case "string":
					cmd.PersistentFlags().StringVarP(&s, *f.Name, *f.ShortName, *f.Default, *f.Description)
				case "bool":
					cmd.PersistentFlags().StringVarP(&s, *f.Name, *f.ShortName, *f.Default, *f.Description)
				default:
					fmt.Printf("Unknown type '%s' for command line option", *f.Type)
					os.Exit(1)
				}
			}
		}
	}
}

func (r ResourceClient) processArgs(cmd *cobra.Command) (interface{}, error) {
	// process required args and error out if any required parameters aren't set

	// look up the operation
	var operation *CmdOperation
	for _, o := range r.CmdConfig.Operations {
		if o.Name.String() == cmd.Name() {
			operation = o
			break
		}
	}

	// iterate through args to determine if required args were provided	and build protobuf message
	var s service.Resource
	var d map[string]interface{}

	s.Type = *r.CmdConfig.Name
	var k string
	var b bool
	for _, opt := range operation.Options {
		switch *opt.Type {
		case "string":
			k, _ = cmd.Flags().GetString(*opt.Name)
			if len(k) == 0 && *opt.Required {
				fmt.Printf("'%s' is a required parameter!", *opt.Name)
				os.Exit(1)
			}
		case "bool":
			b, _ = cmd.Flags().GetBool(*opt.Name)
			if !b && *opt.Required {
				fmt.Printf("'%s' is a required parameter!", *opt.Name)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown type '%s' in configuration", *opt.Type)
			os.Exit(1)
		}

		d[*opt.Name] = k
	}

	// generate protobuf message
	details, err := util.GenerateProtoAny(d)
	if err != nil {
		return s, err
	}

	s.Details = details

	return s, nil
}

func (r ResourceClient) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it to create!")

	message, err := r.processArgs(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resource, ok := message.(service.Resource)
	if !ok {
		// TODO
	}
	// do create request
	resourceId, err := r.Client.CreateResource(context.Background(), &resource)  // TODO: grpc.CallOption

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resourceId)
}

func (r ResourceClient) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it to get!")

	message, err := r.processArgs(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	// do get request
	resourceId, ok := message.(service.ResourceId)
	if !ok {
		// TODO
	}
	resource, err := r.doGet(&resourceId)
	if err != nil {
		fmt.Println("Error retrieving resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resource)
}

func (r ResourceClient) doGet(resourceId *service.ResourceId) (*service.Resource, error) {
	resource, err := r.Client.GetResource(context.Background(), resourceId)
	return resource, err
}

func (r ResourceClient) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it to list!")

	queryString, _ := cmd.Flags().GetString("query")  // TODO: special handling for query?  it's sort of a globally required option given the functionality of the database...
	query := service.Query{
		Query: queryString,
	}

	// do list request
	resourceList, err := r.Client.ListResources(context.Background(), &query)

	if err != nil {
		fmt.Println("Error listing resources: %s", err)
		os.Exit(1)
	}

	// TODO: print formatted response
	fmt.Println(resourceList)
}	

func (r ResourceClient) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it to update!")

	message, err := r.processArgs(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	input, ok := message.(service.Resource)
	if !ok {
		// TODO
	}

	// lookup existing resource 
	_, err = r.doGet(&service.ResourceId{Id: input.Id, Type: input.Type,})

	if err != nil {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)

	} else {
		_, err = r.Client.UpdateResource(context.Background(), &input)
		if err != nil {
			fmt.Printf("Error updating resource: %s", err)
			os.Exit(1)
		}
	}

	// TODO: return formatted response
	fmt.Printf("Successfully updated resource '%s'", input.Id)
}

func (r ResourceClient) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	message, err := r.processArgs(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resourceId, ok := message.(service.ResourceId)
	if !ok {
		// TODO
	}
	// lookup existing resource 
	_, err = r.doGet(&service.ResourceId{Id: resourceId.Id, Type: resourceId.Type})

	if err != nil {
		if err != nil {
			_, err = r.Client.DeleteResource(context.Background(), &resourceId)
		} else {
			fmt.Printf("Error deleting resource: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)
	}

	// TODO: return formatted response
	fmt.Printf("Successfully deleted resource '%s'", resourceId.Id)
} 