package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIErrorPageProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIErrorPageProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIErrorPageProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIErrorPageProfileExists("avi_errorpageprofile.testErrorPageProfile"),
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testErrorPageProfile", "name", "test-epp-abc"),
				),
			},
			{
				Config: testAccAVIErrorPageProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIErrorPageProfileExists("avi_errorpageprofile.testErrorPageProfile"),
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testErrorPageProfile", "name", "test-epp-updated"),
				),
			},
		},
	})

}

func testAccCheckAVIErrorPageProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ErrorPageProfile ID is set")
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

func testAccCheckAVIErrorPageProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_errorpageprofile" {
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
			return fmt.Errorf("AVI ErrorPageProfile still exists")
		}
	}
	return nil
}

const testAccAVIErrorPageProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpageprofile" "testErrorPageProfile" {
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "test-epp-abc"
}
`

const testAccAVIErrorPageProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpageprofile" "testErrorPageProfile" {
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "test-epp-updated"
}
`
