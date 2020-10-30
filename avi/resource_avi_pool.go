/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"errors"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
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
		"apic_epg_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"capacity_estimation_ttfb_thresh": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
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
		"conn_pool_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConnPoolPropertiesSchema(),
		},
		"connection_ramp_duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"default_server_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  80,
		},
		"delete_server_on_dns_refresh": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"graceful_disable_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"health_monitor_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"host_check_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"inline_health_monitor": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ipaddrgroup_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"lb_algorithm_hash": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS",
		},
		"lookup_server_by_name": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"max_concurrent_connections_per_server": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_conn_rate_per_server": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"min_health_monitors_up": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"min_servers_up": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
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
		"request_queue_depth": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  128,
		},
		"request_queue_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"resolve_pool_by_dns": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"rewrite_host_header_to_server_name": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header_to_sni": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
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
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
		"use_service_port": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
	err := ApiRead(d, meta, "pool", s)
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
	if ignore_servers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err := errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
		ignoreServers = true
	}
	err := ApiCreateOrUpdate(d, meta, "pool", s)
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

	if ignore_servers, ok := d.GetOk("ignore_servers"); ok && ignore_servers.(bool) {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err = errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
		err = ApiCreateOrUpdate(d, meta, "pool", s, true)
		if err == nil {
			err = ResourceAviPoolRead(d, meta)
		}
		d.Set("servers", nil)
		d.Set("ignore_servers", true)
	} else {
		err = ApiCreateOrUpdate(d, meta, "pool", s)
		if err == nil {
			err = ResourceAviPoolRead(d, meta)
		}
	}

	return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	if ignore_servers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err := errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
	}
	objType := "pool"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[ERROR] resourceAviPoolDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
