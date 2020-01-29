package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAVIDataSourcePoolGroupDeploymentPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSPoolGroupDeploymentPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy", "name", "test-pgpp-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSPoolGroupDeploymentPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroupdeploymentpolicy" "testPoolGroupDeploymentPolicy" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-pgpp-test-abc"
}

data "avi_poolgroupdeploymentpolicy" "testPoolGroupDeploymentPolicy" {
    name= "${avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy.name}"
}
`
