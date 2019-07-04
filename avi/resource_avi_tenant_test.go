package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVITenantBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVITenantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITenantConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITenantExists("avi_tenant.testTenant"),
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "name", "test-admin-abc"),
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "local", "true"),
				),
			},
			{
				Config: testAccAVITenantupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITenantExists("avi_tenant.testTenant"),
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "name", "test-admin-updated"),
					resource.TestCheckResourceAttr(
						"avi_tenant.testTenant", "local", "true"),
				),
			},
			{
				ResourceName:      "avi_tenant.testTenant",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVITenantConfig,
			},
		},
	})

}

func testAccCheckAVITenantExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Tenant ID is set")
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

func testAccCheckAVITenantDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_tenant" {
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
			return fmt.Errorf("AVI Tenant still exists")
		}
	}
	return nil
}

const testAccAVITenantConfig = `
resource "avi_tenant" "testTenant" {
	"local" = true
	"name" = "test-admin-abc"
}
`

const testAccAVITenantupdatedConfig = `
resource "avi_tenant" "testTenant" {
	"local" = true
	"name" = "test-admin-updated"
}
`
