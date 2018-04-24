package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIVSVipBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIVSVipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIVSVipConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVSVipExists("avi_vsvip.testvsvip"),
					resource.TestCheckResourceAttr(
						"avi_vsvip.testvsvip", "name", "vsvip-test")),
			},
			{
				Config: testAccUpdatedAVIVSVipConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVSVipExists("avi_vsvip.testvsvip"),
					resource.TestCheckResourceAttr(
						"avi_vsvip.testvsvip", "name", "vsvip-abc")),
			},
		},
	})

}

func testAccCheckAVIVSVipExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No VSVip ID is set")
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

func testAccCheckAVIVSVipDestroy(s *terraform.State) error {
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
			return fmt.Errorf("AVI VSVip still exists")
		}
	}
	return nil
}

const testAccAVIVSVipConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}

resource "avi_vsvip" "testvsvip" {
	name = "vsvip-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
	vrf_context_ref= "${data.avi_vrfcontext.global_vrf.id}"
	vip= [{
		vip_id= "1"
		avi_allocated_fip= false
		auto_allocate_ip= false
		enabled= false
		auto_allocate_floating_ip= false
		avi_allocated_vip= false
		ip_address= {
			type= "V4"
			addr= "1.2.3.1"
		}
	}]
}
`

const testAccUpdatedAVIVSVipConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}

resource "avi_vsvip" "testvsvip" {
	name = "vsvip-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
	vrf_context_ref= "${data.avi_vrfcontext.global_vrf.id}"
	vip= [{
		vip_id= "1"
		avi_allocated_fip= false
		auto_allocate_ip= false
		enabled= false
		auto_allocate_floating_ip= false
		avi_allocated_vip= false
		ip_address= {
			type= "V4"
			addr= "1.2.3.1"
		}
	}]
}
`
