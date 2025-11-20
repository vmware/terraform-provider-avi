// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceHTTPPolicySetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSHTTPPolicySetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_httppolicyset.testHTTPPolicySet", "name", "test-http-policyset"),
					resource.TestCheckResourceAttr(
						"avi_httppolicyset.testHTTPPolicySet", "is_internal_policy", "false"),
				),
			},
		},
	})

}

const testAccAVIDSHTTPPolicySetConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_httppolicyset" "testHTTPPolicySet" {
	is_internal_policy = false
	name = "test-http-policyset"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_httppolicyset" "testHTTPPolicySet" {
    name= "${avi_httppolicyset.testHTTPPolicySet.name}"
}
`
