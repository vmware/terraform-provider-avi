// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceDnsPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSDnsPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testDnsPolicy", "name", "test-dp-abc"),
				),
			},
		},
	})

}

//nolint
const testAccAVIDSDnsPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_dnspolicy" "testDnsPolicy" {
	name = "test-dp-abc"
	description = "test dns policy"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_dnspolicy" "testDnsPolicy" {
    name= "${avi_dnspolicy.testDnsPolicy.name}"
}
`
