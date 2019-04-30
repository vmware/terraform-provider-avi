package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIVSDataScriptSetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIVSDataScriptSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIVSDataScriptSetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVSDataScriptSetExists("avi_vsdatascriptset.testvsdatascriptset"),
					resource.TestCheckResourceAttr(
						"avi_vsdatascriptset.testvsdatascriptset", "name", "vsd-test")),
			},
			{
				Config: testAccUpdatedAVIVSDataScriptSetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVSDataScriptSetExists("avi_vsdatascriptset.testvsdatascriptset"),
					resource.TestCheckResourceAttr(
						"avi_vsdatascriptset.testvsdatascriptset", "name", "vsd-abc")),
			},
		},
	})

}

func testAccCheckAVIVSDataScriptSetExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No VS DataScript Set ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIVSDataScriptSetDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_vsdatascriptset" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI VS DataScript Set still exists")
		}
	}
	return nil
}

const testAccAVIVSDataScriptSetConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_vsdatascriptset" "testvsdatascriptset" {
	name = "vsd-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
    ipgroup_refs= []
    pool_group_refs= []
    pool_refs= []
    string_group_refs= []
    protocol_parser_refs= []
}
`

const testAccUpdatedAVIVSDataScriptSetConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_vsdatascriptset" "testvsdatascriptset" {
	name = "vsd-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
    ipgroup_refs= []
    pool_group_refs= []
    pool_refs= []
    string_group_refs= []
    protocol_parser_refs= []
}
`
