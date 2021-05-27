// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceVSDataScriptSetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSVSDataScriptSetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_vsdatascriptset.testVSDataScriptSet", "name", "test-vsd-abc"),
				),
			},
		},
	})

}

const testAccAVIDSVSDataScriptSetConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_vsdatascriptset" "testVSDataScriptSet" {
	name = "test-vsd-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	ipgroup_refs = []
	pool_group_refs = []
	pool_refs = []
	string_group_refs = []
	protocol_parser_refs = []
}

data "avi_vsdatascriptset" "testVSDataScriptSet" {
    name= "${avi_vsdatascriptset.testVSDataScriptSet.name}"
}
`
