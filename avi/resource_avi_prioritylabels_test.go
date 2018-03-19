package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIPriorityLabelsBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIPriorityLabelsConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPriorityLabelsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPriorityLabelsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPriorityLabelsExists("avi_prioritylabels.testprioritylabels"),
					resource.TestCheckResourceAttr(
						"avi_prioritylabels.testprioritylabels", "name", "pl-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPriorityLabelsExists("avi_prioritylabels.testprioritylabels"),
					resource.TestCheckResourceAttr(
						"avi_prioritylabels.testprioritylabels", "name", "pl-abc")),
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
			return fmt.Errorf("No Priority Labels ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Priority Labels still exists")
		}
	}
	return nil
}

const testAccAVIPriorityLabelsConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_prioritylabels" "testprioritylabels" {
	name = "pl-%s"
	description = "test priority labels"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}
`
