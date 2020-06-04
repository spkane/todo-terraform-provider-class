package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/myuser/todo-terraform-provider-class/terraform-provider-todo/todo"
)

// main is the entrypoint to the terraform plugin
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: todo.Provider})
}
