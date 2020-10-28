package main

import (
	// Upstream Terraform Plugin Library
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	// Our local Terraform Provider code
	"github.com/myuser/todo-terraform-provider-class/provider-parts/08-complete/todo"
)

// main is the entrypoint to the terraform plugin
func main() {
	// see: https://github.com/hashicorp/terraform-plugin-sdk/blob/master/plugin/serve.go
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: todo.Provider})
}
