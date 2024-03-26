package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIIpAddrGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIIpAddrGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIIpAddrGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIpAddrGroupExists("avi_ipaddrgroup.testIpAddrGroup"),
					resource.TestCheckResourceAttr(
						"avi_ipaddrgroup.testIpAddrGroup", "name", "test-Internal-abc"),
				),
			},
			{
				Config: testAccAVIIpAddrGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIpAddrGroupExists("avi_ipaddrgroup.testIpAddrGroup"),
					resource.TestCheckResourceAttr(
						"avi_ipaddrgroup.testIpAddrGroup", "name", "test-Internal-updated"),
				),
			},
			{
				ResourceName:      "avi_ipaddrgroup.testIpAddrGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIIpAddrGroupConfig,
			},
		},
	})

}

//nolint
func testAccCheckAVIIpAddrGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI IpAddrGroup ID is set")
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

//nolint
func testAccCheckAVIIpAddrGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_ipaddrgroup" {
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
			return fmt.Errorf("AVI IpAddrGroup still exists")
		}
	}
	return nil
}

//nolint
const testAccAVIIpAddrGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipaddrgroup" "testIpAddrGroup" {
	name = "test-Internal-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	prefixes {
	ip_addr {
		type = "V4"
		addr = "10.0.0.0"
	}
	mask = "8"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "192.168.0.0"
	}
	mask = "16"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "172.16.0.0"
	}
	mask = "12"
}
}
`

//nolint
const testAccAVIIpAddrGroupupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipaddrgroup" "testIpAddrGroup" {
	name = "test-Internal-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	prefixes {
	ip_addr {
		type = "V4"
		addr = "10.0.0.0"
	}
	mask = "8"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "192.168.0.0"
	}
	mask = "16"
}
prefixes {
	ip_addr {
		type = "V4"
		addr = "172.16.0.0"
	}
	mask = "12"
}
}
`
