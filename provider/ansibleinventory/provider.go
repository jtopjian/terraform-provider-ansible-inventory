package ansibleinventory

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"ansibleinventory_group": resourceAnsibleInventoryGroup(),
			"ansibleinventory_host":  resourceAnsibleInventoryHost(),
		},
	}
}
