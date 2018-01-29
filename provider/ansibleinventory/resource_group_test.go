package ansibleinventory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAnsibleInventoryGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAnsibleInventoryGroup_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ansibleinventory_group.group_1", "id", "group_1"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_group.group_1", "children.0", "foo"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_group.group_1", "children.2", "baz"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_group.group_1", "vars.foo", "bar"),
					resource.TestCheckResourceAttr(
						"ansibleinventory_group.group_1", "vars.bar", "2"),
				),
			},
		},
	})
}

const testAccAnsibleInventoryGroup_basic = `
  resource "ansibleinventory_group" "group_1" {
    name = "group_1"
    children = ["foo", "bar", "baz"]
    vars {
      foo = "bar"
      bar = 2
    }
  }
`
