package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jtopjian/terraform-provider-ansible-inventory/provider/ansible"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ansible.Provider})
}
