/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip_or_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vm_hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_mor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
			},
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"customer_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_enforcement": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"expired_burst_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_on": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"reasons": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceMemoryUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sub_element": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"tenancy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_ONE",
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
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"action_success": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"end_time_str": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_se_assigned": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_requested": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scale_se": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time_str": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_placement_resolution_info": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceGslbDownloadStatusSchema(),
			},
			"gslb_download": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"instanceport": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nodeid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbRuntimeSchema(),
			},
			"watermark": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"match_value": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cif": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"https_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"last_login_timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_password_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logged_in": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"vnicdb_cmd_history_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
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
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_email": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
				Computed: true,
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
			},
			"pool_network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_network": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
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
				Default:  180,
			},
			"keepalive_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"local_as": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"md5_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"peer_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"remote_as": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shutdown": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"egress_service_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"iam_assume_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publish_vip_to_public_zone": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_access_key": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"all_response_headers": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_response_time": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"client_browser": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cipher_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCipherListSchema(),
			},
			"client_dest_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_device": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_insights": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_ip": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_log_filter_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_os": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"compression_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"connection_error_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnErrorInfoSchema(),
			},
			"data_transfer_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"datascript_error_trace": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDataScriptErrorTraceSchema(),
			},
			"datascript_log": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers_received_from_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers_sent_to_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http2_stream_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_request_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"paa_log": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirected_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"referer": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_timestamp": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"request_content_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_headers": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_served_locally_remote_site_down": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"request_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_content_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_headers": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_time_first_byte": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_time_last_byte": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rewritten_uri_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewritten_uri_query": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"server_conn_src_ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_connection_reused": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_dest_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_response_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_time_first_byte": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_time_last_byte": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_rtt": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_side_redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_src_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"spdy_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udf": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_agent": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vs_ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"waf_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafLogSchema(),
			},
			"xff": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"org_display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_entity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSamlServiceProviderNodeSchema(),
			},
			"tech_contact_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tech_contact_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceMetricLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"report_timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"time_series": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Default:  false,
			},
			"ignore_incomplete_request_body_error": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"learning_params": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"region_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"usable_network_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
				Optional: true,
			},
			"inband": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPInBandManagementSchema(),
			},
			"one_arm": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPOneArmModeSchema(),
			},
			"two_arm": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceDnsClientIpMatchSchema(),
			},
			"geo_location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsGeoLocationMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsTransportProtocolMatchSchema(),
			},
			"query_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsQueryNameMatchSchema(),
			},
			"query_type": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceVipAutoscaleConfigurationSchema(),
			},
			"policy": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"http_container_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"nuage_controller": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
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

func ResourcePerformanceLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_concurrent_connections": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_throughput": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Default:  100,
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
				Elem:     ResourceCPUUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
				Elem:     ResourceURIParamTokenSchema(),
			},
			"type": {
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
				Elem:     ResourceOperationalStatusSchema(),
			},
			"pool_id": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceCloudTenantsDeletedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenants": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudTenantCleanupSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"created_on": {
				Type:     schema.TypeString,
				Optional: true,
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
			"last_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_on": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"valid_until": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"performance_reason": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"resources_penalty": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resources_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_penalty": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"security_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_resource_prefix": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EQUAL",
			},
			"toleration_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpCommunitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:     schema.TypeList,
				Optional: true,
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

func ResourceServerAutoScaleOutCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_config_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceCloudStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cntr_public_ip": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"secret_access_key": {
				Type:     schema.TypeString,
				Required: true,
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
			},
			"datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			},
			"management_ip_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"management_network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_se_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_url": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_fail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_ok": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_fails": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"old_flv": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
		},
	}
}

func ResourceDnsQueryNameMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceNuageSDNControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nuage_organization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_vsd_host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_policy_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_zone": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"device": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmtaddr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_info_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceMatchSchema(),
			},
			"source_ip": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pending_vcenter_reqs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sevm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			},
			"vm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"used_ip_count": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsEdnsOptionSchema(),
			},
			"udp_payload_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceGslbGeoDbEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:     schema.TypeSet,
				Optional: true,
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

func ResourceTCPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dsr_profile": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTPS",
			},
			"query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302",
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

func ResourceConfigActionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disc_end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disc_start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_datacenter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_progress": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_clusters": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_dcs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcenter_req_pending": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"vcenter_template_se_location": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"writer_count": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAwsConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asg_poll_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"ebs_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceAwsEncryptionSchema(),
			},
			"secret_access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sqs_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"result": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Elem:     ResourceFlowtableProfileSchema(),
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
				Elem:     ResourceOperationalStatusSchema(),
			},
			"server_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"server_port": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"dns_vs_states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPerDnsStateSchema(),
			},
			"gs_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsGsStatusSchema(),
			},
			"retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"max": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ts": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mean": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_ts": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_samples": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"trend": {
				Type:     schema.TypeFloat,
				Optional: true,
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
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Elem:     ResourceConfigInfoSchema(),
			},
			"gap_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"geo_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"ghm_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"glb_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"gpki_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"gs_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"mm_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"repl_queue": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
			},
			"sync_info": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"total": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceAlertFilterSchema(),
			},
			"event_match_filter": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceHealthMonitorTcpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_half_open": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tcp_request": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_response": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"auto_expand": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datacenter_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dvs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_nw": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_ports": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"switch_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vlan_range": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanRangeSchema(),
			},
			"vm_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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

func ResourceSeLicensedBandwdithExceededEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_pkts_dropped": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"db_num_client_resp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_num_db_queries": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_num_db_resp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_num_oom": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_queue_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_rum_queries": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_rum_rows": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"marathon_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:8080",
			},
			"marathon_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_port_range": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"public_port_range": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Optional: true,
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
			},
			"percent_servers_up": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventServerDetailsSchema(),
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHTTPCookieDataSchema(),
			},
			"hdr": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"curr_in_service_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
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
			},
			"execution_datestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"val": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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

func ResourceCloudSyncServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceHttpCookiePersistenceKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aes_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hmac_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceVipAutoscaleZonesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHTTPLocalFileSchema(),
			},
			"redirect": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"request": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"authentication_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceGslbRuntimeSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
			},
			"third_party_site": {
				Type:     schema.TypeSet,
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

func ResourceConfigUserAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"roles": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenants": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"nsname": {
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
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nhop_inst": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nhop_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceL4RuleActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceL4RuleMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"network_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAzureNetworkInfoSchema(),
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"disk": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hyper_threading": {
				Type:     schema.TypeBool,
				Optional: true,
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

func ResourceOpenStackLbPluginOpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"elapsed": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
		},
	}
}

func ResourceIPPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_ref": {
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
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"nat_ip_range": {
				Type:     schema.TypeSet,
				Optional: true,
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
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shm": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema(),
			},
			"ssl_error_code": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"version": {
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

func ResourceLinuxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"banner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cis_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"motd": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"cookie": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"loc_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLocationHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProtocolMatchSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceHTTPStatusMatchSchema(),
			},
			"version": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
		},
	}
}

func ResourceTencentZoneNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"usable_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"remote_attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"contrail_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			},
			"nuage_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_virtualip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nuage_vsd_host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"gs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_value_to_se": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"sp_pools": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbServiceSitePersistencePoolSchema(),
			},
			"vip_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vserver_l4_metrics": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVserverL4MetricsObjSchema(),
			},
			"vserver_l7_metrics": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_prefix_len": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_prefix_len": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_ip": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_ports": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_secgrp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_svrgrp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceDnsTransportProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"dst_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"dst_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"input_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proto": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"src_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceTimeStampSchema(),
			},
			"sync_state": {
				Type:     schema.TypeString,
				Optional: true,
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
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"network_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pci_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"l4_policy_set_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_nsx_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_poll_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
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

func ResourceSeIpAddedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSnmpV3UserParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks",
			},
			"auth_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_AUTH_MD5",
			},
			"priv_passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks",
			},
			"priv_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_PRIV_DES",
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"uuid": {
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
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"gslb_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsmember": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGslbPoolMemberDetailsSchema(),
			},
			"ha_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"host_hdr": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
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
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distinguished_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"locality": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"usable_subnet_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"partition_passwd": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"partition_serial_number": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_used": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"process": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGcpInfoSchema(),
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"migrate_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
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

func ResourceMetricsDbQueueHealthyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"instanceport": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nodeid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runtime": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbRuntimeSchema(),
			},
			"watermark": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceWafPolicyWhitelistRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourcePermissionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceTencentCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"nuage_controller": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceURIParamSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamQuerySchema(),
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
				Elem:     ResourceCfgStateSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"send_event": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"send_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"services_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDiskUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"performance_rating": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_score": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClientLogStreamingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
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
			"assign_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assign_tenant": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthMatchAttributeSchema(),
			},
			"group_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthMatchGroupMembershipSchema(),
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"role_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_failed_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_scanners_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_to_uri_failed_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"client_ip_to_uri_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"custom_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_scanners_requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Optional: true,
				Computed: true,
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
			},
			"crypto_user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crypto_user_password": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"https_port": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
			},
			"glb_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rxed_site_hs": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteHealthStatusSchema(),
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_cfg": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeCfgSchema(),
			},
			"site_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
			},
			"site_stats": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeStatsSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"api_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
		},
	}
}

func ResourceGCPCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_account_keyfile_data": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"heap_config_hard_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"heap_config_soft_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"heap_config_usage": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"heap_conn_usage": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shm_config_hard_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shm_config_soft_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shm_config_usage": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shm_conn_usage": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceAlertSyslogServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSLOG_LEGACY",
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

func ResourceVserverL4MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexc": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"apdexrtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_application_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bytes_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connections_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_app_error": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bad_rst_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_ip_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_fake_session": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_abort": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_error": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_timeout": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_malformed_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_non_syn_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_cookie_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_custom_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_hdr_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_bad_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_unknown_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop_bad": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_bad_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_unknown_rl_drop": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_slow_uri": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_small_window_stress": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_ssl_error": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_syn_flood": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_total_req": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_tx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_zero_window_stress": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errored_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_l4_client_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_req": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_network_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_new_established_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pkts_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_syns": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_pkts": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_active_se": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_open_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_bytes_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_pkts_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_bytes_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_pkts_absolute": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_application_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connection_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connections_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_bandwidth": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_rx_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_network_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_pkts_dos_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_policy_drops": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_conn_duration": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_dropped_user_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connections_dropped": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dup_ack_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_conns": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_req": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_out_of_orders": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_packet_dropped_user_bandwidth_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rtt_valid_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_sack_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_flow_control": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_timeout_retransmits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_zero_window_size_events": {
				Type:     schema.TypeFloat,
				Optional: true,
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
			},
			"client_cert": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_priv_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_major_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_minor_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceAzureUserPassCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMatchReplacePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_string": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"timestamp": {
				Type:     schema.TypeFloat,
				Optional: true,
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
				Optional: true,
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
				Elem:     ResourceCC_AgentPropertiesSchema(),
			},
			"controller_props": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerPropertiesSchema(),
			},
			"flavor_props": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudFlavorSchema(),
			},
			"flavor_regex_filter": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"event": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_trace": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"value": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
		},
	}
}

func ResourceDnsPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"disable_se_reboot": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_id": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourcePoolGroupMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority_label": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"service_engine_vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_network_attachment": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_hs_db_writes": {
				Type:     schema.TypeBool,
				Optional: true,
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

func ResourceSecMgrThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"command_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_variables": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"capture_filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCaptureFiltersSchema(),
			},
			"capture_params": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"fault": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"reason_code_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceControllerUpgradeStateSchema(),
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_patch": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rollback": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeStatusSummarySchema(),
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"last_alert_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ops": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceAlertMgrDebugFilterSchema(),
			},
			"autoscale_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleMgrDebugFilterSchema(),
			},
			"cloud_connector_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudConnectorDebugFilterSchema(),
			},
			"hs_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMgrDebugFilterSchema(),
			},
			"mesos_metrics_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosMetricsDebugFilterSchema(),
			},
			"metrics_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsMgrDebugFilterSchema(),
			},
			"metricsapi_srv_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsApiSrvDebugFilterSchema(),
			},
			"se_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrDebugFilterSchema(),
			},
			"se_rpc_proxy_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeRpcProxyDebugFilterSchema(),
			},
			"state_cache_mgr_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStateCacheMgrDebugFilterSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"flow_mirror_add": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_mirror_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_mirror_del": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_probe": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_probe_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipc_batched": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipc_rx_req": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipc_rx_res": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipc_tx_req": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipc_tx_res": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vs_hb": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceGeoLocationSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"propogate_state": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"alert_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"sample_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"skip_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
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
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_ip_addr": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Type:     schema.TypeString,
				Required: true,
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
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_standby_se": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"status_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"gslb_service_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceVsApicExtensionSchema(),
			},
			"controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"datapath_debug": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"last_key_rotation_time": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"prev_metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_db": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redis_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sec_mgr_info": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Computed: true,
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
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourcePoolDeploymentUpdateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"evaluation_duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceWafDataFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"from_client": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rcv_timestamp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rx_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleMatchDataSchema(),
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"phase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceSAMLSPConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookie_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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

func ResourceVIMgrSEVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureInfoSchema(),
			},
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deletion_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"discovery_response": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk_gb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"guest_nic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_discovery": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"segroup_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_group_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vcenter_rm_cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_vm": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_vappname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vappvendor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_intf_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
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

func ResourceDnsRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
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
				Elem:     ResourceRmAddNetworksEventDetailsSchema(),
			},
			"all_seupgrade_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAllSeUpgradeEventDetailsSchema(),
			},
			"anomaly_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnomalyEventDetailsSchema(),
			},
			"apic_agent_bd_vrf_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentBridgeDomainVrfChangeSchema(),
			},
			"apic_agent_generic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentGenericEventDetailsSchema(),
			},
			"apic_agent_vs_network_error": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentVsNetworkErrorSchema(),
			},
			"avg_uptime_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAvgUptimeChangeDetailsSchema(),
			},
			"aws_asg_deletion_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSASGDeleteSchema(),
			},
			"aws_asg_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSASGNotifDetailsSchema(),
			},
			"aws_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSSetupSchema(),
			},
			"azure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureSetupSchema(),
			},
			"azure_mp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureMarketplaceSchema(),
			},
			"bind_vs_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmBindVsSeEventDetailsSchema(),
			},
			"bm_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBMSetupSchema(),
			},
			"bootup_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeBootupFailEventDetailsSchema(),
			},
			"burst_checkout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBurstLicenseDetailsSchema(),
			},
			"cc_cluster_vip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudClusterVipSchema(),
			},
			"cc_dns_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudDnsUpdateSchema(),
			},
			"cc_health_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudHealthSchema(),
			},
			"cc_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudGenericSchema(),
			},
			"cc_ip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudIpChangeSchema(),
			},
			"cc_parkintf_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipParkingIntfSchema(),
			},
			"cc_scaleset_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCCScaleSetNotifDetailsSchema(),
			},
			"cc_se_vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSeVmChangeSchema(),
			},
			"cc_sync_services_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSyncServicesSchema(),
			},
			"cc_tenant_del_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudTenantsDeletedSchema(),
			},
			"cc_vip_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipUpdateSchema(),
			},
			"cc_vnic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVnicChangeSchema(),
			},
			"cluster_config_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterConfigFailedEventSchema(),
			},
			"cluster_leader_failover_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterLeaderFailoverEventSchema(),
			},
			"cluster_node_add_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeAddEventSchema(),
			},
			"cluster_node_db_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeDbFailedEventSchema(),
			},
			"cluster_node_remove_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeRemoveEventSchema(),
			},
			"cluster_node_shutdown_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeShutdownEventSchema(),
			},
			"cluster_node_started_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeStartedEventSchema(),
			},
			"cluster_service_critical_failure_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceCriticalFailureEventSchema(),
			},
			"cluster_service_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceFailedEventSchema(),
			},
			"cluster_service_restored_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceRestoredEventSchema(),
			},
			"cluster_warm_reboot_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterWarmRebootEventSchema(),
			},
			"cntlr_host_list_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraCntlrHostUnreachableListSchema(),
			},
			"config_action_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigActionDetailsSchema(),
			},
			"config_create_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigCreateDetailsSchema(),
			},
			"config_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigDeleteDetailsSchema(),
			},
			"config_password_change_request_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserPasswordChangeRequestSchema(),
			},
			"config_se_grp_flv_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigSeGrpFlvUpdateSchema(),
			},
			"config_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUpdateDetailsSchema(),
			},
			"config_user_authrz_rule_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserAuthrzByRuleSchema(),
			},
			"config_user_login_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLoginSchema(),
			},
			"config_user_logout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLogoutSchema(),
			},
			"config_user_not_authrz_rule_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserNotAuthrzByRuleSchema(),
			},
			"container_cloud_batch_setup": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudBatchSetupSchema(),
			},
			"container_cloud_setup": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudSetupSchema(),
			},
			"container_cloud_sevice": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudServiceSchema(),
			},
			"cs_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudStackSetupSchema(),
			},
			"delete_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmDeleteSeEventDetailsSchema(),
			},
			"disable_se_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDisableSeMigrateEventDetailsSchema(),
			},
			"disc_summary": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraDiscSummaryDetailsSchema(),
			},
			"dns_sync_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSVsSyncInfoSchema(),
			},
			"docker_ucp_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerUCPSetupSchema(),
			},
			"dos_attack_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosAttackEventDetailsSchema(),
			},
			"gcp_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPSetupSchema(),
			},
			"glb_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbStatusSchema(),
			},
			"gs_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceStatusSchema(),
			},
			"host_unavail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostUnavailEventDetailsSchema(),
			},
			"hs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreDetailsSchema(),
			},
			"ip_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeIpFailEventDetailsSchema(),
			},
			"license_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseDetailsSchema(),
			},
			"license_expiry_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseExpiryDetailsSchema(),
			},
			"marathon_service_port_conflict_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMarathonServicePortConflictSchema(),
			},
			"memory_balancer_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMemoryBalancerInfoSchema(),
			},
			"mesos_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosSetupSchema(),
			},
			"metric_threshold_up_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricThresoldUpDetailsSchema(),
			},
			"metrics_db_disk_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbDiskEventDetailsSchema(),
			},
			"metrics_db_queue_full_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbQueueFullEventDetailsSchema(),
			},
			"metrics_db_queue_healthy_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbQueueHealthyEventDetailsSchema(),
			},
			"mgmt_nw_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraMgmtNwChangeDetailsSchema(),
			},
			"modify_networks_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmModifyNetworksEventDetailsSchema(),
			},
			"network_subnet_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetInfoSchema(),
			},
			"nw_subnet_clash_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetClashSchema(),
			},
			"nw_summarized_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSummarizedInfoSchema(),
			},
			"oci_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOCISetupSchema(),
			},
			"os_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackClusterSetupSchema(),
			},
			"os_ip_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackIpChangeSchema(),
			},
			"os_lbaudit_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbProvAuditCheckSchema(),
			},
			"os_lbplugin_op_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbPluginOpSchema(),
			},
			"os_se_vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSeVmChangeSchema(),
			},
			"os_sync_services_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSyncServicesSchema(),
			},
			"os_vnic_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackVnicChangeSchema(),
			},
			"pool_deployment_failure_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentFailureInfoSchema(),
			},
			"pool_deployment_success_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentSuccessInfoSchema(),
			},
			"pool_deployment_update_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentUpdateInfoSchema(),
			},
			"pool_server_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraPoolServerDeleteDetailsSchema(),
			},
			"rebalance_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceMigrateEventDetailsSchema(),
			},
			"rebalance_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleinEventDetailsSchema(),
			},
			"rebalance_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleoutEventDetailsSchema(),
			},
			"reboot_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmRebootSeEventDetailsSchema(),
			},
			"scheduler_action_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSchedulerActionDetailsSchema(),
			},
			"se_bgp_peer_state_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeBgpPeerStateChangeDetailsSchema(),
			},
			"se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrEventDetailsSchema(),
			},
			"se_dupip_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeDupipEventDetailsSchema(),
			},
			"se_gateway_heartbeat_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatFailedDetailsSchema(),
			},
			"se_gateway_heartbeat_success_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatSuccessDetailsSchema(),
			},
			"se_geo_db_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGeoDbDetailsSchema(),
			},
			"se_hb_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHBEventDetailsSchema(),
			},
			"se_hm_gs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGSDetailsSchema(),
			},
			"se_hm_gsgroup_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGslbPoolDetailsSchema(),
			},
			"se_hm_pool_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventPoolDetailsSchema(),
			},
			"se_hm_vs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventVsDetailsSchema(),
			},
			"se_ip6_dad_failed_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIP6DadFailedEventDetailsSchema(),
			},
			"se_ip_added_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpAddedEventDetailsSchema(),
			},
			"se_ip_removed_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpRemovedEventDetailsSchema(),
			},
			"se_ipfailure_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpfailureEventDetailsSchema(),
			},
			"se_licensed_bandwdith_exceeded_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeLicensedBandwdithExceededEventDetailsSchema(),
			},
			"se_memory_limit_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMemoryLimitEventDetailsSchema(),
			},
			"se_persistence_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePersistenceEventDetailsSchema(),
			},
			"se_pool_lb_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePoolLbEventDetailsSchema(),
			},
			"se_thresh_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeThreshEventDetailsSchema(),
			},
			"se_version_check_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVersionCheckFailedEventSchema(),
			},
			"se_vnic_down_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicDownEventDetailsSchema(),
			},
			"se_vnic_tx_queue_stall_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicTxQueueStallEventDetailsSchema(),
			},
			"se_vnic_up_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicUpEventDetailsSchema(),
			},
			"se_vs_fault_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVsFaultEventDetailsSchema(),
			},
			"semigrate_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMigrateEventDetailsSchema(),
			},
			"server_autoscale_failed_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleFailedInfoSchema(),
			},
			"server_autoscalein_complete_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInCompleteInfoSchema(),
			},
			"server_autoscalein_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInInfoSchema(),
			},
			"server_autoscaleout_complete_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutCompleteInfoSchema(),
			},
			"server_autoscaleout_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutInfoSchema(),
			},
			"seupgrade_disrupted_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeVsDisruptedEventDetailsSchema(),
			},
			"seupgrade_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeEventDetailsSchema(),
			},
			"seupgrade_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeMigrateEventDetailsSchema(),
			},
			"seupgrade_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleinEventDetailsSchema(),
			},
			"seupgrade_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleoutEventDetailsSchema(),
			},
			"spawn_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSpawnSeEventDetailsSchema(),
			},
			"ssl_expire_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExpireDetailsSchema(),
			},
			"ssl_export_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExportDetailsSchema(),
			},
			"ssl_renew_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewDetailsSchema(),
			},
			"ssl_renew_failed_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewFailedDetailsSchema(),
			},
			"switchover_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverEventDetailsSchema(),
			},
			"switchover_fail_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverFailEventDetailsSchema(),
			},
			"sync_services_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSyncServicesSchema(),
			},
			"system_upgrade_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSystemUpgradeDetailsSchema(),
			},
			"tencent_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTencentSetupSchema(),
			},
			"unbind_vs_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmUnbindVsSeEventDetailsSchema(),
			},
			"vca_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVCASetupSchema(),
			},
			"vcenter_connectivity_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterConnectivityStatusSchema(),
			},
			"vcenter_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterBadCredentialsSchema(),
			},
			"vcenter_disc_failure": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterDiscoveryFailureSchema(),
			},
			"vcenter_network_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterNetworkLimitSchema(),
			},
			"vcenter_obj_delete_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterObjDeleteDetailsSchema(),
			},
			"vip_autoscale": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVipScaleDetailsSchema(),
			},
			"vip_dns_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSRegisterInfoSchema(),
			},
			"vm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVmDetailsSchema(),
			},
			"vs_awaitingse_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsAwaitingSeEventDetailsSchema(),
			},
			"vs_error_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsErrorEventDetailsSchema(),
			},
			"vs_fsm_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsFsmEventDetailsSchema(),
			},
			"vs_initialplacement_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsInitialPlacementEventDetailsSchema(),
			},
			"vs_migrate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateEventDetailsSchema(),
			},
			"vs_pool_nw_fltr_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsPoolNwFilterEventDetailsSchema(),
			},
			"vs_scalein_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleInEventDetailsSchema(),
			},
			"vs_scaleout_details": {
				Type:     schema.TypeSet,
				Optional: true,
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

func ResourceFailActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_rsp": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPLocalResponseSchema(),
			},
			"redirect": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPRedirectSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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
				Computed: true,
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
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_network_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"standby": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_assigned": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"match_element": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_element_criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafExclusionTypeSchema(),
			},
			"uri_match_criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafExclusionTypeSchema(),
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
		},
	}
}

func ResourceSeUpgradeErrorsSchema() *schema.Resource {
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
			},
			"num_se_group": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
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
				Required: true,
			},
			"to_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": {
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
				Optional: true,
				Elem:     ResourceRateLimiterActionSchema(),
			},
			"burst_sz": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
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
			},
			"http_header": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
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
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip6_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"max_resp_time": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDNSConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"authoritative_domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dns_over_tcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			},
			"error_response": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_ERROR_RESPONSE_NONE",
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
				Elem:     ResourceFullClientLogsSchema(),
			},
			"metrics_realtime_update": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceVsEvStatusSchema(),
			},
			"first_se_assigned_time": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceTimeStampSchema(),
			},
			"last_scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"migrate_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_request": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"progress_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"requested_resource": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceMemoryUsageSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"queue": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceServiceMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
			},
			"source_port": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceIngAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"answer_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authentic_data": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"checking_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"identifier": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nameserver_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opcode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opt_record": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsOptRecordSchema(),
			},
			"query_or_response": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"question_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recursion_desired": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"rl_profile": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fallback_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"query_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Elem:     ResourceHSMThalesRFSSchema(),
			},
			"sluna": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMSafenetLunaSchema(),
			},
			"type": {
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceVirtualServiceResourceSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceMicroServiceContainerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"task_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"common_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distinguished_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_refreshed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_interval": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceGeoLocationSchema(),
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"site_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"capture_ipc": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCaptureIPCSchema(),
			},
			"dst_port_end": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dst_port_start": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"eth_proto": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ETH_TYPE_IPV4",
			},
			"ip_proto": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IP_TYPE_TCP",
			},
			"src_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_ack": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcp_fin": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcp_push": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcp_syn": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"se_rum_sampling_nav_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_nav_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_res_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"se_rum_sampling_res_percent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
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
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsvip_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceEmailConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_username": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
			},
			"num_file_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_rxed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_txed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_bad_responses": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_events_generated": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_skip_outstanding_requests": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
				Elem:     ResourceDnsRuleActionAllowDropSchema(),
			},
			"gslb_site_selection": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionGslbSiteSelectionSchema(),
			},
			"pool_switching": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionPoolSwitchingSchema(),
			},
			"response": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"shell_server_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"snmp_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"ssh_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"sysint_access": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
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
			},
			"geolocation_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_criteria": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"external_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prst_hdr_val": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"pid": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceTCPFastPathProfileSchema(),
			},
			"tcp_proxy_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTCPProxyProfileSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"udp_fast_path_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUDPFastPathProfileSchema(),
			},
			"udp_proxy_profile": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"event_cache": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceCfgStateSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceL4RuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceL4RulePortMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceL4RuleProtocolMatchSchema(),
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"auth_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"keystone_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceAuthenticationActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthenticationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDebugVirtualServiceCaptureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
				Type:     schema.TypeString,
				Required: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"intf_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			"max_retransmissions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"max_segment_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_syn_retransmissions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"min_rexmt_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"alertconfig_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"available_capacity": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"load": {
				Type:     schema.TypeFloat,
				Optional: true,
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
				Elem:     ResourceGslbServiceRuntimeSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"primary": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"standby": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"se_ref2": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"entity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_grace_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_first_n": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logging_freq": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_cluster_map_check": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_metrics_db_writes": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"var": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cfg_version_in_flight": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSnmpConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"large_trap_payload": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snmp_v3_config": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"num_servers": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_servers_up": {
				Type:     schema.TypeInt,
				Optional: true,
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
			"spdy_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_client_certificate_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLClientCertificateActionSchema(),
			},
			"ssl_client_certificate_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CLIENT_CERTIFICATE_NONE",
			},
			"ssl_everywhere_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			},
			"request_uri_path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"headers_sent_to_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"servers_tried": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"new_resource_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_resource_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_str": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
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
				Optional: true,
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

func ResourceSamlServiceProviderNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"latency_request_header_phase": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_body_phase": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_header_phase": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_records_in_response": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Optional: true,
			},
			"vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceTimeStampSchema(),
			},
			"num_partial_updates": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partial_update_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"request_header_value": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"sip_callid_hdr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sip_contact_hdr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sip_from_hdr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sip_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSipMessageSchema(),
			},
			"sip_to_hdr": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"metric_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_min_scale": {
				Type:     schema.TypeFloat,
				Optional: true,
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
			},
			"obj_id_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceengine_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"statistics": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricStatisticsSchema(),
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmAddVnicSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceSePoolLbEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cloud_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"udp_request": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_response": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"heat_scale_up_url": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnet": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"addr_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dclass": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mesos_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mesos_url": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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
			"udp_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
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
			},
			"val": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"client_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_log_filter_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"dns_fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"dns_qtype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRequestSchema(),
			},
			"dns_response": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsResponseSchema(),
			},
			"ds_log": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbpool_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"microservice": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mss": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_syn_retransmit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_transaction": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_protocol": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"significance": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceSipLogSchema(),
			},
			"sni_hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vs_ip6": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"dns_fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"server": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"attribute_value_list": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"answer_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authoritative_answer": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fallback_algorithm_used": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_wildcard": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nameserver_records_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opcode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opt_record": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsOptRecordSchema(),
			},
			"query_or_response": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"question_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsResourceRecordSchema(),
			},
			"recursion_available": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"recursion_desired": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"response_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"truncation": {
				Type:     schema.TypeBool,
				Optional: true,
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

func ResourceGslbPoolMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"public_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbIpAddrSchema(),
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mac_addr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
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
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafPSMLocationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMRuleSchema(),
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
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"elapsed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"series": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDataSeriesSchema(),
			},
			"start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"step": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stop": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"server_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
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
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_threshold": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceAlertMetricThresholdSchema(),
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
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"remote_mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"entity_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
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
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"prev_in_service_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"webhook_result": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Optional: true,
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
			},
			"fip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"total_records": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"match_value_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"paranoia_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_PARANOIA_LEVEL_LOW",
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"subnet6": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Optional: true,
				Computed: true,
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
			},
			"connected": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceAzureSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alb_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vnet_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			},
			"azure_userpass": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vs_rt": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_content": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pass_phrase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"alertconfig_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"available_capacity": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"load": {
				Type:     schema.TypeFloat,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  49,
			},
			"server": {
				Type:     schema.TypeList,
				Optional: true,
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
			},
			"exceptions": {
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
				Optional: true,
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
			},
			"total_records": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     ResourceClusterNodeSchema(),
			},
			"previous_leader_node": {
				Type:     schema.TypeSet,
				Optional: true,
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
			"firewall_target_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gcs_bucket_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcs_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_se_group_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"network_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPNetworkConfigSchema(),
			},
			"region_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Optional: true,
			},
			"backend_data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"frontend_data_vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"frontend_data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"frontend_data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceStringMatchSchema(),
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
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceSSLCertificateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_signing_request": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"chain_verified": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"issuer": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCertificateDescriptionSchema(),
			},
			"key_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyParamsSchema(),
			},
			"not_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_signed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeRuntimePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			},
			"mcache_fetch_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_in_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_in_max_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_in_min_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_out_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Optional: true,
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
			},
			"vs_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"max_nics": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
			},
			"vrf_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"encryption_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Elem:     ResourceSeUpgradeErrorsSchema(),
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
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_errors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsErrorSchema(),
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
			},
			"details_summary": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_details": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"obj_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"key_number": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"geo_db_profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_db_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"expiry_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
		},
	}
}

func ResourceGslbSiteDnsVsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"shared_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceDockerRegistrySchema(),
			},
			"east_west_placement_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"enable_event_subscription": {
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
			},
			"se_include_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_pod_tolerations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePodTolerationSchema(),
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"shard_prefix": {
				Type:     schema.TypeString,
				Optional: true,
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
				Default:  false,
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
			},
			"is_controller": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"local_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pub_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pub_key_pem": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceUpgradeTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"use_dedicated_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			},
			"fqdn": {
				Type:     schema.TypeList,
				Optional: true,
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

func ResourceConnErrorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_syn_retransmit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retransmits": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_num_window_shrink": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_out_of_orders": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_retransmits": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_rx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_tx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_zero_window_size_events": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeouts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_pkts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zero_window_size_events": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"is_null": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num_samples": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_high": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_low": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"value_str": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value_str_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
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
			},
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"offer": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publisher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skus": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnet_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrSchema(),
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceGeoLocationSchema(),
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLExportDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceOshiftDockerRegistryMetaDataSchema(),
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
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
			"se_repository_push": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"user": {
				Type:     schema.TypeSet,
				Optional: true,
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
			"disrupted_vs_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enqueue_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_no_vs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_vs_not_scaledout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_vs_scaledout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_disrupted": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_group_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_reboot_in_progress_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"thread": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_migrate_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scalein_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scaleout_in_progress_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"worker": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"new_vrf": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_vrf": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
				Computed: true,
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
			},
			"server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLdapDirectorySettingsSchema(),
			},
			"user_bind": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
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
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceGslbGeoLocationSchema(),
			},
			"member_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_PASSIVE_MEMBER",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
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
			},
			"burst_cores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_apps": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"throughput": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0",
			},
			"node_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceAuthorizationActionSchema(),
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthorizationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"cookie": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"method": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProtocolMatchSchema(),
			},
			"query": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceQueryMatchSchema(),
			},
			"version": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"se_log_disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_sys_disk_path": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsInfoSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"event_cache": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEventCacheSchema(),
			},
			"hs_state": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_changed_time": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_of_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SITE_STATE_NULL",
			},
			"state_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sw_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Not-Initialized",
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
			},
			"previous_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_object": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHostHdrMatchSchema(),
			},
			"methods": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
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
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"maintenance_mode": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mgmt_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"num_cpu_packages": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_cpu_threads": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCdpLldpInfoSchema(),
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine_start_ts": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantined": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"quarantined_periods": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_fail_cnt": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_success_cnt": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
			},
			"apic_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"apic_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_username": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"linux_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": {
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
				Optional: true,
				Elem:     ResourceNatPolicyActionSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNatMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"disk_gb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enhanced_nw": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_recommended": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"max_ip6s_per_nic": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ips_per_nic": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_nics": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"capture_params": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceDebugIpAddrSchema(),
			},
			"dns_options": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Optional: true,
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
				Elem:     ResourceAttackMitigationActionSchema(),
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"str_value": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"http_response_code": {
				Type:     schema.TypeList,
				Optional: true,
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
			},
			"ssl_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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
			},
			"cores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourcePoolServerSchema(),
			},
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redirect_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
			},
			"rewrite_url_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRewriteURLActionSchema(),
			},
			"switching_action": {
				Type:     schema.TypeSet,
				Optional: true,
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
				Computed: true,
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

func ResourceGslbObjectInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbObjSchema(),
			},
			"object_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pb_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceHTTPRewriteLocHdrActionSchema(),
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"match": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": {
				Type:     schema.TypeSet,
				Optional: true,
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
			"force": {
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
				Default:  false,
			},
			"rollback": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_suspended": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"suspend_on_failure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"test": {
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
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenancy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"health_monitor": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resp_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"containers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMicroServiceContainerSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_list": {
				Type:     schema.TypeBool,
				Optional: true,
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
			},
			"failure_code": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"microservice": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceMatchSchema(),
			},
			"vs_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
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
			},
			"client_request_body_sent": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
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

func ResourceSeUpgradeScaleoutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scaleout_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleoutParamsSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Required: true,
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
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"docker_ucp_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"failed_hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_deploy_method_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ucp_nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_vm": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cpu_reservation": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpu_shares": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creation_in_progress": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"guest_nic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem_shares": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_reservation": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_cpu": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ovf_avisetype_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"powerstate": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vcenter_rm_cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_vm": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_vappname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vappvendor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vnic_discovered": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vm_lb_weight": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"num_dcs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
				Computed: true,
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
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceSamlIdentityProviderSettingsSchema(),
			},
			"sp": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"fault_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_engine": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"avg_application_response_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_error_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errors_excluded": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_frustrated_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_http_headers_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_http_headers_count": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_http_params_count": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_params_per_req": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_post_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_post_compression_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pre_compression_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_reqs_per_session": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_1xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_2xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_3xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx_avi_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx_avi_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_satisfactory_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_dsa": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_ecdsa": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_rsa": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_failed_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_network_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_protocol_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_new": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_reused": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_timedout": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_dh": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_ecdh": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_rsa": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_non_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_pfs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_ssl30": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls10": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls11": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls12": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tolerated_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_http2_requests": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_requests": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_uri_length": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_concurrent_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ssl_open_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_cache_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_cacheable_hits": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_get_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_post_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_response_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_ssl_failed_connections": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_evaluated": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_matched": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_apdexr": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"ssl_protocol_strength": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_application_response_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_blocking_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_browser_rendering_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dns_lookup_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dom_content_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_errors": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_http_headers_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_http_headers_count": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_http_params_count": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_rum_samples": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_download_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_load_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_bytes": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket1": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket2": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_reqs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_redirection_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_reqs_finished_sessions": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_reqs_with_params": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_1xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_2xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_3xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_4xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_5xx": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rum_client_data_transfer_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_rtt": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_service_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_responses": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_uri_length": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_attacks": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_disabled": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_body_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_header_phase": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waiting_time": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceSecureChannelConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bypass_secure_channel_must_checks": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sslkeyandcertificate_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip6": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"port_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
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

func ResourceSSLExpireDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days_left": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     ResourceSSLKeyECParamsSchema(),
			},
			"rsa_params": {
				Type:     schema.TypeSet,
				Optional: true,
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
			},
			"sub_domain": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"obj_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
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
			},
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"se_cookie": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"mesos_slave": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_entity": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceL4RuleProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
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
