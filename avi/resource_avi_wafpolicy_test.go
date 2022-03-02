/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIWAFPolicyBasic(t *testing.T) {
	//updatedConfig := fmt.Sprintf(testAccAVIWAFPolicyConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIWAFPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIWAFPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIWAFPolicyExists("avi_wafpolicy.testwafpolicy"),
					resource.TestCheckResourceAttr(
						"avi_wafpolicy.testwafpolicy", "name", "wp")),
			},
			{
				Config: updatetestAccAVIWAFPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIWAFPolicyExists("avi_wafpolicy.testwafpolicy"),
					resource.TestCheckResourceAttr(
						"avi_wafpolicy.testwafpolicy", "name", "wp-test")),
			},
		},
	})

}

func testAccCheckAVIWAFPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No WAF Policy ID is set")
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

func testAccCheckAVIWAFPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_wafpolicy" {
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
			return fmt.Errorf("AVI WAF Policy still exists")
		}
	}
	return nil
}

const testAccAVIWAFPolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

data "avi_wafprofile" "wafprofile" {
	name = "System-WAF-Profile"
}
resource "avi_wafpolicy" "testwafpolicy" {
	name = "wp"
	mode= "WAF_MODE_DETECTION_ONLY"
	paranoia_level= "WAF_PARANOIA_LEVEL_LOW"
	waf_profile_ref= data.avi_wafprofile.wafprofile.id
	tenant_ref= data.avi_tenant.default_tenant.id
}
`

const updatetestAccAVIWAFPolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

data "avi_wafprofile" "wafprofile" {
	name = "System-WAF-Profile"
}
resource "avi_wafpolicy" "testwafpolicy" {
	name = "wp-test"
	mode= "WAF_MODE_DETECTION_ONLY"
	paranoia_level= "WAF_PARANOIA_LEVEL_LOW"
	waf_profile_ref= data.avi_wafprofile.wafprofile.id
	tenant_ref= data.avi_tenant.default_tenant.id
}
`
