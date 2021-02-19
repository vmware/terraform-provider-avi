package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIPriorityLabelsBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPriorityLabelsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPriorityLabelsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPriorityLabelsExists("avi_prioritylabels.testPriorityLabels"),
					resource.TestCheckResourceAttr(
						"avi_prioritylabels.testPriorityLabels", "name", "test-pl-test-abc"),
				),
			},
			{
				Config: testAccAVIPriorityLabelsupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPriorityLabelsExists("avi_prioritylabels.testPriorityLabels"),
					resource.TestCheckResourceAttr(
						"avi_prioritylabels.testPriorityLabels", "name", "test-pl-updated"),
				),
			},
			{
				ResourceName:      "avi_prioritylabels.testPriorityLabels",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIPriorityLabelsConfig,
			},
		},
	})

}

func testAccCheckAVIPriorityLabelsExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI PriorityLabels ID is set")
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

func testAccCheckAVIPriorityLabelsDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_prioritylabels" {
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
			return fmt.Errorf("AVI PriorityLabels still exists")
		}
	}
	return nil
}

const testAccAVIPriorityLabelsConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_prioritylabels" "testPriorityLabels" {
	name = "test-pl-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "test priority labels"
}
`

const testAccAVIPriorityLabelsupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_prioritylabels" "testPriorityLabels" {
	name = "test-pl-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "test priority labels"
}
`
