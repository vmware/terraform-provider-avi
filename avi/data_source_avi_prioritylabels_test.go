package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourcePriorityLabelsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSPriorityLabelsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_prioritylabels.testPriorityLabels", "name", "test-pl-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSPriorityLabelsConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_prioritylabels" "testPriorityLabels" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-pl-test-abc"
	description = "test priority labels"
}

data "avi_prioritylabels" "testPriorityLabels" {
    name= "${avi_prioritylabels.testPriorityLabels.name}"
}
`
