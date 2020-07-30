package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIRoleBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIRoleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIRoleExists("avi_role.testRole"),
					resource.TestCheckResourceAttr(
						"avi_role.testRole", "name", "test-Application-Admin-abc"),
				),
			},
			{
				Config: testAccAVIRoleupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIRoleExists("avi_role.testRole"),
					resource.TestCheckResourceAttr(
						"avi_role.testRole", "name", "test-Application-Admin-updated"),
				),
			},
			{
				ResourceName:      "avi_role.testRole",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIRoleConfig,
			},
		},
	})

}

func testAccCheckAVIRoleExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Role ID is set")
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

func testAccCheckAVIRoleDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_role" {
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
			return fmt.Errorf("AVI Role still exists")
		}
	}
	return nil
}

const testAccAVIRoleConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_role" "testRole" {
	name = "test-Application-Admin-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_CONTROLLER"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_CLOUD"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_INTERNAL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VIRTUALSERVICE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOLGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOLGROUPDEPLOYMENTPOLICY"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_PRIORITYLABELS"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_HEALTHMONITOR"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_NETWORKPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_APPLICATIONPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_HTTPPOLICYSET"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_IPADDRGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_STRINGGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_MICROSERVICEGROUP"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SSLPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SSLKEYANDCERTIFICATE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_NETWORKSECURITYPOLICY"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_APPLICATIONPERSISTENCEPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ANALYTICSPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_IPAMDNSPROVIDERPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VSDATASCRIPTSET"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_TENANT"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_PKIPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_CERTIFICATEMANAGEMENTPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_AUTHPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SERVICEENGINE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SERVICEENGINEGROUP"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_NETWORK"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_SYSTEMCONFIGURATION"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VRFCONTEXT"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_USER"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_ROLE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ALERT"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ALERTCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_ALERTEMAILCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_ALERTSYSLOGCONFIG"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ACTIONGROUPCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SNMPTRAPPROFILE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_UPGRADE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_REBOOT"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_TECHSUPPORT"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_TRAFFIC_CAPTURE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_GSLB"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_GSLBGEODBPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_GSLBSERVICE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_TRAFFICCLONEPROFILE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_SE_TOKEN"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_WAFPOLICY"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_WAFPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_USER_CREDENTIAL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_AUTOSCALE"
}
}
`

const testAccAVIRoleupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_role" "testRole" {
	name = "test-Application-Admin-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_CONTROLLER"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_CLOUD"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_INTERNAL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VIRTUALSERVICE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOLGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_POOLGROUPDEPLOYMENTPOLICY"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_PRIORITYLABELS"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_HEALTHMONITOR"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_NETWORKPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_APPLICATIONPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_HTTPPOLICYSET"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_IPADDRGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_STRINGGROUP"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_MICROSERVICEGROUP"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SSLPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SSLKEYANDCERTIFICATE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_NETWORKSECURITYPOLICY"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_APPLICATIONPERSISTENCEPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ANALYTICSPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_IPAMDNSPROVIDERPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VSDATASCRIPTSET"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_TENANT"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_PKIPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_CERTIFICATEMANAGEMENTPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_AUTHPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SERVICEENGINE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SERVICEENGINEGROUP"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_NETWORK"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_SYSTEMCONFIGURATION"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_VRFCONTEXT"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_USER"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_ROLE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ALERT"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ALERTCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_ALERTEMAILCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_ALERTSYSLOGCONFIG"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_ACTIONGROUPCONFIG"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_SNMPTRAPPROFILE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_UPGRADE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_REBOOT"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_TECHSUPPORT"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_TRAFFIC_CAPTURE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_GSLB"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_GSLBGEODBPROFILE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_GSLBSERVICE"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_TRAFFICCLONEPROFILE"
}
privileges {
	type = "NO_ACCESS"
	resource = "PERMISSION_SE_TOKEN"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_WAFPOLICY"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_WAFPROFILE"
}
privileges {
	type = "READ_ACCESS"
	resource = "PERMISSION_USER_CREDENTIAL"
}
privileges {
	type = "WRITE_ACCESS"
	resource = "PERMISSION_AUTOSCALE"
}
}
`
