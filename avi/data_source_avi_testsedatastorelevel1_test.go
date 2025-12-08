// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceTestSeDatastoreLevel1Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSTestSeDatastoreLevel1Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_testsedatastorelevel1.testTestSeDatastoreLevel1", "name", "test-se-datastore-lvl1"),
				),
			},
		},
	})

}

const testAccAVIDSTestSeDatastoreLevel1Config = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_testsedatastorelevel1" "testTestSeDatastoreLevel1" {
	name = "test-se-datastore-lvl1"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_testsedatastorelevel1" "testTestSeDatastoreLevel1" {
    name= "${avi_testsedatastorelevel1.testTestSeDatastoreLevel1.name}"
}
`
