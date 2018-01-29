package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jtopjian/terraform-provider-ansible-inventory/provider/ansibleinventory"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ansibleinventory.Provider})
}
