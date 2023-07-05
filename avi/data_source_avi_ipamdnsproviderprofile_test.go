// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceIpamDnsProviderProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIpamDnsProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "name", "test-ipam-abc-abc"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "allocate_ip_in_vrf", "false"),
				),
			},
		},
	})

}

// nolint
const testAccAVIDSIpamDnsProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
	name = "test-ipam-abc-abc"
	allocate_ip_in_vrf = false
	internal_profile {
		ttl = "31"
	}
	type = "IPAMDNS_TYPE_INTERNAL"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
    name= "${avi_ipamdnsproviderprofile.testIpamDnsProviderProfile.name}"
}
`
