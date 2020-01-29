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
						"avi_poolgroup.testPoolGroup", "min_servers", "0"),
					resource.TestCheckResourceAttr(
						"avi_poolgroup.testPoolGroup", "implicit_priority_labels", "false"),
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
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
	min_servers = "0"
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "pg-test-abc"
	implicit_priority_labels = false
}

data "avi_poolgroup" "testPoolGroup" {
    name= "${avi_poolgroup.testPoolGroup.name}"
}
`
