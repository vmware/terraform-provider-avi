package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAVIDataSourceCertificateManagementProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSCertificateManagementProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_certificatemanagementprofile.testCertificateManagementProfile", "name", "test-cert-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSCertificateManagementProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_certificatemanagementprofile" "testCertificateManagementProfile" {
	script_path = "test script path"
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-cert-test-abc"
}

data "avi_certificatemanagementprofile" "testCertificateManagementProfile" {
    name= "${avi_certificatemanagementprofile.testCertificateManagementProfile.name}"
}
`
