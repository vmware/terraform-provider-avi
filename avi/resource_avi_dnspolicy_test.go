package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIDnsPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIDnsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDnsPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIDnsPolicyExists("avi_dnspolicy.testDnsPolicy"),
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testDnsPolicy", "name", "test-dp-abc"),
				),
			},
			{
				Config: testAccAVIDnsPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIDnsPolicyExists("avi_dnspolicy.testDnsPolicy"),
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testDnsPolicy", "name", "test-dp-updated"),
				),
			},
			{
				ResourceName:      "avi_dnspolicy.testDnsPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIDnsPolicyConfig,
			},
		},
	})

}

func testAccCheckAVIDnsPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI DnsPolicy ID is set")
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

func testAccCheckAVIDnsPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_dnspolicy" {
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
			return fmt.Errorf("AVI DnsPolicy still exists")
		}
	}
	return nil
}

const testAccAVIDnsPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_dnspolicy" "testDnsPolicy" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-dp-abc"
	description = "test dns policy"
}
`

const testAccAVIDnsPolicyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_dnspolicy" "testDnsPolicy" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-dp-updated"
	description = "test dns policy"
}
`
