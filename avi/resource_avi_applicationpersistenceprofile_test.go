package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

func TestAVIApplicationPersistenceProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIApplicationPersistenceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIApplicationPersistenceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationPersistenceProfileExists("avi_applicationpersistenceprofile.testApplicationPersistenceProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "name", "test-System-Persistence-Client-IP-abc"),
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "is_federated", "false"),
				),
			},
			{
				Config: testAccAVIApplicationPersistenceProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationPersistenceProfileExists("avi_applicationpersistenceprofile.testApplicationPersistenceProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "name", "test-System-Persistence-Client-IP-updated"),
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "is_federated", "false"),
				),
			},
			{
				ResourceName:      "avi_applicationpersistenceprofile.testApplicationPersistenceProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIApplicationPersistenceProfileConfig,
			},
		},
	})

}

func testAccCheckAVIApplicationPersistenceProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ApplicationPersistenceProfile ID is set")
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

func testAccCheckAVIApplicationPersistenceProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_applicationpersistenceprofile" {
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
			return fmt.Errorf("AVI ApplicationPersistenceProfile still exists")
		}
	}
	return nil
}

const testAccAVIApplicationPersistenceProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_applicationpersistenceprofile" "testApplicationPersistenceProfile" {
	name = "test-System-Persistence-Client-IP-abc"
	persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
	tenant_ref = data.avi_tenant.default_tenant.id
	server_hm_down_recovery = "HM_DOWN_PICK_NEW_SERVER"
	is_federated = false
	ip_persistence_profile {
		ip_persistent_timeout = "5"
	}
}
`

const testAccAVIApplicationPersistenceProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_applicationpersistenceprofile" "testApplicationPersistenceProfile" {
	name = "test-System-Persistence-Client-IP-updated"
	persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
	tenant_ref = data.avi_tenant.default_tenant.id
	server_hm_down_recovery = "HM_DOWN_PICK_NEW_SERVER"
	is_federated = false
	ip_persistence_profile {
		ip_persistent_timeout = "5"
	}
}
`
