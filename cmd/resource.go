package cmd

import (
	"os"
	"fmt"
	"reflect"
	"context"
	"github.com/spf13/cobra"
	cmdconfig "leveler/cmdconfig"
	service "leveler/grpc"
	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
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

func (r ResourceClient) AddOptions(cmd *cobra.Command) {
	var def string
	for _, o := range r.CmdConfig.Operations {
		// process string options
		if o.Name.String() == cmd.Name() {
			for _, f := range o.StringOptions {
				if f.Default == nil {
					def = ""
				} else {
					def = *f.Default
				}

				if *f.Required {
					cmd.Args = func()

				} else {
					cmd.PersistentFlags().StringVarP(new(string), *f.Name, *f.Name[0], def, *f.Description)
				}
			}

			// process bool options
			for _, f := range o.BoolOptions {
				cmd.PersistentFlags().BoolVarP(new(bool), *f.Name, *f.Name[0], *f.Default, *f.Description)
			}

			// process int64 options
			for _, f := range o.Int64Options {
				cmd.PersistentFlags().Int64VarP(new(int64), *f.Name, *f.Name[0], *f.Default, *f.Description)
			}
		}
	}
}

func (r ResourceClient) getId(cmd *cobra.Command) (string, error) {

}

func (r ResourceClient) processDetails(operation *cmdconfig.CmdOperation, cmd *cobra.Command) ([]*any.Any, error) {
	// process required args and error out if any required parameters aren't set
	var d = []*any.Any{}

	var k string
	var b bool
	var i int64
	var err error
	// process string options
	for _, opt := range operation.StringOptions {
		k, err = cmd.Flags().GetString(*opt.Name)
		if (err != nil || len(k) == 0) && *opt.Required {
			fmt.Printf("'%s' is a required parameter!", *opt.Name)
			os.Exit(1)
		} else if err != nil || len(k) == 0 {
			continue
		}

		detail := &service.StringDetail{
			Name: *opt.Name,
			Value: k,
		}

		a, err := ptypes.MarshalAny(detail) 
		if err != nil {
			return d, nil
		}

		d = append(d, a)
	}

	// process bool options
	for _, opt := range operation.BoolOptions {
		b, err = cmd.Flags().GetBool(*opt.Name)
		if err != nil && *opt.Required {
			fmt.Printf("'%s' is a required parameter!", *opt.Name)
			os.Exit(1)
		} else if err != nil {
			continue
		}

		detail := &service.BoolDetail{
			Name: *opt.Name,
			Value: b,
		}

		a, err := ptypes.MarshalAny(detail) 
		if err != nil {
			return d, nil
		}

		d = append(d, a)
	}

	// process int64 options
	for _, opt := range operation.Int64Options {
		i, err = cmd.Flags().GetInt64(*opt.Name)
		if err != nil && *opt.Required {
			fmt.Printf("'%s' is a required parameter!", *opt.Name)
			os.Exit(1)
		} else if err != nil {
			continue
		}

		detail := &service.Int64Detail{
			Name: *opt.Name,
			Value: i,
		}

		a, err := ptypes.MarshalAny(detail) 
		if err != nil {
			return d, nil
		}

		d = append(d, a)
	}

	return d, nil
}

func (r ResourceClient) lookupOperation(op string, cmd *cobra.Command) *cmdconfig.CmdOperation {
	// look up the operation in the cmdconfig
	var operation *cmdconfig.CmdOperation
	for _, o := range r.CmdConfig.Operations {
		if o.Name.String() == op {
			operation = o
			break
		}
	}

	return operation
}

func (r ResourceClient) CreateRequest(cmd *cobra.Command) {
	var err error
	var s = &service.Resource{
		Type: cmd.Name(),
	}

	// op := r.lookupOperation("create", cmd)

	s.Details, err = r.processDetails(op, cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resource, ok := message.(*service.Resource)
	if !ok {
		fmt.Printf("Type assertion error (can't cast to type Resource) -- %v", message)
		os.Exit(1)
	}
	// do create request
	resourceId, err := r.Client.CreateResource(context.Background(), resource)  // TODO: grpc.CallOption

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resourceId.Id)
}

func (r ResourceClient) GetRequest(cmd *cobra.Command) {
	var err error 

	op := r.lookupOperation("get", cmd)
	id, err := r.getId(cmd)
	if err != nil {
		fmt.Println("Couldn't process positional arg (id): %v", err)
		os.Exit(1)
	}

	var s = &service.Resource{
		Type: cmd.Name(),
		Id: id
	}

	// do get request
	resource, ok := message.(*service.Resource)
	if !ok {
		fmt.Printf("Type assertion error (can't cast to type Resource) -- %v", message)
		os.Exit(1)
	}
	resource, err := r.doGet(resource)
	if err != nil {
		fmt.Println("Error retrieving resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resource)
}

func (r ResourceClient) doGet(resourceMetadata *service.ResourceMetadata) (*service.Resource, error) {
	resource, err := r.Client.GetResource(context.Background(), resourceMetadata)
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

	message, err := r.processArgs("update",cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	input, ok := message.(*service.Resource)
	if !ok {
		fmt.Println("Type assertion error (can't cast to type Resource) -- %v", message)
		os.Exit(1)
	}

	// lookup existing resource 
	_, err = r.doGet(&service.ResourceMetadata{Id: input.Metadata.Id, Type: input.Metadata.Type,})

	if err != nil {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)

	} else {
		_, err = r.Client.UpdateResource(context.Background(), input)
		if err != nil {
			fmt.Printf("Error updating resource: %s", err)
			os.Exit(1)
		}
	}

	// TODO: return formatted response
	fmt.Printf("Successfully updated resource '%s'", input.Metadata.Id)
}

func (r ResourceClient) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	message, err := r.processArgs("delete", cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resourceMetadata, ok := message.(*service.ResourceMetadata)
	if !ok {
		fmt.Printf("Type assertion error (can't cast to type ResourceMetadata) -- %v", message)
		os.Exit(1)
	}
	// lookup existing resource 
	_, err = r.doGet(resourceMetadata)

	if err != nil {
		if err != nil {
			_, err = r.Client.DeleteResource(context.Background(), resourceMetadata)
		} else {
			fmt.Printf("Error deleting resource: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)
	}

	// TODO: return formatted response
	fmt.Printf("Successfully deleted resource '%s'", resourceMetadata.Id)
} 