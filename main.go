package main

import (
	"github.com/nholuongutterraform-plugin-sdk/plugin"
	"github.com/nholuongut/terraform-provider-scaffolding/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.New})
}
