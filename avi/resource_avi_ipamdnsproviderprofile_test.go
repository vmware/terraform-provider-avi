package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIIpamDnsProviderProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIIpamDnsProviderProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIIpamDnsProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIpamDnsProviderProfileExists("avi_ipamdnsproviderprofile.testIpamDnsProviderProfile"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "name", "test-ipam-abc-abc"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "allocate_ip_in_vrf", "false"),
				),
			},
			{
				Config: testAccAVIIpamDnsProviderProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIpamDnsProviderProfileExists("avi_ipamdnsproviderprofile.testIpamDnsProviderProfile"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "name", "test-ipam-updated"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "allocate_ip_in_vrf", "false"),
				),
			},
			{
				ResourceName:      "avi_ipamdnsproviderprofile.testIpamDnsProviderProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIIpamDnsProviderProfileConfig,
			},
		},
	})

}

// nolint
func testAccCheckAVIIpamDnsProviderProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI IpamDnsProviderProfile ID is set")
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

// nolint
func testAccCheckAVIIpamDnsProviderProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_ipamdnsproviderprofile" {
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
			return fmt.Errorf("AVI IpamDnsProviderProfile still exists")
		}
	}
	return nil
}

// nolint
const testAccAVIIpamDnsProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
	name = "test-ipam-abc-abc"
	allocate_ip_in_vrf = false
	internal_profile {
		ttl = "31"
	}
	type = "IPAMDNS_TYPE_INTERNAL"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

// nolint
const testAccAVIIpamDnsProviderProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
	name = "test-ipam-updated"
	allocate_ip_in_vrf = false
	internal_profile {
		ttl = "31"
	}
	type = "IPAMDNS_TYPE_INTERNAL"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
