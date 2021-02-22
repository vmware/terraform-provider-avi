package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	name = "test-pgpp-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_poolgroupdeploymentpolicy" "testPoolGroupDeploymentPolicy" {
    name= "${avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy.name}"
}
`
