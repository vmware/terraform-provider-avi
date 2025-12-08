// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceNatPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSNatPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_natpolicy.testNatPolicy", "name", "NAT-Policy"),
				),
			},
		},
	})

}

const testAccAVIDSNatPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_natpolicy" "testNatPolicy" {
	rules {
	enable = true
	match {
		services {
			destination_port {
				match_criteria = "IS_IN"
				ports = ["9000"]
			}
			source_port {
				match_criteria = "IS_IN"
				ports = ["8000"]
			}
			protocol {
				match_criteria = "IS_IN"
				protocol = "PROTOCOL_ICMP"
			}
		}
	}
	action {
		nat_info {
	nat_ip {
		addr = "192.168.10.10"
		type = "V4"
	}
	nat_ip_range {
		begin {
			addr = "192.168.10.5"
			type = "V4"
		}
		end {
			addr = "192.168.10.15"
			type = "V4"
		}
	}
}
		type = "NAT_POLICY_ACTION_TYPE_DYNAMIC_IP_PORT"
	}
	name = "natrule"
	index = "0"
}
	name = "NAT-Policy"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_natpolicy" "testNatPolicy" {
    name= "${avi_natpolicy.testNatPolicy.name}"
}
`
