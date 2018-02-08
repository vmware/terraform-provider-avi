package avi

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVITenantCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITenantConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITenantExists("avi_tenant.test_tenant"),
					resource.TestCheckResourceAttr(
						"avi_tenant.test_tenant", "name", "tenant-1")),
			},
		},
	})

}

func testAccCheckAVITenantExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Tenant ID is set")
		}
		return nil
	}
}

const testAccAVITenantConfig = `
resource "avi_tenant" "test_tenant"{
	name= "tenant-1"
}
`
