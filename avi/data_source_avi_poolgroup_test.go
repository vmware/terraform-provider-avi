package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourcePoolGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSPoolGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "name", "pg-test-abc"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "implicit_priority_labels", "false"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "min_servers", "0"),
				),
			},
		},
	})

}

const testAccAVIDSPoolGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroup" "testPoolGroup" {
	name = "pg-test-abc"
	implicit_priority_labels = false
	min_servers = "0"
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_poolgroup" "testPoolGroup" {
    name= "${avi_poolgroup.testPoolGroup.name}"
}
`
