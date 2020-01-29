package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceMicroServiceGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSMicroServiceGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_microservicegroup.testMicroServiceGroup", "name", "msg-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSMicroServiceGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_microservicegroup" "testMicroServiceGroup" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "msg-test-abc"
	service_refs = []
}

data "avi_microservicegroup" "testMicroServiceGroup" {
    name= "${avi_microservicegroup.testMicroServiceGroup.name}"
}
`
