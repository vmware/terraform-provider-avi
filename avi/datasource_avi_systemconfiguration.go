// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemConfigurationRead,
		Schema: map[string]*schema.Schema{
			"admin_auth_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAdminAuthConfigurationSchema(),
			},
			"common_criteria_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"controller_analytics_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerAnalyticsPolicySchema(),
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceEmailConfigurationSchema(),
			},
			"enable_cors": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fips_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"global_tenant_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTenantConfigurationSchema(),
			},
			"host_key_algorithm_exclude": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kex_algorithm_exclude": {
				Type:     schema.TypeString,
				Computed: true,
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
			"sddcmanager_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
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
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
