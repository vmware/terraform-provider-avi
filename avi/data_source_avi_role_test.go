// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

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

data "avi_role" "testRole" {
    name= "${avi_role.testRole.name}"
}
`
