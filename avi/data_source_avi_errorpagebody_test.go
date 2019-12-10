package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
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
	error_page_body = "<!DOCTYPE html><html><head></head><body><div><p> Please contact our technical support</p></div></body></html>"
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-Custom-Error-Page-abc"
}

data "avi_errorpagebody" "testErrorPageBody" {
    name= "${avi_errorpagebody.testErrorPageBody.name}"
}
`
