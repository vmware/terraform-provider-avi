// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceL4PolicySetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSL4PolicySetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "name", "L4Policy-set"),
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "is_internal_policy", "false"),
				),
			},
		},
	})

}

const testAccAVIDSL4PolicySetConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_l4policyset" "testL4PolicySet" {
	is_internal_policy = false
	l4_connection_policy {
	}
	name = "L4Policy-set"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_l4policyset" "testL4PolicySet" {
    name= "${avi_l4policyset.testL4PolicySet.name}"
}
`
