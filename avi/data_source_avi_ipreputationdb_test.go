// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceIPReputationDBBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIPReputationDBConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipreputationdb.testIPReputationDB", "name", "ip-reputation-db"),
				),
			},
		},
	})

}

const testAccAVIDSIPReputationDBConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipreputationdb" "testIPReputationDB" {
	name = "ip-reputation-db"
	tenant_ref = data.avi_tenant.default_tenant.id
	vendor = "IP_REPUTATION_VENDOR_WEBROOT"
	description = "IP reputation database"
}

data "avi_ipreputationdb" "testIPReputationDB" {
    name= "${avi_ipreputationdb.testIPReputationDB.name}"
}
`
