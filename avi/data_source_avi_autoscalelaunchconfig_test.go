// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	name = "test-default-autoscalelaunchconfig-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	image_id = "default"
	use_external_asg = true
}

data "avi_autoscalelaunchconfig" "testAutoScaleLaunchConfig" {
    name= "${avi_autoscalelaunchconfig.testAutoScaleLaunchConfig.name}"
}
`
