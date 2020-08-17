package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIUserAccountProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIUserAccountProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIUserAccountProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIUserAccountProfileExists("avi_useraccountprofile.testUserAccountProfile"),
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
			{
				Config: testAccAVIUserAccountProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIUserAccountProfileExists("avi_useraccountprofile.testUserAccountProfile"),
					resource.TestCheckResourceAttr(
						"avi_useraccountprofile.testUserAccountProfile", "name", "test-Default-User-Account-Profile-updated"),
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
			{
				ResourceName:      "avi_useraccountprofile.testUserAccountProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIUserAccountProfileConfig,
			},
		},
	})

}

func testAccCheckAVIUserAccountProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI UserAccountProfile ID is set")
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

func testAccCheckAVIUserAccountProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_useraccountprofile" {
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
			return fmt.Errorf("AVI UserAccountProfile still exists")
		}
	}
	return nil
}

const testAccAVIUserAccountProfileConfig = `
resource "avi_useraccountprofile" "testUserAccountProfile" {
	max_concurrent_sessions = "0"
	name = "test-Default-User-Account-Profile-abc"
	account_lock_timeout = "30"
	max_login_failure_count = "20"
	max_password_history_count = "0"
	credentials_timeout_threshold = "0"
}
`

const testAccAVIUserAccountProfileupdatedConfig = `
resource "avi_useraccountprofile" "testUserAccountProfile" {
	max_concurrent_sessions = "0"
	name = "test-Default-User-Account-Profile-updated"
	account_lock_timeout = "30"
	max_login_failure_count = "20"
	max_password_history_count = "0"
	credentials_timeout_threshold = "0"
}
`
