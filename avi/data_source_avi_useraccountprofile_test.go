// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceUserAccountProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSUserAccountProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "name", "test-Default-User-Account-Profile-abc"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "max_concurrent_sessions", "0"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "account_lock_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "max_login_failure_count", "20"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "max_password_history_count", "0"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "credentials_timeout_threshold", "0"),
				),
			},
		},
	})

}

const testAccAVIDSUserAccountProfileConfig = `
resource "avi_useraccountprofile" "testUserAccountProfile" {
	name = "test-Default-User-Account-Profile-abc"
	max_concurrent_sessions = "0"
	account_lock_timeout = "30"
	max_login_failure_count = "20"
	max_password_history_count = "0"
	credentials_timeout_threshold = "0"
}

data "avi_useraccountprofile" "testUserAccountProfile" {
    name= "${avi_useraccountprofile.testUserAccountProfile.name}"
}
`
