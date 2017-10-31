package leveler

import (
	"github.com/spf13/cobra"
	resources "leveler/resources"
)

var resourceList = []resources.Resource{&resources.Action{}, &resources.Requirement{}, &resources.Role{}, &resources.Host{}}

func CreateCommand(resource resources.Resource) *cobra.Command {

	var create = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.CreateRequest(cmd)
		},
	}

	resource.AddFlags("create", create)

	return create	
}

func GetCommand(resource resources.Resource) *cobra.Command {

	var get = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.GetRequest(cmd)
		},
	}

	resource.AddFlags("get", get)

	return get
}

func ListCommand(resource resources.Resource) *cobra.Command {

	var list = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.ListRequest(cmd)
		},
	}

	resource.AddFlags("list", list)

	return list
}

func UpdateCommand(resource resources.Resource) *cobra.Command {

	var update = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.UpdateRequest(cmd)
		},
	}

	resource.AddFlags("update", update)

	return update
}

func DeleteCommand(resource resources.Resource) *cobra.Command {

	var delete = &cobra.Command{
		Use:   resource.Usage(),
		Short: resource.ShortDescription(),
		Long: resource.LongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			resource.DeleteRequest(cmd)
		},
	}

	resource.AddFlags("delete", delete)

	return delete
}

func AddCreateCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(CreateCommand(r))
	}
}

func AddListCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(ListCommand(r))
	}
}

func AddGetCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(GetCommand(r))
	}
}

func AddUpdateCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(UpdateCommand(r))
	}
}

func AddDeleteCommands(parent *cobra.Command) {
	for _, r := range resourceList {
		parent.AddCommand(DeleteCommand(r))
	}
}