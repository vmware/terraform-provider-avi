package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceVSDataScriptSetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSVSDataScriptSetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_vsdatascriptset.testVSDataScriptSet", "name", "test-vsd-abc"),
				),
			},
		},
	})

}

const testAccAVIDSVSDataScriptSetConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_vsdatascriptset" "testVSDataScriptSet" {
	name = "test-vsd-abc"
	pool_refs = []
	tenant_ref = data.avi_tenant.default_tenant.id
	protocol_parser_refs = []
	pool_group_refs = []
	string_group_refs = []
	ipgroup_refs = []
}

data "avi_vsdatascriptset" "testVSDataScriptSet" {
    name= "${avi_vsdatascriptset.testVSDataScriptSet.name}"
}
`
