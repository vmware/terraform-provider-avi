package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVITestSeDatastoreLevel3Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVITestSeDatastoreLevel3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITestSeDatastoreLevel3Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITestSeDatastoreLevel3Exists("avi_testsedatastorelevel3.testTestSeDatastoreLevel3"),
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel3.testTestSeDatastoreLevel3", "name", "test-se-datastore-l3"),
				),
			},
			{
				Config: testAccAVITestSeDatastoreLevel3updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITestSeDatastoreLevel3Exists("avi_testsedatastorelevel3.testTestSeDatastoreLevel3"),
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel3.testTestSeDatastoreLevel3", "name", "test-se-datastore-l3-updated"),
				),
			},
			{
				ResourceName:      "avi_testsedatastorelevel3.testTestSeDatastoreLevel3",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVITestSeDatastoreLevel3Config,
			},
		},
	})

}

func testAccCheckAVITestSeDatastoreLevel3Exists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI TestSeDatastoreLevel3 ID is set")
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

func testAccCheckAVITestSeDatastoreLevel3Destroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_testsedatastorelevel3" {
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
			return fmt.Errorf("AVI TestSeDatastoreLevel3 still exists")
		}
	}
	return nil
}

const testAccAVITestSeDatastoreLevel3Config = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel3" "testTestSeDatastoreLevel3" {
	name = "test-se-datastore-l3"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVITestSeDatastoreLevel3updatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel3" "testTestSeDatastoreLevel3" {
	name = "test-se-datastore-l3-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
