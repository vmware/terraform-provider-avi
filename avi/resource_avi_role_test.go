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
	privileges {
	resource = "PERMISSION_CONTROLLER"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_CLOUD"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_INTERNAL"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_VIRTUALSERVICE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOL"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOLGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOLGROUPDEPLOYMENTPOLICY"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_PRIORITYLABELS"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_HEALTHMONITOR"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORKPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_APPLICATIONPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_HTTPPOLICYSET"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_IPADDRGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_STRINGGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_MICROSERVICEGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SSLPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SSLKEYANDCERTIFICATE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORKSECURITYPOLICY"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_APPLICATIONPERSISTENCEPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ANALYTICSPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_IPAMDNSPROVIDERPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_VSDATASCRIPTSET"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_TENANT"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_PKIPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_CERTIFICATEMANAGEMENTPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_AUTHPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SERVICEENGINE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SERVICEENGINEGROUP"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORK"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SYSTEMCONFIGURATION"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_VRFCONTEXT"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_USER"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_ROLE"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERT"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTCONFIG"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTEMAILCONFIG"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTSYSLOGCONFIG"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_ACTIONGROUPCONFIG"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SNMPTRAPPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_UPGRADE"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_REBOOT"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_TECHSUPPORT"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_TRAFFIC_CAPTURE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLB"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLBGEODBPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLBSERVICE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_TRAFFICCLONEPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SE_TOKEN"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_WAFPOLICY"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_WAFPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_USER_CREDENTIAL"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_AUTOSCALE"
	type = "WRITE_ACCESS"
}
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-Application-Admin-abc"
}
`

const testAccAVIRoleupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_role" "testRole" {
	privileges {
	resource = "PERMISSION_CONTROLLER"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_CLOUD"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_INTERNAL"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_VIRTUALSERVICE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOL"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOLGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_POOLGROUPDEPLOYMENTPOLICY"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_PRIORITYLABELS"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_HEALTHMONITOR"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORKPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_APPLICATIONPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_HTTPPOLICYSET"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_IPADDRGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_STRINGGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_MICROSERVICEGROUP"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SSLPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SSLKEYANDCERTIFICATE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORKSECURITYPOLICY"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_APPLICATIONPERSISTENCEPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ANALYTICSPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_IPAMDNSPROVIDERPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_VSDATASCRIPTSET"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_TENANT"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_PKIPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_CERTIFICATEMANAGEMENTPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_AUTHPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SERVICEENGINE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SERVICEENGINEGROUP"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_NETWORK"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_SYSTEMCONFIGURATION"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_VRFCONTEXT"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_USER"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_ROLE"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERT"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTCONFIG"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTEMAILCONFIG"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_ALERTSYSLOGCONFIG"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_ACTIONGROUPCONFIG"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SNMPTRAPPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_UPGRADE"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_REBOOT"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_TECHSUPPORT"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_TRAFFIC_CAPTURE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLB"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLBGEODBPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_GSLBSERVICE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_TRAFFICCLONEPROFILE"
	type = "WRITE_ACCESS"
}
privileges {
	resource = "PERMISSION_SE_TOKEN"
	type = "NO_ACCESS"
}
privileges {
	resource = "PERMISSION_WAFPOLICY"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_WAFPROFILE"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_USER_CREDENTIAL"
	type = "READ_ACCESS"
}
privileges {
	resource = "PERMISSION_AUTOSCALE"
	type = "WRITE_ACCESS"
}
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-Application-Admin-updated"
}
`
