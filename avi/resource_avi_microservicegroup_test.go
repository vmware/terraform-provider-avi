package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
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
					testAccCheckAVIMicroServiceGroupExists("avi_microservicegroup.testMicroServiceGroup"),
					resource.TestCheckResourceAttr(
						"avi_microservicegroup.testMicroServiceGroup", "name", "msg-test-abc"),
				),
			},
			{
				Config: testAccAVIMicroServiceGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIMicroServiceGroupExists("avi_microservicegroup.testMicroServiceGroup"),
					resource.TestCheckResourceAttr(
						"avi_microservicegroup.testMicroServiceGroup", "name", "msg-updated"),
				),
			},
			{
				ResourceName:      "avi_microservicegroup.testMicroServiceGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIMicroServiceGroupConfig,
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
			return fmt.Errorf("No AVI MicroServiceGroup ID is set")
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
			return fmt.Errorf("AVI MicroServiceGroup still exists")
		}
	}
	return nil
}

const testAccAVIMicroServiceGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_microservicegroup" "testMicroServiceGroup" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "msg-test-abc"
	service_refs = []
}
`

const testAccAVIMicroServiceGroupupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_microservicegroup" "testMicroServiceGroup" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "msg-updated"
	service_refs = []
}
`
