// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudRead,
		Schema: map[string]*schema.Schema{
			"apic_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAPICConfigurationSchema(),
			},
			"apic_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"autoscale_polling_interval": {
				Type:     schema.TypeInt,
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
			"custom_tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCustomTagSchema(),
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dns_provider_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_resolution_on_se": {
				Type:     schema.TypeBool,
				Computed: true,
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
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_vip_static_routes": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"gcp_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGCPConfigurationSchema(),
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
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
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsx_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceNsxConfigurationSchema(),
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
			"oshiftk8s_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOShiftK8SConfigurationSchema(),
			},
			"prefer_static_routes": {
				Type:     schema.TypeBool,
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
				Type:     schema.TypeBool,
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
			"vtype": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
