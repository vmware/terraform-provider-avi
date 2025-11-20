package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIAvailabilityZoneBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAvailabilityZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAvailabilityZoneConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAvailabilityZoneExists("avi_availabilityzone.testAvailabilityZone"),
					resource.TestCheckResourceAttr(
						"avi_availabilityzone.testAvailabilityZone", "name", "az-1"),
				),
			},
			{
				Config: testAccAVIAvailabilityZoneupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAvailabilityZoneExists("avi_availabilityzone.testAvailabilityZone"),
					resource.TestCheckResourceAttr(
						"avi_availabilityzone.testAvailabilityZone", "name", "az-1"),
				),
			},
			{
				ResourceName:      "avi_availabilityzone.testAvailabilityZone",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIAvailabilityZoneConfig,
			},
		},
	})

}

func testAccCheckAVIAvailabilityZoneExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI AvailabilityZone ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIAvailabilityZoneDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_availabilityzone" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI AvailabilityZone still exists")
		}
	}
	return nil
}

const testAccAVIAvailabilityZoneConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_availabilityzone" "testAvailabilityZone" {
	name = "az-1"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
}
`

const testAccAVIAvailabilityZoneupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_availabilityzone" "testAvailabilityZone" {
	name = "az-1"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
}
`
