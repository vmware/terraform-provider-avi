// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudRead,
		Schema: map[string]*schema.Schema{
			"autoscale_polling_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAwsConfigurationSchema(),
			},
			"azure_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAzureConfigurationSchema(),
			},
			"cloudstack_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCloudStackConfigurationSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"custom_tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCustomTagSchema(),
			},
			"dhcp_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_provider_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_resolution_on_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_resolvers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsResolverSchema(),
			},
			"docker_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDockerConfigurationSchema(),
			},
			"east_west_dns_provider_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"east_west_ipam_provider_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_vip_on_all_interfaces": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_vip_static_routes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGCPConfigurationSchema(),
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipam_provider_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"linuxserver_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLinuxServerConfigurationSchema(),
			},
			"maintenance_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"mtu": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsxt_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceNsxtConfigurationSchema(),
			},
			"obj_name_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"openstack_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOpenStackConfigurationSchema(),
			},
			"prefer_static_routes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"rancher_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRancherConfigurationSchema(),
			},
			"se_group_template_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_based_dns_registration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vca_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcevCloudAirConfigurationSchema(),
			},
			"vcenter_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcevCenterConfigurationSchema(),
			},
			"vmc_deployment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
