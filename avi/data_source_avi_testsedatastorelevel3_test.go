// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceTestSeDatastoreLevel3Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSTestSeDatastoreLevel3Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel3.testTestSeDatastoreLevel3", "name", "test-se-datastore-l3"),
				),
			},
		},
	})

}

const testAccAVIDSTestSeDatastoreLevel3Config = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel3" "testTestSeDatastoreLevel3" {
	name = "test-se-datastore-l3"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_testsedatastorelevel3" "testTestSeDatastoreLevel3" {
    name= "${avi_testsedatastorelevel3.testTestSeDatastoreLevel3.name}"
}
`
