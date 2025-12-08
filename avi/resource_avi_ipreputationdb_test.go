package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIIPReputationDBBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIIPReputationDBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIIPReputationDBConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIPReputationDBExists("avi_ipreputationdb.testIPReputationDB"),
					resource.TestCheckResourceAttr(
						"avi_ipreputationdb.testIPReputationDB", "name", "ip-reputation-db"),
				),
			},
			{
				Config: testAccAVIIPReputationDBupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIPReputationDBExists("avi_ipreputationdb.testIPReputationDB"),
					resource.TestCheckResourceAttr(
						"avi_ipreputationdb.testIPReputationDB", "name", "ip-reputation-db-updated"),
				),
			},
			{
				ResourceName:      "avi_ipreputationdb.testIPReputationDB",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIIPReputationDBConfig,
			},
		},
	})

}

func testAccCheckAVIIPReputationDBExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI IPReputationDB ID is set")
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

func testAccCheckAVIIPReputationDBDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_ipreputationdb" {
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
			return fmt.Errorf("AVI IPReputationDB still exists")
		}
	}
	return nil
}

const testAccAVIIPReputationDBConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipreputationdb" "testIPReputationDB" {
	name = "ip-reputation-db"
	tenant_ref = data.avi_tenant.default_tenant.id
	vendor = "IP_REPUTATION_VENDOR_WEBROOT"
	description = "IP reputation database"
}
`

const testAccAVIIPReputationDBupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipreputationdb" "testIPReputationDB" {
	name = "ip-reputation-db-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	vendor = "IP_REPUTATION_VENDOR_WEBROOT"
	description = "IP reputation database"
}
`
