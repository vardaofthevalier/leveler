package cmd

import (
	"os"
	"github.com/spf13/cobra"
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

func (r *resource) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it to create!")

	// process required args and error out if any required parameters aren't set
	// generate protobuf message to send to the server
	// send the message to the server
	// if there are errors, report them
	// otherwise print the response

	// name is required
	// name, _ := cmd.Flags().GetString("name")
	// if len(name) == 0 {
	// 	fmt.Println("'name' is a required parameter!")
	// 	os.Exit(1)
	// }

	// // description is optional
	// desc, _ := cmd.Flags().GetString("description")

	// // command is required
	// command, _:= cmd.Flags().GetString("command")
	// if len(command) == 0 {
	// 	fmt.Println("'command' is a required parameter!")
	// 	os.Exit(1)
	// }

	// // shell is optional (default == /bin/bash)
	// shell, _ := cmd.Flags().GetString("shell")

	// a := service.Action{
	// 	Name: name,
	// 	Description: desc,
	// 	Command: command,
	// 	Shell: shell,
	// }

	// // do create request
	// actionId, err := action.Client.CreateAction(context.Background(), &a)  // TODO: grpc.CallOption

	// if err != nil {
	// 	fmt.Printf("Error creating action: %s", err)
	// 	os.Exit(1)
	// }

	// // TODO: return formatted response
	// fmt.Println(actionId)
}

func (r *resource) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it to get!")
	// id is required
	id, _ := cmd.Flags().GetString("id")
	if len(id) == 0 {
		fmt.Println("'id' is a required parameter!")
		os.Exit(1)
	}

	a := service.ActionId{
		Id: id,
	}

	fmt.Println(a)

	// do get request
	actionData, err := action.doGet(&a)
	if err != nil {
		fmt.Println("Error retrieving action: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(actionData)
}

func (r *resource) doGet(actionId *service.ActionId) (*service.Action, error) {
	a, err := action.Client.GetAction(context.Background(), actionId)
	return a, err
}

func (r *resource) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it to list!")

	queryString, _ := cmd.Flags().GetString("query")
	query := service.Query{
		Query: queryString,
	}

	// do list request
	actionList, err := action.Client.ListActions(context.Background(), &query)

	if err != nil {
		fmt.Println("Error listing actions: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(actionList)
}	

func (r *resource) UpdateRequest(cmd *cobra.Command) {
	fmt.Println("made it to update!")

	// id required
	id, _ := cmd.Flags().GetString("id")
	if len(id) == 0 {
		fmt.Println("'id' is a required parameter!")
		os.Exit(1)
	}

	// name is optional
	name, _ := cmd.Flags().GetString("name")

	// description is optional
	desc, _ := cmd.Flags().GetString("description")

	// command is optional
	command, _:= cmd.Flags().GetString("command")

	// shell is optional
	shell, _ := cmd.Flags().GetString("shell")

	a := service.Action{
		Id: id,
		Name: name,
		Description: desc,
		Command: command,
		Shell: shell,
	}

	fmt.Println(a)

	// lookup existing resource 
	actionData, err := action.doGet(&service.ActionId{a.Id})

	if err != nil {
		actionData, err = action.Client.UpdateAction(context.Background(), &a)
		if err != nil {
			fmt.Printf("Error updating action: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested action doesn't exist -- nothing to do")
		os.Exit(0)
	}

	// TODO: return formatted response
	fmt.Println(actionData)
}

func (r *resource) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	// id is required
	id, _ := cmd.Flags().GetString("id")
	if len(id) == 0 {
		fmt.Println("'id' is a required parameter!")
		os.Exit(1)
	}

	a := service.ActionId{
		Id: id,
	}

	fmt.Println(a)

	// lookup existing resource 
	actionData, err := action.doGet(&a)

	if err != nil {
		_, err = action.Client.DeleteAction(context.Background(), &a)
		if err != nil {
			fmt.Printf("Error deleting action: %s", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Requested action doesn't exist -- nothing to do")
		os.Exit(0)
	}

	// TODO: return formatted response
	fmt.Println(actionData)
} 