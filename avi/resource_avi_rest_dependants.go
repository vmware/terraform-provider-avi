// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceACSubjectInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
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

func ResourceACUserIdentitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
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

func ResourceALBServicesAccountSchema() *schema.Resource {
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
			"users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceALBServicesAccountUserSchema(),
			},
		},
	}
}

func ResourceALBServicesAccountUserSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceALBServicesCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asset_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_attachments": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceALBServicesCaseAttachmentSchema(),
			},
			"case_created_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceALBServicesUserSchema(),
			},
			"created_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deployment_environment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fr_business_justification": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fr_current_solution": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fr_timing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fr_use_cases": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_modified_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
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

func ResourceALBServicesCaseAttachmentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attachment_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attachment_size": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attachment_url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceALBServicesJobParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
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

func ResourceALBServicesStatusDetailsSchema() *schema.Resource {
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

func ResourceALBServicesUserSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_accounts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceALBServicesAccountSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAPICConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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

func ResourceAbPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ratio": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAdaptReplEventInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigVersionStatusSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recommendation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAdminAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_local_user_login": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"remote_auth_configurations": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceRemoteAuthConfigurationSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"throttle_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"timestamp": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceAlertMetricThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceAlertObjectListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceAlertRuleMetricSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceAlertSyslogServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anon_auth": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "514",
				ValidateFunc: validateInteger,
			},
			"tls_enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"udp": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"num_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceAlternateAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"mapping_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthMappingRuleSchema(),
			},
		},
	}
}

func ResourceAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"learning_log_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLearningLogPolicySchema(),
			},
			"metrics_realtime_update": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsRealTimeUpdateSchema(),
			},
			"significant_log_throttle": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"udf_log_throttle": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAnomalyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deviation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceApiVersionDeprecatedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version_used": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_supported_api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
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

func ResourceApicAgentBridgeDomainVrfChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceApicAgentGenericEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceApicAgentVsNetworkErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
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

func ResourceAppLearningConfidenceOverrideSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"confid_high_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9500",
				ValidateFunc: validateInteger,
			},
			"confid_low_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "7500",
				ValidateFunc: validateInteger,
			},
			"confid_probable_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9000",
				ValidateFunc: validateInteger,
			},
			"confid_very_high_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9999",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAppLearningParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_per_uri_learning": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"learn_from_authenticated_clients_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"max_params": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"max_uris": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "500",
				ValidateFunc: validateInteger,
			},
			"min_hits_to_learn": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"sampling_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"trusted_ipgroup_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAppSignatureConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_signature_sync_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1440",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAppSignatureEventDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_successful_updated_time": {
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
		},
	}
}

func ResourceApplicationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
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

func ResourceApplicationLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"auth_status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_STATUS_NO_AUTHENTICATION",
			},
			"avg_ingress_latency_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"avg_ingress_latency_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"body_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NOT_UPDATED",
			},
			"bot_management_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotManagementLogSchema(),
			},
			"cache_disabled_by_ds": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"cache_hit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"cacheable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"client_device": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_fingerprints": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceClientFingerprintsSchema(),
			},
			"client_insights": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"client_src_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compression_percentage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"conn_est_time_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"conn_est_time_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"connection_error_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConnErrorInfoSchema(),
			},
			"critical_error_encountered": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"grpc_method_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grpc_service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grpc_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"grpc_status_reason_phrase": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"icap_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIcapLogSchema(),
			},
			"jwt_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceJwtLogSchema(),
			},
			"log_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"max_ingress_latency_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_ingress_latency_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"ntlm_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNtlmLogSchema(),
			},
			"oauth_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthLogSchema(),
			},
			"ocsp_status_resp_sent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"oob_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOutOfBandRequestLogSchema(),
			},
			"paa_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePaaLogSchema(),
			},
			"persistence_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"persistent_session_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"request_content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"request_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"request_served_locally_remote_site_down": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"request_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_time_first_byte": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_time_last_byte": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"saml_auth_response_received": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"saml_auth_session_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"saml_authentication_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"saml_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSamlLogSchema(),
			},
			"saml_session_cookie_valid": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"server_conn_src_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_conn_src_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_connection_reused": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"server_dest_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"server_push_initiated": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"server_pushed_request": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"server_response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_response_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_response_time_first_byte": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_response_time_last_byte": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_side_redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_src_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_ssl_session_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_ssl_session_reused": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"servers_tried": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"service_engine": {
				Type:     schema.TypeString,
				Required: true,
			},
			"session_id": {
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
			"source_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"source_ip6": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"udf": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"vh_match_rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtualservice": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceAttachIpStatusEventDetailsSchema() *schema.Resource {
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
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAttackDnsAmplificationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"record_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceAttackMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"amplification": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAttackDnsAmplificationSchema(),
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_resp_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAttackMitigationActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deny": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceAuditComplianceEventInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"core_archive": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detailed_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_generated_by_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeString,
				Required: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subjects": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceACSubjectInfoSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_identities": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceACUserIdentitySchema(),
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAuthAttributeMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_value_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceStringMatchSchema(),
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
			"assign_userprofile": {
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
			"default_tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthMatchGroupMembershipSchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"is_superuser": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
			"userprofile_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userprofile_ref": {
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
				Required: true,
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

func ResourceAuthMatchGroupMembershipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAuthProfileHTTPClientParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_expiration_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
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

func ResourceAuthTacacsPlusAttributeValuePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mandatory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceAuthnRuleMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rule_action": {
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

func ResourceAuthorizationActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"status_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALLOW_ACCESS",
			},
		},
	}
}

func ResourceAuthorizationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceJWTMatchSchema(),
			},
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

func ResourceAuthorizationRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceAuthorizationActionSchema(),
			},
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"match": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceAuthorizationMatchSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAuthzRuleMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rule_action": {
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

func ResourceAutoScaleMesosSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"force": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceAutoScaleMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_aws_autoscale_integration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"intelligent_autoscale_period": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pool_ref": {
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

func ResourceAvgUptimeChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "600",
				ValidateFunc: validateInteger,
			},
			"ebs_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAwsEncryptionSchema(),
			},
			"free_elasticips": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"iam_assume_role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_vip_to_public_zone": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "us-west-1",
			},
			"route53_integration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"use_iam_roles": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_sns_sqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"des_id": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_enhanced_ha": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_managed_disks": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"use_standard_alb": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceBOTLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"hdrs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"mapping_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceBfdProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"minrx": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"mintx": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"multi": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceBgpPeerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_snat_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"advertise_vip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"advertisement_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"bfd": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"connect_timer": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"ebgp_multihop": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"hold_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ibgp_local_as_override": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"keepalive_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_as": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"shutdown": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceBgpProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hold_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "180",
				ValidateFunc: validateInteger,
			},
			"ibgp": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"ip_communities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpCommunitySchema(),
			},
			"keepalive_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"local_as": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"local_preference": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_as_path_prepend": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"peers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpPeerSchema(),
			},
			"routing_options": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpRoutingOptionsSchema(),
			},
			"send_community": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"shutdown": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceBgpRoutingOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_default_route": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"advertise_learned_routes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"learn_only_default_route": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"learn_routes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"max_learn_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "50",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceBotAllowListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBotAllowRuleSchema(),
			},
		},
	}
}

func ResourceBotAllowRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceMatchTargetSchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceBotClassMatcherSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_classes": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"op": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IS_IN",
			},
		},
	}
}

func ResourceBotClassificationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_defined_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceBotConfigIPLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ip_location_db_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_cloud_providers_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_search_engines_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceBotConfigIPReputationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ip_reputation_db_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_ip_reputation_mapping_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceBotConfigUserAgentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"use_tls_fingerprint": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceBotDetectionMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classifications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBotClassificationSchema(),
			},
			"match_operation": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceBotEvaluationResultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"component": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"confidence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identification": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotIdentificationSchema(),
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceBotIdentificationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identifier": {
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

func ResourceBotManagementLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classification": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotClassificationSchema(),
			},
			"results": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBotEvaluationResultSchema(),
			},
		},
	}
}

func ResourceBotMappingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classification": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceBotClassificationSchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"match": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceBotMappingRuleMatchTargetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceBotMappingRuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class_matcher": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotClassMatcherSchema(),
			},
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"component_matcher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"identifier_matcher": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
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
			"type_matcher": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotTypeMatcherSchema(),
			},
		},
	}
}

func ResourceBotTypeMatcherSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"op": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IS_IN",
			},
		},
	}
}

func ResourceBuildInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"build_no": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceBurstLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceCC_AgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"async_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"async_retries_delay": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"poll_duration_target": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"poll_fast_target": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"poll_slow_target": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "240",
				ValidateFunc: validateInteger,
			},
			"vcenter_host_ping_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"vcenter_inventory_max_object_updates": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"vcenter_max_datastore_go_routines": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"vcenter_reconcile_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600",
				ValidateFunc: validateInteger,
			},
			"vnic_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"vnic_retries_delay": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceCC_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rpc_poll_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"rpc_queue_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
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

func ResourceCPUUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"num_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceCRSDeploymentFailureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crs_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCRSDetailsSchema(),
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCRSDeploymentSuccessSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crs_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCRSDetailsSchema(),
			},
		},
	}
}

func ResourceCRSDetailsSchema() *schema.Resource {
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

func ResourceCRSUpdateDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crs_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCRSDetailsSchema(),
			},
		},
	}
}

func ResourceCaptureFileSizeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"absolute_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"percentage_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"dst_port_start": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"src_port_range_end": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"tcp_ack": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"tcp_fin": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"tcp_push": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"tcp_syn": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceCaptureIPCSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flow_del_probe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"flow_mirror_add": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"flow_mirror_all": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"flow_mirror_del": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"flow_probe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"flow_probe_all": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ipc_batched": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ipc_rx_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ipc_rx_res": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ipc_tx_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ipc_tx_res": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"vs_hb": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceCaseConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_auto_case_creation_on_controller_failure": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_auto_case_creation_on_se_failure": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_cleanup_of_attached_files": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourceCentralLicenseRefreshDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceCentralLicenseSubscriptionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message": {
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

func ResourceCfgStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cfg_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cfg_version_in_flight": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceChildProcessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceClientFingerprintsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filtered_tls_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_tls_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"normalized_tls_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_client_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTlsClientInfoSchema(),
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

func ResourceClientLogConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_significant_log_collection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourceClientLogFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"enabled": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceClientLogStreamingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_server": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_server_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "514",
				ValidateFunc: validateInteger,
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
			"marker_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"max_logs_per_second": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
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

func ResourceCloudConnectorDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_se_reboot": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"se_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceCloudFlavorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disk_gb": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"enhanced_nw": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_recommended": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"max_ip6s_per_nic": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_ips_per_nic": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_nics": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ram_mb": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"ip_mask": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
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

func ResourceCloudRouteNotifDetailsSchema() *schema.Resource {
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
			"route_table": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_secgrp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_svrgrp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceClusterHAConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_vsphere_ha": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vmg_name": {
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

func ResourceClusterNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"categories": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceControllerInterfaceSchema(),
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
			"static_routes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticRouteSchema(),
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

func ResourceClusterNodeDbFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceClusterWarmRebootEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceCompressionProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buf_num": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"buf_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"compressible_content_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compression": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCompressionFilterSchema(),
			},
			"hash_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16384",
				ValidateFunc: validateInteger,
			},
			"level_aggressive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"level_normal": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"max_low_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"min_high_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "200",
				ValidateFunc: validateInteger,
			},
			"min_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"mobile_str_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_accept_encoding_header": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"window_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
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

func ResourceConfigInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queue": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema(),
			},
			"reader_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"writer_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceConfigPbAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"userprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"client_type": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceConfigVersionStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceConnErrorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_syn_retransmit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_window_shrink": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"out_of_orders": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"retransmits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"rx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_num_window_shrink": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_out_of_orders": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_retransmits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_rx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_timeouts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_tx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_zero_window_size_events": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"timeouts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"tx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"zero_window_size_events": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceConnPoolPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"upstream_connpool_conn_idle_tmo": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60000",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_conn_life_tmo": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "600000",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_conn_max_reuse": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_server_max_cache": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceConnectionEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceConnectionLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"average_turntime": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"avg_ingress_latency_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"avg_ingress_latency_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"client_dest_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"client_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"client_src_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"conn_est_time_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"conn_est_time_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"connection_ended": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
			"dns_tcp_conn_close_from_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"max_ingress_latency_be": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_ingress_latency_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"network_security_policy_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_syn_retransmit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_transaction": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_window_shrink": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ocsp_status_resp_sent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"out_of_orders": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"persistence_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"retransmits": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"rx_bytes": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"rx_pkts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_conn_src_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_conn_src_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_dest_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_out_of_orders": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_retransmits": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_rtt": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_rx_bytes": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_rx_pkts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_src_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_timeouts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_total_bytes": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_total_pkts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_tx_bytes": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_tx_pkts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"server_zero_window_size_events": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"timeouts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"total_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"total_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"total_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"tx_bytes": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"tx_pkts": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"udf": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"vcpu_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"virtualservice": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vs_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zero_window_size_events": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceContainerCloudSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_access": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceContentLibConfigSchema() *schema.Resource {
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
		},
	}
}

func ResourceContentRewriteProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rewritable_content_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsp_rewrite_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRspContentRewriteRuleSchema(),
			},
		},
	}
}

func ResourceControllerAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_event_thresholds": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsEventThresholdSchema(),
			},
		},
	}
}

func ResourceControllerCloudLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clouds": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceControllerDiscontinuousTimeChangeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_servers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceControllerFaultsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_scheduler_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"cluster_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"deprecated_api_version_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"license_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"migration_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"sslprofile_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceControllerInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_controller_mem_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceControllerInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"gateway6": {
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
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_ip_or_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceControllerInternalAuthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"symmetric_jwks_keys": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceJWSKeySchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"customer_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_enforcement": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"expired_burst_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"initialized": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sockets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceControllerLicenseReconcileDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"new_available_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"new_consumed_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"new_escrow_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"new_remaining_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"old_available_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"old_consumed_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"old_escrow_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"old_remaining_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceControllerLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bot_limits": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBOTLimitsSchema(),
			},
			"certificates_per_virtualservice": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"controller_cloud_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceControllerCloudLimitsSchema(),
			},
			"controller_sizing_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceControllerSizingLimitsSchema(),
			},
			"default_routes_per_vrfcontext": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"gateway_mon_per_vrf": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ipaddress_limits": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIPAddrLimitsSchema(),
			},
			"ips_per_ipgroup": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"l7_limits": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL7limitsSchema(),
			},
			"poolgroups_per_virtualservice": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pools_per_poolgroup": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pools_per_virtualservice": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"routes_per_vrfcontext": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"rules_per_nat_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"rules_per_networksecuritypolicy": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"servers_per_pool": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"sni_children_per_parent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"strings_per_stringgroup": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vs_bgp_scaleout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vs_l2_scaleout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"waf_limits": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceWAFLimitsSchema(),
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

func ResourceControllerSizeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_cpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"min_memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceControllerSizingCloudLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clouds": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceControllerSizingLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_sizing_cloud_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceControllerSizingCloudLimitsSchema(),
			},
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_clouds": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_east_west_virtualservices": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_servers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_serviceengines": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_tenants": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_virtualservices": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_virtualservices_rt_metrics": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vrfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceControllerVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fips_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
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

func ResourceCumulativeLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_ses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_bandwidth_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sockets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCustomIpamSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceCustomParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_dynamic": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_sensitive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceDNSQueryErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceDNSVsSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_records": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDSRequestLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ds_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": {
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
			"http_response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"servers_tried": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"source_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"total_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
		},
	}
}

func ResourceDataNetworkConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tier1_segment_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtTier1SegmentConfigSchema(),
			},
			"transport_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tz_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_segments": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceDatabaseEventInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"component": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDbAppLearningInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceURIInfoSchema(),
			},
			"vs_uuid": {
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
			"logmanager_debug_filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLogManagerDebugFilterSchema(),
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

func ResourceDebugSeAgentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"log_every_n": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceDebugSeCpuSharesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"shares": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceDebugSeFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"faults": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeFaultSchema(),
			},
			"se_malloc_fail_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_malloc_fail_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_mbuf_cl_sanity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_shm_malloc_fail_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_shm_malloc_fail_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_waf_alloc_fail_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_waf_learning_alloc_fail_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDebugServiceEngineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"benchmark_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_BENCHMARK_MODE_DROP",
			},
			"benchmark_layer": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_BENCHMARK_LAYER_NONE",
			},
			"benchmark_option": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_BENCHMARK_REFLECT_SWAP_L4",
			},
			"benchmark_rss_hash": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_BENCHMARK_DISABLE_HASH",
			},
			"capture": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_rpc_timing_profiler": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"selogagent_debug": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugSeAgentSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_memory": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugTraceMemorySchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceDebugTraceMallocTypesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"malloc_type_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDebugTraceMemorySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"trace_malloc_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugTraceMallocTypesSchema(),
			},
			"trace_shm_malloc_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugTraceShmMallocTypesSchema(),
			},
		},
	}
}

func ResourceDebugTraceShmMallocTypesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"shm_malloc_type_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDebugVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
			"latency_audit_filters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCaptureFiltersSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"objsync": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDebugVirtualServiceObjSyncSchema(),
			},
			"resync_flows": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"enable_ssl_session_key_capture": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"file_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"num_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pcap_ng": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"pkt_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDebugVirtualServiceObjSyncSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"trigger_initial_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceDebugVrfContextSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_buffer_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"command_buffer_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32768",
				ValidateFunc: validateInteger,
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVrfSchema(),
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

func ResourceDetachIpStatusEventDetailsSchema() *schema.Resource {
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
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceDnsAttackSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack_vector": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"max_mitigation_age": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"mitigation_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAttackMitigationActionSchema(),
			},
			"threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceDnsClientIpMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"use_edns_client_subnet_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceDnsClientPortMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ports": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourcePortMatchGenericSchema(),
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

func ResourceDnsEdnsOptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr_family": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_prefix_len": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"source_prefix_len": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"subnet_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"subnet_ip6": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_records_in_response": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_A",
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceDnsOptRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dnssec_ok": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsEdnsOptionSchema(),
			},
			"udp_payload_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceDnsRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"service_locator": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsSrvRdataSchema(),
			},
			"ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceDnsRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"answer_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"authentic_data": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"checking_disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"client_location": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"nameserver_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"question_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"recursion_desired": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceDnsResolverSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fixed_ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"min_ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"nameserver_ips": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"resolver_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_mgmt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dclass": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"priority": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceDnsResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"answer_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"authoritative_answer": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"fallback_algorithm_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"is_wildcard": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"nameserver_records_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"question_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsResourceRecordSchema(),
			},
			"recursion_available": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"recursion_desired": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"response_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"truncation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"log": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
			"gs_group_selection": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsRuleActionGsGroupSelectionSchema(),
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

func ResourceDnsRuleActionAllowDropSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"reset_conn": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceDnsRuleActionGsGroupSelectionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"site_name": {
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

func ResourceDnsRuleActionResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authoritative": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceDnsRuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsClientIpMatchSchema(),
			},
			"client_port_numbers": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDnsClientPortMatchSchema(),
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

func ResourceDnsServiceApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aaaa_empty_response": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"admin_email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "hostmaster",
			},
			"client_dns_tcp_request_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"close_tcp_connection_post_response": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"dns_over_tcp_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"edns": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"edns_client_subnet_prefix_len": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"num_dns_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
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
			"pass_through": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"record_ttl": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceDnsSrvRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"priority": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default.host",
			},
			"weight": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceDockerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_frontend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_se_creation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"feproxy_container_port_as_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "25",
				ValidateFunc: validateInteger,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"services_accessible_all_interfaces": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_controller_image": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceDockerUCPSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"docker_ucp_access": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceDosAttackEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attack_count": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
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

func ResourceDosThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": {
				Type:     schema.TypeString,
				Required: true,
			},
			"max_value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"min_value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"email_timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"from_email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin@avicontroller.net",
			},
			"from_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mail_server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "localhost",
			},
			"mail_server_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "25",
				ValidateFunc: validateInteger,
			},
			"smtp_type": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceErrorPageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceEventCacheSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_state": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"exceptions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adaptrepl_event": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAdaptReplEventInfoSchema(),
			},
			"add_networks_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmAddNetworksEventDetailsSchema(),
			},
			"albservices_case_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceALBServicesCaseSchema(),
			},
			"albservices_file_upload_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceALBServicesFileUploadSchema(),
			},
			"albservices_status_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceALBServicesStatusDetailsSchema(),
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
			"api_version_deprecated": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceApiVersionDeprecatedSchema(),
			},
			"app_signature_event_data": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAppSignatureEventDataSchema(),
			},
			"attach_ip_status_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAttachIpStatusEventDetailsSchema(),
			},
			"avg_uptime_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAvgUptimeChangeDetailsSchema(),
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
			"central_license_refresh_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCentralLicenseRefreshDetailsSchema(),
			},
			"central_license_subscription_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCentralLicenseSubscriptionDetailsSchema(),
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
			"cloud_route_notif_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudRouteNotifDetailsSchema(),
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
			"connection_event": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConnectionEventDetailsSchema(),
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
			"controller_discontinuous_time_change_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceControllerDiscontinuousTimeChangeEventDetailsSchema(),
			},
			"controller_license_reconcile_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceControllerLicenseReconcileDetailsSchema(),
			},
			"crs_deployment_failure": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCRSDeploymentFailureSchema(),
			},
			"crs_deployment_success": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCRSDeploymentSuccessSchema(),
			},
			"crs_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCRSDetailsSchema(),
			},
			"crs_update_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCRSUpdateDetailsSchema(),
			},
			"cs_infra_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCloudStackSetupSchema(),
			},
			"database_event_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDatabaseEventInfoSchema(),
			},
			"delete_se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRmDeleteSeEventDetailsSchema(),
			},
			"detach_ip_status_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDetachIpStatusEventDetailsSchema(),
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
			"dns_query_error": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDNSQueryErrorSchema(),
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
			"false_positive_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFalsePositiveDetailsSchema(),
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
			"generic_audit_compliance_event_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuditComplianceEventInfoSchema(),
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
			"ip_threat_db_event_data": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIPThreatDBEventDataSchema(),
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
			"license_tier_switch_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLicenseTierSwitchDetiailsSchema(),
			},
			"license_transaction_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLicenseTransactionDetailsSchema(),
			},
			"log_agent_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLogAgentEventDetailSchema(),
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
			"metrics_db_sync_failure_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsDbSyncFailureEventDetailsSchema(),
			},
			"metrics_grpc_auth_failure_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsGRPCAuthFailureDetailsSchema(),
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
			"nsxt_endpoint_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtSIEndpointDetailsSchema(),
			},
			"nsxt_img_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtImageDetailsSchema(),
			},
			"nsxt_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtSetupSchema(),
			},
			"nsxt_policy_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtSIPolicyDetailsSchema(),
			},
			"nsxt_rule_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtSIRuleDetailsSchema(),
			},
			"nsxt_service_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtSIServiceDetailsSchema(),
			},
			"nsxt_t1_seg_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtT1SegDetailsSchema(),
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
			"os_api_ver_check_failure": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpenStackApiVersionCheckFailureSchema(),
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
			"se_bgp_peer_down_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeBgpPeerDownDetailsSchema(),
			},
			"se_bgp_peer_state_change_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeBgpPeerStateChangeDetailsSchema(),
			},
			"se_debug_mode_event_detail": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeDebugModeEventDetailSchema(),
			},
			"se_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeMgrEventDetailsSchema(),
			},
			"se_discontinuous_time_change_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeDiscontinuousTimeChangeEventDetailsSchema(),
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
			"se_hb_recovered_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHbRecoveredEventDetailsSchema(),
			},
			"se_high_egress_proc_latency_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHighEgressProcLatencyEventDetailsSchema(),
			},
			"se_high_ingress_proc_latency_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeHighIngressProcLatencyEventDetailsSchema(),
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
			"se_ntp_synchronization_failed": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeNtpSynchronizationFailedSchema(),
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
			"se_reconcile_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeReconcileDetailsSchema(),
			},
			"se_thresh_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeThreshEventDetailsSchema(),
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
			"se_vs_del_flows_disrupted": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeVsDelFlowsDisruptedSchema(),
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
			"sec_mgr_data_event": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecMgrDataEventSchema(),
			},
			"sec_mgr_ua_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecMgrUAEventDetailsSchema(),
			},
			"secure_key_exchange_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecureKeyExchangeDetailsSchema(),
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
			"ssl_ignored_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLIgnoredDetailsSchema(),
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
			"ssl_revoked_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSSLRevokedDetailsSchema(),
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
			"vcenter_cluster_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVcenterClusterDetailsSchema(),
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
			"vcenter_img_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVcenterImageDetailsSchema(),
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
			"vcenter_tag_event_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVcenterTagEventDetailsSchema(),
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
			"vip_symmetry_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVipSymmetryDetailsSchema(),
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
			"vs_switchover_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsSwitchoverEventDetailsSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"internal": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EVENT_INTERNAL",
			},
			"is_security_event": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceFTPProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deactivate_active": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"deactivate_passive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceFalsePositiveDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"false_positive_results": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceFalsePositiveResultSchema(),
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceFalsePositiveLearningConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_apps_supported": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"min_monitor_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10080",
				ValidateFunc: validateInteger,
			},
			"min_trans_per_application": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5000000",
				ValidateFunc: validateInteger,
			},
			"min_trans_per_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceFalsePositiveResultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"confidence": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"false_positive": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"fp_result_header": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceFalsePositiveResultHeaderSchema(),
			},
			"http_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_request_header_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHeaderInfoInURISchema(),
			},
			"params_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceParamsInURISchema(),
			},
			"rule_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceRuleInfoSchema(),
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri_result_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceFalsePositiveResultHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_analysis_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_data_received_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_data_received_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_analysis_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transactions_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceFbGsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
		},
	}
}

func ResourceFbPoolInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
		},
	}
}

func ResourceFbSeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
		},
	}
}

func ResourceFbVsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "80",
				ValidateFunc: validateInteger,
			},
			"subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
			},
			"token": {
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

func ResourceFlowtableProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"icmp_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"tcp_closed_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"tcp_connection_setup_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"tcp_half_closed_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"tcp_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"tcp_reset_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"udp_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceFullClientLogsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"enabled": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"throttle": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
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
			"gcp_service_account_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"management_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGCPSeGroupConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_data_vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backend_data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backend_data_vpc_subnet_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceGCPTwoArmModeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_data_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_data_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"management_vpc_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_vpc_subnet_name": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceGCPVIPRoutesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_se_group_subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"route_priority": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"gateway_monitor_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"gateway_monitor_success_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "15",
				ValidateFunc: validateInteger,
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

func ResourceGcpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"machine_type": {
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

func ResourceGeoDBFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"file_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGeoDBMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elements": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceGeoDBMappingElementSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGeoDBMappingElementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
			"values": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGeoLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latitude": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"longitude": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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

func ResourceGeoMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_operation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"values": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceGslbDnsGeoUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"partial_update_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
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

func ResourceGslbDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_active": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceGslbDnsSeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceGslbDnsUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"uuid": {
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

func ResourceGslbGeoDbEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceGslbGeoDbFileSchema(),
			},
			"priority": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceGslbGeoDbFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_id_checksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"timestamp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"consistent_hash_mask6": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"fallback_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manual_resume": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"members": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceGslbPoolMemberSchema(),
			},
			"min_health_monitors_up": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
			"preference_order": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"public_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbIpAddrSchema(),
			},
			"ratio": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"resolve_fqdn_to_v6": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
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
			"health_monitor_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_value_to_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceGslbReplicationStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"acknowledged_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pending_object_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"received_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"dns_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceGslbServiceDownResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fallback_cname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

func ResourceGslbServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"send_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceGslbServiceSitePersistencePoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_http2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_servers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_servers_up": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"hm_proxies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"hm_shard_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "443",
				ValidateFunc: validateInteger,
			},
			"ratio": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"suspend_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"last_fail_obj": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigVersionStatusSchema(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recommendation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_state": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceGslbSiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"glb_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_monitor_info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replication_stats": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbReplicationStatsSchema(),
			},
			"rxed_site_hs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteHealthStatusSchema(),
			},
			"send_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"gjwt_info": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"event_cache": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceEventCacheSchema(),
			},
			"hs_state": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceGslbSiteRuntimeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_file_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_file_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gap_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_geo_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_ghm_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gjwt_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_glb_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gpki_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_cr_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_del_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gs_upd_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_cert_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_cert_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_cert_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_cr_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_del_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_gssl_upd_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_health_msgs_rxed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_health_msgs_txed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_of_bad_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_of_events_generated": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_of_skip_outstanding_requests": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_of_timeouts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceGslbSubDomainPlacementRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"placement_allowed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceGslbThirdPartySiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbThirdPartySiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"health_monitor_info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"site_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
			},
		},
	}
}

func ResourceHSMAwsCloudHsmSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_config": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"mgmt_config": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"session_minor_number": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceHSMSafenetLunaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_group_num": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"is_ha": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceHSMSafenetLunaServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"priority": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"remote_ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"remote_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9004",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9004",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceHTTP2ApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_http2_server_push": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"http2_initial_window_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "64",
				ValidateFunc: validateInteger,
			},
			"max_http2_concurrent_pushes_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"max_http2_concurrent_streams_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"max_http2_control_frames_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"max_http2_empty_data_frames_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"max_http2_header_field_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"max_http2_queued_frames_to_client_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"max_http2_requests_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceHTTP2PoolPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_http2_control_frames_per_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_http2_header_field_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceHTTPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_dots_in_header_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"cache_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHttpCacheConfigSchema(),
			},
			"client_body_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30000",
				ValidateFunc: validateInteger,
			},
			"client_header_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"client_max_body_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"client_max_header_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "12",
				ValidateFunc: validateInteger,
			},
			"client_max_request_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "48",
				ValidateFunc: validateInteger,
			},
			"collect_client_tls_fingerprint": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"compression_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceCompressionProfileSchema(),
			},
			"connection_multiplexing_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"detect_ntlm_app": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"disable_keepalive_posts_msie6": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"disable_sni_hostname_check": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_chunk_merge": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_fire_and_forget": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_request_body_buffering": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_request_body_metrics": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"fwd_close_hdr_for_bound_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"hsts_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"hsts_max_age": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "365",
				ValidateFunc: validateInteger,
			},
			"hsts_subdomains_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"http2_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTP2ApplicationProfileSchema(),
			},
			"http_to_https": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"http_upstream_buffer_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"httponly_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"keepalive_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"keepalive_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30000",
				ValidateFunc: validateInteger,
			},
			"max_bad_rps_cip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_bad_rps_cip_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_bad_rps_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_header_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "256",
				ValidateFunc: validateInteger,
			},
			"max_keepalive_requests": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"max_response_headers_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "48",
				ValidateFunc: validateInteger,
			},
			"max_rps_cip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_rps_cip_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_rps_unknown_cip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_rps_unknown_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_rps_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"pass_through_x_accel_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_accept_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30000",
				ValidateFunc: validateInteger,
			},
			"reset_conn_http_on_ssl_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"respond_with_100_continue": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"secure_cookie_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"server_side_redirect_to_https": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"true_client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTrueClientIPConfigSchema(),
			},
			"use_app_keepalive_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_true_client_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"websockets_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"x_forwarded_proto_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"xff_alternate_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "X-Forwarded-For",
			},
			"xff_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"xff_update": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "REPLACE_XFF_HEADERS",
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
			"hdr_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceHTTPHdrValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_sensitive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
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
			"file_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceHTTPRedirectActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"keep_query": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceHTTPRequestRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"enable": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"hdr_action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"log": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceHTTPResponseRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"enable": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"hdr_action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"loc_hdr_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHTTPRewriteLocHdrActionSchema(),
			},
			"log": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"path": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceURIParamSchema(),
			},
			"port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceHTTPSecurityActionRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterActionSchema(),
			},
			"per_client_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"per_uri_path": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"rate_limiter": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRateLimiterSchema(),
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"log": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceHTTPServerReselectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"num_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"retry_nonidempotent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"retry_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceHTTPStatusRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"begin": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"end": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceHeaderInfoInURISchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"header_field_name": {
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

func ResourceHealthMonitorAuthInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
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

func ResourceHealthMonitorFtpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FTP_PASSIVE_MODE",
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

func ResourceHealthMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exact_http_request": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"http_request": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET / HTTP/1.0",
			},
			"http_request_body": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"response_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceHealthMonitorImapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"folder": {
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

func ResourceHealthMonitorLdapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": {
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

func ResourceHealthMonitorPop3Schema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssl_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceHealthMonitorSSLAttributesSchema(),
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

func ResourceHealthMonitorSctpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sctp_request": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHealthMonitorSmtpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domainname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mail_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recipients_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sender_id": {
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

func ResourceHealthMonitorTcpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_half_open": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceHealthScoreDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"previous_value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resources_penalty": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"resources_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_penalty": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"security_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"step": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceHorizonProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"blast_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8443",
				ValidateFunc: validateInteger,
			},
			"pcoip_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4172",
				ValidateFunc: validateInteger,
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

func ResourceHttpCacheConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"age_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"aggressive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"date_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"default_expire": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "600",
				ValidateFunc: validateInteger,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"heuristic_expire": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ignore_request_cache_control": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"max_cache_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_object_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4194304",
				ValidateFunc: validateInteger,
			},
			"mime_types_block_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_block_lists": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"query_cacheable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"uri_non_cacheable": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePathMatchSchema(),
			},
			"xcache_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourceHttpCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_send_cookie": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"http_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_persistent_cookie": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_nics": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceIPAddrLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address_group_per_match_criteria": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ip_address_prefix_per_match_criteria": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ip_address_range_per_match_criteria": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ip_addresses_per_match_criteria": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceIPPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_mask": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ip_persistent_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceIPReputationServiceStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_successful_update_check": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
		},
	}
}

func ResourceIPReputationTypeMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bot_identification": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceBotIdentificationSchema(),
			},
			"ip_reputation_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceIPReputationTypeMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_operation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reputation_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIPThreatDBEventDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIcapLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIcapRequestLogSchema(),
			},
		},
	}
}

func ResourceIcapNSXDefenderLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"score": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"status_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIcapNsxDefenderConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"status_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://user.lastline.com/portal#/analyst/task/$uuid/overview",
			},
		},
	}
}

func ResourceIcapOPSWATLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threat_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIcapViolationSchema(),
			},
		},
	}
}

func ResourceIcapRequestLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"complete_body_sent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"http_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"icap_absolute_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_headers_received_from_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_headers_sent_to_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"icap_server_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"icap_server_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"modified_content_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"nsx_defender_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIcapNSXDefenderLogSchema(),
			},
			"opswat_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIcapOPSWATLogSchema(),
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
			"source_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIcapViolationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolution": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threat_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceImageCloudDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_data_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceImageCloudSpecificDataSchema(),
			},
			"cloud_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceImageCloudSpecificDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
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

func ResourceImageEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceImageEventMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nodes_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceImageEventSchema(),
			},
			"sub_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceImageEventSchema(),
			},
			"task_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceImageUploadOpsStatusSchema() *schema.Resource {
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

func ResourceInternalGatewayMonitorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"gateway_monitor_failure_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"gateway_monitor_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"gateway_monitor_success_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "15",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceInventoryConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_search_info": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceIpReputationConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_reputation_file_object_expiry_duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3",
				ValidateFunc: validateInteger,
			},
			"ip_reputation_sync_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_standard_alb": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"virtual_network_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"usable_alloc_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomIpamSubnetSchema(),
			},
			"usable_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIpamDnsGCPProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_se_group_subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vpc_network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"usable_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpamUsableNetworkSchema(),
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

func ResourceIpamUsableNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceKeyValueTupleSchema(),
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceJWSKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alg": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HS256",
			},
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"kid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kty": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "oct",
			},
		},
	}
}

func ResourceJWTClaimMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bool_match": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"int_match": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"is_mandatory": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"string_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceStringMatchSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"validate": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceJWTMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceJWTClaimMatchSchema(),
			},
			"token_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceJWTValidationParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"audience": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceJWTValidationVsConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"audience": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jwt_location": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "JWT_LOCATION_AUTHORIZATION_HEADER",
			},
			"jwt_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceJwtLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authn_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthnRuleMatchSchema(),
			},
			"authz_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthzRuleMatchSchema(),
			},
			"is_jwt_verified": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"token_payload": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceKeyValueTupleSchema() *schema.Resource {
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

func ResourceKniPortRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"range": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourcePortRangeSchema(),
			},
		},
	}
}

func ResourceL1FMandatoryTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mandatory_message": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceL2FMandatoryTestCaseSchema(),
			},
			"mandatory_messages": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceL2FMandatoryTestCaseSchema(),
			},
			"mandatory_string": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mandatory_strings": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceL1FSensitiveTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sensitive_message": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL2FSensitiveTestCaseSchema(),
			},
			"sensitive_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceL2FSensitiveTestCaseSchema(),
			},
			"sensitive_string": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceL1StringLengthTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"string_length_message": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL2StringLengthTestCaseSchema(),
			},
			"string_length_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceL2StringLengthTestCaseSchema(),
			},
			"test_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_strings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceL2FMandatoryTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mandatory_message": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceSingleOptionalFieldMessageSchema(),
			},
			"mandatory_messages": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceSingleOptionalFieldMessageSchema(),
			},
			"mandatory_string": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mandatory_strings": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceL2FSensitiveTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sensitive_message": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSingleOptionalSensitiveFieldMessageSchema(),
			},
			"sensitive_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSingleOptionalSensitiveFieldMessageSchema(),
			},
			"sensitive_string": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceL2StringLengthTestCaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"string_length_message": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSingleOptionalStringFieldSchema(),
			},
			"string_length_messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSingleOptionalStringFieldSchema(),
			},
			"test_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_strings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceL4PoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"l4_policy_set_ref": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceL4SSLApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssl_stream_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceL7limitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_policies_per_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_compression_filters": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_custom_str": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_matches_per_rule": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_rules_per_evh_host": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_rules_per_http_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_strgroups_per_match": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_cache_mime": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_groups_cache_mime": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_groups_no_cache_mime": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_groups_no_cache_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_no_cache_mime": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"str_no_cache_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceLDAPVSConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"realm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_auth_ldap_bind_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_cache_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_connect_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_conns_per_server": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_reconnect_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_request_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_servers_failover_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceLatencyAuditPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_est_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_est_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"latency_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"latency_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "389",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceLearningLogPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceLicenseExpiryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"burst_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cpu_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_ses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sockets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"throughput": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceLicenseInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"service_cores": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier": {
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

func ResourceLicenseServiceUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_units": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOrgServiceUnitsSchema(),
			},
		},
	}
}

func ResourceLicenseTierSwitchDetiailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_tier": {
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

func ResourceLicenseTierUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usage": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLicenseUsageSchema(),
			},
		},
	}
}

func ResourceLicenseTransactionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"licensed_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"operation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overdraft": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceLicenseUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"consumed": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"escrow": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"remaining": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"motd": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_log_disk_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_log_disk_size_gb": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"se_sys_disk_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_sys_disk_size_gb": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceLogAgentEventDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tcp_detail": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceLogAgentTCPClientEventDetailSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLogAgentTCPClientEventDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceLogManagerDebugFilterSchema() *schema.Resource {
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

func ResourceManagementNetworkConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"overlay_segment": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTier1LogicalRouterInfoSchema(),
			},
			"transport_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tz_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_segment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vs_name_tag_framework": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceMatchReplacePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bot_detection_result": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceBotDetectionMatchSchema(),
			},
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
			"geo_matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGeoMatchSchema(),
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
			"ip_reputation_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIPReputationTypeMatchSchema(),
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
			"source_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"tls_fingerprint_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTlsFingerprintMatchSchema(),
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

func ResourceMatchesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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

func ResourceMemberInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceMemoryBalancerInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"child": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceChildProcessInfoSchema(),
			},
			"controller_memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"controller_memory_usage_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"debug_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"memory_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"pid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"process": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_trend": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceMemoryUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"total": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceMesosConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_vses_are_feproxy": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"app_sync_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"container_port_match_http_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_frontend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_gs_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_se_creation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"feproxy_bridge_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cbr1",
			},
			"feproxy_container_port_as_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"feproxy_route_publish": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFeProxyRoutePublishConfigSchema(),
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "25",
				ValidateFunc: validateInteger,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi/se",
			},
			"services_accessible_all_interfaces": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_bridge_ip_as_vip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_container_ip_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_controller_image": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_vips_for_east_west_services": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateFloat,
			},
			"memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceMetricLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_timestamp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"metric_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"report_timestamp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"step": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"time_series": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceMetricsQueryResponseSchema(),
			},
			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceMetricStatisticsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_sample": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_ts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mean": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"min": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"min_ts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_samples": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"sum": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"trend": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceMetricThresoldUpDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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

func ResourceMetricsDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_response_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"blocking_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"browser_rendering_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"client_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"connection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"dns_lookup_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"dom_content_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"is_null": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"num_samples": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"page_download_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"page_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"prediction_interval_high": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"prediction_interval_low": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"redirection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"rum_client_data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"server_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"service_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"metrics_sum_agg_invalid": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourceMetricsDbDiskEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_deleted_tables": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metrics_free_sz": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"metrics_quota": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceMetricsDbQueueFullEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"instanceport": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"low": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceMetricsDbQueueHealthyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"instanceport": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"low": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceMetricsDbRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"db_client_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_num_client_queries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_num_client_resp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_num_db_queries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_num_db_resp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_num_oom": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_queue_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_rum_queries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"db_rum_rows": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceMetricsDbSyncFailureEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"include_derivation_metrics": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"second_order_derivation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_backend_derivation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceMetricsEventThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_event_threshold_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reset_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"watermark_thresholds": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceMetricsGRPCAuthFailureDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"peer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceMetricsMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"debug_skip_eval_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"stop": {
				Type:     schema.TypeString,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"enabled": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"task_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceMustChecksInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"check_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"details": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_details": {
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceNTPServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_number": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceNetworkProfileUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sctp_fast_path_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSCTPFastPathProfileSchema(),
			},
			"sctp_proxy_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSCTPProxyProfileSchema(),
			},
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

func ResourceNetworkRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_uuids": {
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

func ResourceNetworkSecurityMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
			},
			"client_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourcePortMatchGenericSchema(),
			},
			"geo_matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGeoMatchSchema(),
			},
			"ip_reputation_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIPReputationTypeMatchSchema(),
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

func ResourceNetworkSecurityPolicyActionRLParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_size": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"max_rate": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"log": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceNetworkSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"used": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceNsxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceNsxtClustersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceNsxtConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"automate_dfw_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"data_network_config": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceDataNetworkConfigSchema(),
			},
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"enforcementpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"management_network_config": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceManagementNetworkConfigSchema(),
			},
			"nsxt_credentials_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nsxt_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
	}
}

func ResourceNsxtCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtDatastoresSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ds_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceNsxtHostsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceNsxtImageDetailsSchema() *schema.Resource {
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
			"image_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vc_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtSIEndpointDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"targetips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tier1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtSIPolicyDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectto": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtSIRuleDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destexclude": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"dests": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceNsxtSIServiceDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtSetupSchema() *schema.Resource {
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transportzone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtT1SegSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"segment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceNsxtT1SegDetailsSchema() *schema.Resource {
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
			"t1seg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNsxtT1SegSchema(),
			},
		},
	}
}

func ResourceNsxtTier1SegmentAutomaticModeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nsxt_segment_subnet": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"num_se_per_segment": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16",
				ValidateFunc: validateInteger,
			},
			"tier1_lr_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceNsxtTier1SegmentConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"automatic": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtTier1SegmentAutomaticModeSchema(),
			},
			"manual": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtTier1SegmentManualModeSchema(),
			},
			"segment_config_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TIER1_SEGMENT_MANUAL",
			},
		},
	}
}

func ResourceNsxtTier1SegmentManualModeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tier1_lrs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceTier1LogicalRouterInfoSchema(),
			},
		},
	}
}

func ResourceNtlmLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ntlm_detected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ntlm_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8443",
				ValidateFunc: validateInteger,
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

func ResourceOAuthAppSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"oidc_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOIDCConfigSchema(),
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceOAuthProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"end_session_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"introspection_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jwks_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"jwks_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oauth_resp_buffer_sz": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "102400",
				ValidateFunc: validateInteger,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"userinfo_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOAuthResourceServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ACCESS_TOKEN_TYPE_JWT",
			},
			"introspection_data_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"jwt_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceJWTValidationParamsSchema(),
			},
			"opaque_token_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOpaqueTokenValidationParamsSchema(),
			},
		},
	}
}

func ResourceOAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOAuthAppSettingsSchema(),
			},
			"auth_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_server": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOAuthResourceServerSchema(),
			},
		},
	}
}

func ResourceOAuthVSConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"logout_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oauth_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOAuthSettingsSchema(),
			},
			"post_logout_redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceOCSPConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failed_ocsp_jobs_retry_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600",
				ValidateFunc: validateInteger,
			},
			"max_tries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"ocsp_req_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "86400",
				ValidateFunc: validateInteger,
			},
			"ocsp_resp_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"responder_url_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"url_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OCSP_RESPONDER_URL_FAILOVER",
			},
		},
	}
}

func ResourceOCSPResponseInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cert_status": {
				Type:     schema.TypeString,
				Required: true,
			},
			"next_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_resp_from_responder_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ocsp_response": {
				Type:     schema.TypeString,
				Required: true,
			},
			"revocation_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revocation_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"this_update": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOIDCConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oidc_enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"profile": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"userinfo": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceOShiftK8SConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"auto_assign_fqdn": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"default_service_as_east_west_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"disable_auto_backend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_frontend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_gs_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_se_creation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_route_ingress_hardening": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"override_service_ports": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"se_restart_force": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_volume": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/avi",
			},
			"secure_egress_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_not_ready_addresses": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"use_controller_image": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_resource_definition_as_ssot": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_scheduling_disabled_nodes": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_service_cluster_ip_as_ew_vip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceOauthLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authn_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthnRuleMatchSchema(),
			},
			"authz_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceAuthzRuleMatchSchema(),
			},
			"is_session_cookie_expired": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"jwks_subrequest": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthSubRequestLogSchema(),
			},
			"oauth_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"token_exchange_subrequest": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthSubRequestLogSchema(),
			},
			"token_introspection_subrequest": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthSubRequestLogSchema(),
			},
			"token_refresh_subrequest": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthSubRequestLogSchema(),
			},
			"userinfo_subrequest": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOauthSubRequestLogSchema(),
			},
		},
	}
}

func ResourceOauthSubRequestLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_request_log": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSubRequestLogSchema(),
			},
		},
	}
}

func ResourceObjSyncConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"objsync_cpu_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"objsync_hub_elect_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"objsync_reconcile_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceOmittedWafLogStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_elements": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceOpaqueTokenValidationParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_secret": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceOpenStackApiVersionCheckFailureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cc_name": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"anti_affinity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"auth_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_drive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"contrail_disable_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"contrail_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contrail_plugin": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"custom_se_image_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePropertySchema(),
			},
			"enable_os_object_caching": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_tagging": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"external_networks": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"free_floatingips": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"insecure": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"keystone_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"map_admin_to_cloudadmin": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"neutron_rbac": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"tenant_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"use_admin_url": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"use_internal_endpoints": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_keystone_auth": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vip_port_in_admin_tenant": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceOperationsConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inventory_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceInventoryConfigSchema(),
			},
		},
	}
}

func ResourceOpsHistorySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ops": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_upgrade_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeUpgradeEventsSchema(),
			},
			"seg_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupStatusSchema(),
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeOpsStateSchema(),
			},
			"statediff_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventMapSchema(),
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceOrgServiceUnitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available_service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"used_service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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

func ResourceOutOfBandRequestLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ds_req_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDSRequestLogSchema(),
			},
		},
	}
}

func ResourceOverallInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"free_percent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mount_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"used": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourcePGDeploymentRuleResultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"result": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"rule": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourcePGDeploymentRuleSchema(),
			},
		},
	}
}

func ResourcePaaLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_hit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"client_request_body_sent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"request_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePaaRequestLogSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"servers_tried": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"uri_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceParamInURISchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"param_name": {
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

func ResourceParamInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"param_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"param_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"param_size_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceParamSizeClassSchema(),
			},
			"param_type_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceParamTypeClassSchema(),
			},
		},
	}
}

func ResourceParamSizeClassSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"len": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceParamTypeClassSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceParamsInURISchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"param_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceParamInURISchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"size": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourcePatchInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"patch_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reboot": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"reboot_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRebootDataSchema(),
			},
		},
	}
}

func ResourcePatchSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"match_decoded_string": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourcePerformanceLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_concurrent_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_throughput": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourcePermissionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subresource": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSubResourceSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourcePlacementScopeConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clusters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClusterHAConfigSchema(),
			},
			"nsxt_clusters": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtClustersSchema(),
			},
			"nsxt_datastores": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtDatastoresSchema(),
			},
			"nsxt_hosts": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceNsxtHostsSchema(),
			},
			"vcenter_folder": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_ref": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"value": {
				Type:     schema.TypeString,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ratio": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourcePortMatchGenericSchema() *schema.Resource {
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
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
		},
	}
}

func ResourcePortRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"start": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourcePortalConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_basic_authentication": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"api_force_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "24",
				ValidateFunc: validateInteger,
			},
			"disable_remote_cli_shell": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_swagger": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_clickjacking_protection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_http": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_https": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"http_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"https_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"minimum_password_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8",
				ValidateFunc: validateInteger,
			},
			"password_strength_check": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"redirect_to_https": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourcePortalFeatureOptInSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_appsignature_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_ip_reputation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_pulse_case_management": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_pulse_waf_management": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_user_agent_db_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceProactiveSupportDefaultsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attach_core_dump": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"attach_tech_support": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"case_severity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Severity 5",
			},
		},
	}
}

func ResourceProcessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_process_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"current_process_mem_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"intimation_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"memory_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"memory_trend_usage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"match_decoded_string": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourceRancherConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_sync_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"container_port_match_http_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"coredump_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_frontend_service_sync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"disable_auto_se_creation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"feproxy_container_port_as_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "25",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_container_ip_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"use_controller_image": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRateLimiterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_sz": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000000000",
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
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

func ResourceRateLimiterPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"msf_num_stages": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"msf_stage_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16384",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"fine_grain": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceRebootDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"patch_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reboot": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRemoteAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_mapping_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_profile_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Default:  "LITERAL_STRING",
			},
			"val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceReplicationPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checkpoint_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "REPLICATION_MODE_CONTINUOUS",
			},
		},
	}
}

func ResourceReportDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_ref": {
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
			"system_readiness": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeReadinessCheckObjSchema(),
			},
		},
	}
}

func ResourceReportEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"start_time": {
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

func ResourceReportOpsStateSchema() *schema.Resource {
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

func ResourceReportSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"previews": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"title": {
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
			"source_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrMatchSchema(),
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

func ResourceResumeSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceRmAddVnicSchema() *schema.Resource {
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
			"subnet": {
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
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"primary": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"standby": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"flavor_name": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceRoleFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"match_label": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"match_operation": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ROLE_FILTER_EQUALS",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRoleFilterMatchLabelSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceRoleMatchOperationMatchLabelSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_label": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"match_operation": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ROLE_FILTER_EQUALS",
			},
		},
	}
}

func ResourceRollbackControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRollbackPatchControllerParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRollbackPatchSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRollbackPatchSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRollbackSeGroupParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRollbackSystemParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceRoutingServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_backend_networks": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_auto_gateway": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_routing": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_vip_on_all_interfaces": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"enable_vmac": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"floating_intf_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_intf_ip6_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_intf_ip6_se_2_addresses": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"nat_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_by_linux_ipstack": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceRspContentRewriteRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pairs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSearchReplacePairSchema(),
			},
		},
	}
}

func ResourceRuleInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchesSchema(),
			},
			"rule_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSAMLSPConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"acs_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"authn_req_acs_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SAML_AUTHN_REQ_ACS_TYPE_NONE",
			},
			"cookie_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
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
			"sp_metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_idp_session_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceSCPoolServerStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_server": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSCTPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_init_chunk_protection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSCTPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie_expiration_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"heartbeat_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"max_retransmissions_association": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"max_retransmissions_init_chunks": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8",
				ValidateFunc: validateInteger,
			},
			"number_of_streams": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"receive_window": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "256",
				ValidateFunc: validateInteger,
			},
			"reset_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceSEBandwidthLimitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSETimeTrackerPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"egress_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_TT_AUDIT_OFF",
			},
			"egress_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"event_gen_window": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"ingress_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_TT_AUDIT_OFF",
			},
			"ingress_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSSHSeDeploymentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"days_until_expire": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "365",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceSSLClientCertificateActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"close_connection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLClientRequestHeaderSchema(),
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

func ResourceSSLExpireDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days_left": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
		},
	}
}

func ResourceSSLIgnoredDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
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

func ResourceSSLKeyRSAParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exponent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "65537",
				ValidateFunc: validateInteger,
			},
			"key_size": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_2048_BITS",
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

func ResourceSSLRevokedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
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

func ResourceSaasLicensingInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"reserve_service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceSaasLicensingStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"expired": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserve_service_units": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceSamlAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attr_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attr_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSamlAuthnRuleMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"saml_authn_matched_rule_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_authn_matched_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSamlAuthzRuleMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"saml_authz_matched_rule_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_authz_matched_rule_name": {
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

func ResourceSamlLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_saml_authentication_used": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"saml_attribute_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSamlAttributeSchema(),
			},
			"saml_auth_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_authn_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSamlAuthnRuleMatchSchema(),
			},
			"saml_authz_rule_match": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSamlAuthzRuleMatchSchema(),
			},
			"saml_session_cookie_expired": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"userid": {
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

func ResourceScaleStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_success": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"end_time_str": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_se_assigned": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se_requested": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"placement_read_fail_cnt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"reason": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceScheduledScalingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"cron_expression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"desired_capacity": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"end_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_max_step": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"start_date": {
				Type:     schema.TypeString,
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

func ResourceSeAgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_echo_miss_aggressive_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"controller_echo_miss_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"controller_echo_rpc_aggressive_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"controller_echo_rpc_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"controller_heartbeat_miss_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "6",
				ValidateFunc: validateInteger,
			},
			"controller_heartbeat_timeout_sec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "12",
				ValidateFunc: validateInteger,
			},
			"controller_registration_timeout_sec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30",
				ValidateFunc: validateInteger,
			},
			"controller_rpc_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"cpustats_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"ctrl_reg_pending_max_wait_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "150",
				ValidateFunc: validateInteger,
			},
			"debug_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"dp_aggressive_deq_interval_msec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"dp_aggressive_enq_interval_msec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"dp_batch_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"dp_deq_interval_msec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"dp_enq_interval_msec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"dp_max_wait_rsp_time_sec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"dp_reg_pending_max_wait_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "75",
				ValidateFunc: validateInteger,
			},
			"headless_timeout_sec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"ignore_docker_mac_change": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ns_helper_deq_interval_msec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"sdb_flush_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sdb_pipeline_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sdb_scan_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"se_grp_change_disruptive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"seagent_statecache_properties": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeAgentStateCachePropertiesSchema(),
			},
			"send_se_ready_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"states_flush_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"vnic_dhcp_ip_check_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "6",
				ValidateFunc: validateInteger,
			},
			"vnic_dhcp_ip_max_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"vnic_ip_delete_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"vnic_probe_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"vnic_rpc_retry_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"vnicdb_cmd_history_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "256",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeAgentStateCachePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sc_batch_buffer_flush_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"sc_cfg_q_batch_dequeue_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sc_cfg_q_max_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"sc_dns_q_batch_dequeue_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sc_dns_q_max_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"sc_shard_cleanup_max_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "900",
				ValidateFunc: validateInteger,
			},
			"sc_state_ring_batch_dequeue_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
			"sc_states_flush_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sc_stream_check_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5000",
				ValidateFunc: validateInteger,
			},
			"sc_thread_q_batch_dequeue_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"sc_thread_q_max_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"sc_thread_sleep_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeBgpPeerDownDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_ip": {
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

func ResourceSeBootupCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buf_num": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"buf_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"hash_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16384",
				ValidateFunc: validateInteger,
			},
			"level_aggressive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"level_normal": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"window_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeBootupPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_backend_portend": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30720",
				ValidateFunc: validateInteger,
			},
			"docker_backend_portstart": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20480",
				ValidateFunc: validateInteger,
			},
			"fair_queueing_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"geo_db_granularity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"l7_conns_per_core": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16384",
				ValidateFunc: validateInteger,
			},
			"l7_resvd_listen_conns_per_core": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "256",
				ValidateFunc: validateInteger,
			},
			"log_agent_debug_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"log_agent_trace_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"se_dp_compression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeBootupCompressionPropertiesSchema(),
			},
			"se_emulated_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_ip_encap_ipc": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_l3_encap_ipc": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_log_buffer_app_blocking_dequeue": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_log_buffer_applog_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"se_log_buffer_chunk_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1024",
				ValidateFunc: validateInteger,
			},
			"se_log_buffer_conn_blocking_dequeue": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_log_buffer_connlog_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1024",
				ValidateFunc: validateInteger,
			},
			"se_log_buffer_events_blocking_dequeue": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"se_log_buffer_events_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "512",
				ValidateFunc: validateInteger,
			},
			"ssl_sess_cache_per_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4096",
				ValidateFunc: validateInteger,
			},
			"tcp_syncache_hashsize": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8192",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeDebugModeEventDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
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

func ResourceSeDiscontinuousTimeChangeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drift_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"from_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_servers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
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

func ResourceSeFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arg": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"num_skips": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceSeGroupAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_event_thresholds": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsEventThresholdSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_suspended": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"notes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se_with_no_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se_with_vs_not_scaledout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se_with_vs_scaledout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vs_disrupted": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeGroupVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fips_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
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

func ResourceSeHBEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hb_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"remote_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reporting_se_ref": {
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

func ResourceSeHbRecoveredEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hb_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"remote_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reporting_se_ref": {
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

func ResourceSeHighEgressProcLatencyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dispatcher_core": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"event_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"flow_core": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"max_proxy_to_disp_queing_delay": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_name": {
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

func ResourceSeHighIngressProcLatencyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dispatcher_core": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"event_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"flow_core": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_disp_to_proxy_queing_delay": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_dispatcher_proc_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeHmEventShmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average_response_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeLicensedBandwdithExceededEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_pkts_dropped": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_on_cloud": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"active_on_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"admin_down_requested": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"attach_ip_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"cloud_programming_done": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"cloud_programming_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"detach_ip_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"floating_intf_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"floating_intf_ip6_addresses": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_primary": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"is_standby": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2001",
				ValidateFunc: validateInteger,
			},
			"mgmt_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mgmt_ip6": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"scaleout_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_programming_done": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"se_ready_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sec_idx": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"snat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"vip6_subnet_mask": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
			},
			"vlan_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"vnic": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"heap_config_soft_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"heap_config_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"heap_conn_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shm_config_hard_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"shm_config_soft_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"shm_config_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"shm_conn_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"new_mgmt_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_mgmt_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"old_mgmt_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"old_mgmt_ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"vsphere_ha_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"vsphere_ha_inprogress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceSeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeNtpSynchronizationFailedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ntp_servers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_name": {
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

func ResourceSePersistenceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeRateLimitersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arp_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"default_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"flow_probe_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "250",
				ValidateFunc: validateInteger,
			},
			"icmp_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"icmp_rsp_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"rst_rl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeReconcileDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"new_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"old_service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores_per_socket": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"disk": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"hyper_threading": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"hypervisor_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"memory": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"num_datapath_processes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vcpus": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"sockets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeRuntimeCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_low_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"min_high_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "200",
				ValidateFunc: validateInteger,
			},
			"min_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"mobile_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSeRuntimePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_ssh_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"app_headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppHdrSchema(),
			},
			"baremetal_dispatcher_handles_flows": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"connections_lossy_log_rate_limiter_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"connections_udfnf_log_rate_limiter_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"disable_flow_probes": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"dos_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"downstream_send_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600000",
				ValidateFunc: validateInteger,
			},
			"dp_aggressive_hb_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"dp_aggressive_hb_timeout_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"dp_hb_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"dp_hb_timeout_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"dupip_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"dupip_timeout_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"enable_hsm_log": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"feproxy_vips_enable_proxy_arp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"flow_table_batch_push_frequency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"global_mtu": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"http_rum_console_log": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"http_rum_min_content_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "64",
				ValidateFunc: validateInteger,
			},
			"lbaction_num_requests_to_dispatch": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"lbaction_rq_per_request_max_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "22",
				ValidateFunc: validateInteger,
			},
			"log_agent_compress_logs": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"log_agent_file_sz_appl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"log_agent_file_sz_conn": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"log_agent_file_sz_debug": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"log_agent_file_sz_event": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"log_agent_log_storage_min_sz": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1024",
				ValidateFunc: validateInteger,
			},
			"log_agent_max_concurrent_rsync": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1024",
				ValidateFunc: validateInteger,
			},
			"log_agent_max_storage_excess_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "110",
				ValidateFunc: validateInteger,
			},
			"log_agent_max_storage_ignore_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateFloat,
			},
			"log_agent_min_storage_per_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"log_agent_sleep_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"log_agent_unknown_vs_timer": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1800",
				ValidateFunc: validateInteger,
			},
			"log_message_max_file_list_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "64",
				ValidateFunc: validateInteger,
			},
			"ngx_free_connection_stack": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"persistence_mem_max": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"scaleout_udp_per_pkt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"se_auth_ldap_bind_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_cache_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_connect_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_conns_per_server": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_reconnect_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_request_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
			},
			"se_auth_ldap_servers_failover_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_dp_compression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeRuntimeCompressionPropertiesSchema(),
			},
			"se_dp_hm_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_dp_if_state_poll_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"se_dp_log_nf_enqueue_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "70",
				ValidateFunc: validateInteger,
			},
			"se_dp_log_udf_enqueue_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "90",
				ValidateFunc: validateInteger,
			},
			"se_dump_core_on_assert": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_handle_interface_routes": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_hb_persist_fudge_bits": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3",
				ValidateFunc: validateInteger,
			},
			"se_mac_error_threshold_to_disable_promiscious": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"se_memory_poison": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"se_metrics_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60000",
				ValidateFunc: validateInteger,
			},
			"se_metrics_rt_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"se_metrics_rt_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1000",
				ValidateFunc: validateInteger,
			},
			"se_packet_buffer_max": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"se_random_tcp_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"tcp_syncache_max_retransmit_default": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
			"upstream_connect_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600000",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_cache_thresh": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "-1",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_conn_idle_thresh_tmo": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "-1",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_core_max_cache": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "-1",
				ValidateFunc: validateInteger,
			},
			"upstream_connpool_enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"upstream_keepalive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"upstream_read_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600000",
				ValidateFunc: validateInteger,
			},
			"upstream_send_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600000",
				ValidateFunc: validateInteger,
			},
			"user_defined_metric_age": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSeThreshEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"thresh": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSeUpgradeEventsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se_group": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"sub_tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceSeUpgradeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disruptive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"patch": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"patch_rollback": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"resume_from_suspend": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"rollback": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skip_suspended": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"suspend_on_failure": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceSeVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fips_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
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

func ResourceSeVipInterfaceListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_portchannel": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceSeVsDelFlowsDisruptedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deleted_vs_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_vs_flows_disrupted": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"reporting_se_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceSeVsPktBufHighEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"virtual_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSearchReplacePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"replacement_string": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceReplaceStringVarSchema(),
			},
			"search_string": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceSearchStringVarSchema(),
			},
		},
	}
}

func ResourceSearchStringVarSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SEARCH_LITERAL_STRING",
			},
			"val": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSecMgrDataEventSchema() *schema.Resource {
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

func ResourceSecMgrThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSecMgrUAEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"free_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"uuid": {
				Type:     schema.TypeString,
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

func ResourceSecureChannelMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_controller": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"local_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marked_for_delete": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceSecureChannelTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expiry_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"in_use": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSecureKeyExchangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ctlr_mgmt_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ctlr_public_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"follower_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"leader_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
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

func ResourceSecurityMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_adaptive_config": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"entity_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceKeyValueTupleSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceSensitiveLogProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"header_field_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSensitiveFieldRuleSchema(),
			},
			"uri_query_field_rules": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"preference_order": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"prst_hdr_val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratio": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"resolve_server_by_dns": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"rewrite_host_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"server_node": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"verify_network": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vm_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServerAutoScaleFailedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_scalein_servers": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"num_servers_up": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceServerAutoScaleInCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nscalein": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"load": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"num_scalein_servers": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"num_servers_up": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
			"scheduled_desired_capacity": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"load": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"num_scaleout_servers": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"num_servers_up": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
			"scheduled_desired_capacity": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceServerConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"def_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"propogate_state": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"timer_exists": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_http2": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_ssl": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"horizon_internal_ports": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_active_ftp_data_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"port_range_end": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceServiceEngineCloudLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrfs_per_serviceengine": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceServiceEngineLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_virtualservices_per_serviceengine": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ew_virtualservices_per_serviceengine": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ns_virtualservices_per_serviceengine": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_logical_intf_per_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_phy_intf_per_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_virtualservices_rt_metrics": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vlan_intf_per_phy_intf": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vlan_intf_per_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"serviceengine_cloud_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServiceEngineCloudLimitsSchema(),
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
			"protocol": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceL4RuleProtocolMatchSchema(),
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"service_port_range_end": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"service_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceServiceengineFaultsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"debug_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1024",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSingleLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addons": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"burst_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"capacity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"ccu": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"cpu_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"service_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateFloat,
			},
			"sockets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"unit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SERVICE_UNIT",
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

func ResourceSingleOptionalFieldMessageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"optional_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSingleOptionalSensitiveFieldMessageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"optional_sensitive_string": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
		},
	}
}

func ResourceSingleOptionalStringFieldSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"test_string": {
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

func ResourceSipMessageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"from_client": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rcv_timestamp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"rx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"tx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceSipServiceApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"transaction_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "162",
				ValidateFunc: validateInteger,
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

func ResourceSpGslbServiceInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fqdns": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceStatediffEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FB_INIT",
			},
			"task_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceStaticIpAllocInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"obj_info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceStaticIpRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"range": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "STATIC_IPS_FOR_VIP_AND_SE",
			},
		},
	}
}

func ResourceStaticIpRangeRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allocated_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticIpAllocInfoSchema(),
			},
			"free_ip_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"total_ip_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "STATIC_IPS_FOR_VIP_AND_SE",
			},
		},
	}
}

func ResourceStaticRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceStreamingSyslogConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"facility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "16",
				ValidateFunc: validateInteger,
			},
			"filtered_log_severity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AviVantage",
			},
			"msg_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NILVALUE",
			},
			"non_significant_log_severity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "6",
				ValidateFunc: validateInteger,
			},
			"proc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NILVALUE",
			},
			"significant_log_severity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
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
			"num_retries": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSubRequestLogSchema() *schema.Resource {
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
			"http_response_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"response_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"source_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"total_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
		},
	}
}

func ResourceSubResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exclude_subresources": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"subresources": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"static_ip_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticIpRangeSchema(),
			},
		},
	}
}

func ResourceSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_range_runtimes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticIpRangeRuntimeSchema(),
			},
			"prefix": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     ResourceIpAddrPrefixSchema(),
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

func ResourceSupportedMigrationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_host_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"controller_min_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8",
				ValidateFunc: validateInteger,
			},
			"controller_min_docker_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1.6.1",
			},
			"controller_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"controller_min_memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "24",
				ValidateFunc: validateInteger,
			},
			"controller_min_total_disk": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
			},
			"max_active_versions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"min_supported_api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"podman_controller_host_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "24",
				ValidateFunc: validateInteger,
			},
			"podman_se_host_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "12",
				ValidateFunc: validateInteger,
			},
			"rollback_controller_disk_space": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"rollback_se_disk_space": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"se_host_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"se_min_cores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"se_min_free_disk_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"se_min_memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"se_min_total_disk": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"versions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"reason": {
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

func ResourceTCPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ftp_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFTPProfileSchema(),
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_protocol_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"session_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceTCPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aggressive_congestion_avoidance": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"auto_window_growth": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"automatic": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"cc_algo": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CC_ALGO_NEW_RENO",
			},
			"congestion_recovery_scaling_factor": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"idle_connection_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "600",
				ValidateFunc: validateInteger,
			},
			"idle_connection_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KEEP_ALIVE",
			},
			"ignore_time_wait": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ip_dscp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"keepalive_in_halfclose_state": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"max_retransmissions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8",
				ValidateFunc: validateInteger,
			},
			"max_segment_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_syn_retransmissions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "8",
				ValidateFunc: validateInteger,
			},
			"min_rexmt_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"nagles_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"reassembly_queue_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"receive_window": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "64",
				ValidateFunc: validateInteger,
			},
			"reorder_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"slow_start_scaling_factor": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"time_wait_delay": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2000",
				ValidateFunc: validateInteger,
			},
			"use_interface_mtu": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "49",
				ValidateFunc: validateInteger,
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

func ResourceTcpAttacksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceTenantConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_in_provider_context": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"tenant_access_to_provider_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"tenant_vrf": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceTenantLabelSchema() *schema.Resource {
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

func ResourceTier1LogicalRouterInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"locale_service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1_lr_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceTimeStampSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secs": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"usecs": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceTimeTrackerPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"be_conn_est_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"be_conn_est_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"fe_conn_est_audit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fe_conn_est_threshold": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"ingress_sig_log": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceTlsClientInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cipher_suites": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"client_hello_tls_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"point_formats": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"supported_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"tls_extensions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"uses_grease": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceTlsFingerprintMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fingerprints": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"match_operation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceTrueClientIPConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LEFT",
			},
			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"index_in_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"session_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"snat": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceUDPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"session_idle_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceURIInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"param_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceParamInfoSchema(),
			},
			"uri_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"uri_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceURIParamQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_query": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceURIParamTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"start_index": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceUdpAttacksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceUpgradeEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"sub_tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"rebooted": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceUpgradeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"se_group_options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSeGroupOptionsSchema(),
			},
			"se_group_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_patch_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_warnings": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"system": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceUpgradeReadinessCheckObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMustChecksInfoSchema(),
			},
			"checks_completed": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch_image_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceUpgradeOpsStateSchema(),
			},
			"total_checks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"upgrade_ops": {
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
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"prechecks_only": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceUserActivitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"concurrent_sessions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"failed_login_attempts": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"login_failure_timestamps": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func ResourceUserAgentCacheConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"batch_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"controller_cache_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "300000",
				ValidateFunc: validateInteger,
			},
			"max_age": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "86400",
				ValidateFunc: validateInteger,
			},
			"max_last_hit_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "86400",
				ValidateFunc: validateInteger,
			},
			"max_upstream_queries": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"max_wait_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"num_entries_upstream_update": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
			"percent_reserved_for_bad_bots": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"percent_reserved_for_browsers": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "50",
				ValidateFunc: validateInteger,
			},
			"percent_reserved_for_good_bots": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20",
				ValidateFunc: validateInteger,
			},
			"percent_reserved_for_outstanding": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10",
				ValidateFunc: validateInteger,
			},
			"se_cache_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "20000",
				ValidateFunc: validateInteger,
			},
			"upstream_update_interval": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "3600",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceUserAgentDBConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_batch_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "500",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceUserRoleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_tenants": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceVCenterCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressSensitiveFieldDiffs,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVHMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     ResourceVHMatchRuleSchema(),
			},
		},
	}
}

func ResourceVHMatchRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": {
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

func ResourceVIMgrGuestNicRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_internal_network": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"connected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"del_pending": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "connected",
			},
			"cpu_hz": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"maintenance_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_cpu_packages": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_cpu_threads": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"quarantined_periods": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"se_fail_cnt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_success_cnt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVIMgrIPSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fip_available": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"ref_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"se_ref_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"uuid": {
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
			"auto_expand": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"host_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_nw": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ip_subnet": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema(),
			},
			"logical_switch_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmtnw": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_ports": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_vmgroup": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"deletion_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"discovery_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovery_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"disk_gb": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcp_se_project_id": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"last_discovery": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_mb": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"vcenter_host_connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_host_ha_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_instance_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_ref": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"vcenter_url": {
				Type:     schema.TypeString,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vsphere_ha_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"vsphere_ha_inprogress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"cpu_reservation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"cpu_shares": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"creation_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mem_shares": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"memory_reservation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_cpu": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"vm_lb_weight": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVSDataScriptsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"vs_datascript_set_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVcenterClusterDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vc_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceVcenterDatastoreSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datastore_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceVcenterImageDetailsSchema() *schema.Resource {
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
			"image_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vc_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVcenterTagEventDetailsSchema() *schema.Resource {
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
			"vm_id": {
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

func ResourceVinfraDiscSummaryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clusters": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_dcs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_hosts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_nws": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vms": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"vcenter": {
				Type:     schema.TypeString,
				Required: true,
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

func ResourceVinfraVcenterBadCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"previous_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVinfraVcenterNetworkLimitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_reason": {
				Type:     schema.TypeString,
				Required: true,
			},
			"current": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"limit": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocate_floating_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"auto_allocate_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"avi_allocated_vip": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"discovered_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
			"prefix_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
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

func ResourceVipAutoscalePolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_cooldown": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "60",
				ValidateFunc: validateInteger,
			},
			"max_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "5",
				ValidateFunc: validateInteger,
			},
			"min_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1",
				ValidateFunc: validateInteger,
			},
			"suspend": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceVipAutoscaleZonesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_capable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"subnet_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"migrate_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"migrate_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsMigrateParamsSchema(),
			},
			"migrate_scalein_pending": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"migrate_scaleout_pending": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"num_additional_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"prev_metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"progress_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"scalein_request": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceVsScaleinParamsSchema(),
			},
			"scaleout_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warmstart_resync_done": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"warmstart_resync_sent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceVipSeAssignedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_on_cloud": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"active_on_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"admin_down_requested": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"attach_ip_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"connected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"detach_ip_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"mgmt_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"mgmt_ip6": {
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
			"oper_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"primary": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scalein_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"scaleout_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"se_ready_in_progress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"snat_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"standby": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceVipSymmetryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_num_se_assigned": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"max_num_se_requested": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"min_num_se_assigned": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"min_num_se_requested": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vsvip_name": {
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

func ResourceVirtualServiceResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_standby_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceVirtualServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"key_rotation_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"marked_for_delete": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"metrics_mgr_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_additional_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"one_plus_one_ha": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"redis_port": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"sec_mgr_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceSecurityMgrRuntimeSchema(),
			},
			"self_se_election": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceVirtualserviceFaultsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"debug_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"pool_server_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"scaleout_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"shared_vip_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ssl_cert_expiry_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"ssl_cert_status_faults": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceVlanInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"if_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6_autocfg_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"is_mgmt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"vlan_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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

func ResourceVlanRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"start": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceVsApicExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceVsAwaitingSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"awaitingse_timeout": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVsMigrateParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"to_host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_new_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVsScaleinParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"from_se_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scalein_primary": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
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

func ResourceVsScaleoutParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_up": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"new_vcpus": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"to_host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"to_new_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceVsSwitchoverEventDetailsSchema() *schema.Resource {
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceVserverL4MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexc": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"apdexrtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_application_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_bandwidth": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_bytes_policy_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_complete_conns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_connections_dropped": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_app_error": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_bad_rst_flood": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_bandwidth": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_conn": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_conn_ip_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_conn_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_fake_session": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_http_abort": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_http_error": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_http_timeout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_malformed_flood": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_non_syn_flood": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_cookie_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_custom_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_hdr_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_rl_drop_bad": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_scan_bad_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_scan_unknown_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_uri_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_ip_uri_rl_drop_bad": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_uri_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_uri_rl_drop_bad": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_uri_scan_bad_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_req_uri_scan_unknown_rl_drop": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_rx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_slow_uri": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_small_window_stress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_ssl_error": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_syn_flood": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_total_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_tx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dos_zero_window_stress": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_errored_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_half_open_conns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_l4_client_latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_lossy_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_lossy_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_network_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_new_established_conns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_pkts_policy_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_policy_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_rx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_rx_bytes_dropped": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_rx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_rx_pkts_dropped": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_syns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_total_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_total_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_tx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_tx_pkts": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_connection_estb_time_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_num_active_se": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_open_conns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_rx_bytes_absolute": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_rx_pkts_absolute": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_tx_bytes_absolute": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_tx_pkts_absolute": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_application_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_connection_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_connections_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_dos_bandwidth": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_dos_rx_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_network_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_pkts_dos_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_policy_drops": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_conn_duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_conn_est_time_exceeded_flows_fe": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_connection_dropped_user_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_connection_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_connections_dropped": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_dup_ack_retransmits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_end_to_end_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_end_to_end_rtt_bucket1": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_end_to_end_rtt_bucket2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_finished_conns": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_ingress_latency_exceeded_flows": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_lossy_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_lossy_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_out_of_orders": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_packet_dropped_user_bandwidth_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_rtt_valid_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_sack_retransmits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_server_flow_control": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_timeout_retransmits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_zero_window_size_events": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceVserverL7MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexr": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_application_response_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_blocking_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_browser_rendering_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_cache_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_cache_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_cacheable_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_cacheable_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_client_data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_client_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_client_txn_latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_complete_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_connection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dns_lookup_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_dom_content_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_error_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_errors_excluded": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_frustrated_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_http_headers_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_http_headers_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_http_params_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_page_download_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_page_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_params_per_req": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_post_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_post_compression_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_pre_compression_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_redirection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_reqs_per_session": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_1xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_2xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_3xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_4xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_4xx_avi_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_5xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_resp_5xx_avi_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_rum_client_data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_satisfactory_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_server_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_service_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_auth_dsa": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_auth_ecdsa": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_auth_rsa": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ecdsa_non_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ecdsa_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_failed_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshake_network_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshake_protocol_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshakes_new": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshakes_non_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshakes_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshakes_reused": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_handshakes_timedout": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_kx_dh": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_kx_ecdh": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_kx_rsa": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_rsa_non_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_rsa_pfs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ver_ssl30": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ver_tls10": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ver_tls11": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ver_tls12": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_ssl_ver_tls13": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_tolerated_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_total_http2_requests": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_total_requests": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_uri_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_evaluated": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_evaluated_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_evaluated_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_evaluated_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_evaluated_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_flagged": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_flagged_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_flagged_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_flagged_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_flagged_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_latency_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_latency_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_latency_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_latency_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_matched": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_matched_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_matched_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_matched_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_matched_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_rejected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_rejected_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_rejected_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_rejected_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waf_rejected_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"avg_waiting_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_concurrent_sessions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"max_ssl_open_sessions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"node_obj_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_cache_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_cacheable_hits": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_get_reqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_post_reqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_response_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_ssl_failed_connections": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_evaluated": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_flagged": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_matched": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"pct_waf_rejected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"rum_apdexr": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"ssl_protocol_strength": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_application_response_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_blocking_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_browser_rendering_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_client_data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_client_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_connection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_dns_lookup_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_dom_content_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_errors": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_finished_sessions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_get_client_txn_latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_get_client_txn_latency_bucket1": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_get_client_txn_latency_bucket2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_get_reqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_http_headers_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_http_headers_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_http_params_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_num_page_load_time_bucket1": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_num_page_load_time_bucket2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_num_rum_samples": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_other_client_txn_latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_other_client_txn_latency_bucket1": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_other_client_txn_latency_bucket2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_other_reqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_page_download_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_page_load_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_post_bytes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_post_client_txn_latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_post_client_txn_latency_bucket1": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_post_client_txn_latency_bucket2": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_post_reqs": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_redirection_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_reqs_finished_sessions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_reqs_with_params": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_resp_1xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_resp_2xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_resp_3xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_resp_4xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_resp_5xx": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_rum_client_data_transfer_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_server_rtt": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_service_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_total_responses": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_uri_length": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_attacks": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_disabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_evaluated_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_evaluated_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_evaluated_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_evaluated_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_flagged": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_flagged_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_flagged_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_flagged_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_flagged_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_latency_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_latency_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_latency_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_latency_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_matched_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_matched_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_matched_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_matched_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_rejected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_rejected_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_rejected_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_rejected_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waf_rejected_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
			"sum_waiting_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateFloat,
			},
		},
	}
}

func ResourceVssPlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_nonaffinity": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "2",
				ValidateFunc: validateInteger,
			},
			"num_subcores": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "4",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceWAFLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_allowed_content_types": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_allowed_request_content_type_charsets": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_allowlist_policy_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_applications": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_data_files": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_pre_post_crs_groups": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_psm_groups": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_psm_match_elements": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_psm_match_rules_per_loc": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_psm_total_locations": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_restricted_extensions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_restricted_headers": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_rule_tags": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_rules_per_rulegroup": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"num_static_extensions": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceWafAllowlistLogSchema() *schema.Resource {
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

func ResourceWafApplicationSignatureAppVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_changed_ruleset_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_of_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceWafApplicationSignaturesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"provider_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolved_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleSchema(),
			},
			"rule_overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleOverridesSchema(),
			},
			"ruleset_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selected_applications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"allowed_request_content_type_charsets": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "32",
				ValidateFunc: validateInteger,
			},
			"content_type_mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafContentTypeMappingSchema(),
			},
			"cookie_format_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"ignore_incomplete_request_body_error": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"max_execution_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "50",
				ValidateFunc: validateInteger,
			},
			"regex_match_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "30000",
				ValidateFunc: validateInteger,
			},
			"regex_recursion_limit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "10000",
				ValidateFunc: validateInteger,
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
			"send_status_header": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"server_response_max_body_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "128",
				ValidateFunc: validateInteger,
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
			"status_header_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "X-WAF-Result",
			},
			"xml_xxe_protection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
		},
	}
}

func ResourceWafContentTypeMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_op": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EQUALS",
			},
			"request_body_parser": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafCrsConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_auto_download_waf_signatures": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enable_waf_signatures_notifications": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_DATAFILE_PM_FROM_FILE",
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

func ResourceWafLearningSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceWafLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowlist_configured": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"allowlist_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafAllowlistLogSchema(),
			},
			"allowlist_processed": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"application_rule_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleLogSchema(),
			},
			"application_rules_configured": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"application_rules_processed": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"latency_request_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"latency_request_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"latency_response_body_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"latency_response_header_phase": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"memory_allocated": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"omitted_app_rule_stats": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOmittedWafLogStatsSchema(),
			},
			"omitted_signature_stats": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceOmittedWafLogStatsSchema(),
			},
			"psm_configured": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"psm_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMLogSchema(),
			},
			"psm_processed": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"rule_logs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleLogSchema(),
			},
			"rules_configured": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"rules_processed": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceWafPSMMatchElementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"excluded": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"match_case": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_op": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EQUALS",
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

func ResourceWafPSMRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
			},
			"match_value_pattern": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_value_string_group_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_value_string_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func ResourceWafPolicyAllowlistSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPolicyAllowlistRuleSchema(),
			},
		},
	}
}

func ResourceWafPolicyAllowlistRuleSchema() *schema.Resource {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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
			"sampling_percent": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "100",
				ValidateFunc: validateInteger,
			},
		},
	}
}

func ResourceWafPolicyRequiredDataFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
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

func ResourceWafPolicyWhitelistSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceWafPolicyWhitelistRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
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

func ResourceWafRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
			},
			"is_sensitive": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"paranoia_level": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phase": {
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

func ResourceWafRuleGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInteger,
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

func ResourceWafRuleGroupOverridesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
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
			"rule_overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleOverridesSchema(),
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
			"omitted_match_elements": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInteger,
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

func ResourceWafRuleMatchDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_internal": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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

func ResourceWafRuleOverridesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"exclude_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafWhitelistLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceWebApplicationSignatureServiceStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_successful_update_check": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"upstream_sync_timestamp": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
		},
	}
}

func ResourcepostsnapshotSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gssnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbGsInfoSchema(),
			},
			"poolsnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbPoolInfoSchema(),
			},
			"sesnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbSeInfoSchema(),
			},
			"vssnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbVsInfoSchema(),
			},
		},
	}
}

func ResourcepresnapshotSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gssnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbGsInfoSchema(),
			},
			"poolsnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbPoolInfoSchema(),
			},
			"sesnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbSeInfoSchema(),
			},
			"vssnapshot": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceFbVsInfoSchema(),
			},
		},
	}
}

func ResourcevCenterConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_lib": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     ResourceContentLibConfigSchema(),
			},
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_nsx_environment": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
			"use_content_lib": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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

func ResourcevNICSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregator_chgd": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"can_se_dp_takeover": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"connected": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBool,
			},
			"del_pending": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"delete_vnic": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"dhcp_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"dp_deletion_done": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"if_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_autocfg_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
			},
			"is_asm": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_avi_internal_network": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_hsm": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_mgmt": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"is_portchannel": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
			},
			"link_up": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "true",
				ValidateFunc: validateBool,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "1500",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0",
				ValidateFunc: validateInteger,
			},
			"vrf_ref": {
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "false",
				ValidateFunc: validateBool,
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
