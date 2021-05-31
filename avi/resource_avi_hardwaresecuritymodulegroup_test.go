package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIHardwareSecurityModuleGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIHardwareSecurityModuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIHardwareSecurityModuleGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHardwareSecurityModuleGroupExists("avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup"),
					resource.TestCheckResourceAttr(
						"avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup", "name", "test-hsmg-test-abc"),
				),
			},
			{
				Config: testAccAVIHardwareSecurityModuleGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHardwareSecurityModuleGroupExists("avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup"),
					resource.TestCheckResourceAttr(
						"avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup", "name", "test-hsmg-updated"),
				),
			},
			{
				ResourceName:      "avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIHardwareSecurityModuleGroupConfig,
			},
		},
	})

}

func testAccCheckAVIHardwareSecurityModuleGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI HardwareSecurityModuleGroup ID is set")
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

func testAccCheckAVIHardwareSecurityModuleGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_hardwaresecuritymodulegroup" {
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
			return fmt.Errorf("AVI HardwareSecurityModuleGroup still exists")
		}
	}
	return nil
}

const testAccAVIHardwareSecurityModuleGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_hardwaresecuritymodulegroup" "testHardwareSecurityModuleGroup" {
	name = "test-hsmg-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	hsm {
		type = "HSM_TYPE_THALES_NETHSM"
		nethsm {
			remote_ip {
				addr = "10.10.15.1"
				type = "V4"
			}
			remote_port = "9005"
			esn = "580A-F79E-BCD9"
			priority = "100"
			module_id = "0"
			keyhash = "198644ebcba88ba1421ae0c34cdd541edf01deb8"
		}
	}
}
`

const testAccAVIHardwareSecurityModuleGroupupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_hardwaresecuritymodulegroup" "testHardwareSecurityModuleGroup" {
	name = "test-hsmg-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	hsm {
		type = "HSM_TYPE_THALES_NETHSM"
		nethsm {
			remote_ip {
				addr = "10.10.15.1"
				type = "V4"
			}
			remote_port = "9005"
			esn = "580A-F79E-BCD9"
			priority = "100"
			module_id = "0"
			keyhash = "198644ebcba88ba1421ae0c34cdd541edf01deb8"
		}
	}
}
`
