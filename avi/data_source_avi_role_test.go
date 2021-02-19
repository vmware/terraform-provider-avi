package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceRoleBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSRoleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_role.testRole", "name", "test-Application-Admin-abc"),
				),
			},
		},
	})

}

const testAccAVIDSRoleConfig = `
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

data "avi_role" "testRole" {
    name= "${avi_role.testRole.name}"
}
`
