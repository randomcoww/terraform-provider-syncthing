package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/randomcoww/terraform-provider-syncthingdevice/syncthingdevice"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: syncthingdevice.Provider})
}
