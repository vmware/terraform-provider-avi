// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceAvailabilityZoneBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAvailabilityZoneConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_availabilityzone.testAvailabilityZone", "name", "az-1"),
				),
			},
		},
	})

}

const testAccAVIDSAvailabilityZoneConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_availabilityzone" "testAvailabilityZone" {
	name = "az-1"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
}

data "avi_availabilityzone" "testAvailabilityZone" {
    name= "${avi_availabilityzone.testAvailabilityZone.name}"
}
`
