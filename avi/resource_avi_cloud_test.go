package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVICloudBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVICloudDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVICloudConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICloudExists("avi_cloud.testCloud"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "name", "test-Default-Cloud-abc"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "enable_vip_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "prefer_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "mtu", "1500"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "apic_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "dhcp_enabled", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "state_based_dns_registration", "true"),
				),
			},
			{
				Config: testAccAVICloudupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICloudExists("avi_cloud.testCloud"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "name", "test-Default-Cloud-updated"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "enable_vip_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "prefer_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "mtu", "1500"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "apic_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "dhcp_enabled", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "state_based_dns_registration", "true"),
				),
			},
			{
				ResourceName:      "avi_cloud.testCloud",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVICloudConfig,
			},
		},
	})

}

func testAccCheckAVICloudExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Cloud ID is set")
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

func testAccCheckAVICloudDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_cloud" {
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
			return fmt.Errorf("AVI Cloud still exists")
		}
	}
	return nil
}

const testAccAVICloudConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_cloud" "testCloud" {
	vtype = "CLOUD_NONE"
	license_tier = "ENTERPRISE_18"
	name = "test-Default-Cloud-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	enable_vip_static_routes = false
	prefer_static_routes = false
	license_type = "LIC_CORES"
	mtu = "1500"
	apic_mode = false
	dhcp_enabled = false
	state_based_dns_registration = true
}
`

const testAccAVICloudupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_cloud" "testCloud" {
	vtype = "CLOUD_NONE"
	license_tier = "ENTERPRISE_18"
	name = "test-Default-Cloud-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	enable_vip_static_routes = false
	prefer_static_routes = false
	license_type = "LIC_CORES"
	mtu = "1500"
	apic_mode = false
	dhcp_enabled = false
	state_based_dns_registration = true
}
`
