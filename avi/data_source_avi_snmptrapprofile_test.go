// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceSnmpTrapProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSSnmpTrapProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testSnmpTrapProfile", "name", "test-snmp-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSSnmpTrapProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_snmptrapprofile" "testSnmpTrapProfile" {
	name = "test-snmp-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	trap_servers {
	community = "abc"
	ip_addr {
		addr = "10.10.10.2"
		type = "V4"
	}
	port = "162"
	version = "SNMP_VER2"
}
}

data "avi_snmptrapprofile" "testSnmpTrapProfile" {
    name= "${avi_snmptrapprofile.testSnmpTrapProfile.name}"
}
`
