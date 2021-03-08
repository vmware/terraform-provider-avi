package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIPoolGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPoolGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupExists("avi_poolgroup.testPoolGroup"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "name", "pg-test-abc"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "implicit_priority_labels", "false"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "min_servers", "0"),
				),
			},
			{
				Config: testAccAVIPoolGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupExists("avi_poolgroup.testPoolGroup"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "name", "pg-updated"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "implicit_priority_labels", "false"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "min_servers", "0"),
				),
			},
			{
				ResourceName:      "avi_poolgroup.testPoolGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIPoolGroupConfig,
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
			return fmt.Errorf("No AVI PoolGroup ID is set")
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

func testAccCheckAVIPoolGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_poolgroup" {
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
			return fmt.Errorf("AVI PoolGroup still exists")
		}
	}
	return nil
}

const testAccAVIPoolGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroup" "testPoolGroup" {
	name = "pg-test-abc"
	implicit_priority_labels = false
	min_servers = "0"
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVIPoolGroupupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroup" "testPoolGroup" {
	name = "pg-updated"
	implicit_priority_labels = false
	min_servers = "0"
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
