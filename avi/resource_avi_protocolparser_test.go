package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIProtocolParserBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIProtocolParserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIProtocolParserConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIProtocolParserExists("avi_protocolparser.testProtocolParser"),
					resource.TestCheckResourceAttr(
						"avi_protocolparser.testProtocolParser", "name", "custom-http-parser"),
				),
			},
			{
				Config: testAccAVIProtocolParserupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIProtocolParserExists("avi_protocolparser.testProtocolParser"),
					resource.TestCheckResourceAttr(
						"avi_protocolparser.testProtocolParser", "name", "custom-http-parser-updated"),
				),
			},
			{
				ResourceName:      "avi_protocolparser.testProtocolParser",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIProtocolParserConfig,
			},
		},
	})

}

func testAccCheckAVIProtocolParserExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ProtocolParser ID is set")
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

func testAccCheckAVIProtocolParserDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_protocolparser" {
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
			return fmt.Errorf("AVI ProtocolParser still exists")
		}
	}
	return nil
}

const testAccAVIProtocolParserConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_protocolparser" "testProtocolParser" {
	name = "custom-http-parser"
	parser_code = <<EOF

function parse_request(req)
  avi.log('Parsing HTTP request for inspection')
  return req
 end
EOF
	description = "Custom HTTP parser with simple logging"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVIProtocolParserupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_protocolparser" "testProtocolParser" {
	name = "custom-http-parser-updated"
	parser_code = <<EOF

function parse_request(req)
  avi.log('Parsing HTTP request for inspection')
  return req
 end
EOF
	description = "Custom HTTP parser with simple logging"
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
