package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIPoolGroupDeploymentPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPoolGroupDeploymentPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolGroupDeploymentPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupDeploymentPolicyExists("avi_poolgroupdeploymentpolicy.testpoolgroupdeploymentpolicy"),
					resource.TestCheckResourceAttr(
						"avi_poolgroupdeploymentpolicy.testpoolgroupdeploymentpolicy", "name", "pgpp-test")),
			},
			{
				Config: testAccUpdatedAVIPoolGroupDeploymentPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupDeploymentPolicyExists("avi_poolgroupdeploymentpolicy.testpoolgroupdeploymentpolicy"),
					resource.TestCheckResourceAttr(
						"avi_poolgroupdeploymentpolicy.testpoolgroupdeploymentpolicy", "name", "pgpp-abc")),
			},
		},
	})

}

func testAccCheckAVIPoolGroupDeploymentPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No PoolGroup Deployment Policy ID is set")
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

func testAccCheckAVIPoolGroupDeploymentPolicyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_poolgroupdeploymentpolicy" {
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
			return fmt.Errorf("AVI PoolGroup Deployment Policy still exists")
		}
	}
	return nil
}

const testAccAVIPoolGroupDeploymentPolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
resource "avi_poolgroupdeploymentpolicy" "testpoolgroupdeploymentpolicy" {
	name = "pgpp-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIPoolGroupDeploymentPolicyConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_poolgroupdeploymentpolicy" "testpoolgroupdeploymentpolicy" {
	name = "pgpp-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
