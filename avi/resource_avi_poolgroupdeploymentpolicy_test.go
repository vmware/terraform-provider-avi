package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
					testAccCheckAVIPoolGroupDeploymentPolicyExists("avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy"),
					resource.TestCheckResourceAttr(
						"avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy", "name", "test-pgpp-test-abc"),
				),
			},
			{
				Config: testAccAVIPoolGroupDeploymentPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolGroupDeploymentPolicyExists("avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy"),
					resource.TestCheckResourceAttr(
						"avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy", "name", "test-pgpp-updated"),
				),
			},
			{
				ResourceName:      "avi_poolgroupdeploymentpolicy.testPoolGroupDeploymentPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIPoolGroupDeploymentPolicyConfig,
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
			return fmt.Errorf("No AVI PoolGroupDeploymentPolicy ID is set")
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
			return fmt.Errorf("AVI PoolGroupDeploymentPolicy still exists")
		}
	}
	return nil
}

const testAccAVIPoolGroupDeploymentPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroupdeploymentpolicy" "testPoolGroupDeploymentPolicy" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-pgpp-test-abc"
}
`

const testAccAVIPoolGroupDeploymentPolicyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_poolgroupdeploymentpolicy" "testPoolGroupDeploymentPolicy" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-pgpp-updated"
}
`
