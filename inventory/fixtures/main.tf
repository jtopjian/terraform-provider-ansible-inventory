resource "ansible_group" "group_1" {
  name     = "group_1"
  children = ["group_2"]

  vars {
    foo = "bar"
  }
}

resource "ansible_group" "group_2" {
  name = "group_2"
}

resource "ansible_host" "host_1" {
  name   = "host_1"
  groups = ["group_1"]

  vars {
    ansible_user = "ubuntu"
    ansible_host = "1.2.3.4"
  }
}
