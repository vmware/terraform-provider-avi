package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceHardwareSecurityModuleGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSHardwareSecurityModuleGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup", "name", "test-hsmg-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSHardwareSecurityModuleGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_hardwaresecuritymodulegroup" "testHardwareSecurityModuleGroup" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-hsmg-test-abc"
	hsm {
		type = "HSM_TYPE_THALES_NETHSM"
		nethsm {
			remote_port = "9005"
			priority = "100"
			keyhash = "198644ebcba88ba1421ae0c34cdd541edf01deb8"
			module_id = "0"
			esn = "580A-F79E-BCD9"
			remote_ip {
				type = "V4"
				addr = "10.10.15.1"
			}
		}
	}

}

data "avi_hardwaresecuritymodulegroup" "testHardwareSecurityModuleGroup" {
    name= avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup.name
}
`
