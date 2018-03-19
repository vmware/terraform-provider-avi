package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIDNSPolicyBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIDNSPolicyConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIDNSPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDNSPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIDNSPolicyExists("avi_dnspolicy.testdnspolicy"),
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testdnspolicy", "name", "dp-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIDNSPolicyExists("avi_dnspolicy.testdnspolicy"),
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testdnspolicy", "name", "dp-abc")),
			},
		},
	})

}

func testAccCheckAVIDNSPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No DNS Policy ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIDNSPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_dnspolicy" {
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
			return fmt.Errorf("AVI DNS Policy still exists")
		}
	}
	return nil
}

const testAccAVIDNSPolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_dnspolicy" "testdnspolicy" {
	name = "dp-%s"
	description = "test dns policy"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
