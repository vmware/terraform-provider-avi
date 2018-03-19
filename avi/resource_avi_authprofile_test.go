package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIAuthProfileBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIAuthProfileConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAuthProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAuthProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAuthProfileExists("avi_authprofile.testauthprofile"),
					resource.TestCheckResourceAttr(
						"avi_authprofile.testauthprofile", "name", "ap-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAuthProfileExists("avi_authprofile.testauthprofile"),
					resource.TestCheckResourceAttr(
						"avi_authprofile.testauthprofile", "name", "ap-abc")),
			},
		},
	})

}

func testAccCheckAVIAuthProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Auth Profile ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIAuthProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_authprofile" {
			continue
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Auth Profile still exists")
		}
	}
	return nil
}

const testAccAVIAuthProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_authprofile" "testauthprofile" {
	name = "ap-%s"
	http= {
		cache_expiration_time= 5
		group_member_is_full_dn= false
	}
	ldap= {
		security_mode= "AUTH_LDAP_SECURE_NONE"
		settings= {
			group_search_scope= "AUTH_LDAP_SCOPE_SUBTREE"
			group_member_is_full_dn= true
			user_search_scope= "AUTH_LDAP_SCOPE_ONE"
			user_id_attribute= "gg"
			group_member_attribute= "member"
			group_filter= "(objectClass=*)"
			ignore_referrals= false
			password= "avi123"
			admin_bind_dn= "avinetworkstest.com"
		}
		bind_as_administrator= true
		server= [
			"10.0.0.1"
		]
		user_bind= {
			token= "<user>"
		}
		full_name_attribute= "name"
		email_attribute= "email"
		port= 389
	}
	type= "AUTH_PROFILE_LDAP"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
