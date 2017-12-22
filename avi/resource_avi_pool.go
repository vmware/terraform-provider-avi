/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourcePoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"a_pool": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ab_pool": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAbPoolSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ab_priority": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"apic_epg_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"application_persistence_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"autoscale_launch_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"autoscale_networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"autoscale_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"capacity_estimation": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"capacity_estimation_ttfb_thresh": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud",
		},
		"connection_ramp_duration": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"default_server_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  80,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"domain_name": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"east_west": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"external_autoscale_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"fail_action": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceFailActionSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"fewest_tasks_feedback_delay": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"graceful_disable_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"health_monitor_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"host_check_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"inline_health_monitor": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ipaddrgroup_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"lb_algorithm": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_LEAST_CONNECTIONS",
		},
		"lb_algorithm_consistent_hash_hdr": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"lb_algorithm_core_nonaffinity": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"lb_algorithm_hash": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS",
		},
		"lookup_server_by_name": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"max_concurrent_connections_per_server": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_conn_rate_per_server": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRateProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkFilterSchema(),
		},
		"nsx_securitygroup": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pki_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"placement_networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePlacementNetworkSchema(),
		},
		"prst_hdr_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"request_queue_depth": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  128,
		},
		"request_queue_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header_to_server_name": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header_to_sni": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"server_auto_scale": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"server_count": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"server_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"server_reselect": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPServerReselectSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServerSchema(),
		},
		"sni_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ssl_key_and_certificate_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ssl_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_service_port": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vrf_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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
	}
}

func ResourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/pool/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := ApiCreateOrUpdate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	return err
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := ApiCreateOrUpdate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "pool"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviPoolDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
