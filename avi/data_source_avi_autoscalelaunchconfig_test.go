package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceAutoScaleLaunchConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAutoScaleLaunchConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "name", "test-default-autoscalelaunchconfig-abc"),
					resource.TestCheckResourceAttr(
						"avi_autoscalelaunchconfig.testAutoScaleLaunchConfig", "use_external_asg", "true"),
				),
			},
		},
	})

}

const testAccAVIDSAutoScaleLaunchConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_autoscalelaunchconfig" "testAutoScaleLaunchConfig" {
	"image_id" = "default"
	"use_external_asg" = true
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "test-default-autoscalelaunchconfig-abc"
}

data "avi_autoscalelaunchconfig" "testAutoScaleLaunchConfig" {
    name= "${avi_autoscalelaunchconfig.testAutoScaleLaunchConfig.name}"
}
`
