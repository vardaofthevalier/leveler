package leveler

import (
	"os"
	"fmt"
	"context"
	"github.com/spf13/cobra"
	endpoints "leveler/endpoints"
)

type Action struct {
	Client endpoints.ActionEndpointClient
}

// CLIENT FUNCTIONS

func (action Action) Usage() string {
	return "action"
}

func (action Action) ShortDescription() string {
	return "Perform an operation on an Action resource"
}

func (action Action) LongDescription() string {
	return `TODO`
}

func (action Action) AddFlags(operation string, cmd *cobra.Command) {

	var id string
	var query string
	var name string
	var description string
	var command string
	var shell string

	switch operation {
	case "create":
		cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "A descriptive name for the action")
		cmd.PersistentFlags().StringVarP(&description, "description", "d", "", "A description for the goal of the action")
		cmd.PersistentFlags().StringVarP(&command, "command", "c", "", "The command that defines the action")
		cmd.PersistentFlags().StringVarP(&shell, "shell", "s", "/bin/bash", "The shell to use for running the command")
		
	case "get":
		cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "The unique ID for the action of interest")

	case "list": 
		cmd.PersistentFlags().StringVarP(&query, "query", "q", "", "A query for filtering the list of results")

	case "update":
		cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "The unique ID for the action of interest")
		cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "A descriptive name for the action")
		cmd.PersistentFlags().StringVarP(&description, "description", "d", "", "A description for the goal of the action")
		cmd.PersistentFlags().StringVarP(&command, "command", "c", "", "The command that defines the action")
		cmd.PersistentFlags().StringVarP(&shell, "shell", "s", "/bin/bash", "The shell to use for running the command")

	case "delete": 
		cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "The unique ID for the action of interest")

	case "apply":
		cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "The unique ID for the action of interest")
	}
}

func (action Action) CreateRequest(cmd *cobra.Command) {
	fmt.Println("made it to create!")

	// name is required
	name, _ := cmd.Flags().GetString("name")
	if len(name) == 0 {
		fmt.Println("'name' is a required parameter!")
		os.Exit(1)
	}

	// description is optional
	desc, _ := cmd.Flags().GetString("description")

	// command is required
	command, _:= cmd.Flags().GetString("command")
	if len(command) == 0 {
		fmt.Println("'command' is a required parameter!")
		os.Exit(1)
	}

	// shell is optional (default == /bin/bash)
	shell, _ := cmd.Flags().GetString("shell")

	a := endpoints.Action{
		Name: name,
		Description: desc,
		Command: command,
		Shell: shell,
	}

	// do create request
	actionId, err := action.Client.CreateAction(context.Background(), &a)  // TODO: grpc.CallOption

	if err != nil {
		fmt.Printf("Error creating action: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(actionId)
}

func (action Action) GetRequest(cmd *cobra.Command) {
	fmt.Println("made it to get!")
	// id is required
	id, _ := cmd.Flags().GetString("id")
	if len(id) == 0 {
		fmt.Println("'id' is a required parameter!")
		os.Exit(1)
	}

	a := endpoints.ActionId{
		Id: id,
	}

	fmt.Println(a)

	// do get request
	actionData, err := action.doGet(&a)
	if err != nil {
		fmt.Println("Error creating action: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(actionData)
}

func (action Action) doGet(actionId *endpoints.ActionId) (*endpoints.Action, error) {
	a, err := action.Client.GetAction(context.Background(), actionId)
	return a, err
}

func (action Action) ListRequest(cmd *cobra.Command) {
	fmt.Println("made it to list!")

	queryString, _ := cmd.Flags().GetString("query")
	query := endpoints.Query{
		Query: queryString,
	}

	// TODO: validate query string

	// do list request
	actionList, err := action.Client.ListActions(context.Background(), &query)

	if err != nil {
		fmt.Println("Error listing actions: %s", err)
		os.Exit(1)
	}

	// TODO: return formatted response
	fmt.Println(actionList)
}	

func (action Action) UpdateRequest(cmd *cobra.Command) {
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

	a := endpoints.Action{
		Id: id,
		Name: name,
		Description: desc,
		Command: command,
		Shell: shell,
	}

	fmt.Println(a)

	// lookup existing resource 
	actionData, err := action.doGet(&endpoints.ActionId{a.Id})

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

func (action Action) DeleteRequest(cmd *cobra.Command) {
	fmt.Println("made it to delete!")

	// id is required
	id, _ := cmd.Flags().GetString("id")
	if len(id) == 0 {
		fmt.Println("'id' is a required parameter!")
		os.Exit(1)
	}

	a := endpoints.ActionId{
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


