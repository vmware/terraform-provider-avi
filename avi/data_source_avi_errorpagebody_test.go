// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceErrorPageBodyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSErrorPageBodyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_errorpagebody.testErrorPageBody", "name", "test-Custom-Error-Page-abc"),
				),
			},
		},
	})

}

const testAccAVIDSErrorPageBodyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpagebody" "testErrorPageBody" {
	name = "test-Custom-Error-Page-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	error_page_body = "<!DOCTYPE html><html><head></head><body><div><p> Please contact our technical support</p></div></body></html>"
}

data "avi_errorpagebody" "testErrorPageBody" {
    name= "${avi_errorpagebody.testErrorPageBody.name}"
}
`
