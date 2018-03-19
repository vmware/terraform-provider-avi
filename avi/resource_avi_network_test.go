package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVINetworkBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVINetworkConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINetworkConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkExists("avi_network.testnetwork"),
					resource.TestCheckResourceAttr(
						"avi_network.testnetwork", "name", "network-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkExists("avi_network.testnetwork"),
					resource.TestCheckResourceAttr(
						"avi_network.testnetwork", "name", "network-abc")),
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
			return fmt.Errorf("No Network ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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

data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}
resource "avi_network" "testnetwork" {
	name = "network-%s"
	vrf_context_ref="${data.avi_vrfcontext.global_vrf.id}"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
	exclude_discovered_subnets= false
	synced_from_se= true
    dhcp_enabled= true
	vcenter_dvs= true
}
`
