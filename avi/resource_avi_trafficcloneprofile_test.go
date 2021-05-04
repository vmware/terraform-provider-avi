package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVITrafficCloneProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVITrafficCloneProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITrafficCloneProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITrafficCloneProfileExists("avi_trafficcloneprofile.testTrafficCloneProfile"),
					resource.TestCheckResourceAttr(
						"avi_trafficcloneprofile.testTrafficCloneProfile", "name", "test-tp-test-abc"),
				),
			},
			{
				Config: testAccAVITrafficCloneProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITrafficCloneProfileExists("avi_trafficcloneprofile.testTrafficCloneProfile"),
					resource.TestCheckResourceAttr(
						"avi_trafficcloneprofile.testTrafficCloneProfile", "name", "test-tp-updated"),
				),
			},
			{
				ResourceName:      "avi_trafficcloneprofile.testTrafficCloneProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVITrafficCloneProfileConfig,
			},
		},
	})

}

func testAccCheckAVITrafficCloneProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI TrafficCloneProfile ID is set")
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

func testAccCheckAVITrafficCloneProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_trafficcloneprofile" {
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
			return fmt.Errorf("AVI TrafficCloneProfile still exists")
		}
	}
	return nil
}

const testAccAVITrafficCloneProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_trafficcloneprofile" "testTrafficCloneProfile" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-tp-test-abc"
}
`

const testAccAVITrafficCloneProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_trafficcloneprofile" "testTrafficCloneProfile" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-tp-updated"
}
`
