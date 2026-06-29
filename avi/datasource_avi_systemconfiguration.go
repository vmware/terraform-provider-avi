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
			"allow_legacy_sha1_ntp_auth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_private_ips": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avi_email_login_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_security_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCertificateSecurityPolicySchema(),
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
			"enable_host_header_check": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_license_quota": {
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
			"intelligent_assist_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kex_algorithm_exclude": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"legacy_ssl_support": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_quota": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceQuotaConfigSchema(),
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
			"password_policy_managed_at_ops": {
				Type:     schema.TypeString,
				Computed: true,
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
			"rekey_time_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rekey_volume_limit": {
				Type:     schema.TypeString,
				Computed: true,
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
			"service_auth_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceServiceAuthConfigurationSchema(),
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
			"sync_kex_host_to_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_syslog_to_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"syslog_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"telemetry_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTelemetryConfigurationSchema(),
			},
			"trusted_host_profiles_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"truststore_pkiprofile_ref": {
				Type:     schema.TypeString,
				Computed: true,
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
