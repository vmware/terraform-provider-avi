package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIClusterCloudDetailsBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIClusterCloudDetailsConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIClusterCloudDetailsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIClusterCloudDetailsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIClusterCloudDetailsExists("avi_clusterclouddetails.testclusterclouddetails"),
					resource.TestCheckResourceAttr(
						"avi_clusterclouddetails.testclusterclouddetails", "name", "ccd-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIClusterCloudDetailsExists("avi_clusterclouddetails.testclusterclouddetails"),
					resource.TestCheckResourceAttr(
						"avi_clusterclouddetails.testclusterclouddetails", "name", "ccd-abc")),
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
			return fmt.Errorf("No Cluster Cloud Details ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Cluster Cloud Details still exists")
		}
	}
	return nil
}

const testAccAVIClusterCloudDetailsConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_clusterclouddetails" "testclusterclouddetails" {
	name = "ccd-%s"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
