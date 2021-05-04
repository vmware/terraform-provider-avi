package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceTenantBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSTenantConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "name", "test-admin-abc"),
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "local", "true"),
				),
			},
		},
	})

}

const testAccAVIDSTenantConfig = `
resource "avi_tenant" "testTenant" {
	local = true
	name = "test-admin-abc"
}

data "avi_tenant" "testTenant" {
    name= "${avi_tenant.testTenant.name}"
}
`
