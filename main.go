package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/joeltio/terraform-provider-ospath/ospath"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ospath.Provider})
}
