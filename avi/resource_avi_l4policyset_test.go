package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIL4PolicySetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIL4PolicySetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIL4PolicySetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIL4PolicySetExists("avi_l4policyset.testL4PolicySet"),
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "name", "L4Policy-set"),
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "is_internal_policy", "false"),
				),
			},
			{
				Config: testAccAVIL4PolicySetupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIL4PolicySetExists("avi_l4policyset.testL4PolicySet"),
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "name", "L4Policy-set-updated"),
					resource.TestCheckResourceAttr(
						"avi_l4policyset.testL4PolicySet", "is_internal_policy", "false"),
				),
			},
			{
				ResourceName:      "avi_l4policyset.testL4PolicySet",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIL4PolicySetConfig,
			},
		},
	})

}

func testAccCheckAVIL4PolicySetExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI L4PolicySet ID is set")
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

func testAccCheckAVIL4PolicySetDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_l4policyset" {
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
			return fmt.Errorf("AVI L4PolicySet still exists")
		}
	}
	return nil
}

const testAccAVIL4PolicySetConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_l4policyset" "testL4PolicySet" {
	is_internal_policy = false
	l4_connection_policy {
	}
	name = "L4Policy-set"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVIL4PolicySetupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_l4policyset" "testL4PolicySet" {
	is_internal_policy = false
	l4_connection_policy {
	}
	name = "L4Policy-set-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
