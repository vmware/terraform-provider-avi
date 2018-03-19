package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVICustomipamdnsProfileBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVICustomipamdnsProfileConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVICustomipamdnsProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVICustomipamdnsProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICustomipamdnsProfileExists("avi_customipamdnsprofile.testipam"),
					resource.TestCheckResourceAttr(
						"avi_customipamdnsprofile.testipam", "name", "ipam-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICustomipamdnsProfileExists("avi_customipamdnsprofile.testipam"),
					resource.TestCheckResourceAttr(
						"avi_customipamdnsprofile.testipam", "name", "ipam-abc")),
			},
		},
	})

}

func testAccCheckAVICustomipamdnsProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Custom ipamdns Profile ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVICustomipamdnsProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_customipamdnsprofile" {
			continue
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Custom ipamdns Profile still exists")
		}
	}
	return nil
}

const testAccAVICustomipamdnsProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_customipamdnsprofile" "testipam" {
	name = "ipam-%s"
	script_uri = "/"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
