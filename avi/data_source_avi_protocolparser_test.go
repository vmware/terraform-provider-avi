// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceProtocolParserBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSProtocolParserConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_protocolparser.testProtocolParser", "name", "custom-http-parser"),
				),
			},
		},
	})

}

const testAccAVIDSProtocolParserConfig = `
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

data "avi_protocolparser" "testProtocolParser" {
    name= "${avi_protocolparser.testProtocolParser.name}"
}
`
