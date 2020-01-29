package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
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
