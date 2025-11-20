// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceTestSeDatastoreLevel2Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSTestSeDatastoreLevel2Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel2.testTestSeDatastoreLevel2", "name", "test-se-datastore-lvl2"),
				),
			},
		},
	})

}

const testAccAVIDSTestSeDatastoreLevel2Config = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel2" "testTestSeDatastoreLevel2" {
	name = "test-se-datastore-lvl2"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_testsedatastorelevel2" "testTestSeDatastoreLevel2" {
    name= "${avi_testsedatastorelevel2.testTestSeDatastoreLevel2.name}"
}
`
