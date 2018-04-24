package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
					testAccCheckAVIErrorPageProfileExists("avi_errorpageprofile.testerrorpageprofile"),
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testerrorpageprofile", "name", "epp-test")),
			},
			{
				Config: testAccUpdatedAVIErrorPageProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIErrorPageProfileExists("avi_errorpageprofile.testerrorpageprofile"),
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testerrorpageprofile", "name", "epp-abc")),
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
			return fmt.Errorf("No Error Page Profile ID is set")
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
			return fmt.Errorf("AVI Error Page Profile still exists")
		}
	}
	return nil
}

const testAccAVIErrorPageProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_errorpageprofile" "testerrorpageprofile" {
	name = "epp-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIErrorPageProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_errorpageprofile" "testerrorpageprofile" {
	name = "epp-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
