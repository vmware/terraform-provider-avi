package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIAlertConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAlertConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAlertConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertConfigExists("avi_alertconfig.testalertconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testalertconfig", "name", "ac-test")),
			},
			{
				Config: testAccUpdatedAVIAlertConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertConfigExists("avi_alertconfig.testalertconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testalertconfig", "name", "ac-abc")),
			},
		},
	})

}

func testAccCheckAVIAlertConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Alert Config ID is set")
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

func testAccCheckAVIAlertConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_alertconfig" {
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
			return fmt.Errorf("AVI Alert Config still exists")
		}
	}
	return nil
}

const testAccAVIAlertConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertconfig" "testalertconfig" {
	name = "ac-test"
	category= "REALTIME"
	expiry_time= 86400
	enabled= true
	summary= "System-CC-Alert System Alert Triggered"
	rolling_window= 300
	source= "EVENT_LOGS"
	threshold= 1
	throttle= 0
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	alert_rule= {
		operator= "OPERATOR_OR"
	}
}
`

const testAccUpdatedAVIAlertConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertconfig" "testalertconfig" {
	name = "ac-abc"
	category= "REALTIME"
	expiry_time= 86400
	enabled= true
	summary= "System-CC-Alert System Alert Triggered"
	rolling_window= 300
	source= "EVENT_LOGS"
	threshold= 1
	throttle= 0
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	alert_rule= {
		operator= "OPERATOR_OR"
	}
}
`
