package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIAutoScaleLaunchConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAutoScaleLaunchConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAutoScaleLaunchConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAutoScaleLaunchConfigExists("avi_autoscalelaunchconfig.testAutoScaleLaunchConfig"),
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "name", "test-default-autoscalelaunchconfig-abc"),
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "use_external_asg", "true"),
				),
			},
			{
				Config: testAccAVIAutoScaleLaunchConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAutoScaleLaunchConfigExists("avi_autoscalelaunchconfig.testAutoScaleLaunchConfig"),
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "name", "test-default-autoscalelaunchconfig-updated"),
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "use_external_asg", "true"),
				),
			},
			{
				ResourceName:      "avi_autoscalelaunchconfig.testAutoScaleLaunchConfig",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIAutoScaleLaunchConfigConfig,
			},
		},
	})

}

func testAccCheckAVIAutoScaleLaunchConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI AutoScaleLaunchConfig ID is set")
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

func testAccCheckAVIAutoScaleLaunchConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_autoscalelaunchconfig" {
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
			return fmt.Errorf("AVI AutoScaleLaunchConfig still exists")
		}
	}
	return nil
}

const testAccAVIAutoScaleLaunchConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_autoscalelaunchconfig" "testAutoScaleLaunchConfig" {
	name = "test-default-autoscalelaunchconfig-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	image_id = "default"
	use_external_asg = true
}
`

const testAccAVIAutoScaleLaunchConfigupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_autoscalelaunchconfig" "testAutoScaleLaunchConfig" {
	name = "test-default-autoscalelaunchconfig-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	image_id = "default"
	use_external_asg = true
}
`
