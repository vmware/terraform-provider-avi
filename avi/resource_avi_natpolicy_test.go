package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVINatPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINatPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINatPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINatPolicyExists("avi_natpolicy.testNatPolicy"),
					resource.TestCheckResourceAttr(
						"avi_natpolicy.testNatPolicy", "name", "NAT-Policy"),
				),
			},
			{
				Config: testAccAVINatPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINatPolicyExists("avi_natpolicy.testNatPolicy"),
					resource.TestCheckResourceAttr(
						"avi_natpolicy.testNatPolicy", "name", "NAT-Policy-Updated"),
				),
			},
			{
				ResourceName:      "avi_natpolicy.testNatPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVINatPolicyConfig,
			},
		},
	})

}

func testAccCheckAVINatPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI NatPolicy ID is set")
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

func testAccCheckAVINatPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_natpolicy" {
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
			return fmt.Errorf("AVI NatPolicy still exists")
		}
	}
	return nil
}

const testAccAVINatPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_natpolicy" "testNatPolicy" {
	rules {
	enable = true
	match {
		services {
			destination_port {
				match_criteria = "IS_IN"
				ports = ["9000"]
			}
			source_port {
				match_criteria = "IS_IN"
				ports = ["8000"]
			}
			protocol {
				match_criteria = "IS_IN"
				protocol = "PROTOCOL_ICMP"
			}
		}
	}
	action {
		nat_info {
	nat_ip {
		addr = "192.168.10.10"
		type = "V4"
	}
	nat_ip_range {
		begin {
			addr = "192.168.10.5"
			type = "V4"
		}
		end {
			addr = "192.168.10.15"
			type = "V4"
		}
	}
}
		type = "NAT_POLICY_ACTION_TYPE_DYNAMIC_IP_PORT"
	}
	name = "natrule"
	index = "0"
}
	name = "NAT-Policy"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVINatPolicyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_natpolicy" "testNatPolicy" {
	rules {
	enable = true
	match {
		services {
			destination_port {
				match_criteria = "IS_IN"
				ports = ["9000"]
			}
			source_port {
				match_criteria = "IS_IN"
				ports = ["8000"]
			}
			protocol {
				match_criteria = "IS_IN"
				protocol = "PROTOCOL_ICMP"
			}
		}
	}
	action {
		nat_info {
	nat_ip {
		addr = "192.168.10.10"
		type = "V4"
	}
	nat_ip_range {
		begin {
			addr = "192.168.10.5"
			type = "V4"
		}
		end {
			addr = "192.168.10.15"
			type = "V4"
		}
	}
}
		type = "NAT_POLICY_ACTION_TYPE_DYNAMIC_IP_PORT"
	}
	name = "natrule"
	index = "0"
}
	name = "NAT-Policy-Updated"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
