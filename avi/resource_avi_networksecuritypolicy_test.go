package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVINetworkSecurityPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINetworkSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINetworkSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkSecurityPolicyExists("avi_networksecuritypolicy.testNetworkSecurityPolicy"),
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testNetworkSecurityPolicy", "name", "ns-abc-abc"),
				),
			},
			{
				Config: testAccAVINetworkSecurityPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkSecurityPolicyExists("avi_networksecuritypolicy.testNetworkSecurityPolicy"),
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testNetworkSecurityPolicy", "name", "ns-updated"),
				),
			},
			{
				ResourceName:      "avi_networksecuritypolicy.testNetworkSecurityPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVINetworkSecurityPolicyConfig,
			},
		},
	})

}

func testAccCheckAVINetworkSecurityPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI NetworkSecurityPolicy ID is set")
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

func testAccCheckAVINetworkSecurityPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_networksecuritypolicy" {
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
			return fmt.Errorf("AVI NetworkSecurityPolicy still exists")
		}
	}
	return nil
}

const testAccAVINetworkSecurityPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "ns-abc-abc"
	"description" = "test network policy"
}
`

const testAccAVINetworkSecurityPolicyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "ns-updated"
	"description" = "test network policy"
}
`
