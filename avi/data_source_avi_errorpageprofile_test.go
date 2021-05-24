// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceErrorPageProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSErrorPageProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testErrorPageProfile", "name", "test-epp-abc"),
				),
			},
		},
	})

}

const testAccAVIDSErrorPageProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpageprofile" "testErrorPageProfile" {
	name = "test-epp-abc"
	error_pages {
	index = "1"
	enable = false
	match {
		match_criteria = "IS_IN"
		status_codes = ["400","404"]
	}
}
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_errorpageprofile" "testErrorPageProfile" {
    name= "${avi_errorpageprofile.testErrorPageProfile.name}"
}
`
