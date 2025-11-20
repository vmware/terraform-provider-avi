// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceSecurityPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_securitypolicy.testSecurityPolicy", "name", "my-security-policy"),
				),
			},
		},
	})

}

const testAccAVIDSSecurityPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_securitypolicy" "testSecurityPolicy" {
	name = "my-security-policy"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "Security policy for L7 rules"
}

data "avi_securitypolicy" "testSecurityPolicy" {
    name= "${avi_securitypolicy.testSecurityPolicy.name}"
}
`
