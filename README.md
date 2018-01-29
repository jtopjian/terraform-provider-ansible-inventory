Terraform Ansible Inventory
===========================

This repository provides a two-part solution to managing an Ansible inventory
with Terraform:

1. A Terraform provider with resources to manage Ansible groups and hosts.
2. An Ansible dynamic inventory binary to generate an inventory from Terraform.

Quickstart
----------

Create a Terraform manifest similar to the following:

```hcl
resource "openstack_compute_instance_v2" "ansible" {
  count       = 3
  name        = "${format("ansible-%02d", count.index+1)}"
  image_name  = "Ubuntu 16.04"
  flavor_name = "${var.flavor}"
  key_pair    = "${var.key_name}"

  security_groups = ["AllowAll"]

  network {
    uuid = "${var.network_id}"
  }
}

resource "ansibleinventory_group" "group_1" {
  name = "group_1"
}

resource "ansibleinventory_host" "ansible" {
  count  = 3
  name   = "${format("ansible-%02d", count.index+1)}"
  groups = ["${ansibleinventory_group.group_1.name}"]

  vars {
    ansible_host = "${replace(element(openstack_compute_instance_v2.ansible.*.access_ip_v6, count.index), "/[][]/", "")}"
    ansible_user = "ubuntu"
  }
}
```

Then do:

```shell
$ terraform init
$ terraform apply
$ ln -s /path/to/terraform-ansible-inventory hosts
$ ansible -m ping all
```

Terraform Resources
-------------------

### ansibleinventory_group

* `name` - (Required) - The name of the group.
* `children` - (Optional) - A list of children of the group.
* `vars` - (Optional) - Arbitrary key/value variables.

### ansibleinventory_host

* `name` - (Required) - The name of the host.
* `groups` - (Optional) - A list of groups that the host belongs to.
* `vars` - (Optional) - Arbitrary key/value variables.

FAQ
---

### Why?

Methods of using Terraform to provide an inventory to Ansible already exist, so
why build another? Because I found them too limiting.

Existing solutions which parsed the Terraform state overlooked customizations
such as considering IPv6 over IPv4 or using other network interfaces beyond the
first detected. Existing solutions which used templating required too much
boilerplate templating to be useful.

In my opinion, the best way to provide an inventory to Ansible is to give the
user full control over what data is provided to Ansible. To do this easily, the
user needs an easy way to model data. This is where the Terraform Provider comes
in. `terraform-provider-ansibleinventory` provides a bridge between a user's
Terraform resources and the data which Ansible needs for an inventory.

> With that said, [terraform-inventory](https://github.com/adammck/terraform-inventory)
> is a very worthwhile project. The inventory code contained in this repo was
> loosely based on it.

### Can You Write the Inventory to Disk?

No. The `terraform-ansible-inventory` binary does not support writing an
inventory to disk. This is intentional. By writing the inventory to disk, you
are creating a second source of truth of your infrastructure state. This creates
a risk of the two sources diverging.

The Terraform state is the single source of truth.

Building From Source
--------------------

Pre-compiled releases are coming soon. Until then, do the following:

```shell
$ go get github.com/jtopjian/terraform-provider-ansible-inventory
$ cd $GOPATH/src/github.com/jtopjian/terraform-provider-ansible-inventory
$ cd provider
$ make build
$ sudo mv $GOPATH/bin/terraform-provider-ansibleinventory /path/to/dir/with/terraform
$ cd ../inventory
$ make build
$ ln -s $GOPATH/bin/terraform-inventory /path/to/ansible/hosts
```
