// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceNetworkSecurityPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSNetworkSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testNetworkSecurityPolicy", "name", "networkw-security-policy"),
				),
			},
		},
	})

}

const testAccAVIDSNetworkSecurityPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
	rules {
	log = false
	rl_param {
		max_rate = "1000"
		burst_size = "2000"
	}
	age = "0"
	match {
		vs_port {
			match_criteria = "IS_IN"
			ports = ["9090"]
		}
	}
	action = "NETWORK_SECURITY_POLICY_ACTION_TYPE_ALLOW"
	name = "networkw-security-policy-rule"
	index = "0"
	enable = false
}
	name = "networkw-security-policy"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
    name= "${avi_networksecuritypolicy.testNetworkSecurityPolicy.name}"
}
`
