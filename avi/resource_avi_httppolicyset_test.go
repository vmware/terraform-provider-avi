package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIHTTPPolicySetBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIHTTPPolicySetConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIHTTPPolicySetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIHTTPPolicySetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHTTPPolicySetExists("avi_httppolicyset.testhttppolicyset"),
					resource.TestCheckResourceAttr(
						"avi_httppolicyset.testhttppolicyset", "name", "policy-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHTTPPolicySetExists("avi_httppolicyset.testhttppolicyset"),
					resource.TestCheckResourceAttr(
						"avi_httppolicyset.testhttppolicyset", "name", "policy-abc")),
			},
		},
	})

}

func testAccCheckAVIHTTPPolicySetExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No HTTP Policy Set ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIHTTPPolicySetDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_httppolicyset" {
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
			return fmt.Errorf("AVI HTTP Policy Set still exists")
		}
	}
	return nil
}

const testAccAVIHTTPPolicySetConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}

resource "avi_pool" "testpool" {
	name = "p42"
	server_count = 1
	servers {
		hostname= "10.90.64.66"
		ip= {
		  type= "V4"
		  addr= "10.90.64.66"
		}
		port= 8080
	}
	fail_action= {
		type= "FAIL_ACTION_CLOSE_CONN"
	}
	vrf_ref="${data.avi_vrfcontext.global_vrf.id}"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}

resource "avi_httppolicyset" "testhttppolicyset" {
	name = "policy-%s"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	is_internal_policy = false
	http_request_policy= {
		rules= [{
         	index= 1
      		enable= true
			name= "rule-1"
			match= {
				hdrs= [{
					match_case= "INSENSITIVE"
					hdr= "User-Agent"
					value= [
						"Backup_Pool_Redirect"
					]
					match_criteria= "HDR_CONTAINS"
				}]
			}
			switching_action= {
				action= "HTTP_SWITCHING_SELECT_POOL"
				pool_ref = "${avi_pool.testpool.id}"
			}
		}]
	}
}
`
