package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVISecurityPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISecurityPolicyExists("avi_securitypolicy.testSecurityPolicy"),
					resource.TestCheckResourceAttr(
						"avi_securitypolicy.testSecurityPolicy", "name", "my-security-policy"),
				),
			},
			{
				Config: testAccAVISecurityPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISecurityPolicyExists("avi_securitypolicy.testSecurityPolicy"),
					resource.TestCheckResourceAttr(
						"avi_securitypolicy.testSecurityPolicy", "name", "my-security-policy-updated"),
				),
			},
			{
				ResourceName:      "avi_securitypolicy.testSecurityPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVISecurityPolicyConfig,
			},
		},
	})

}

func testAccCheckAVISecurityPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI SecurityPolicy ID is set")
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

func testAccCheckAVISecurityPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_securitypolicy" {
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
			return fmt.Errorf("AVI SecurityPolicy still exists")
		}
	}
	return nil
}

const testAccAVISecurityPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_securitypolicy" "testSecurityPolicy" {
	name = "my-security-policy"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "Security policy for L7 rules"
}
`

const testAccAVISecurityPolicyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_securitypolicy" "testSecurityPolicy" {
	name = "my-security-policy-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "Security policy for L7 rules"
}
`
