// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceBotIPReputationTypeMappingBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSBotIPReputationTypeMappingConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_botipreputationtypemapping.testBotIPReputationTypeMapping", "name", "bot_ip_reputation_map_31"),
				),
			},
		},
	})

}

const testAccAVIDSBotIPReputationTypeMappingConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_botipreputationtypemapping" "testBotIPReputationTypeMapping" {
	name = "bot_ip_reputation_map_31"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_botipreputationtypemapping" "testBotIPReputationTypeMapping" {
    name= "${avi_botipreputationtypemapping.testBotIPReputationTypeMapping.name}"
}
`
