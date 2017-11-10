package cmd

import (
	"os"
	"github.com/spf13/cobra"
	service "leveler/grpc"
	util "leveler/util"
	proto "github.com/golang/protobuf/proto"
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

type resource struct {
	Client service.ResourceEndpointClient
	CmdConfig CmdConfig
}

// CLIENT FUNCTIONS

func (r *resource) Usage() string {
	return r.CmdConfig.Usage
}

func (r *resource) ShortDescription() string {
	return r.CmdConfig.ShortDescription
}

func (r *resource) LongDescription() string {
	return r.CmdConfig.LongDescription
}

func (r *resource) AddFlags(operation string, cmd *cobra.Command) {
	// TODO: find the correct operation within the resource
	for o := range Options {
		switch o {
		case "string":
			var s string
			cmd.PersistentFlags().StringVarP(&s, o.Name, o.ShortName, o.Default, o.Description)
		case "bool":
			var b bool
			cmd.PersistentFlags().StringVarP(&b, o.Name, o.ShortName, o.Default, o.Description)
		default:
			fmt.Printf("Unknown type '%s' for command line option", o)
			os.Exit(1)
		}
	}
}

func (r *resource) processArgs(cmd *cobra.Command) (*proto.Message, error) {
	// process required args and error out if any required parameters aren't set

	// look up the operation
	var operation CmdOperation
	for _, o := range r.CmdConfig.Operations {
		if o.Name == cmd.Name() {
			operation = o
			break
		}
	}

	// iterate through args to determine if required args were provided	and build protobuf message
	var s *server.Resource
	var d map[string]interface{}

	s.Type = r.CmdConfig.Name
	for _, opt := range operation.Options {
		switch opt.Type {
		case "string":
			k, _ := cmd.Flags.GetString(opt.Name)
		case "bool":
			k, _ := cmd.Flags.GetBool(opt.Name)
		default:
			fmt.Printf("Unknown type '%s' in configuration", opt.Type)
			os.Exit(1)
		}

		if len(k) == 0 && opt.Required {
			fmt.Printf("'%s' is a required parameter!")
			os.Exit(1)
		}

		d[opt.Name] = k
	}

	// generate protobuf message
	details, err := util.GenerateProtoAny(d)
	if err != nil {
		return s, err
	}

	s.Details = details

	return s, nil
}

func (r *resource) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it to create!")

	message, err := r.processArgs()
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	// do create request
	resourceId, err := r.Client.CreateAction(context.Background(), &message)  // TODO: grpc.CallOption

	if err != nil {
		fmt.Printf("Error creating resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resourceId)
}

func (r *resource) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it to get!")

	message, err := r.processArgs()
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	// do get request
	resource, err := action.doGet(&message)
	if err != nil {
		fmt.Println("Error retrieving resource: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(resource)
}

func (r *resource) doGet(resourceId *service.ResourceId) {
	resource, err := r.Client.GetAction(context.Background(), resourceId)
	return resource, err
}

func (r *resource) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it to list!")

	queryString, _ := cmd.Flags().GetString("query")  // TODO: special handling for query?  it's sort of a globally required option given the functionality of the database...
	query := service.Query{
		Query: queryString,
	}

	// do list request
	resourceList, err := r.Client.ListActions(context.Background(), &query)

	if err != nil {
		fmt.Println("Error listing actions: %s", err)
		os.Exit(1)
	}

	// TODO: print formatted response
	fmt.Println(resourceList)
}	

func (r *resource) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it to update!")

	message, err := r.processArgs()
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	// lookup existing resource 
	_, err = action.doGet(&service.ResourceId{message.Id})

	if err != nil {
		fmt.Println("Requested action doesn't exist -- nothing to do")
		os.Exit(0)

	} else {
		_, err = action.Client.UpdateResource(context.Background(), &message)
		if err != nil {
			fmt.Printf("Error updating resource: %s", err)
			os.Exit(1)
		}
	}

	// TODO: return formatted response
	fmt.Println(resource)
}

func (r *resource) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	message, err := r.processArgs()
	if err != nil {
		fmt.Printf("Couldn't process args: %v", err)
		os.Exit(1)
	}

	// lookup existing resource 
	_, err = action.doGet(&a)

	if err != nil {
		if err != nil {
			_, err = r.Client.DeleteAction(context.Background(), &message)
		} else {
			fmt.Printf("Error deleting resource: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested resource doesn't exist -- nothing to do")
		os.Exit(0)
	}

	// TODO: return formatted response
	fmt.Printf("Successfully deleted resource '%s'", message.Id)
} 