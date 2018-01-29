package ansibleinventory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAnsibleInventoryHost_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAnsibleInventoryHost_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ansibleinventory_host.host_1", "id", "host_1"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_host.host_1", "groups.0", "group_1"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_host.host_1", "vars.ansible_host", "1.2.3.4"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_host.host_1", "vars.ansible_user", "ubuntu"),
				),
			},
		},
	})
}

const testAccAnsibleInventoryHost_basic = `
	resource "ansibleinventory_group" "group_1" {
		name = "group_1"
	}

  resource "ansibleinventory_host" "host_1" {
    name = "host_1"
    groups = ["group_1"]

    vars {
      ansible_host = "1.2.3.4"
      ansible_user = "ubuntu"
    }
  }
`
