package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

data "avi_hardwaresecuritymodulegroup" "testHardwareSecurityModuleGroup" {
    name= "${avi_hardwaresecuritymodulegroup.testHardwareSecurityModuleGroup.name}"
}
`
