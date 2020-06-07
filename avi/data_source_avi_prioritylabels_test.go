package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
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
	name = "test-pl-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "test priority labels"
}

data "avi_prioritylabels" "testPriorityLabels" {
    name= "${avi_prioritylabels.testPriorityLabels.name}"
}
`
