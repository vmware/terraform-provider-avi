package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIAlertScriptConfigBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIAlertScriptConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAlertScriptConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAlertScriptConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertScriptConfigExists("avi_alertscriptconfig.testalertscriptconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertscriptconfig.testalertscriptconfig", "name", "asc-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertScriptConfigExists("avi_alertscriptconfig.testalertscriptconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertscriptconfig.testalertscriptconfig", "name", "asc-abc")),
			},
		},
	})

}

func testAccCheckAVIAlertScriptConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Alert Script Config ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIAlertScriptConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_alertscriptconfig" {
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
			return fmt.Errorf("AVI Alert Script Config still exists")
		}
	}
	return nil
}

const testAccAVIAlertScriptConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertscriptconfig" "testalertscriptconfig" {
	name = "asc-%s"
	action_script= "test script"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
