package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIMicroServiceGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIMicroServiceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIMicroServiceGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIMicroServiceGroupExists("avi_microservicegroup.testmicroservicegroup"),
					resource.TestCheckResourceAttr(
						"avi_microservicegroup.testmicroservicegroup", "name", "msg-test")),
			},
			{
				Config: testAccUpdatedAVIMicroServiceGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIMicroServiceGroupExists("avi_microservicegroup.testmicroservicegroup"),
					resource.TestCheckResourceAttr(
						"avi_microservicegroup.testmicroservicegroup", "name", "msg-abc")),
			},
		},
	})

}

func testAccCheckAVIMicroServiceGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Micro Service Group ID is set")
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

func testAccCheckAVIMicroServiceGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_microservicegroup" {
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
			return fmt.Errorf("AVI Micro Service Group still exists")
		}
	}
	return nil
}

const testAccAVIMicroServiceGroupConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_microservicegroup" "testmicroservicegroup" {
	name = "msg-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIMicroServiceGroupConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_microservicegroup" "testmicroservicegroup" {
	name = "msg-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
