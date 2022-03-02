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

func TestAVIAlertSyslogConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAlertSyslogConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAlertSyslogConfig,
				//ImportState: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertSyslogConfigExists("avi_alertsyslogconfig.testalertsyslogconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertsyslogconfig.testalertsyslogconfig", "name", "asyc-test")),
			},
			{
				Config: testAccUpdatedAVIAlertSyslogConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertSyslogConfigExists("avi_alertsyslogconfig.testalertsyslogconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertsyslogconfig.testalertsyslogconfig", "name", "asyc-abc")),
			},
		},
	})

}

func testAccCheckAVIAlertSyslogConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Alert Syslog Config ID is set")
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

func testAccCheckAVIAlertSyslogConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_alertsyslogconfig" {
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
			return fmt.Errorf("AVI Alert Syslog Config still exists")
		}
	}
	return nil
}

const testAccAVIAlertSyslogConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertsyslogconfig" "testalertsyslogconfig" {
	name = "asyc-test"
	description= "test alert syslog"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIAlertSyslogConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertsyslogconfig" "testalertsyslogconfig" {
	name = "asyc-abc"
	description= "test alert syslog"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
