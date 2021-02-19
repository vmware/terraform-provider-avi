package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourcePoolBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSPoolConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "name", "test-Pool-abc"),
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "enabled", "false"),
				),
			},
			{
				Config: testAccAVIDSPoolConfigIgnoreServers,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_pool.testPool2", "name", "test-Pool-2"),
					//ToDo: Fix ignore_servers diff issue
					resource.TestCheckResourceAttr(
						"avi_pool.testPool2", "ignore_servers", "false"),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})

}

const testAccAVIDSPoolConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
data "avi_healthmonitor" "default_monitor" {
    name= "System-HTTP"
}
resource "avi_pool" "testPool" {
    ignore_servers = false
    name = "test-Pool-abc"
    cloud_ref = data.avi_cloud.default_cloud.id
    tenant_ref = data.avi_tenant.default_tenant.id
	enabled = false
    servers {
	   ip {
		  type = "V4"
		  addr = "10.90.64.66"
	   }
	   hostname = "10.90.64.66"
	   ratio = "1"
	   port = "8080"
	   enabled = true
    }
    health_monitor_refs = [data.avi_healthmonitor.default_monitor.id]
    fail_action {
        type = "FAIL_ACTION_CLOSE_CONN"
	}
}
data "avi_pool" "testPool" {
    name= avi_pool.testPool.name
	ignore_servers= false
}
`

const testAccAVIDSPoolConfigIgnoreServers = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
data "avi_healthmonitor" "default_monitor" {
    name= "System-HTTP"
}
resource "avi_pool" "testPool2" {
    //ToDo: Fix ignore_servers diff issue
    ignore_servers = false
    name = "test-Pool-2"
    cloud_ref = data.avi_cloud.default_cloud.id
    tenant_ref = data.avi_tenant.default_tenant.id
	enabled = false
    health_monitor_refs = [data.avi_healthmonitor.default_monitor.id]
    fail_action {
        type = "FAIL_ACTION_CLOSE_CONN"
	}
}
data "avi_pool" "testPool2" {
    name= avi_pool.testPool2.name
	ignore_servers= false
    depends_on = [avi_pool.testPool2]
}
`
