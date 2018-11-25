package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/randomcoww/terraform-provider-syncthing/syncthing"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: syncthing.Provider})
}
