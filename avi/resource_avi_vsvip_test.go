package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIVsVipBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIVsVipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIVsVipConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVsVipExists("avi_vsvip.testVsVip"),
					resource.TestCheckResourceAttr(
						"avi_vsvip.testVsVip", "name", "test-vsvip-test-abc"),
				),
			},
			{
				Config: testAccAVIVsVipupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVsVipExists("avi_vsvip.testVsVip"),
					resource.TestCheckResourceAttr(
						"avi_vsvip.testVsVip", "name", "test-vsvip-updated"),
				),
			},
			{
				ResourceName:      "avi_vsvip.testVsVip",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIVsVipConfig,
			},
		},
	})

}

func testAccCheckAVIVsVipExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI VsVip ID is set")
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

func testAccCheckAVIVsVipDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_vsvip" {
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
			return fmt.Errorf("AVI VsVip still exists")
		}
	}
	return nil
}

const testAccAVIVsVipConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_vsvip" "testVsVip" {
	name = "test-vsvip-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	vip {
	vip_id = "1"
	avi_allocated_fip = false
	auto_allocate_ip = false
	enabled = false
	auto_allocate_floating_ip = false
	avi_allocated_vip = false
	ip_address {
		type = "V4"
		addr = "1.2.3.1"
	}
}
}
`

const testAccAVIVsVipupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_vsvip" "testVsVip" {
	name = "test-vsvip-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	vip {
	vip_id = "1"
	avi_allocated_fip = false
	auto_allocate_ip = false
	enabled = false
	auto_allocate_floating_ip = false
	avi_allocated_vip = false
	ip_address {
		type = "V4"
		addr = "1.2.3.1"
	}
}
}
`
