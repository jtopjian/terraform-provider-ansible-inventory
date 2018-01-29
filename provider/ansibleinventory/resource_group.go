package ansibleinventory

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAnsibleInventoryGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAnsibleInventoryGroupCreate,
		Read:   resourceAnsibleInventoryGroupRead,
		Delete: resourceAnsibleInventoryGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"children": {
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

func resourceAnsibleInventoryGroupRead(d *schema.ResourceData, _ interface{}) error {
	return nil
}

func resourceAnsibleInventoryGroupCreate(d *schema.ResourceData, _ interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
	return nil
}

func resourceAnsibleInventoryGroupDelete(d *schema.ResourceData, _ interface{}) error {
	return nil
}
