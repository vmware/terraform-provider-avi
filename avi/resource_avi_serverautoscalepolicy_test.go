package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIServerAutoScalePolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIServerAutoScalePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIServerAutoScalePolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIServerAutoScalePolicyBExists("avi_serverautoscalepolicy.testserverautoscalepolicy"),
					resource.TestCheckResourceAttr(
						"avi_serverautoscalepolicy.testserverautoscalepolicy", "name", "ssp-test")),
			},
			{
				Config: testAccUpdatedAVIServerAutoScalePolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIServerAutoScalePolicyBExists("avi_serverautoscalepolicy.testserverautoscalepolicy"),
					resource.TestCheckResourceAttr(
						"avi_serverautoscalepolicy.testserverautoscalepolicy", "name", "ssp-abc")),
			},
		},
	})

}

func testAccCheckAVIServerAutoScalePolicyBExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Server Auto Scale Policy ID is set")
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

func testAccCheckAVIServerAutoScalePolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_serverautoscalepolicy" {
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
			return fmt.Errorf("AVI Server Auto Scale Policy still exists")
		}
	}
	return nil
}

const testAccAVIServerAutoScalePolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_serverautoscalepolicy" "testserverautoscalepolicy" {
	name = "ssp-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIServerAutoScalePolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_serverautoscalepolicy" "testserverautoscalepolicy" {
	name = "ssp-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
