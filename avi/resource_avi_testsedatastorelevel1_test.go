package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVITestSeDatastoreLevel1Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVITestSeDatastoreLevel1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITestSeDatastoreLevel1Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITestSeDatastoreLevel1Exists("avi_testsedatastorelevel1.testTestSeDatastoreLevel1"),
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel1.testTestSeDatastoreLevel1", "name", "test-se-datastore-lvl1"),
				),
			},
			{
				Config: testAccAVITestSeDatastoreLevel1updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITestSeDatastoreLevel1Exists("avi_testsedatastorelevel1.testTestSeDatastoreLevel1"),
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel1.testTestSeDatastoreLevel1", "name", "test-se-datastore-lvl1-updated"),
				),
			},
			{
				ResourceName:      "avi_testsedatastorelevel1.testTestSeDatastoreLevel1",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVITestSeDatastoreLevel1Config,
			},
		},
	})

}

func testAccCheckAVITestSeDatastoreLevel1Exists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI TestSeDatastoreLevel1 ID is set")
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

func testAccCheckAVITestSeDatastoreLevel1Destroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_testsedatastorelevel1" {
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
			return fmt.Errorf("AVI TestSeDatastoreLevel1 still exists")
		}
	}
	return nil
}

const testAccAVITestSeDatastoreLevel1Config = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel1" "testTestSeDatastoreLevel1" {
	name = "test-se-datastore-lvl1"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVITestSeDatastoreLevel1updatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel1" "testTestSeDatastoreLevel1" {
	name = "test-se-datastore-lvl1-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
