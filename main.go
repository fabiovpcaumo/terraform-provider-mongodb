package main

import (
	"flag"

	"github.com/fabiovpcaumo/terraform-provider-mongodb/mongodb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/fabiovpcaumo/mongodb",
		ProviderFunc: mongodb.Provider,
	}

	plugin.Serve(opts)
}
