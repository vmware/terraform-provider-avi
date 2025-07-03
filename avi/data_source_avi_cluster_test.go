/***************************************************************************
 * ========================================================================
 * Copyright (c) 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved. Broadcom Confidential.
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
