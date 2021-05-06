// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceClusterCloudDetailsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSClusterCloudDetailsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_clusterclouddetails.testClusterCloudDetails", "name", "test-ccd-abc"),
				),
			},
		},
	})

}

const testAccAVIDSClusterCloudDetailsConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_clusterclouddetails" "testClusterCloudDetails" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-ccd-abc"
}

data "avi_clusterclouddetails" "testClusterCloudDetails" {
    name= "${avi_clusterclouddetails.testClusterCloudDetails.name}"
}
`
