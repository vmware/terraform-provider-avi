// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceIpAddrGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIpAddrGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipaddrgroup.testIpAddrGroup", "name", "test-Internal-abc"),
				),
			},
		},
	})

}

//nolint
const testAccAVIDSIpAddrGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipaddrgroup" "testIpAddrGroup" {
	name = "test-Internal-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	prefixes {
	ip_addr {
		type = "V4"
		addr = "10.0.0.0"
	}
	mask = "8"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "192.168.0.0"
	}
	mask = "16"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "172.16.0.0"
	}
	mask = "12"
}
}

data "avi_ipaddrgroup" "testIpAddrGroup" {
    name= "${avi_ipaddrgroup.testIpAddrGroup.name}"
}
`
