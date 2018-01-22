package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIPoolBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIPoolConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testpool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testpool", "name", "p42-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testpool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testpool", "name", "p42-abc")),
			},
		},
	})

}

func testAccCheckAVIPoolExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI POOL ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIPoolDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_pool" {
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
			return fmt.Errorf("AVI Pool still exists")
		}
	}
	return nil
}

const testAccAVIPoolConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
  
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}
resource "avi_pool" "testpool" {
	name = "p42-%s"
	server_count = 1
	servers {
		hostname= "10.90.64.66"
		ip= {
		  type= "V4"
		  addr= "10.90.64.66"
		}
		port= 8080
	}
	fail_action= {
		type= "FAIL_ACTION_CLOSE_CONN"
	}
	vrf_ref="${data.avi_vrfcontext.global_vrf.id}"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}
`
