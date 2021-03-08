package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVINetworkBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINetworkConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkExists("avi_network.testNetwork"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "name", "test-network-test-abc"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "exclude_discovered_subnets", "false"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "synced_from_se", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "dhcp_enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "vcenter_dvs", "true"),
				),
			},
			{
				Config: testAccAVINetworkupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkExists("avi_network.testNetwork"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "name", "test-network-updated"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "exclude_discovered_subnets", "false"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "synced_from_se", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "dhcp_enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "vcenter_dvs", "true"),
				),
			},
			{
				ResourceName:      "avi_network.testNetwork",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVINetworkConfig,
			},
		},
	})

}

func testAccCheckAVINetworkExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Network ID is set")
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

func testAccCheckAVINetworkDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_network" {
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
			return fmt.Errorf("AVI Network still exists")
		}
	}
	return nil
}

const testAccAVINetworkConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_network" "testNetwork" {
	name = "test-network-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	exclude_discovered_subnets = false
	synced_from_se = true
	dhcp_enabled = true
	vcenter_dvs = true
}
`

const testAccAVINetworkupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_network" "testNetwork" {
	name = "test-network-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	exclude_discovered_subnets = false
	synced_from_se = true
	dhcp_enabled = true
	vcenter_dvs = true
}
`
