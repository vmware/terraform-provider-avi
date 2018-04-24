package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVINetworkSecuritypolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINetworkSecuritypolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINetworkSecuritypolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkSecuritypolicyExists("avi_networksecuritypolicy.testnetworksecuritypolicy"),
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testnetworksecuritypolicy", "name", "ns-test")),
			},
			{
				Config: testAccUpdatedAVINetworkSecuritypolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkSecuritypolicyExists("avi_networksecuritypolicy.testnetworksecuritypolicy"),
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testnetworksecuritypolicy", "name", "ns-abc")),
			},
		},
	})

}

func testAccCheckAVINetworkSecuritypolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network Security policy ID is set")
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

func testAccCheckAVINetworkSecuritypolicyDestroy(s *terraform.State) error {
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
			return fmt.Errorf("AVI Network Security policy still exists")
		}
	}
	return nil
}

const testAccAVINetworkSecuritypolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_networksecuritypolicy" "testnetworksecuritypolicy" {
	name = "ns-test"
	description= "test network policy"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVINetworkSecuritypolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_networksecuritypolicy" "testnetworksecuritypolicy" {
	name = "ns-abc"
	description= "test network policy"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
