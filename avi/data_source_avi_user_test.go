// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceUserBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSUserConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "name", "test-tf-user"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "is_superuser", "true"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "local", "true"),
				),
			},
		},
	})

}

const testAccAVIDSUserConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_role" "default-system-admin-role" {
    name= "System-Admin"
}
data "avi_useraccountprofile" "default-user-account-profile" {
    name= "Default-User-Account-Profile"
}
resource "avi_user" "testUser" {
	username = "test-tf-user"
	user_profile_ref = data.avi_useraccountprofile.default-user-account-profile.id
	name = "test-tf-user"
	access {
	all_tenants = false
	tenant_ref = data.avi_tenant.default_tenant.id
	role_ref = data.avi_role.default-system-admin-role.id
}
	is_superuser = true
	full_name = "System Administrator"
	password = "pbkdf2_sha256$100000$vwZd950E3jSj$tC/x4hJBolHm2Ki4uVNbMW59ZQcC95/p5UZUWjmTuFs="
	local = true
	email = ""
	default_tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_user" "testUser" {
    name= "${avi_user.testUser.name}"
}
`
