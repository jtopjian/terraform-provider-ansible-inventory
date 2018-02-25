package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedState = State{
	Modules: []Module{
		Module{
			Resources: map[string]Resource{
				"ansible_host.host_1": Resource{
					Type: "ansible_host",
					Primary: Primary{
						ID: "host_1",
						Attributes: map[string]string{
							"id":                "host_1",
							"name":              "host_1",
							"groups.#":          "1",
							"groups.0":          "group_1",
							"vars.%":            "2",
							"vars.ansible_host": "1.2.3.4",
							"vars.ansible_user": "ubuntu",
						},
					},
				},
				"ansible_group.group_1": Resource{
					Type: "ansible_group",
					Primary: Primary{
						ID: "group_1",
						Attributes: map[string]string{
							"id":         "group_1",
							"name":       "group_1",
							"children.#": "1",
							"children.0": "group_2",
							"vars.%":     "1",
							"vars.foo":   "bar",
						},
					},
				},
				"ansible_group.group_2": Resource{
					Type: "ansible_group",
					Primary: Primary{
						ID: "group_2",
						Attributes: map[string]string{
							"id":   "group_2",
							"name": "group_2",
						},
					},
				},
			},
		},
	},
}

var expectedInventory = map[string]interface{}{
	"group_1": map[string]interface{}{
		"hosts":    []string{"host_1"},
		"children": []string{"group_2"},
		"vars": map[string]interface{}{
			"foo": "bar",
		},
	},
	"group_2": map[string]interface{}{
		"vars": map[string]interface{}{},
	},
	"_meta": map[string]interface{}{
		"hostvars": map[string]interface{}{
			"host_1": map[string]interface{}{
				"ansible_host": "1.2.3.4",
				"ansible_user": "ubuntu",
			},
		},
	},
}

func TestState_basic(t *testing.T) {
	actual, err := getState("fixtures")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &expectedState, actual)

	expectedGroups := []string{"group_1", "group_2"}
	actualGroups, err := actual.GetGroups()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedGroups, actualGroups)

	expectedHosts := []string{"host_1"}
	actualHosts, err := actual.GetHostsForGroup(expectedGroups[0])
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedHosts, actualHosts)

	expectedVars := map[string]interface{}{
		"ansible_host": "1.2.3.4",
		"ansible_user": "ubuntu",
	}

	actualVars, err := actual.GetVarsForHost(expectedHosts[0])
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedVars, actualVars)

	actualInventory, err := actual.BuildInventory()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedInventory, actualInventory)
}
