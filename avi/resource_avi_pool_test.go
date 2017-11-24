package avi

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIPoolCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckAVIPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testpool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testpool", "name", "p42-%s")),
			},
		},
	})

}

func TestAVIPoolUpdate(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIPoolConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckAVIPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testpool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testpool", "name", "p42-abc")),
			},
		},
	})

}

func testAccCheckAVIPoolExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI POOL ID is set")
		}
		return nil
	}

}

// func testAccCheckAVIPoolDestroy(s *terraform.State) error {
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "avi_pool" {
// 			continue
// 		}
// 		if rs.Primary.ID != "" {
// 			return fmt.Errorf("AVI POOL ID is set")
// 		}
// 	}
// 	return nil
// }

const testAccAVIPoolConfig = `
resource "avi_healthmonitor" "test_hm" {
	name= "hm1"
	type= "HEALTH_MONITOR_HTTP"
}

resource "avi_pool" "testpool" {
	name = "p42-%s",
	health_monitor_refs= ["${avi_healthmonitor.test_hm.id}"]
	servers {
		ip= {
		  type= "V4",
		  addr= "10.90.64.66",
		}
		port= 8080
	}
	servers {
		ip= {
			type= "V4",
			addr= "10.90.64.70",
		}
	}
	fail_action= {
		type= "FAIL_ACTION_CLOSE_CONN"
	}
}
`
