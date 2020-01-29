/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemConfigurationRead,
		Schema: map[string]*schema.Schema{
			"admin_auth_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAdminAuthConfigurationSchema(),
			},
			"default_license_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDNSConfigurationSchema(),
			},
			"dns_virtualservice_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"docker_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceEmailConfigurationSchema(),
			},
			"global_tenant_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTenantConfigurationSchema(),
			},
			"linux_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLinuxConfigurationSchema(),
			},
			"mgmt_ip_access_control": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceMgmtIpAccessControlSchema(),
			},
			"ntp_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceNTPConfigurationSchema(),
			},
			"portal_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortalConfigurationSchema(),
			},
			"proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"secure_channel_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSecureChannelConfigurationSchema(),
			},
			"snmp_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSnmpConfigurationSchema(),
			},
			"ssh_ciphers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssh_hmacs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"welcome_workflow_complete": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
