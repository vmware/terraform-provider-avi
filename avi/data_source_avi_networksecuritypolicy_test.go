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
						"avi_networksecuritypolicy.testNetworkSecurityPolicy", "name", "ns-abc-abc"),
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
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "ns-abc-abc"
	description = "test network policy"
}

data "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
    name= "${avi_networksecuritypolicy.testNetworkSecurityPolicy.name}"
}
`
