package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceHealthMonitorBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSHealthMonitorConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "name", "test-System-HTTP-abc"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "receive_timeout", "4"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "is_federated", "false"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "failed_checks", "3"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "send_interval", "10"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testHealthMonitor", "successful_checks", "3"),
				),
			},
		},
	})

}

const testAccAVIDSHealthMonitorConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_healthmonitor" "testHealthMonitor" {
	receive_timeout = "4"
	name = "test-System-HTTP-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	is_federated = false
	failed_checks = "3"
	send_interval = "10"
	http_monitor {
		exact_http_request = false
		http_request = "HEAD / HTTP/1.0"
		http_response_code = ["HTTP_2XX","HTTP_3XX"]
	}
	type = "HEALTH_MONITOR_HTTP"
	successful_checks = "3"

}

data "avi_healthmonitor" "testHealthMonitor" {
    name= avi_healthmonitor.testHealthMonitor.name
}
`
