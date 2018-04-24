package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVISNMPTrapProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISNMPTrapProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISNMPTrapProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISNMPTrapProfileExists("avi_snmptrapprofile.testsnmptrapprofile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testsnmptrapprofile", "name", "snmp-test")),
			},
			{
				Config: testAccUpdatedAVISNMPTrapProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISNMPTrapProfileExists("avi_snmptrapprofile.testsnmptrapprofile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testsnmptrapprofile", "name", "snmp-abc")),
			},
		},
	})

}

func testAccCheckAVISNMPTrapProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No SNMP Trap Profile ID is set")
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

func testAccCheckAVISNMPTrapProfileDestroy(s *terraform.State) error {
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
			return fmt.Errorf("AVI SNMP Trap Profile still exists")
		}
	}
	return nil
}

const testAccAVISNMPTrapProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_snmptrapprofile" "testsnmptrapprofile" {
	name = "snmp-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVISNMPTrapProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_snmptrapprofile" "testsnmptrapprofile" {
	name = "snmp-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
