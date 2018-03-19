package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIPoolGroupBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIPoolGroupConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPoolGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupExists("avi_poolgroup.testpoolgroup"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testpoolgroup", "name", "pg-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupExists("avi_poolgroup.testpoolgroup"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testpoolgroup", "name", "pg-abc")),
			},
		},
	})

}

func testAccCheckAVIPoolGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI POOL Group ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIPoolGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_poolgroup" {
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
			return fmt.Errorf("AVI Pool Group still exists")
		}
	}
	return nil
}

const testAccAVIPoolGroupConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_poolgroup" "testpoolgroup" {
	name = "pg-%s"
	implicit_priority_labels= false
	min_servers= 0
	fail_action= {
		type= "FAIL_ACTION_CLOSE_CONN"
	}
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}
`
