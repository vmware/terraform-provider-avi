/***************************************************************************
 * ========================================================================
 * Copyright 2023 VMware, Inc. All rights reserved. VMware Confidential
 * ========================================================================
 */

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceClusterBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.avi_cluster.testCluster", "name", "cluster-0-1"),
				),
			},
		},
	})

}

const testAccAVIDSClusterConfig = `
data "avi_cluster" "testCluster" {
	name = "cluster-0-1"
}
`
