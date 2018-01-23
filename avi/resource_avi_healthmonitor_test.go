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
	updatedConfig := fmt.Sprintf(testAccAVIHealthMonitorConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIHealthMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIHealthMonitorConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHealthMonitorExists("avi_healthmonitor.testhm"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testhm", "name", "hm-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIHealthMonitorExists("avi_healthmonitor.testhm"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testhm", "name", "hm-abc")),
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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

resource "avi_healthmonitor" "testhm" {
	name = "hm-%s"
	type = "HEALTH_MONITOR_HTTP"
	http_monitor {
		http_request = "GET / HTTP/1.0"
		http_response_code = ["HTTP_2XX"]
	}
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
