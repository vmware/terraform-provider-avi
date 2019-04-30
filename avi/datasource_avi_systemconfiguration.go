/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemConfigurationRead,
		Schema: map[string]*schema.Schema{
			"admin_auth_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAdminAuthConfigurationSchema(),
			},
			"default_license_tier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENTERPRISE_18"},
			"dns_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSConfigurationSchema(),
			},
			"dns_virtualservice_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"docker_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"email_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEmailConfigurationSchema(),
			},
			"global_tenant_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTenantConfigurationSchema(),
			},
			"linux_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLinuxConfigurationSchema(),
			},
			"mgmt_ip_access_control": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMgmtIpAccessControlSchema(),
			},
			"ntp_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNTPConfigurationSchema(),
			},
			"portal_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortalConfigurationSchema(),
			},
			"proxy_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"secure_channel_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSecureChannelConfigurationSchema(),
			},
			"snmp_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpConfigurationSchema(),
			},
			"ssh_ciphers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssh_hmacs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"welcome_workflow_complete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
