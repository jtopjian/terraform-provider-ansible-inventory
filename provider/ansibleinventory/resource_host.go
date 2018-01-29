package ansibleinventory

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAnsibleInventoryHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnsibleInventoryHostCreate,
		Read:   resourceAnsibleInventoryHostRead,
		Delete: resourceAnsibleInventoryHostDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"vars": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAnsibleInventoryHostRead(d *schema.ResourceData, _ interface{}) error {
	return nil
}

func resourceAnsibleInventoryHostCreate(d *schema.ResourceData, _ interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
	return nil
}

func resourceAnsibleInventoryHostDelete(d *schema.ResourceData, _ interface{}) error {
	return nil
}
