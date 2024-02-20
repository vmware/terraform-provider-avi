// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ignore_servers": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"analytics_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePoolAnalyticsPolicySchema(),
		},
		"analytics_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"append_port": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "NON_DEFAULT_80_443",
		},
		"application_persistence_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"autoscale_launch_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"autoscale_networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"autoscale_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"capacity_estimation": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"capacity_estimation_ttfb_thresh": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_ref": {
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
		"conn_pool_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConnPoolPropertiesSchema(),
		},
		"connection_ramp_duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"default_server_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "80",
			ValidateFunc: validateInteger,
		},
		"delete_server_on_dns_refresh": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"domain_name": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"east_west": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"enable_http2": {
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
		"external_autoscale_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"fail_action": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFailActionSchema(),
		},
		"fewest_tasks_feedback_delay": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"graceful_disable_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"gslb_sp_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"health_monitor_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"horizon_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHorizonProfileSchema(),
		},
		"host_check_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"http2_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTP2PoolPropertiesSchema(),
		},
		"ignore_server_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"inline_health_monitor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"ipaddrgroup_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"lb_algo_rr_per_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"lb_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_LEAST_CONNECTIONS",
		},
		"lb_algorithm_consistent_hash_hdr": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"lb_algorithm_core_nonaffinity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"lb_algorithm_hash": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS",
		},
		"lookup_server_by_name": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"max_concurrent_connections_per_server": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"max_conn_rate_per_server": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"min_health_monitors_up": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"min_servers_up": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkFilterSchema(),
		},
		"nsx_securitygroup": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pki_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"placement_networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePlacementNetworkSchema(),
		},
		"pool_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "POOL_TYPE_GENERIC_APP",
		},
		"request_queue_depth": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "128",
			ValidateFunc: validateInteger,
		},
		"request_queue_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"resolve_pool_by_dns": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"rewrite_host_header_to_server_name": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"rewrite_host_header_to_sni": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"routing_pool": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"server_disable_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DISALLOW_NEW_CONNECTION",
		},
		"server_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"server_reselect": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPServerReselectSchema(),
		},
		"server_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"servers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServerSchema(),
		},
		"service_metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"sni_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"sp_gs_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSpGslbServiceInfoSchema(),
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
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tier1_lr": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_service_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"use_service_ssl_mode": {
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
		"vrf_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolCreate,
		Read:   ResourceAviPoolRead,
		Update: resourceAviPoolUpdate,
		Delete: resourceAviPoolDelete,
		Schema: ResourcePoolSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolImporter,
		},
	}
}

func ResourcePoolImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolSchema()
	return ResourceImporter(d, m, "pool", s)
}

func ResourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := APIRead(d, meta, "pool", s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviPoolRead in reading object %v\n", err)
	} else {
		if ignoreServers, ok := d.GetOk("ignore_servers"); ok {
			log.Printf("[DEBUG] ResourceAviPoolRead ignore_servers %v so clearing servers\n", ignoreServers)
			d.Set("servers", nil)
			d.Set("ignore_servers", ignoreServers.(bool))
		}
	}
	return err
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	ignoreServers := false
	if ignServers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignServers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err := errors.New("Error Invalid Plan. cannot set ignore_servers and servers together")
			return err
		}
		ignoreServers = true
	}
	err := APICreate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}

	if ignoreServers {
		log.Printf("[DEBUG] ignore_servers %v so clearing servers\n", ignoreServers)
		d.Set("servers", nil)
		d.Set("ignore_servers", ignoreServers)
	}

	return err
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	var err error
	s := ResourcePoolSchema()

	if ignoreServers, ok := d.GetOk("ignore_servers"); ok && ignoreServers.(bool) {
		if servers, ok := d.GetOk("servers"); ok && ignoreServers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err = errors.New("Error Invalid Plan. cannot set ignore_servers and servers together")
			return err
		}
		err = APIUpdate(d, meta, "pool", s, true)
		if err == nil {
			err = ResourceAviPoolRead(d, meta)
		}
		d.Set("servers", nil)
		d.Set("ignore_servers", true)
	} else {
		err = APIUpdate(d, meta, "pool", s)
		if err == nil {
			err = ResourceAviPoolRead(d, meta)
		}
	}

	return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if ignoreServers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignoreServers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err = errors.New("Error Invalid Plan. cannot set ignore_servers and servers together")
			return err
		}
	}
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "pool")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
