// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceCloudBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSCloudConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "name", "test-Default-Cloud-abc"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "dhcp_enabled", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "enable_vip_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "state_based_dns_registration", "true"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "prefer_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "apic_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "mtu", "1500"),
				),
			},
		},
	})

}

const testAccAVIDSCloudConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_cloud" "testCloud" {
	name = "test-Default-Cloud-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	dhcp_enabled = false
	vtype = "CLOUD_NONE"
	license_tier = "ENTERPRISE"
	enable_vip_static_routes = false
	state_based_dns_registration = true
	prefer_static_routes = false
	license_type = "LIC_CORES"
	apic_mode = false
	mtu = "1500"
}

data "avi_cloud" "testCloud" {
    name= "${avi_cloud.testCloud.name}"
}
`
