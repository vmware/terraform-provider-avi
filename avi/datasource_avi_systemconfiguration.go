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
				Set: func(v interface{}) int {
					return 0
				},
			},
			"default_license_tier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENTERPRISE_18"},
			"dns_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
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
				Set: func(v interface{}) int {
					return 0
				},
			},
			"global_tenant_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTenantConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"linux_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLinuxConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mgmt_ip_access_control": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMgmtIpAccessControlSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ntp_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNTPConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"portal_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortalConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proxy_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"secure_channel_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSecureChannelConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"snmp_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
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
		},
	}
}
