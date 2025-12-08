package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIBotIPReputationTypeMappingBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIBotIPReputationTypeMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIBotIPReputationTypeMappingConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIBotIPReputationTypeMappingExists("avi_botipreputationtypemapping.testBotIPReputationTypeMapping"),
					resource.TestCheckResourceAttr(
						"avi_botipreputationtypemapping.testBotIPReputationTypeMapping", "name", "bot_ip_reputation_map_31"),
				),
			},
			{
				Config: testAccAVIBotIPReputationTypeMappingupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIBotIPReputationTypeMappingExists("avi_botipreputationtypemapping.testBotIPReputationTypeMapping"),
					resource.TestCheckResourceAttr(
						"avi_botipreputationtypemapping.testBotIPReputationTypeMapping", "name", "bot_ip_reputation_map_31-updated"),
				),
			},
			{
				ResourceName:      "avi_botipreputationtypemapping.testBotIPReputationTypeMapping",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIBotIPReputationTypeMappingConfig,
			},
		},
	})

}

func testAccCheckAVIBotIPReputationTypeMappingExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI BotIPReputationTypeMapping ID is set")
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

func testAccCheckAVIBotIPReputationTypeMappingDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_botipreputationtypemapping" {
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
			return fmt.Errorf("AVI BotIPReputationTypeMapping still exists")
		}
	}
	return nil
}

const testAccAVIBotIPReputationTypeMappingConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_botipreputationtypemapping" "testBotIPReputationTypeMapping" {
	name = "bot_ip_reputation_map_31"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVIBotIPReputationTypeMappingupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_botipreputationtypemapping" "testBotIPReputationTypeMapping" {
	name = "bot_ip_reputation_map_31-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
