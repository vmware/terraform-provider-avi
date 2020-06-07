package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIHealthMonitorBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIHealthMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIHealthMonitorConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHealthMonitorExists("avi_healthmonitor.testHealthMonitor"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "name", "test-System-HTTP-abc"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "is_federated", "false"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "receive_timeout", "4"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "failed_checks", "3"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "send_interval", "10"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "successful_checks", "3"),
				),
			},
			{
				Config: testAccAVIHealthMonitorupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHealthMonitorExists("avi_healthmonitor.testHealthMonitor"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "name", "test-System-HTTP-updated"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "is_federated", "false"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "receive_timeout", "4"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "failed_checks", "3"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "send_interval", "10"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "successful_checks", "3"),
				),
			},
			{
				ResourceName:      "avi_healthmonitor.testHealthMonitor",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIHealthMonitorConfig,
			},
		},
	})

}

func testAccCheckAVIHealthMonitorExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI HealthMonitor ID is set")
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

func testAccCheckAVIHealthMonitorDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_healthmonitor" {
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
			return fmt.Errorf("AVI HealthMonitor still exists")
		}
	}
	return nil
}

const testAccAVIHealthMonitorConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_healthmonitor" "testHealthMonitor" {
	name = "test-System-HTTP-abc"
	is_federated = false
	tenant_ref = data.avi_tenant.default_tenant.id
	receive_timeout = "4"
	failed_checks = "3"
	send_interval = "10"
	http_monitor {
		exact_http_request = false
		http_request = "HEAD / HTTP/1.0"
		response_size = "2050"
		http_response_code = ["HTTP_2XX","HTTP_3XX"]
	}
	successful_checks = "3"
	type = "HEALTH_MONITOR_HTTP"
}
`

const testAccAVIHealthMonitorupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_healthmonitor" "testHealthMonitor" {
	name = "test-System-HTTP-updated"
	is_federated = false
	tenant_ref = data.avi_tenant.default_tenant.id
	receive_timeout = "4"
	failed_checks = "3"
	send_interval = "10"
	http_monitor {
		exact_http_request = false
		http_request = "HEAD / HTTP/1.0"
		response_size = "2050"
		http_response_code = ["HTTP_2XX","HTTP_3XX"]
	}
	successful_checks = "3"
	type = "HEALTH_MONITOR_HTTP"
}
`
