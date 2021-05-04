package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVISnmpTrapProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISnmpTrapProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISnmpTrapProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISnmpTrapProfileExists("avi_snmptrapprofile.testSnmpTrapProfile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testSnmpTrapProfile", "name", "test-snmp-test-abc"),
				),
			},
			{
				Config: testAccAVISnmpTrapProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISnmpTrapProfileExists("avi_snmptrapprofile.testSnmpTrapProfile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testSnmpTrapProfile", "name", "test-snmp-updated"),
				),
			},
			{
				ResourceName:      "avi_snmptrapprofile.testSnmpTrapProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVISnmpTrapProfileConfig,
			},
		},
	})

}

func testAccCheckAVISnmpTrapProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI SnmpTrapProfile ID is set")
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

func testAccCheckAVISnmpTrapProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_snmptrapprofile" {
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
			return fmt.Errorf("AVI SnmpTrapProfile still exists")
		}
	}
	return nil
}

const testAccAVISnmpTrapProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_snmptrapprofile" "testSnmpTrapProfile" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-snmp-test-abc"
}
`

const testAccAVISnmpTrapProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_snmptrapprofile" "testSnmpTrapProfile" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-snmp-updated"
}
`
