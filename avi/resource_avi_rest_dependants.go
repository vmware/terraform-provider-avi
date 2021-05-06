// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceClusterNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"categories": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "node",
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"public_ip_or_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vm_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_mor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugVirtualServiceSeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceControllerLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_burst_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"burst_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cores": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"customer_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_enforcement": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"expired_burst_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_tiers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCumulativeLicenseSchema(),
			},
			"licenses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSingleLicenseSchema(),
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_on": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_until": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHostUnavailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reasons": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPStatusRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"begin": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"end": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceSeGroupVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLocationHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceWafPSMMatchElementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"excluded": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_element": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpamDnsOCIProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_credentials_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenancy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceContainerCloudServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceFloatingIpSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeHmEventGSDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gslb_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDosThresholdProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"thresh_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDosThresholdSchema(),
			},
			"thresh_period": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceVsInitialPlacementEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLdapDirectorySettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_bind_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_filter": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "(objectClass=*)",
			},
			"group_member_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "member",
			},
			"group_member_is_full_dn": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"group_search_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_search_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_SUBTREE",
			},
			"ignore_referrals": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"user_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_search_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_search_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_ONE",
			},
		},
	}
}

func ResourceSeHmEventGslbPoolMemberDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema(),
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"shm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema(),
			},
			"ssl_error_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSEBandwidthLimitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmModifyVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceScaleStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_success": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"end_time_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_se_assigned": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se_requested": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"placement_read_fail_cnt": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason_code_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scale_se": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_placement_resolution_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVipPlacementResolutionInfoSchema(),
			},
		},
	}
}

func ResourceKeyValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGcpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDsrProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dsr_encap_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENCAP_IPINIP",
			},
			"dsr_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DSR_TYPE_L3",
			},
		},
	}
}

func ResourceGslbPerDnsStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"geo_download": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
			},
			"gslb_download": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"placement_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSubDomainPlacementRuntimeSchema(),
			},
			"se_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_dns_vs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceWafWhitelistLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsDbQueueFullEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"instanceport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"low": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodeid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbRuntimeSchema(),
			},
			"watermark": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceFailActionHTTPLocalResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FAIL_HTTP_STATUS_CODE_503",
			},
		},
	}
}

func ResourceWafRuleMatchDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_internal": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match_element": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRebalanceScaleoutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scaleout_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleoutParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCifSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cif": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPSecurityActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"https_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rate_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPSecurityActionRateProfileSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOpenStackSeVmChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceUpgradeOpsEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeOpsParamSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeStatusInfoSchema(),
			},
			"upgrade_ops": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVIDCInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceSSLRenewFailedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafExclusionTypeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SENSITIVE",
			},
			"match_op": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EQUALS",
			},
		},
	}
}

func ResourceGslbThirdPartySiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"site_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
			},
		},
	}
}

func ResourceUserActivitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"concurrent_sessions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"failed_login_attempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"last_login_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_login_timestamp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_password_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logged_in": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"previous_password": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSamlServiceProviderNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"signing_ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_signon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "USER_DEFINED",
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafRuleGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleSchema(),
			},
		},
	}
}

func ResourceIpamDnsInternalProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_service_domain": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsServiceDomainSchema(),
			},
			"dns_virtualservice_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"usable_network_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceApicAgentVsNetworkErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsRrSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cname": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsAAAARdataSchema(),
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsARdataSchema(),
			},
			"nses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsNsRdataSchema(),
			},
			"ttl": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGeoLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNetworkSubnetClashSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_nw": {
				Type:     schema.TypeString,
				Required: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAwsZoneNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usable_network_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceBgpPeerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_snat_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"advertise_vip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"advertisement_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"bfd": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connect_timer": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ebgp_multihop": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"hold_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keepalive_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"local_as": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"md5_secret": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"peer_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"remote_as": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"shutdown": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIpamDnsAwsProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"egress_service_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"iam_assume_role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_vip_to_public_zone": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret_access_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"usable_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_network_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_iam_roles": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vpc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAwsZoneNetworkSchema(),
			},
		},
	}
}

func ResourceStaticRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"next_hop": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"prefix": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"route_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceFailActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_rsp": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFailActionHTTPLocalResponseSchema(),
			},
			"redirect": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFailActionHTTPRedirectSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceApplicationLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"all_request_headers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"all_response_headers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_response_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"body_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NOT_UPDATED",
			},
			"cache_hit": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cacheable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cipher_bytes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_browser": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cipher_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLCipherListSchema(),
			},
			"client_dest_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_device": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_insights": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_ip": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_log_filter_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_os": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_rtt": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_src_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compression_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"connection_error_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConnErrorInfoSchema(),
			},
			"data_transfer_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"datascript_error_trace": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDataScriptErrorTraceSchema(),
			},
			"datascript_log": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers_received_from_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers_sent_to_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2_stream_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_request_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_response_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"microservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"microservice_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"paa_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePaaLogSchema(),
			},
			"persistence_used": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"persistent_session_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirected_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"referer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_timestamp": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"request_content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_headers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"request_served_locally_remote_site_down": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"request_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"response_content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_headers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"response_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"response_time_first_byte": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"response_time_last_byte": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rewritten_uri_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rewritten_uri_query": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_auth_request_generated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"saml_auth_response_received": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"saml_auth_session_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"saml_authentication_used": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"saml_session_cookie_valid": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_conn_src_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_conn_src_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_connection_reused": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_dest_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_response_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_response_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_response_time_first_byte": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_response_time_last_byte": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_rtt": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_side_redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_src_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_ssl_session_reused": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"servers_tried": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_engine": {
				Type:     schema.TypeString,
				Required: true,
			},
			"significance": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"significant": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"significant_log": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sni_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spdy_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cipher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udf": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri_query": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcpu_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtualservice": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vs_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceWafLogSchema(),
			},
			"xff": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSamlServiceProviderSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"org_display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"org_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"org_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_entity_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSamlServiceProviderNodeSchema(),
			},
			"tech_contact_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tech_contact_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServerIdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceInfobloxSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceMetricLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"report_timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"time_series": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsQueryResponseSchema(),
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourceWafConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_http_versions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_methods": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_request_content_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"argument_separator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "&",
			},
			"client_request_max_body_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"confidence_override": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAppLearningConfidenceOverrideSchema(),
			},
			"cookie_format_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"enable_auto_rule_updates": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ignore_incomplete_request_body_error": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"learning_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAppLearningParamsSchema(),
			},
			"max_execution_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"min_confidence": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CONFIDENCE_VERY_HIGH",
			},
			"regex_match_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"regex_recursion_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"request_body_default_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:2,deny,status:403,log,auditlog",
			},
			"request_hdr_default_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:1,deny,status:403,log,auditlog",
			},
			"response_body_default_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:4,deny,status:403,log,auditlog",
			},
			"response_hdr_default_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:3,deny,status:403,log,auditlog",
			},
			"restricted_extensions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"restricted_headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"server_response_max_body_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"static_extensions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status_code_for_rejected_requests": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_RESPONSE_CODE_403",
			},
		},
	}
}

func ResourceIpAddrPortSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceUDPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"session_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceUpgradeOpsStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpamDnsGCPProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_se_group_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"network_host_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usable_network_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_gcp_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAutoScaleMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_aws_autoscale_integration": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"intelligent_autoscale_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceL4ConnectionPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceL4RuleSchema(),
			},
		},
	}
}

func ResourceGCPNetworkConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config": {
				Type:     schema.TypeString,
				Required: true,
			},
			"inband": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPInBandManagementSchema(),
			},
			"one_arm": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPOneArmModeSchema(),
			},
			"two_arm": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPTwoArmModeSchema(),
			},
		},
	}
}

func ResourceDnsRuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsClientIpMatchSchema(),
			},
			"geo_location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsGeoLocationMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsTransportProtocolMatchSchema(),
			},
			"query_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsQueryNameMatchSchema(),
			},
			"query_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsQueryTypeMatchSchema(),
			},
		},
	}
}

func ResourceMetricsDimensionDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dimension": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dimension_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVipAutoscaleGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVipAutoscaleConfigurationSchema(),
			},
			"policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVipAutoscalePolicySchema(),
			},
		},
	}
}

func ResourceRancherConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_sync_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"container_port_match_http_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"enable_event_subscription": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_container_port_as_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_container_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"nuage_controller": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNuageSDNControllerSchema(),
			},
			"rancher_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_deployment_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH",
			},
			"se_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_spawn_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"secret_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"services_accessible_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_container_ip_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeRuntimeCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_low_rtt": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"min_high_rtt": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  200,
			},
			"min_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"mobile_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRollbackSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourcePerformanceLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_concurrent_connections": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_throughput": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDockerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"ca_tls_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_tls_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_port_match_http_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"enable_event_subscription": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_container_port_as_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_container_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"se_deployment_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH",
			},
			"se_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_spawn_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"services_accessible_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ucp_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_container_ip_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceAppLearningParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_per_uri_learning": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"max_params": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"max_uris": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  500,
			},
			"min_hits_to_learn": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"sampling_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"update_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
		},
	}
}

func ResourceCPUUsagePerNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCPUUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceEquivalentLabelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceURIParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tokens": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceURIParamTokenSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRollbackPatchControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSCPoolServerStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_server": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSCServerStateInfoSchema(),
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
		},
	}
}

func ResourceHTTPSecurityActionRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterActionSchema(),
			},
			"per_client_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"per_uri_path": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rate_limiter": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterSchema(),
			},
		},
	}
}

func ResourceCloudTenantsDeletedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenants": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudTenantCleanupSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSingleLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cores": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"created_on": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customer_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enforced_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"expired": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"serial_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_on": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_until": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHealthScoreDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"anomaly_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"performance_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"performance_score": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"previous_value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resources_penalty": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resources_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_penalty": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"security_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sub_resource_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourceSnmpTrapServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"ip_addr": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  162,
			},
			"user": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSnmpV3UserParamsSchema(),
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}

func ResourceClusterNodeStartedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePodTolerationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EQUAL",
			},
			"toleration_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpCommunitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip_begin": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_end": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceTCPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_protocol_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"proxy_protocol_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PROXY_PROTOCOL_VERSION_1",
			},
			"ssl_client_certificate_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CLIENT_CERTIFICATE_NONE",
			},
		},
	}
}

func ResourceDnsCnameRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGCPCloudRouterUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"router_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"api_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cntr_public_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KVM",
			},
			"mgmt_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret_access_key": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceSSHSeDeploymentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceVIMgrClusterRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datacenter_managed_object_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNetworkRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnet_runtime": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSubnetRuntimeSchema(),
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
		},
	}
}

func ResourceVsMigrateParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"to_host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_new_se": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"to_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugFilterUnionSchema(),
			},
			"log_level": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_module": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_level": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcevCenterConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"management_network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"privilege": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_template_se_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudHealthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_fail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_fail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_ok": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_fails": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmRebootSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugVrfContextSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_buffer_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"command_buffer_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32768,
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVrfSchema(),
			},
		},
	}
}

func ResourceConfigSeGrpFlvUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"new_flv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"old_flv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeThreshEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_value": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"thresh": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourcePropertySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHSMSafenetLunaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_group_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"is_ha": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"node_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMSafenetClientInfoSchema(),
			},
			"server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMSafenetLunaServerSchema(),
			},
			"server_pem": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_dedicated_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSidebandProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"sideband_max_request_body_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
		},
	}
}

func ResourcePatchControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_patch_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceNuageSDNControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nuage_organization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nuage_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"nuage_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nuage_vsd_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_policy_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeVersionCheckFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecurityMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVIMgrControllerRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIControllerVnicInfoSchema(),
			},
		},
	}
}

func ResourceCdpLldpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chassis": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmtaddr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_info_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVipAutoscaleConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipAutoscaleZonesSchema(),
			},
		},
	}
}

func ResourceNatMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServiceMatchSchema(),
			},
			"source_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
		},
	}
}

func ResourceVIMgrDCRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"interested_nws": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"interested_vms": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"inventory_state": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nw_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pending_vcenter_reqs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sevm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVsEvStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free_ip_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_alloced": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAllocInfoSchema(),
			},
			"prefix": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"total_ip_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"used_ip_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbThirdPartySiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hm_proxies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsOptRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dnssec_ok": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsEdnsOptionSchema(),
			},
			"udp_payload_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbGeoDbEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceGslbGeoDbFileSchema(),
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceVIControllerVnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"portgroup": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vnic_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIGuestvNicIPAddrSchema(),
			},
		},
	}
}

func ResourceVipAutoscalePolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_cooldown": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"max_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"min_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"suspend": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceConnErrorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_syn_retransmit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"out_of_orders": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmits": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_out_of_orders": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_retransmits": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_rx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_tx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_zero_window_size_events": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zero_window_size_events": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecMgrThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceTCPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dsr_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDsrProfileSchema(),
			},
			"enable_syn_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"session_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
		},
	}
}

func ResourceSeAgentStateCachePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sc_batch_buffer_flush_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"sc_shard_cleanup_max_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
			},
			"sc_state_ring_batch_dequeue_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"sc_states_flush_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sc_stream_check_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5000,
			},
			"sc_thread_q_batch_dequeue_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sc_thread_q_max_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"sc_thread_sleep_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceSSLKeyECParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curve": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_EC_CURVE_SECP256R1",
			},
		},
	}
}

func ResourceControllerUpgradeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_progress": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"rollback": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tasks_completed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceUpgradeTaskSchema(),
			},
		},
	}
}

func ResourceDnsRuleRLActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RL_ACTION_NONE",
			},
		},
	}
}

func ResourceVIMgrVcenterRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apic_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datacenter_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disc_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disc_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovered_datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inventory_progress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inventory_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_clusters": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_dcs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_hosts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_nws": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vcenter_req_pending": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vms": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_connected": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcenter_fullname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_template_se_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOpenStackHypervisorPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hypervisor": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePropertySchema(),
			},
		},
	}
}

func ResourceAuthMatchGroupMembershipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceConfigInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queue": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema(),
			},
			"reader_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"writer_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClusterServiceCriticalFailureEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAwsConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"asg_poll_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"ebs_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAwsEncryptionSchema(),
			},
			"free_elasticips": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"iam_assume_role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_vip_to_public_zone": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "us-west-1",
			},
			"route53_integration": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"s3_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAwsEncryptionSchema(),
			},
			"secret_access_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"sqs_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAwsEncryptionSchema(),
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"use_iam_roles": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_sns_sqs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vpc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAwsZoneConfigSchema(),
			},
		},
	}
}

func ResourcePGDeploymentRuleResultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_value": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"result": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourcePGDeploymentRuleSchema(),
			},
		},
	}
}

func ResourceRoutingServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_backend_networks": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_routing": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_vip_on_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_vmac": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"floating_intf_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_intf_ip_se_2": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"flowtable_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFlowtableProfileSchema(),
			},
			"graceful_restart": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nat_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_by_linux_ipstack": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceUpgradeSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHTTPServerReselectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"num_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"retry_nonidempotent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"retry_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"svr_resp_code": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPReselectRespCodeSchema(),
			},
		},
	}
}

func ResourceSCServerStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"server_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dns_vs_states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPerDnsStateSchema(),
			},
			"gs_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDnsGsStatusSchema(),
			},
			"retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricStatisticsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_sample": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_ts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mean": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"min": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"min_ts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_samples": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sum": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"trend": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAwsEncryptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AWS_ENCRYPTION_MODE_NONE",
			},
		},
	}
}

func ResourceHealthMonitorSSLAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbSiteRuntimeCfgSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fd_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"gap_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"geo_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"ghm_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"glb_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"gpki_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"gs_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"mm_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"repl_queue": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"sync_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteCfgSyncInfoSchema(),
			},
		},
	}
}

func ResourceNetworkSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"total": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"used": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAlertRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_app_log_rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAlertFilterSchema(),
			},
			"event_match_filter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics_rule": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleMetricSchema(),
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPERATOR_AND",
			},
			"sys_event_rule": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleEventSchema(),
			},
		},
	}
}

func ResourceJobEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expires_at": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subjobs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSubJobSchema(),
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
		},
	}
}

func ResourceGslbDnsSeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceOpenStackLbPluginOpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elapsed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prov": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAuthProfileHTTPClientParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_expiration_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"request_header": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"require_user_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGslbServiceDownResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fallback_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"fallback_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_DOWN_RESPONSE_NONE",
			},
		},
	}
}

func ResourceDNSConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsErrorEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVCASetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVIMgrNWRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_vrf_context": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_expand": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dvs": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_nw": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ip_subnet": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema(),
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmtnw": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_ports": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_range": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanRangeSchema(),
			},
			"vm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeVsPktBufHighEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeLicensedBandwdithExceededEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_pkts_dropped": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsDbRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"db_num_client_queries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_num_client_resp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_num_db_queries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_num_db_resp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_num_oom": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_queue_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_rum_queries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_rum_rows": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmDeleteSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTcpAttacksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceClusterServiceFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVsApicExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"txn_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema(),
			},
		},
	}
}

func ResourceAlertRuleEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_details": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventDetailsFilterSchema(),
			},
			"event_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"not_cond": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceMarathonConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"framework_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marathon_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"marathon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:8080",
			},
			"marathon_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_port_range": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"public_port_range": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin",
			},
			"use_token_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_name_tag_framework": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceOshiftSharedVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"virtualservice_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeHmEventPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"percent_servers_up": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventServerDetailsSchema(),
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCompressionFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"devices_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_addr_prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ip_addr_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
			"ip_addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_addrs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"level": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IS_IN",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_agent": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHTTPHdrActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cookie": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPCookieDataSchema(),
			},
			"hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPHdrDataSchema(),
			},
		},
	}
}

func ResourcePoolDeploymentFailureInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_in_service_pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"curr_in_service_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"results": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleResultSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugVsDataplaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloneServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceMemoryUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"total": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSchedulerActionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_uri": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"control_script_output": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"execution_datestamp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduler_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceContainerCloudBatchSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ccs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceContainerCloudSetupSchema(),
			},
		},
	}
}

func ResourceCCScaleSetNotifDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scaleset_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceReplaceStringVarSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsRateLimiterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceDnsRuleRLActionSchema(),
			},
			"rate_limiter_object": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterSchema(),
			},
		},
	}
}

func ResourceMetricsDataSeriesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDataSchema(),
			},
			"header": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceMetricsDataHeaderSchema(),
			},
		},
	}
}

func ResourceNatPolicyActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nat_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNatAddrInfoSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceSSLExportDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGCPVIPILBSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_router_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRebalanceMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceServerAutoScaleInCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nscalein": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scaled_in_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceVIMgrSEVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAzureInfoSchema(),
			},
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_ip_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creation_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"deletion_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"discovery_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovery_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disk_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_nic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"init_vnics": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_discovery": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segroup_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_rm_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_se_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_template_vm": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vcenter_vappname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_vappvendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_vm_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPReselectRespCodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"codes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPStatusRangeSchema(),
			},
			"resp_code_block": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGCPEncryptionKeysSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gcs_bucket_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcs_objects_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_disk_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_image_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAppLearningConfidenceOverrideSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"confid_high_value": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9500,
			},
			"confid_low_value": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  7500,
			},
			"confid_probable_value": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9000,
			},
			"confid_very_high_value": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9999,
			},
		},
	}
}

func ResourcePatchDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"patch_image_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_image_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVipAutoscaleZonesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClusterServiceRestoredEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRateLimiterActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"redirect": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_LOCAL_RESPONSE_STATUS_CODE_429",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RL_ACTION_NONE",
			},
		},
	}
}

func ResourceHSMThalesRFSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}

func ResourceAllSeUpgradeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeParamsSchema(),
			},
		},
	}
}

func ResourceAzureServicePrincipalCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_token": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudSeVmChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceUDPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dsr_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDsrProfileSchema(),
			},
			"per_pkt_loadbalance": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"session_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"snat": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceIngAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceConfigUserAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policies": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenants": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHardwareSecurityModuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloudhsm": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHSMAwsCloudHsmSchema(),
			},
			"nethsm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMThalesNetHsmSchema(),
			},
			"rfs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHSMThalesRFSSchema(),
			},
			"sluna": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHSMSafenetLunaSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGCPSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nhop_inst": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nhop_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudClusterVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceL4RuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RuleActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RuleMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceURIParamQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_query": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceAzureConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cloud_credentials_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_info": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceAzureNetworkInfoSchema(),
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_azure_dns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_enhanced_ha": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_managed_disks": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_standard_alb": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores_per_socket": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disk": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hyper_threading": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_vcpus": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_sync_complete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
		},
	}
}

func ResourceDebugSeFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"faults": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeFaultSchema(),
			},
			"se_malloc_fail_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_malloc_fail_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_mbuf_cl_sanity": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_shm_malloc_fail_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_shm_malloc_fail_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_waf_alloc_fail_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_waf_learning_alloc_fail_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceDnsTxtRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text_str": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsScaleinParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scalein_primary": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIPPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_persistent_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}

func ResourceClusterNodeDbFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClientLogStreamingFormatSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"format": {
				Type:     schema.TypeString,
				Required: true,
			},
			"included_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourcePsmProgramDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClusterWarmRebootEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceVsErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_timestamp": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceConfigUserLogoutSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNetworkSecurityRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"age": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceNetworkSecurityMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rl_param": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNetworkSecurityPolicyActionRLParamSchema(),
			},
		},
	}
}

func ResourceVSDataScriptsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"vs_datascript_set_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNatAddrInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"nat_ip_range": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceSeHmEventServerDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"shm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema(),
			},
			"ssl_error_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugSeDataplaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceControllerVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloudASGNotifDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asg_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLinuxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"banner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cis_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"motd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceResponseMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"cookie": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCookieMatchSchema(),
			},
			"hdrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"loc_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLocationHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceProtocolMatchSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceQueryMatchSchema(),
			},
			"rsp_hdrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPStatusMatchSchema(),
			},
			"version": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchSchema(),
			},
		},
	}
}

func ResourceCustomTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tag_val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTencentZoneNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usable_subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceUpgradeOpsParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_resume_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupResumeOptionsSchema(),
			},
		},
	}
}

func ResourceConfigUserLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"remote_attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSensitiveLogProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"header_field_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSensitiveFieldRuleSchema(),
			},
			"waf_field_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSensitiveFieldRuleSchema(),
			},
		},
	}
}

func ResourceRollbackSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceOpenStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": {
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowed_address_pairs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"anti_affinity": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_drive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"contrail_disable_policy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"contrail_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contrail_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"custom_se_image_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePropertySchema(),
			},
			"external_networks": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"free_floatingips": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KVM",
			},
			"hypervisor_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackHypervisorPropertiesSchema(),
			},
			"img_format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OS_IMG_FMT_AUTO",
			},
			"import_keystone_tenants": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"insecure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"keystone_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"map_admin_to_cloudadmin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mgmt_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name_owner": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"neutron_rbac": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"nuage_organization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nuage_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"nuage_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nuage_virtualip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nuage_vsd_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"port_security": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"privilege": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prov_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"provider_vip_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackVipNetworkSchema(),
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_mapping": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackRoleMappingSchema(),
			},
			"security_groups": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_se": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_admin_url": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_internal_endpoints": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_keystone_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_nuagevip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbPoolMemberRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"datapath_status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberDatapathStatusSchema(),
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_value_to_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_value_to_se": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"oper_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServiceSchema(),
			},
			"site_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_pools": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbServiceSitePersistencePoolSchema(),
			},
			"vip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserver_l4_metrics": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVserverL4MetricsObjSchema(),
			},
			"vserver_l7_metrics": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVserverL7MetricsObjSchema(),
			},
		},
	}
}

func ResourceDnsEdnsOptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr_family": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_prefix_len": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_prefix_len": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"subnet_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"subnet_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudTenantCleanupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_ports": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_secgrp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_svrgrp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsTransportProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePackageDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"build": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBuildInfoSchema(),
			},
			"hash": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePatchInfoSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAuthMatchAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSecurityMgrRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"thresholds": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSecMgrThresholdSchema(),
			},
		},
	}
}

func ResourceGslbSiteCfgSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"errored_objects": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema(),
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"sync_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcevNICSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregator_chgd": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"can_se_dp_takeover": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"del_pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delete_vnic": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dp_deletion_done": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_asm": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_avi_internal_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_hsm": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_mgmt": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_portchannel": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"link_up": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMemberInterfaceSchema(),
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pci_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vlan_interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanInterfaceSchema(),
			},
			"vnic_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema(),
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vrf_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceL4PoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"l4_policy_set_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNsxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_nsx_prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsx_manager_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsx_manager_password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"nsx_manager_username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsx_poll_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
		},
	}
}

func ResourcePatchSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_patch_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_patch_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHealthMonitorSIPSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sip_monitor_transport": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SIP_UDP_PROTO",
			},
			"sip_request_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SIP_OPTIONS",
			},
			"sip_response": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SIP/2.0",
			},
		},
	}
}

func ResourceDnsRuleActionResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authoritative": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"rcode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RCODE_NOERROR",
			},
			"resource_record_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRuleDnsRrSetSchema(),
			},
			"truncation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceIpamDnsAzureProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"azure_serviceprincipal": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			},
			"azure_userpass": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAzureUserPassCredentialsSchema(),
			},
			"egress_service_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usable_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_network_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_enhanced_ha": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_standard_alb": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"virtual_network_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSnmpV3UserParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_passphrase": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "avinetworks",
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_AUTH_MD5",
			},
			"priv_passphrase": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "avinetworks",
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"priv_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_PRIV_DES",
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbDnsUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gslb_geo_db_profile_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gslb_service_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gslb_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"obj_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbObjectInfoSchema(),
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeUpgradeEventsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se_group": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVipPlacementResolutionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsAAAARdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip6_address": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceVinfraVcenterObjDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAwsZoneConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpAddrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceVsScaleInEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceScaleStatusSchema(),
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeHmEventGslbPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gsgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslb_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gsmember": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventGslbPoolMemberDetailsSchema(),
			},
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNatRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceNatPolicyActionSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceNatMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceInternalGatewayMonitorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gateway_monitor_failure_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"gateway_monitor_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"gateway_monitor_success_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
		},
	}
}

func ResourceSSLCertificateDescriptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"common_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distinguished_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locality": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpamDnsTencentProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_credentials_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usable_subnet_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTencentZoneNetworkSchema(),
			},
		},
	}
}

func ResourceHSMSafenetLunaServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"partition_passwd": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"partition_serial_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_cert": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMemoryBalancerInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"child": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceChildProcessInfoSchema(),
			},
			"controller_memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_used": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"process": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudAutoscalingConfigFailureDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeMgrEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGcpInfoSchema(),
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"migrate_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceUpgradeEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"sub_tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceMetricsDbQueueHealthyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"instanceport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"low": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodeid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbRuntimeSchema(),
			},
			"watermark": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafPolicyWhitelistRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraVcenterConnectivityStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceResumeSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupResumeOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeIpAddedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ns": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTencentCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"secret_key": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceMesosConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_vses_are_feproxy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"app_sync_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"container_port_match_http_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_gs_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"enable_event_subscription": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_bridge_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cbr1",
			},
			"feproxy_container_port_as_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_route_publish": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFeProxyRoutePublishConfigSchema(),
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_container_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"marathon_configurations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMarathonConfigurationSchema(),
			},
			"marathon_se_deployment": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMarathonSeDeploymentSchema(),
			},
			"mesos_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:5050",
			},
			"node_availability_zone_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nuage_controller": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNuageSDNControllerSchema(),
			},
			"se_deployment_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "MESOS_SE_CREATE_FLEET",
			},
			"se_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosSeResourcesSchema(),
			},
			"se_spawn_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi/se",
			},
			"services_accessible_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_bridge_ip_as_vip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_container_ip_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_vips_for_east_west_services": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceSensitiveFieldRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOG_FIELD_REMOVE",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAuthorizationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attr_matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthAttributeMatchSchema(),
			},
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
		},
	}
}

func ResourceHTTPRewriteURLActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamQuerySchema(),
			},
		},
	}
}

func ResourceControllerPortalAssetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asset_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flr_state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolRuntimeSchema(),
			},
			"ldr_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"send_event": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"send_status": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"services_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDiskUsagePerNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disk_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDiskUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTencentSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSSLRatingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compatibility_rating": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"performance_rating": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_score": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClientLogStreamingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_server": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"format_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClientLogStreamingFormatSchema(),
			},
			"log_types_to_send": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_ALL",
			},
			"max_logs_per_second": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOG_STREAMING_PROTOCOL_UDP",
			},
			"syslog_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStreamingSyslogConfigSchema(),
			},
		},
	}
}

func ResourceServerAutoScaleFailedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_scalein_servers": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAuthMappingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthMatchAttributeSchema(),
			},
			"group_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthMatchGroupMembershipSchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"object_access_policy_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"policy_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRmSpawnSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"network_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRateLimiterProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_connections_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_failed_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_scanners_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_to_uri_failed_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_to_uri_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"custom_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"http_header_rate_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_failed_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_scanners_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
		},
	}
}

func ResourceVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocate_floating_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_allocate_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_allocate_ip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V4_ONLY",
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avi_allocated_fip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"avi_allocated_vip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"discovered_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"floating_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_subnet6_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"floating_subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ipam_network_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIPNetworkSubnetSchema(),
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"placement_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipPlacementNetworkSchema(),
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHSMAwsCloudHsmSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crypto_user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crypto_user_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"hsm_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVIMgrIPSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fip_available": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fip_subnet_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floatingip_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceFloatingIpSubnetSchema(),
			},
			"ip_subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"primary": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ref_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_ref_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePortalConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_basic_authentication": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"api_force_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  24,
			},
			"disable_remote_cli_shell": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_swagger": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_clickjacking_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_http": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_https": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"http_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"https_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password_strength_check": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"redirect_to_https": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"sslkeyandcertificate_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sslprofile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_uuid_from_input": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceGslbSiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"glb_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rxed_site_hs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteHealthStatusSchema(),
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"site_cfg": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeCfgSchema(),
			},
			"site_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
			},
			"site_stats": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeStatsSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"view_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudStackSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGCPVIPRoutesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_se_group_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceCaptureFileSizeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"absolute_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"percentage_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceAppHdrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr_match_case": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hdr_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hdr_string_op": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafPolicyWhitelistSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPolicyWhitelistRuleSchema(),
			},
		},
	}
}

func ResourceAttackMitigationActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deny": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceSeMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGCPCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_account_keyfile_data": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceSeMemoryLimitEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_memory_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heap_config_hard_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heap_config_soft_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heap_config_usage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heap_conn_usage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shm_config_hard_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"shm_config_soft_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"shm_config_usage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"shm_conn_usage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSiteRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_leader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDnsInfoSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"event_cache": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceEventCacheSchema(),
			},
			"hs_state": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_of_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_NOT_A_MEMBER",
			},
			"rrtoken": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"site_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SITE_STATE_NULL",
			},
			"state_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Not-Initialized",
			},
		},
	}
}

func ResourceAlertSyslogServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anon_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSLOG_LEGACY",
			},
			"pkiprofile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syslog_server": {
				Type:     schema.TypeString,
				Required: true,
			},
			"syslog_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"tls_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"udp": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func ResourceOpenStackIpChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mac_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCertificateAuthoritySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ca_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHSMSafenetClientInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chrystoki_conf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_priv_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"session_major_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"session_minor_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAzureUserPassCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMatchReplacePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_string": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replacement_string": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceReplaceStringVarSchema(),
			},
		},
	}
}

func ResourceGslbSiteHealthStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_gsinfo": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"datapath_gsinfo": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"dns_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDnsInfoSchema(),
			},
			"gap_table": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"geo_table": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"ghm_table": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"glb_table": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"gs_table": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"sw_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAlertMetricThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMarathonSeDeploymentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_image": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fedora",
			},
			"host_os": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "COREOS",
			},
			"resource_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceL4RulePortMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceCloudInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cca_props": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCC_AgentPropertiesSchema(),
			},
			"controller_props": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: ResourceControllerPropertiesSchema(),
				},
			},
			"flavor_props": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudFlavorSchema(),
			},
			"flavor_regex_filter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htypes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vtype": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDataScriptErrorTraceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stack_trace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPHdrDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPHdrValueSchema(),
			},
		},
	}
}

func ResourceStateCacheMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMesosAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_policy_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceCloudConnectorDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_se_reboot": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePartitionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quota": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceIptableRuleSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIptableRuleSchema(),
			},
			"table": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceFullClientLogsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"throttle": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceSePersistenceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePoolGroupMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func ResourceMarathonServicePortConflictSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marathon_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceApicAgentGenericEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"contract_graphs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lif_cif_attachment": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lifs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_engine_vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_network_attachment": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVsScaleoutParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_up": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"new_vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"to_host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_new_se": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"to_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHSMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_hs_db_writes": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsDbDiskEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_deleted_tables": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metrics_free_sz": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"metrics_quota": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceL4RuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RulePortMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RuleProtocolMatchSchema(),
			},
		},
	}
}

func ResourceHealthMonitorExternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"command_parameters": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_variables": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugServiceEngineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"capture_filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCaptureFiltersSchema(),
			},
			"capture_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugVirtualServiceCaptureSchema(),
			},
			"cpu_shares": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeCpuSharesSchema(),
			},
			"debug_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"enable_kdump": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fault": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugSeFaultSchema(),
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeDataplaneSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VM name unknown",
			},
			"seagent_debug": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeAgentSchema(),
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
		},
	}
}

func ResourceOperationalStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason_code_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPER_UNAVAIL",
			},
		},
	}
}

func ResourceSystemUpgradeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceControllerUpgradeStateSchema(),
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"from_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_patch": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rollback": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeStatusSummarySchema(),
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceBurstResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accounted_license_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_alert_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmUnbindVsSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVersionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ds_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ops": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugFilterUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAlertMgrDebugFilterSchema(),
			},
			"autoscale_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAutoScaleMgrDebugFilterSchema(),
			},
			"cloud_connector_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudConnectorDebugFilterSchema(),
			},
			"hs_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHSMgrDebugFilterSchema(),
			},
			"mesos_metrics_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMesosMetricsDebugFilterSchema(),
			},
			"metrics_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsMgrDebugFilterSchema(),
			},
			"metricsapi_srv_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsApiSrvDebugFilterSchema(),
			},
			"se_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeMgrDebugFilterSchema(),
			},
			"se_rpc_proxy_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeRpcProxyDebugFilterSchema(),
			},
			"securitymgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecurityMgrDebugFilterSchema(),
			},
			"state_cache_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStateCacheMgrDebugFilterSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsDebugFilterSchema(),
			},
		},
	}
}

func ResourceCaptureIPCSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flow_del_probe": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_mirror_add": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_mirror_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_mirror_del": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_probe": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_probe_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipc_batched": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipc_rx_req": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipc_rx_res": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipc_tx_req": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipc_tx_res": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vs_hb": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServerConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"def_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addr": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"last_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPER_UNAVAIL",
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"propogate_state": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"timer_exists": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceAlertMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_objid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cfg_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVIMgrInterestedEntitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interested_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceClientInsightsSamplingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"sample_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"skip_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
		},
	}
}

func ResourceControllerPortalAuthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"instance_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jwt_token": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceIpamDnsInfobloxProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"extensible_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"ip_address": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"network_view": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"usable_alloc_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceInfobloxSubnetSchema(),
			},
			"usable_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wapi_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2.0",
			},
		},
	}
}

func ResourceVirtualServiceResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_standby_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCC_VnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"status_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gslb_geo_db_profile_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslb_service_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslb_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVirtualServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_extension": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsApicExtensionSchema(),
			},
			"controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datapath_debug": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugVirtualServiceSchema(),
			},
			"east_west": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gslb_dns_update": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbDnsUpdateSchema(),
			},
			"ipam_dns_records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"is_dns_vs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key_rotation_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"last_key_rotation_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"lif": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"manual_placement": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_additional_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"one_plus_one_ha": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"prev_controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redis_db": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redis_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sec_mgr_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecurityMgrRuntimeSchema(),
			},
			"self_se_election": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tls_ticket_key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTLSTicketSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL",
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vh_child_vs_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_runtime": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipRuntimeSchema(),
			},
			"vs_update_pending": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceSchema(),
			},
		},
	}
}

func ResourceSSLProfileSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceTLSTicketSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aes_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hmac_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceIPNetworkSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTP2ApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http2_initial_window_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"max_http2_concurrent_streams_per_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"max_http2_control_frames_per_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"max_http2_empty_data_frames_per_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"max_http2_header_field_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"max_http2_queued_frames_to_client_per_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"max_http2_requests_per_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
		},
	}
}

func ResourceDebugIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceClientLogFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
		},
	}
}

func ResourcePoolDeploymentUpdateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"evaluation_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"results": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleResultSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafDataFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloudSyncServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCookieMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafPSMLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"group_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleMatchDataSchema(),
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecureChannelAvailableLocalIPsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"free_controller_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"free_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAWSASGDeleteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asgs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_vses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteDnsVsSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hm_proxies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"hm_shard_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"member_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_PASSIVE_MEMBER",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugDnsOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gslb_service_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAppCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encryption_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prst_hdr_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
		},
	}
}

func ResourceAuthorizationActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_RESPONSE_STATUS_CODE_403",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALLOW_ACCESS",
			},
		},
	}
}

func ResourceSeGroupOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_on_error": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SUSPEND_UPGRADE_OPS_ON_ERROR",
			},
			"disruptive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceStreamingSyslogConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"facility": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16,
			},
			"filtered_log_severity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AviVantage",
			},
			"non_significant_log_severity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"significant_log_severity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
		},
	}
}

func ResourceSipServiceApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"transaction_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
		},
	}
}

func ResourceHttpCookiePersistenceKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aes_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPSecurityPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPSecurityRuleSchema(),
			},
		},
	}
}

func ResourceVcenterClustersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceClientLogConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_significant_log_collection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"filtered_log_processing": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
			"non_significant_log_processing": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
			"significant_log_processing": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
		},
	}
}

func ResourceHTTPLocalFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_content": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceConfigUserNotAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policies": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Policies",
			},
			"roles": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Roles",
			},
			"tenants": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Tenants",
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSSLKeyRSAParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exponent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65537,
			},
			"key_size": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_2048_BITS",
			},
		},
	}
}

func ResourceSeVipInterfaceListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_portchannel": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_intf_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_intf_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_intf_mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceRouteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HOST_NAMESPACE",
			},
			"nexthop": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"subnet": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceDnsRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDisableSeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMetricThresoldUpDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"entity_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_networks_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmAddNetworksEventDetailsSchema(),
			},
			"all_seupgrade_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAllSeUpgradeEventDetailsSchema(),
			},
			"anomaly_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAnomalyEventDetailsSchema(),
			},
			"apic_agent_bd_vrf_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceApicAgentBridgeDomainVrfChangeSchema(),
			},
			"apic_agent_generic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceApicAgentGenericEventDetailsSchema(),
			},
			"apic_agent_vs_network_error": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceApicAgentVsNetworkErrorSchema(),
			},
			"avg_uptime_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAvgUptimeChangeDetailsSchema(),
			},
			"avi_cloud_crs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAviCloudCRSDetailsSchema(),
			},
			"avi_cloud_status_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAviCloudStatusDetailsSchema(),
			},
			"aws_asg_deletion_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAWSASGDeleteSchema(),
			},
			"aws_asg_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAWSASGNotifDetailsSchema(),
			},
			"aws_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAWSSetupSchema(),
			},
			"azure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAzureSetupSchema(),
			},
			"azure_mp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAzureMarketplaceSchema(),
			},
			"bind_vs_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmBindVsSeEventDetailsSchema(),
			},
			"bm_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBMSetupSchema(),
			},
			"bootup_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmSeBootupFailEventDetailsSchema(),
			},
			"burst_checkout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBurstLicenseDetailsSchema(),
			},
			"cc_cluster_vip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudClusterVipSchema(),
			},
			"cc_dns_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudDnsUpdateSchema(),
			},
			"cc_health_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudHealthSchema(),
			},
			"cc_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudGenericSchema(),
			},
			"cc_ip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudIpChangeSchema(),
			},
			"cc_parkintf_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudVipParkingIntfSchema(),
			},
			"cc_scaleset_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCCScaleSetNotifDetailsSchema(),
			},
			"cc_se_vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudSeVmChangeSchema(),
			},
			"cc_sync_services_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudSyncServicesSchema(),
			},
			"cc_tenant_del_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudTenantsDeletedSchema(),
			},
			"cc_vip_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudVipUpdateSchema(),
			},
			"cc_vnic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudVnicChangeSchema(),
			},
			"cloud_asg_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudASGNotifDetailsSchema(),
			},
			"cloud_autoscaling_config_failure_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudAutoscalingConfigFailureDetailsSchema(),
			},
			"cluster_config_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterConfigFailedEventSchema(),
			},
			"cluster_leader_failover_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterLeaderFailoverEventSchema(),
			},
			"cluster_node_add_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeAddEventSchema(),
			},
			"cluster_node_db_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeDbFailedEventSchema(),
			},
			"cluster_node_remove_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeRemoveEventSchema(),
			},
			"cluster_node_shutdown_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeShutdownEventSchema(),
			},
			"cluster_node_started_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeStartedEventSchema(),
			},
			"cluster_service_critical_failure_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterServiceCriticalFailureEventSchema(),
			},
			"cluster_service_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterServiceFailedEventSchema(),
			},
			"cluster_service_restored_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterServiceRestoredEventSchema(),
			},
			"cluster_warm_reboot_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterWarmRebootEventSchema(),
			},
			"cntlr_host_list_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraCntlrHostUnreachableListSchema(),
			},
			"config_action_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigActionDetailsSchema(),
			},
			"config_create_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigCreateDetailsSchema(),
			},
			"config_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigDeleteDetailsSchema(),
			},
			"config_password_change_request_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUserPasswordChangeRequestSchema(),
			},
			"config_se_grp_flv_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigSeGrpFlvUpdateSchema(),
			},
			"config_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUpdateDetailsSchema(),
			},
			"config_user_authrz_rule_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUserAuthrzByRuleSchema(),
			},
			"config_user_login_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUserLoginSchema(),
			},
			"config_user_logout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUserLogoutSchema(),
			},
			"config_user_not_authrz_rule_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigUserNotAuthrzByRuleSchema(),
			},
			"container_cloud_batch_setup": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceContainerCloudBatchSetupSchema(),
			},
			"container_cloud_setup": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceContainerCloudSetupSchema(),
			},
			"container_cloud_sevice": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceContainerCloudServiceSchema(),
			},
			"cs_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudStackSetupSchema(),
			},
			"delete_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmDeleteSeEventDetailsSchema(),
			},
			"disable_se_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDisableSeMigrateEventDetailsSchema(),
			},
			"disc_summary": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraDiscSummaryDetailsSchema(),
			},
			"dns_sync_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDNSVsSyncInfoSchema(),
			},
			"docker_ucp_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDockerUCPSetupSchema(),
			},
			"dos_attack_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDosAttackEventDetailsSchema(),
			},
			"gcp_cloud_router_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPCloudRouterUpdateSchema(),
			},
			"gcp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPSetupSchema(),
			},
			"glb_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbStatusSchema(),
			},
			"gs_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbServiceStatusSchema(),
			},
			"host_unavail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostUnavailEventDetailsSchema(),
			},
			"hs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHealthScoreDetailsSchema(),
			},
			"ip_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmSeIpFailEventDetailsSchema(),
			},
			"license_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLicenseDetailsSchema(),
			},
			"license_expiry_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLicenseExpiryDetailsSchema(),
			},
			"marathon_service_port_conflict_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMarathonServicePortConflictSchema(),
			},
			"memory_balancer_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMemoryBalancerInfoSchema(),
			},
			"mesos_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMesosSetupSchema(),
			},
			"metric_threshold_up_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricThresoldUpDetailsSchema(),
			},
			"metrics_db_disk_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbDiskEventDetailsSchema(),
			},
			"metrics_db_queue_full_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbQueueFullEventDetailsSchema(),
			},
			"metrics_db_queue_healthy_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbQueueHealthyEventDetailsSchema(),
			},
			"mgmt_nw_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraMgmtNwChangeDetailsSchema(),
			},
			"modify_networks_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmModifyNetworksEventDetailsSchema(),
			},
			"network_subnet_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNetworkSubnetInfoSchema(),
			},
			"nw_subnet_clash_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNetworkSubnetClashSchema(),
			},
			"nw_summarized_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSummarizedInfoSchema(),
			},
			"oci_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOCISetupSchema(),
			},
			"os_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackClusterSetupSchema(),
			},
			"os_ip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackIpChangeSchema(),
			},
			"os_lbaudit_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackLbProvAuditCheckSchema(),
			},
			"os_lbplugin_op_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackLbPluginOpSchema(),
			},
			"os_se_vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackSeVmChangeSchema(),
			},
			"os_sync_services_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackSyncServicesSchema(),
			},
			"os_vnic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackVnicChangeSchema(),
			},
			"pool_deployment_failure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePoolDeploymentFailureInfoSchema(),
			},
			"pool_deployment_success_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePoolDeploymentSuccessInfoSchema(),
			},
			"pool_deployment_update_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePoolDeploymentUpdateInfoSchema(),
			},
			"pool_server_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraPoolServerDeleteDetailsSchema(),
			},
			"psm_program_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePsmProgramDetailsSchema(),
			},
			"rate_limiter_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateLimiterEventDetailsSchema(),
			},
			"rebalance_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRebalanceMigrateEventDetailsSchema(),
			},
			"rebalance_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRebalanceScaleinEventDetailsSchema(),
			},
			"rebalance_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRebalanceScaleoutEventDetailsSchema(),
			},
			"reboot_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmRebootSeEventDetailsSchema(),
			},
			"scheduler_action_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSchedulerActionDetailsSchema(),
			},
			"se_bgp_peer_state_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeBgpPeerStateChangeDetailsSchema(),
			},
			"se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeMgrEventDetailsSchema(),
			},
			"se_dupip_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeDupipEventDetailsSchema(),
			},
			"se_gateway_heartbeat_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGatewayHeartbeatFailedDetailsSchema(),
			},
			"se_gateway_heartbeat_success_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGatewayHeartbeatSuccessDetailsSchema(),
			},
			"se_geo_db_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGeoDbDetailsSchema(),
			},
			"se_hb_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHBEventDetailsSchema(),
			},
			"se_hm_gs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventGSDetailsSchema(),
			},
			"se_hm_gsgroup_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventGslbPoolDetailsSchema(),
			},
			"se_hm_pool_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventPoolDetailsSchema(),
			},
			"se_hm_vs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHmEventVsDetailsSchema(),
			},
			"se_ip6_dad_failed_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeIP6DadFailedEventDetailsSchema(),
			},
			"se_ip_added_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeIpAddedEventDetailsSchema(),
			},
			"se_ip_removed_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeIpRemovedEventDetailsSchema(),
			},
			"se_ipfailure_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeIpfailureEventDetailsSchema(),
			},
			"se_licensed_bandwdith_exceeded_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeLicensedBandwdithExceededEventDetailsSchema(),
			},
			"se_memory_limit_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeMemoryLimitEventDetailsSchema(),
			},
			"se_persistence_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSePersistenceEventDetailsSchema(),
			},
			"se_pool_lb_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSePoolLbEventDetailsSchema(),
			},
			"se_thresh_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeThreshEventDetailsSchema(),
			},
			"se_version_check_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVersionCheckFailedEventSchema(),
			},
			"se_vnic_down_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVnicDownEventDetailsSchema(),
			},
			"se_vnic_tx_queue_stall_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVnicTxQueueStallEventDetailsSchema(),
			},
			"se_vnic_up_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVnicUpEventDetailsSchema(),
			},
			"se_vs_fault_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVsFaultEventDetailsSchema(),
			},
			"se_vs_pkt_buf_high_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVsPktBufHighEventDetailsSchema(),
			},
			"semigrate_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeMigrateEventDetailsSchema(),
			},
			"server_autoscale_failed_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServerAutoScaleFailedInfoSchema(),
			},
			"server_autoscalein_complete_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServerAutoScaleInCompleteInfoSchema(),
			},
			"server_autoscalein_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServerAutoScaleInInfoSchema(),
			},
			"server_autoscaleout_complete_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServerAutoScaleOutCompleteInfoSchema(),
			},
			"server_autoscaleout_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceServerAutoScaleOutInfoSchema(),
			},
			"seupgrade_disrupted_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeVsDisruptedEventDetailsSchema(),
			},
			"seupgrade_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeEventDetailsSchema(),
			},
			"seupgrade_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeMigrateEventDetailsSchema(),
			},
			"seupgrade_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeScaleinEventDetailsSchema(),
			},
			"seupgrade_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeUpgradeScaleoutEventDetailsSchema(),
			},
			"spawn_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmSpawnSeEventDetailsSchema(),
			},
			"ssl_expire_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLExpireDetailsSchema(),
			},
			"ssl_export_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLExportDetailsSchema(),
			},
			"ssl_renew_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLRenewDetailsSchema(),
			},
			"ssl_renew_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLRenewFailedDetailsSchema(),
			},
			"switchover_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSwitchoverEventDetailsSchema(),
			},
			"switchover_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSwitchoverFailEventDetailsSchema(),
			},
			"sync_services_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudSyncServicesSchema(),
			},
			"system_upgrade_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSystemUpgradeDetailsSchema(),
			},
			"tencent_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTencentSetupSchema(),
			},
			"unbind_vs_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmUnbindVsSeEventDetailsSchema(),
			},
			"upgrade_entry": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeOpsEntrySchema(),
			},
			"upgrade_status_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeStatusInfoSchema(),
			},
			"vca_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVCASetupSchema(),
			},
			"vcenter_connectivity_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVcenterConnectivityStatusSchema(),
			},
			"vcenter_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVcenterBadCredentialsSchema(),
			},
			"vcenter_disc_failure": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVcenterDiscoveryFailureSchema(),
			},
			"vcenter_network_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVcenterNetworkLimitSchema(),
			},
			"vcenter_obj_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVcenterObjDeleteDetailsSchema(),
			},
			"vip_autoscale": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVipScaleDetailsSchema(),
			},
			"vip_dns_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDNSRegisterInfoSchema(),
			},
			"vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVinfraVmDetailsSchema(),
			},
			"vs_awaitingse_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsAwaitingSeEventDetailsSchema(),
			},
			"vs_error_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsErrorEventDetailsSchema(),
			},
			"vs_fsm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsFsmEventDetailsSchema(),
			},
			"vs_initialplacement_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsInitialPlacementEventDetailsSchema(),
			},
			"vs_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateEventDetailsSchema(),
			},
			"vs_pool_nw_fltr_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsPoolNwFilterEventDetailsSchema(),
			},
			"vs_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleInEventDetailsSchema(),
			},
			"vs_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleOutEventDetailsSchema(),
			},
		},
	}
}

func ResourceMetricsApiSrvDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePathMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIpamDnsOpenstackProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"keystone_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmBindVsSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"standby": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVsAwaitingSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"awaitingse_timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafExcludeListEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_element": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_element_criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceWafExclusionTypeSchema(),
			},
			"uri_match_criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceWafExclusionTypeSchema(),
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSwitchoverFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOshiftDockerRegistryMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"registry_service": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "docker-registry",
			},
			"registry_vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceDnsRuleActionGslbSiteSelectionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fallback_site_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"is_site_preferred": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVIPGNameInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterActionSchema(),
			},
			"explicit_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fine_grain": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_header": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_limiter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateLimiterSchema(),
			},
		},
	}
}

func ResourceCPUUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAzureInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceProxyConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudVnicChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_VnicInfoSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeHmEventVsDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip6_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAttackMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_resp_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSwitchoverEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsMissingDataIntervalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMethodMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"methods": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSupportedMigrationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_host_min_free_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"controller_min_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"controller_min_free_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"controller_min_memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  24,
			},
			"controller_min_total_disk": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"max_active_versions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"rollback_controller_disk_space": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"rollback_se_disk_space": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_host_min_free_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_min_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"se_min_free_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_min_memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"se_min_total_disk": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"versions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourcePortRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsServiceApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aaaa_empty_response": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"admin_email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "hostmaster",
			},
			"dns_over_tcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dns_zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsZoneSchema(),
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ecs_stripping_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"edns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"edns_client_subnet_prefix_len": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"error_response": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_ERROR_RESPONSE_NONE",
			},
			"name_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"negative_caching_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"num_dns_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
		},
	}
}

func ResourceSeGroupResumeOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_on_error": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SUSPEND_UPGRADE_OPS_ON_ERROR",
			},
			"disruptive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"skip_suspended": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceMemoryUsageInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mem_usage_on_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMemoryUsagePerNodeSchema(),
			},
		},
	}
}

func ResourceVlanRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_insights": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NO_INSIGHTS",
			},
			"client_insights_sampling": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClientInsightsSamplingSchema(),
			},
			"client_log_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClientLogFilterSchema(),
			},
			"full_client_logs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFullClientLogsSchema(),
			},
			"metrics_realtime_update": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsRealTimeUpdateSchema(),
			},
			"significant_log_throttle": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"udf_log_throttle": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceGslbDnsGeoUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceVipRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ev": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ev_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsEvStatusSchema(),
			},
			"first_se_assigned_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"first_time_placement": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fsm_state_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"fsm_state_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VipFsmMap::Inactive",
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"last_scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceScaleStatusSchema(),
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"migrate_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"migrate_scalein_pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_scaleout_pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_additional_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"prev_metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"progress_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"requested_resource": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceScaleStatusSchema(),
			},
			"scalein_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scalein_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleinParamsSchema(),
			},
			"scaleout_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeListSchema(),
			},
			"supp_runtime_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"user_scaleout_pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warmstart_resync_done": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"warmstart_resync_sent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocols": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAzureClusterInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_credential_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMemoryUsagePerNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mem_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMemoryUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceUpgradeSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceServiceMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchSchema(),
			},
			"source_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchSchema(),
			},
		},
	}
}

func ResourceAuthTacacsPlusAttributeValuePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mandatory": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsSrvRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default.host",
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceGslbStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gslb_runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbRuntimeSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"site": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
			},
			"third_party_site": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbThirdPartySiteRuntimeSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"answer_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authentic_data": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"checking_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"client_location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"identifier": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nameserver_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"opcode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opt_record": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsOptRecordSchema(),
			},
			"query_or_response": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"question_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"recursion_desired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeIpRemovedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ns": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceL4RuleActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"select_pool": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RuleActionSelectPoolSchema(),
			},
		},
	}
}

func ResourceDosRateLimitProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dos_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"rl_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRateLimiterProfileSchema(),
			},
		},
	}
}

func ResourceOpenStackSyncServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMemberInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"if_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_ALGORITHM_ROUND_ROBIN",
			},
			"consistent_hash_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"consistent_hash_mask6": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fallback_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"members": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceGslbPoolMemberSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceDnsQueryTypeMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceClusterNodeRemoveEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsNsRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip6_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"nsname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceScaleStatusSchema(),
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMicroServiceContainerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"task_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCRLSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"body": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"common_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distinguished_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_refreshed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"text": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbPoolMemberDatapathStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"site_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCaptureFiltersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"capture_ipc": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCaptureIPCSchema(),
			},
			"dst_port_end": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst_port_start": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eth_proto": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_proto": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_ack": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tcp_fin": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tcp_push": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tcp_syn": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeBootupPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_backend_portend": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30720,
			},
			"docker_backend_portstart": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20480,
			},
			"fair_queueing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"geo_db_granularity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"l7_conns_per_core": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"l7_resvd_listen_conns_per_core": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
			"log_agent_debug_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_agent_trace_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_dp_compression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeBootupCompressionPropertiesSchema(),
			},
			"se_emulated_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_ip_encap_ipc": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_l3_encap_ipc": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_log_buffer_app_blocking_dequeue": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_applog_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"se_log_buffer_chunk_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_conn_blocking_dequeue": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_connlog_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_events_blocking_dequeue": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_log_buffer_events_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  512,
			},
			"ssl_sess_cache_per_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"ssl_sess_cache_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"tcp_syncache_hashsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8192,
			},
		},
	}
}

func ResourceSeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVipScaleDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vsvip_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arg": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fault_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"function_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_executions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"num_skips": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceEmailConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"auth_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_tls": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"from_email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin@avicontroller.net",
			},
			"mail_server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "localhost",
			},
			"mail_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"smtp_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePortMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ports": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceGatewayMonitorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"gateway_monitor_fail_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"gateway_monitor_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"gateway_monitor_success_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceGslbSiteRuntimeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_file_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_file_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gap_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_geo_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_ghm_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_glb_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gpki_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_gs_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_health_msgs_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_health_msgs_txed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_of_bad_responses": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_of_events_generated": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_of_skip_outstanding_requests": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_of_timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAWSASGNotifDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asg_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_ip_addr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHostHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRmAddVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPVersionMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"versions": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVIGuestvNicIPAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mask": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsRuleActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionAllowDropSchema(),
			},
			"dns_rate_limiter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRateLimiterSchema(),
			},
			"gslb_site_selection": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionGslbSiteSelectionSchema(),
			},
			"pool_switching": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionPoolSwitchingSchema(),
			},
			"response": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionResponseSchema(),
			},
		},
	}
}

func ResourceMgmtIpAccessControlSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"shell_server_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"snmp_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"ssh_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"sysint_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
		},
	}
}

func ResourceDebugSeAgentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"log_every_n": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"log_level": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_module": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trace_level": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbHealthMonitorProxySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proxy_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_HEALTH_MONITOR_PROXY_PRIVATE_MEMBERS",
			},
			"site_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsGeoLocationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"geolocation_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geolocation_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_edns_client_subnet_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_group_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovered_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_orchestration_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prst_hdr_val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"resolve_server_by_dns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_node": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"verify_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vm_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbGeoDbFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_GEODB_FILE_FORMAT_AVI",
			},
		},
	}
}

func ResourceChildProcessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNetworkProfileUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tcp_fast_path_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTCPFastPathProfileSchema(),
			},
			"tcp_proxy_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTCPProxyProfileSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"udp_fast_path_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUDPFastPathProfileSchema(),
			},
			"udp_proxy_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUDPProxyProfileSchema(),
			},
		},
	}
}

func ResourceGslbRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dns_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"event_cache": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceEventCacheSchema(),
			},
			"flr_state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"ldr_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"site": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"third_party_sites": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbThirdPartySiteRuntimeSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVssPlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_nonaffinity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"num_subcores": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
		},
	}
}

func ResourceRollbackPatchSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceAviCloudCRSDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"release_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsARdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceStringMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceClusterNodeShutdownEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbClientIpAddrGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_IP_PUBLIC",
			},
		},
	}
}

func ResourceEventDetailsFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"event_details_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"event_details_value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAuthenticationRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthenticationActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthenticationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDebugVirtualServiceCaptureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture_file_size": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCaptureFileSizeSchema(),
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"enable_ssl_session_key_capture": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pcap_ng": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"pkt_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
		},
	}
}

func ResourceOpenStackRoleMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_role": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_role": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVlanInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"if_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_mgmt": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vnic_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema(),
			},
			"vrf_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcevCloudAirConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"privilege": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_mgmt_network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_orgnization": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"vca_username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_vdc": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHdrPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prst_hdr_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudVipParkingIntfSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intf_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTCPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aggressive_congestion_avoidance": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"automatic": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"cc_algo": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CC_ALGO_NEW_RENO",
			},
			"congestion_recovery_scaling_factor": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"idle_connection_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"idle_connection_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KEEP_ALIVE",
			},
			"ignore_time_wait": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_dscp": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"keepalive_in_halfclose_state": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"max_retransmissions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"max_segment_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_syn_retransmissions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"min_rexmt_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nagles_algorithm": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"reassembly_queue_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"receive_window": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"reorder_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"slow_start_scaling_factor": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"time_wait_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"use_interface_mtu": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceServerAutoScaleInInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alertconfig_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alertconfig_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"available_capacity": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"load": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"num_scalein_servers": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"scalein_server_candidates": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceLdapUserBindSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dn_template": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "<user>",
			},
			"user_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceTenantConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_in_provider_context": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_access_to_provider_se": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_vrf": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceLinuxServerHostSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_attr": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHostAttributesSchema(),
			},
			"host_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDiskUsageInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disk_usage_on_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiskUsagePerNodeSchema(),
			},
		},
	}
}

func ResourceBuildInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"build_no": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbServiceStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gs_runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbServiceRuntimeSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVipSeAssignedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down_requested": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"connected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"primary": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scalein_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"standby": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSipMessageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"from_client": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rcv_timestamp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rx_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tx_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeHBEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hb_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_ref1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPStatusMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPStatusRangeSchema(),
			},
			"status_codes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceSeUpgradeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMetricsMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_hw_training": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_grace_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_first_n": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logging_freq": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_cluster_map_check": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_metrics_db_writes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPHdrValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"var": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCfgStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cfg_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cfg_version_in_flight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"site_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeIpfailureEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSnmpConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"large_trap_payload": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snmp_v3_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSnmpV3ConfigurationSchema(),
			},
			"sys_contact": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "support@avinetworks.com",
			},
			"sys_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}

func ResourceIpAddrRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"begin": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"end": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceAutoScaleMesosSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceGslbServiceSitePersistencePoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_servers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_servers_up": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerConfigSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePatchSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHTTPClientAuthenticationParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_uri_path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeVnicDownEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsServiceDomainSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_dns_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"pass_through": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"record_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePaaRequestLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"headers_received_from_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers_sent_to_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servers_tried": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceConfigUpdateDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_resource_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"old_resource_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDiskUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cntlr_disk_free": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOverallInfoSchema(),
			},
			"cntlr_disk_usage": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePartitionInfoSchema(),
			},
			"se_disk_free": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOverallInfoSchema(),
			},
			"se_disk_usage": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePartitionInfoSchema(),
			},
		},
	}
}

func ResourceAvgUptimeChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRebalanceScaleinEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scalein_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleinParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsRuleDnsRrSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_record_set": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceDnsRrSetSchema(),
			},
			"section": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_MESSAGE_SECTION_ANSWER",
			},
		},
	}
}

func ResourceDnsZoneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNTPConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ntp_authentication_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPAuthenticationKeySchema(),
			},
			"ntp_server_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPServerSchema(),
			},
		},
	}
}

func ResourceWafLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latency_request_body_phase": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latency_request_header_phase": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latency_response_body_phase": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latency_response_header_phase": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"psm_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"psm_executed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"psm_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMLogSchema(),
			},
			"rule_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleLogSchema(),
			},
			"rules_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rules_executed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelist_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"whitelist_executed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"whitelist_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafWhitelistLogSchema(),
			},
		},
	}
}

func ResourceIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_RESPONSE_CONSISTENT_HASH",
			},
			"cname": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_records_in_response": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_A",
			},
		},
	}
}

func ResourceGCPInBandManagementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceFeProxyRoutePublishConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FE_PROXY_ROUTE_PUBLISH_NONE",
			},
			"publisher_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"subnet": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbDnsGsStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"num_partial_updates": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"partial_update_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNetworkFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_filter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpamDnsCustomProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"custom_ipam_dns_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"usable_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceSystemUpgradeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"upgrade_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSystemUpgradeStateSchema(),
			},
		},
	}
}

func ResourceSSLClientRequestHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"request_header": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_header_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSipLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_callid_hdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_contact_hdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_from_hdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSipMessageSchema(),
			},
			"sip_to_hdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsDataHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"derivation_data": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDerivationDataSchema(),
			},
			"dimension_data": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDimensionDataSchema(),
			},
			"entity_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics_min_scale": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"metrics_sum_agg_invalid": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"missing_intervals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsMissingDataIntervalSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_id_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serviceengine_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"statistics": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricStatisticsSchema(),
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRIC_COUNT",
			},
		},
	}
}

func ResourceSeUpgradeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRmSeIpFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmAddVnicSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcevNICNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ctlr_alloc": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"mode": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceUdpAttacksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceSePoolLbEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeBgpPeerStateChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"peer_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"peer_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceContainerCloudSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"master_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"missing_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMesosSeResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "2.0",
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
		},
	}
}

func ResourceVsPoolNwFilterEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHealthMonitorUdpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_request": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAutoScaleOpenStackSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"heat_scale_down_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heat_scale_up_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeRateLimitersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arp_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"default_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"flow_probe_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  250,
			},
			"icmp_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"icmp_rsp_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  500,
			},
			"rst_rl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func ResourceAPICLifsRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocated": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cifs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCifSchema(),
			},
			"contract_graphs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lif_label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"multi_vrf": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transaction_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsResourceRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr6_ip_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addr_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dclass": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mail_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"text_rdata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOpenStackClusterSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keystone_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMesosSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mesos_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"mesos_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"missing_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceFlowtableProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tcp_closed_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"tcp_connection_setup_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"tcp_half_closed_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"tcp_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"tcp_reset_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"udp_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceRateLimiterPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"msf_num_stages": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"msf_stage_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
		},
	}
}

func ResourceAlertSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_script_output": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert_config_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"app_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceApplicationLogSchema(),
			},
			"conn_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionLogSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_pages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventLogSchema(),
			},
			"last_throttle_timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"level": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricLogSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Required: true,
			},
			"related_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"throttle_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"timestamp": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecureChannelMetadataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpAddrPrefixSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mask": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDiscoveredNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceConnectionLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"average_turntime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client_dest_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_log_filter_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_rtt": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_src_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"connection_ended": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"dns_etype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"dns_qtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRequestSchema(),
			},
			"dns_response": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsResponseSchema(),
			},
			"ds_log": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbpool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbservice_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"microservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"microservice_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mss": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_syn_retransmit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_transaction": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"out_of_orders": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_timestamp": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"retransmits": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"rx_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"rx_pkts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_conn_src_ip": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_conn_src_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_dest_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_ip": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_out_of_orders": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_retransmits": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rtt": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rx_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rx_pkts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_src_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_timeouts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_total_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_total_pkts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_tx_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_tx_pkts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_zero_window_size_events": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_engine": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"significance": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"significant": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"significant_log": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sip_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSipLogSchema(),
			},
			"sni_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cipher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_timestamp": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"timeouts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"total_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"tx_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tx_pkts": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"udf": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"vcpu_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtualservice": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vs_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zero_window_size_events": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceCloudDnsUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVSDataScriptSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"evt": {
				Type:     schema.TypeString,
				Required: true,
			},
			"script": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNTPServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceAuthAttributeMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute_value_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
		},
	}
}

func ResourceFailActionBackupPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"answer_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authoritative_answer": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"fallback_algorithm_used": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_wildcard": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nameserver_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"opcode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opt_record": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsOptRecordSchema(),
			},
			"query_or_response": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"question_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsResourceRecordSchema(),
			},
			"recursion_available": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"recursion_desired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"response_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"truncation": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVcenterDatastoreSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datastore_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbPoolMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"public_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbIpAddrSchema(),
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"resolve_fqdn_to_v6": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVserverL4MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexc": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"apdexrtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_application_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_bytes_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_complete_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_connections_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_app_error": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_bad_rst_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_conn": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_conn_ip_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_conn_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_fake_session": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_http_abort": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_http_error": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_http_timeout": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_malformed_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_non_syn_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_cookie_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_custom_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_hdr_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_scan_bad_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_scan_unknown_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_uri_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_ip_uri_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_uri_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_uri_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_uri_scan_bad_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_req_uri_scan_unknown_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_slow_uri": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_small_window_stress": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_ssl_error": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_syn_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_total_req": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_tx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dos_zero_window_stress": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_errored_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_l4_client_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_lossy_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_lossy_req": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_network_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_new_established_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_pkts_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_rx_bytes_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_rx_pkts": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_rx_pkts_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_syns": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_total_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_total_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_tx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_tx_pkts": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_num_active_se": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_open_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_rx_bytes_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_rx_pkts_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_tx_bytes_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_tx_pkts_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_application_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_connection_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_connections_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_dos_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_dos_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_network_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_pkts_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_conn_duration": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_connection_dropped_user_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_connection_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_connections_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_dup_ack_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_end_to_end_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_end_to_end_rtt_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_end_to_end_rtt_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_finished_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_lossy_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_lossy_req": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_out_of_orders": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_packet_dropped_user_bandwidth_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_rtt_valid_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_sack_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_server_flow_control": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_timeout_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_zero_window_size_events": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCPUUsageInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_usage_on_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCPUUsagePerNodeSchema(),
			},
		},
	}
}

func ResourceSeUpgradeVsDisruptedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeUpgradeScaleoutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scaleout_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleoutParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOpenStackLbProvAuditCheckSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elapsed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMetricsQueryResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"series": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDataSeriesSchema(),
			},
			"start": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDNSConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"search_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceAviCloudStatusDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connectivity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"registration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAlertRuleMetricSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_threshold": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceAlertMetricThresholdSchema(),
			},
		},
	}
}

func ResourceClusterNodeAddEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafLearningSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceSeDupipEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPCookieDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafPositiveSecurityModelSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceServicePoolSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_pool_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_port_range_end": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"override_application_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_network_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"port_range_end": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceServerAutoScaleOutCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_config_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nscaleout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scaled_out_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceUpgradeControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeGatewayHeartbeatFailedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePoolDeploymentSuccessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prev_in_service_pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_in_service_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"results": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleResultSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCompressionProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compressible_content_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compression": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCompressionFilterSchema(),
			},
			"remove_accept_encoding_header": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceL4RuleActionSelectPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDNSRegisterInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"total_records": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHostAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attr_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attr_val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafPSMRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_elements": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMMatchElementSchema(),
			},
			"match_value_max_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match_value_pattern": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"paranoia_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_PARANOIA_LEVEL_LOW",
			},
			"rule_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraPoolServerDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVipPlacementNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceAuthenticationPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authn_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthenticationRuleSchema(),
			},
			"default_auth_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVIMgrGuestNicRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_internal_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"connected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"del_pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"guest_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema(),
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Unknown",
			},
			"mac_addr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_vnic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCC_AgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"async_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"async_retries_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"poll_duration_target": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"poll_fast_target": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"poll_slow_target": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  240,
			},
			"vnic_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"vnic_retries_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}

func ResourceCloudGenericSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRebootDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reboot": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPRequestPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPRequestRuleSchema(),
			},
		},
	}
}

func ResourceIptableRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dnat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"dst_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"dst_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"input_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proto": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"src_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAzureSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alb_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceCC_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rpc_poll_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"rpc_queue_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func ResourceQueryMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHTTPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_dots_in_header_name": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cache_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHttpCacheConfigSchema(),
			},
			"client_body_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"client_header_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"client_max_body_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"client_max_header_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"client_max_request_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  48,
			},
			"compression_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCompressionProfileSchema(),
			},
			"connection_multiplexing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_keepalive_posts_msie6": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_sni_hostname_check": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_chunk_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_fire_and_forget": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_request_body_buffering": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_request_body_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fwd_close_hdr_for_bound_connections": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hsts_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hsts_max_age": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"hsts_subdomains_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"http2_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http2_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTP2ApplicationProfileSchema(),
			},
			"http_to_https": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"httponly_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"max_bad_rps_cip": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_bad_rps_cip_uri": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_bad_rps_uri": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_keepalive_requests": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"max_response_headers_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  48,
			},
			"max_rps_cip": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_cip_uri": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_unknown_cip": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_unknown_uri": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_uri": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_accept_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"reset_conn_http_on_ssl_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"respond_with_100_continue": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"secure_cookie_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_side_redirect_to_https": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_client_certificate_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLClientCertificateActionSchema(),
			},
			"ssl_client_certificate_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CLIENT_CERTIFICATE_NONE",
			},
			"use_app_keepalive_timeout": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"websockets_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"x_forwarded_proto_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"xff_alternate_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "X-Forwarded-For",
			},
			"xff_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceVinfraCntlrHostUnreachableListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsFsmEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_rt": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceRuntimeSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOCICredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fingerprint": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"key_content": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"pass_phrase": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServerAutoScaleOutInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alertconfig_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alertconfig_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"available_capacity": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"load": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"num_scaleout_servers": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
		},
	}
}

func ResourceTacacsPlusAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_attrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthTacacsPlusAttributeValuePairSchema(),
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  49,
			},
			"server": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_TACACS_PLUS_SERVICE_LOGIN",
			},
		},
	}
}

func ResourceTimeStampSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"usecs": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceEventCacheSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_state": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"exceptions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceWafRuleLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleMatchDataSchema(),
			},
			"msg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsClientIpMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"use_edns_client_subnet_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHSMThalesNetHsmSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"esn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keyhash": {
				Type:     schema.TypeString,
				Required: true,
			},
			"module_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"remote_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"remote_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}

func ResourceDNSVsSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_records": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClusterConfigFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceClusterLeaderFailoverEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_node": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeSchema(),
			},
			"previous_leader_node": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClusterNodeSchema(),
			},
		},
	}
}

func ResourceGCPConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_credentials_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_keys": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPEncryptionKeysSchema(),
			},
			"firewall_target_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gcs_bucket_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcs_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_config": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceGCPNetworkConfigSchema(),
			},
			"region_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vip_allocation_strategy": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceGCPVIPAllocationSchema(),
			},
			"zones": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceUserRoleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_tenants": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"object_access_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsDerivationDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"derivation_fn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"exclude_derived_metric": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"include_derivation_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"join_tables": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_ids": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result_has_additional_fields": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"second_order_derivation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"skip_backend_derivation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceGCPTwoArmModeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_data_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frontend_data_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frontend_data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frontend_data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsRuleActionPoolSwitchingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGCPOneArmModeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDockerUCPSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"docker_ucp_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"missing_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ucp_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceConnPoolPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"upstream_connpool_conn_idle_tmo": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60000,
			},
			"upstream_connpool_conn_life_tmo": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600000,
			},
			"upstream_connpool_conn_max_reuse": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"upstream_connpool_server_max_cache": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceBurstLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbDownloadStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_DOWNLOAD_NONE",
			},
		},
	}
}

func ResourceRmModifyNetworksEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmModifyVnicSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRollbackPatchSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSSLCertificateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate_signing_request": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"chain_verified": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"days_until_expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"expiry_status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_GOOD",
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLCertificateDescriptionSchema(),
			},
			"key_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLKeyParamsSchema(),
			},
			"not_after": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"not_before": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"self_signed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLCertificateDescriptionSchema(),
			},
			"subject_alt_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"text": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSubJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expires_at": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeRuntimePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_ssh_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"app_headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppHdrSchema(),
			},
			"baremetal_dispatcher_handles_flows": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"connections_lossy_log_rate_limiter_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"connections_udfnf_log_rate_limiter_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"disable_flow_probes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dos_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"downstream_send_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"dp_aggressive_hb_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_aggressive_hb_timeout_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dp_hb_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_hb_timeout_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dupip_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"dupip_timeout_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"enable_hsm_log": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"flow_table_batch_push_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"global_mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"http_rum_console_log": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_rum_min_content_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"lbaction_num_requests_to_dispatch": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"lbaction_rq_per_request_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  22,
			},
			"log_agent_compress_logs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"log_agent_conn_send_buffer_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"log_agent_export_msg_buffer_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  524288,
			},
			"log_agent_export_wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_file_sz_appl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_conn": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_debug": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_event": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_log_storage_min_sz": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_active_adf_files_per_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_max_concurrent_rsync": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_logmessage_proto_sz": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65536,
			},
			"log_agent_max_storage_excess_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  110,
			},
			"log_agent_max_storage_ignore_percent": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "20.0",
			},
			"log_agent_min_storage_per_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_pause_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"log_agent_sleep_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_unknown_vs_timer": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1800,
			},
			"log_message_max_file_list_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"mcache_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"mcache_fetch_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"mcache_store_in_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"mcache_store_in_max_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mcache_store_in_min_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mcache_store_out_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ngx_free_connection_stack": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"persistence_mem_max": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"scaleout_udp_per_pkt": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_auth_ldap_bind_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5000,
			},
			"se_auth_ldap_cache_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100000,
			},
			"se_auth_ldap_connect_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_conns_per_server": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_auth_ldap_reconnect_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_request_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_servers_failover_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_dp_compression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeRuntimeCompressionPropertiesSchema(),
			},
			"se_dp_hm_drops": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_dp_if_state_poll_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"se_dp_log_nf_enqueue_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  70,
			},
			"se_dp_log_udf_enqueue_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  90,
			},
			"se_dump_core_on_assert": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_handle_interface_routes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_hb_persist_fudge_bits": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"se_mac_error_threshold_to_disable_promiscious": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_memory_poison": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60000,
			},
			"se_metrics_rt_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_rt_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_packet_buffer_max": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_random_tcp_drops": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_rate_limiters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeRateLimitersSchema(),
			},
			"service_ip_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"service_port_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"services_accessible_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_parse_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tcp_syncache_max_retransmit_default": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"upstream_connect_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_connpool_cache_thresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_idle_thresh_tmo": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_core_max_cache": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"upstream_connpool_strategy": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_keepalive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"upstream_read_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_send_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"user_defined_metric_age": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
		},
	}
}

func ResourcePlacementNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceDosAttackEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attack_count": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"ipgroup_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"meta_data": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAttackMetaDataSchema(),
			},
			"src_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"urls": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSCVsStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
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
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSamlIdentityProviderSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVinfraVcenterNetworkLimitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_reason": {
				Type:     schema.TypeString,
				Required: true,
			},
			"current": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceHypervisor_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"htype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"max_ips_per_nic": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_nics": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceApplicationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
			"virtualservice_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAuthorizationPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authz_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthorizationRuleSchema(),
			},
		},
	}
}

func ResourceSeGatewayHeartbeatSuccessDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHttpCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_send_cookie": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cookie_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePoolServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsRealTimeUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func ResourceAdminAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_local_user_login": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mapping_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthMappingRuleSchema(),
			},
		},
	}
}

func ResourceAlertFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_string": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceBMSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRateLimiterEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rl_resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rl_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeUpgradeStatusSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_already_upgraded_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_disconnected_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeGroupStatusSchema(),
			},
			"se_ip_missing_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_poweredoff_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_completed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_errors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeUpgradeEventsSchema(),
			},
			"se_upgrade_failed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_in_progress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_not_started": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_completed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_failed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_in_progress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_skip_suspended": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_suspended": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_errors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsErrorSchema(),
			},
		},
	}
}

func ResourceRateLimiterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_sz": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000000000,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func ResourceEventLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"context": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"details_summary": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceEventDetailsSchema(),
			},
			"event_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"event_pages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ignore_event_details_display": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"internal": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EVENT_INTERNAL",
			},
			"is_security_event": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"module": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"related_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"report_timestamp": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNTPAuthenticationKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NTP_AUTH_ALGORITHM_MD5",
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_number": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceSeGeoDbDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geo_db_profile_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geo_db_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expiry_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHealthMonitorDNSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"qtype": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_QUERY_TYPE",
			},
			"query_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rcode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RCODE_NO_ERROR",
			},
			"record_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_A",
			},
			"response_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSiteDnsVsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceErrorPageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"error_page_body_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_redirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPStatusMatchSchema(),
			},
		},
	}
}

func ResourceNetworkSecurityPolicyActionRLParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_rate": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceGslbIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceAbPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceSecureChannelTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expiry_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSecureChannelMetadataSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceContentRewriteProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"req_match_replace_pair": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema(),
			},
			"request_rewrite_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"response_rewrite_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewritable_content_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsp_match_replace_pair": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema(),
			},
		},
	}
}

func ResourceHealthMonitorRadiusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"shared_secret": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOShiftK8SConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"auto_assign_fqdn": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"avi_bridge_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ca_tls_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_tls_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_port_match_http_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"default_service_as_east_west_service": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_auto_backend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_gs_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/run/docker.sock",
			},
			"docker_registry_se": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"enable_event_subscription": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_route_ingress_hardening": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"http_container_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"ing_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIngAttributeSchema(),
			},
			"ing_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIngAttributeSchema(),
			},
			"l4_health_monitoring": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"master_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_availability_zone_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ns_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"ns_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"num_shards": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"override_service_ports": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"persistent_volume_claim": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRouteInfoSchema(),
			},
			"sdn_overlay": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_deployment_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_POD",
			},
			"se_exclude_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_image_pull_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"se_pod_tolerations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePodTolerationSchema(),
			},
			"se_priority_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_restart_batch_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_restart_force": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/avi",
			},
			"secure_egress_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_account_token": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"shard_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shared_virtualservice_namespace": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_not_ready_addresses": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_controller_image": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_resource_definition_as_ssot": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_scheduling_disabled_nodes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_service_cluster_ip_as_ew_vip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_default_gateway": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceSecureChannelMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_controller": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"local_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pub_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pub_key_pem": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SECURE_CHANNEL_NONE",
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceObjectAccessPolicyRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceObjectAccessMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceUpgradeTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIpAllocInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePGDeploymentRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "health.health_score_value",
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CO_GE",
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsQueryNameMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query_domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_RESPONSE_ROUND_ROBIN",
			},
			"cname": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"delegated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip6_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsAAAARdataSchema(),
			},
			"ip_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsARdataSchema(),
			},
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mx_records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsMxRdataSchema(),
			},
			"ns": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsNsRdataSchema(),
			},
			"num_records_in_response": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_locator": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsSrvRdataSchema(),
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"txt_records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsTxtRdataSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wildcard_match": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDnsMxRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceSummarizedSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSAMLSPConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"signing_ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_signon_url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAuthenticationActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "USE_DEFAULT_AUTHENTICATION",
			},
		},
	}
}

func ResourceOverallInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"free_percent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"used": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceMetricsDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_response_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"is_null": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"num_samples": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"prediction_interval_high": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"prediction_interval_low": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"value_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value_str_desc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePoolAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_realtime_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down_requested": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"attach_ip_status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Programming Network reachability to the Virtual Service IP in the Cloud",
			},
			"attach_ip_success": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"floating_intf_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"incarnation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_portchannel": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_primary": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_standby": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2001,
			},
			"scaleout_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sec_idx": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"snat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"vip6_subnet_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"vip_intf_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_intf_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInterfaceListSchema(),
			},
			"vip_intf_mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_subnet_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vnic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema(),
			},
		},
	}
}

func ResourceAWSSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAzureMarketplaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"offer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publisher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skus": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeIP6DadFailedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dad_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbGeoLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"source": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRollbackControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDockerRegistrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oshift_registry": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOshiftDockerRegistryMetaDataSchema(),
			},
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"private": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"registry": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks/se",
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"static_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"static_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceSnmpV3ConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"engine_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSnmpV3UserParamsSchema(),
			},
		},
	}
}

func ResourceBgpProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hold_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"ibgp": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"ip_communities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpCommunitySchema(),
			},
			"keepalive_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"local_as": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"peers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpPeerSchema(),
			},
			"send_community": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"shutdown": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeGroupStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disrupted_vs_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enqueue_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se_with_no_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se_with_vs_not_scaledout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_se_with_vs_scaledout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vs_disrupted": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_already_upgraded_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_disconnected_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ip_missing_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_poweredoff_at_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_reboot_in_progress_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_upgrade_completed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_errors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeUpgradeEventsSchema(),
			},
			"se_upgrade_failed": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_in_progress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_not_started": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_skip_suspended": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_suspended": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_with_no_vs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_with_vs_not_scaledout": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_with_vs_scaledout": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"thread": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_errors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsErrorSchema(),
			},
			"vs_migrate_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scalein_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scaleout_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"worker": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceApicAgentBridgeDomainVrfChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bridge_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_vrf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"old_vrf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceMicroServiceMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLdapAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bind_as_administrator": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"email_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "email",
			},
			"full_name_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "name",
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  389,
			},
			"security_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLdapDirectorySettingsSchema(),
			},
			"user_bind": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLdapUserBindSettingsSchema(),
			},
		},
	}
}

func ResourceVinfraVcenterDiscoveryFailureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeUpgradeScaleinEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scalein_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleinParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbPoolRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
		},
	}
}

func ResourceHealthMonitorTcpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_half_open": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tcp_request": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLicenseExpiryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"burst_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expiry_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_apps": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"throughput": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLogControllerMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0",
			},
			"node_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0",
			},
			"static_mapping": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAuthorizationRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthorizationActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthorizationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"cookie": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCookieMatchSchema(),
			},
			"hdrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceProtocolMatchSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceQueryMatchSchema(),
			},
			"version": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchSchema(),
			},
		},
	}
}

func ResourceAlertObjectListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"objects": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceLinuxServerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceLinuxServerHostSchema(),
			},
			"se_inband_mgmt": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_disk_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_log_disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_sys_disk_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_sys_disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeBootupCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buf_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"buf_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"hash_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"level_aggressive": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"level_normal": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"window_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
		},
	}
}

func ResourceDnsAttacksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attacks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsAttackSchema(),
			},
			"oper_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugVrfSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceConfigActionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameter_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVinfraVcenterBadCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"previous_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafPSMLocationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"methods": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
		},
	}
}

func ResourceObjectAccessMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"label_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label_values": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceCloudMetaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOpenStackVipNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"os_network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_tenant_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVIMgrHostRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cntlr_accessible": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "connected",
			},
			"cpu_hz": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mgmt_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_uuids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_cpu_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_cpu_packages": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_cpu_threads": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCdpLldpInfoSchema(),
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_start_ts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantined": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"quarantined_periods": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_fail_cnt": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_success_cnt": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAppInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_hdr_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"app_hdr_value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAPICConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_admin_tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "common",
			},
			"apic_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apic_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"apic_password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"apic_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"context_aware": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SINGLE_CONTEXT",
			},
			"se_tunnel_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceSeVnicUpEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAuthenticationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
		},
	}
}

func ResourceCloudFlavorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disk_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enhanced_nw": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_recommended": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"max_ip6s_per_nic": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_ips_per_nic": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_nics": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"meta": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudMetaSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ram_mb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"capture_filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCaptureFiltersSchema(),
			},
			"capture_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugVirtualServiceCaptureSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debug_hm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEBUG_VS_HM_NONE",
			},
			"debug_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"dns_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugDnsOptionsSchema(),
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVsDataplaneSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resync_flows": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugVirtualServiceSeParamsSchema(),
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
		},
	}
}

func ResourceHTTPRewriteLocHdrActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"keep_query": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHTTPPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_policy_set_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsAttackSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack_vector": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"max_mitigation_age": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"mitigation_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAttackMitigationActionSchema(),
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceURIParamTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"str_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHealthMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exact_http_request": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_request": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET / HTTP/1.0",
			},
			"http_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_response_code": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"maintenance_code": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"maintenance_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHealthMonitorSSLAttributesSchema(),
			},
		},
	}
}

func ResourceVsSeVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lif": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLRenewDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeRpcProxyDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"queue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCumulativeLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cores": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPSwitchingActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePoolServerSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPRequestRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"hdr_action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redirect_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
			},
			"rewrite_url_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPRewriteURLActionSchema(),
			},
			"switching_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPSwitchingActionSchema(),
			},
		},
	}
}

func ResourceVcenterHostsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceGCPVIPAllocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ilb": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPVIPILBSchema(),
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ROUTES",
			},
			"routes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGCPVIPRoutesSchema(),
			},
		},
	}
}

func ResourceGslbObjectInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbObjSchema(),
			},
			"object_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pb_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPResponseRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"hdr_action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"loc_hdr_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPRewriteLocHdrActionSchema(),
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceResponseMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHTTPSecurityRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPSecurityActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsScaleOutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceScaleStatusSchema(),
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLClientCertificateActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"close_connection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLClientRequestHeaderSchema(),
			},
		},
	}
}

func ResourceSeUpgradeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disruptive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"patch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"patch_rollback": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"resume_from_suspend": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rollback": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_suspended": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"suspend_on_failure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOCISetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenancy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudVipUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeHmEventShmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average_response_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_monitor": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resp_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDnsRuleActionAllowDropSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"reset_conn": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceVinfraMgmtNwChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"existing_nw": {
				Type:     schema.TypeString,
				Required: true,
			},
			"new_nw": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSummarizedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSummarizedSubnetInfoSchema(),
			},
		},
	}
}

func ResourceEventMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nodes_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceUpgradeEventSchema(),
			},
			"sub_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceUpgradeEventSchema(),
			},
			"task": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePatchInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reboot": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"reboot_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRebootDataSchema(),
			},
		},
	}
}

func ResourceHTTPResponsePolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPResponseRuleSchema(),
			},
		},
	}
}

func ResourceMicroServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"containers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMicroServiceContainerSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_list": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"orchestrator_name": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceSeAgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_echo_miss_aggressive_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"controller_echo_miss_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"controller_echo_rpc_aggressive_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_echo_rpc_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_heartbeat_miss_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"controller_heartbeat_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"controller_registration_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"controller_rpc_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"cpustats_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"ctrl_reg_pending_max_wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  150,
			},
			"debug_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dp_aggressive_deq_interval_msec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_aggressive_enq_interval_msec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_batch_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_deq_interval_msec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_enq_interval_msec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_max_wait_rsp_time_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"dp_reg_pending_max_wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  75,
			},
			"headless_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ignore_docker_mac_change": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ns_helper_deq_interval_msec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"sdb_flush_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_pipeline_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_scan_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_grp_change_disruptive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"seagent_statecache_properties": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeAgentStateCachePropertiesSchema(),
			},
			"send_se_ready_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"states_flush_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"vnic_dhcp_ip_check_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"vnic_dhcp_ip_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"vnic_ip_delete_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"vnic_probe_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"vnic_rpc_retry_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"vnicdb_cmd_history_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
		},
	}
}

func ResourceNetworkSecurityMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"microservice": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMicroServiceMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchSchema(),
			},
		},
	}
}

func ResourceConfigUserPasswordChangeRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePaaLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_hit": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"client_request_body_sent": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"request_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePaaRequestLogSchema(),
			},
		},
	}
}

func ResourceHTTPRedirectActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"keep_query": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302",
			},
		},
	}
}

func ResourceWafPSMLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceWafPSMLocationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMRuleSchema(),
			},
		},
	}
}

func ResourceRmSeBootupFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVIMgrVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_ip_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_vm": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cpu_reservation": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cpu_shares": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creation_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"guest_nic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"init_vnics": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem_shares": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_reservation": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_cpu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ovf_avisetype_field": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ver": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_rm_cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_se_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_template_vm": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vcenter_vappname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_vappvendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_vm_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_vnic_discovered": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vm_lb_weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeVnicTxQueueStallEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"queue": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVinfraDiscSummaryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clusters": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_dcs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_hosts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_nws": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_vms": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHttpCacheConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"age_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"aggressive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"date_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"heuristic_expire": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ignore_request_cache_control": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_cache_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_object_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4194304,
			},
			"mime_types_black_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_black_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"min_object_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"query_cacheable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uri_non_cacheable": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"xcache_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceOpenStackVnicChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraVmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDebugSeCpuSharesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"shares": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceSamlSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"idp": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSamlIdentityProviderSettingsSchema(),
			},
			"sp": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceSamlServiceProviderSettingsSchema(),
			},
		},
	}
}

func ResourceSeVsFaultEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fault_object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fault_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_engine": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVserverL7MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexr": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_application_response_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_cache_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_cache_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_cacheable_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_cacheable_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_complete_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_error_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_errors_excluded": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_frustrated_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_http_headers_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_http_headers_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_http_params_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_params_per_req": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_post_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_post_compression_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_pre_compression_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_reqs_per_session": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_1xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_2xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_3xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_4xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_4xx_avi_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_5xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_resp_5xx_avi_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_satisfactory_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_auth_dsa": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_auth_ecdsa": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_auth_rsa": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ecdsa_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ecdsa_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_failed_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshake_network_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshake_protocol_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshakes_new": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshakes_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshakes_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshakes_reused": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_handshakes_timedout": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_kx_dh": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_kx_ecdh": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_kx_rsa": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_rsa_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_rsa_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ver_ssl30": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ver_tls10": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ver_tls11": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ver_tls12": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_ssl_ver_tls13": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_tolerated_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_total_http2_requests": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_total_requests": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_uri_length": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_evaluated": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_evaluated_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_evaluated_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_evaluated_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_evaluated_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_flagged_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_flagged_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_flagged_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_flagged_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_latency_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_latency_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_latency_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_latency_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_matched": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_matched_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_matched_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_matched_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_matched_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_rejected_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_rejected_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_rejected_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waf_rejected_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"avg_waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_concurrent_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"max_ssl_open_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_cache_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_cacheable_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_get_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_post_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_response_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_ssl_failed_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_evaluated": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_matched": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"pct_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"rum_apdexr": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"ssl_protocol_strength": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_application_response_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_finished_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_get_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_get_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_get_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_get_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_http_headers_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_http_headers_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_http_params_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_num_page_load_time_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_num_page_load_time_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_num_rum_samples": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_other_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_other_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_other_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_other_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_post_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_post_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_post_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_post_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_post_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_reqs_finished_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_reqs_with_params": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_resp_1xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_resp_2xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_resp_3xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_resp_4xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_resp_5xx": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_total_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_uri_length": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_evaluated_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_evaluated_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_evaluated_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_evaluated_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_flagged_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_flagged_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_flagged_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_flagged_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_latency_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_latency_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_latency_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_latency_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_matched_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_matched_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_matched_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_matched_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_rejected_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_rejected_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_rejected_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waf_rejected_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"sum_waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecureChannelConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sslkeyandcertificate_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceCloudIpChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip6_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"ip_mask": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"mac_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSSLCipherListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"identified_ciphers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"unidentified_ciphers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDosThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": {
				Type:     schema.TypeString,
				Required: true,
			},
			"max_value": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"min_value": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceFailActionHTTPRedirectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTPS",
			},
			"query": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302",
			},
		},
	}
}

func ResourceSSLExpireDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days_left": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSSLKeyParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ec_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLKeyECParamsSchema(),
			},
			"rsa_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLKeyRSAParamsSchema(),
			},
		},
	}
}

func ResourceCustomParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_dynamic": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_sensitive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSubDomainPlacementRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"placement_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sub_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transition_ops": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_NONE",
			},
		},
	}
}

func ResourceAnomalyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deviation": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"metric_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_timestamp": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EXPONENTIAL_WEIGHTED_MOVING_AVG",
			},
			"models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourcePermissionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMesosMetricsDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mesos_master": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mesos_slave": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics_collection_frq": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
		},
	}
}

func ResourceConfigDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAzureNetworkInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"management_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceL4RuleProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceConfigCreateDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmAddNetworksEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmAddVnicSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
