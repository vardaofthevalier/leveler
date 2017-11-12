package cmd

import (
	"os"
	"fmt"
	"context"
	"github.com/spf13/cobra"
	cmdconfig "leveler/resources"
	service "leveler/resources"
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
	// process string options
	for _, f := range r.CmdConfig.Spec.StringOptions {
		cmd.PersistentFlags().StringVarP(new(string), f.Name, string(f.Name[0]), f.Default, f.Description)
	}

	// process bool options
	for _, f := range r.CmdConfig.Spec.BoolOptions {
		cmd.PersistentFlags().BoolVarP(new(bool), f.Name, string(f.Name[0]), f.Default, f.Description)
	}

	// process int64 options
	for _, f := range r.CmdConfig.Spec.Int64Options {
		cmd.PersistentFlags().Int64VarP(new(int64), f.Name, string(f.Name[0]), f.Default, f.Description)
	}
}

func (r ResourceClient) getId(cmd *cobra.Command) string {
	return cmd.Flags().Arg(0)
}

func (r ResourceClient) processFlags(cmd *cobra.Command) ([]*any.Any, error) {
	// process required args and error out if any required parameters aren't set
	var d = []*any.Any{}

	var k string
	var b bool
	var i int64
	var err error
	// process string options
	for _, opt := range r.CmdConfig.Spec.StringOptions {
		k, err = cmd.Flags().GetString(opt.Name)
		if (err != nil || len(k) == 0) && opt.Required {
			fmt.Printf("'%s' is a required parameter!", opt.Name)
			os.Exit(1)
		} else if err != nil || len(k) == 0 {
			continue
		}

		detail := &service.StringDetail{
			Name: opt.Name,
			Value: k,
		}

		a, err := ptypes.MarshalAny(detail) 
		if err != nil {
			return d, nil
		}

		d = append(d, a)
	}

	// process bool options
	for _, opt := range r.CmdConfig.Spec.BoolOptions {
		b, err = cmd.Flags().GetBool(opt.Name)
		if err != nil && opt.Required {
			fmt.Printf("'%s' is a required parameter!", opt.Name)
			os.Exit(1)
		} else if err != nil {
			continue
		}

		detail := &service.BoolDetail{
			Name: opt.Name,
			Value: b,
		}

		a, err := ptypes.MarshalAny(detail) 
		if err != nil {
			return d, nil
		}

		d = append(d, a)
	}

	// process int64 options
	for _, opt := range r.CmdConfig.Spec.Int64Options {
		i, err = cmd.Flags().GetInt64(opt.Name)
		if err != nil && opt.Required {
			fmt.Printf("'%s' is a required parameter!", opt.Name)
			os.Exit(1)
		} else if err != nil {
			continue
		}

		detail := &service.Int64Detail{
			Name: opt.Name,
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

func (r ResourceClient) CreateRequest(cmd *cobra.Command) {
	var err error
	var s = &service.Resource{
		Type: cmd.Name(),
	}

	s.Details, err = r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	resourceId, err := r.Client.CreateResource(context.Background(), s)  // TODO: grpc.CallOption

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resourceId.Id)
}

func (r ResourceClient) GetRequest(cmd *cobra.Command) {
	id := r.getId(cmd)

	var s = &service.Resource{
		Type: cmd.Name(),
		Id: id,
	}

	resource, err := r.doGet(s)
	if err != nil {
		fmt.Println("Error retrieving resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Printf("%v", resource)
}

func (r ResourceClient) doGet(resource *service.Resource) (*service.Resource, error) {
	resource, err := r.Client.GetResource(context.Background(), resource)
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
	fmt.Printf("%v", resourceList)
}	

func (r ResourceClient) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it to update!")

	var err error
	var s = &service.Resource{
		Type: cmd.Name(),
	}

	s.Details, err = r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	_, err = r.doGet(s)

	if err != nil {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)

	} else {
		_, err = r.Client.UpdateResource(context.Background(), s)
		if err != nil {
			fmt.Printf("Error updating resource: %s", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Successfully updated resource '%s'", s.Id)
}

func (r ResourceClient) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	var err error
	var s = &service.Resource{
		Type: cmd.Name(),
	}

	s.Details, err = r.processFlags(cmd)
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	_, err = r.doGet(s)

	if err != nil {
		if err != nil {
			_, err = r.Client.DeleteResource(context.Background(), s)
		} else {
			fmt.Printf("Error deleting resource: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)
	}

	fmt.Printf("Successfully deleted resource '%s'", s.Id)
} 