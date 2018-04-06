package main

import (
	"github.com/avinetworks/terraform-provider-avi/avi"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: avi.Provider})
}
