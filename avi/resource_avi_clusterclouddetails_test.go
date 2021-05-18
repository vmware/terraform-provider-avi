package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIClusterCloudDetailsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIClusterCloudDetailsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIClusterCloudDetailsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIClusterCloudDetailsExists("avi_clusterclouddetails.testClusterCloudDetails"),
					resource.TestCheckResourceAttr(
						"avi_clusterclouddetails.testClusterCloudDetails", "name", "test-ccd-abc"),
				),
			},
			{
				Config: testAccAVIClusterCloudDetailsupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIClusterCloudDetailsExists("avi_clusterclouddetails.testClusterCloudDetails"),
					resource.TestCheckResourceAttr(
						"avi_clusterclouddetails.testClusterCloudDetails", "name", "test-ccd-updated"),
				),
			},
			{
				ResourceName:      "avi_clusterclouddetails.testClusterCloudDetails",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIClusterCloudDetailsConfig,
			},
		},
	})

}

func testAccCheckAVIClusterCloudDetailsExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ClusterCloudDetails ID is set")
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

func testAccCheckAVIClusterCloudDetailsDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_clusterclouddetails" {
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
			return fmt.Errorf("AVI ClusterCloudDetails still exists")
		}
	}
	return nil
}

const testAccAVIClusterCloudDetailsConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_clusterclouddetails" "testClusterCloudDetails" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-ccd-abc"
}
`

const testAccAVIClusterCloudDetailsupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_clusterclouddetails" "testClusterCloudDetails" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-ccd-updated"
}
`
