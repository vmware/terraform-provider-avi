package avi

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVICloudCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVICloudConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICloudExists("avi_cloud.test_cloud"),
					resource.TestCheckResourceAttr(
						"avi_cloud.test_cloud", "name", "Default-Cloud-1")),
			},
		},
	})

}

func testAccCheckAVICloudExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Cloud ID is set")
		}
		return nil
	}

}

const testAccAVICloudConfig = `
resource "avi_cloud" "test_cloud"{
	name = "Default-Cloud-1"
	vtype= "CLOUD_NONE"
}
`
