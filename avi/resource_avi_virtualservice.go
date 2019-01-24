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

func ResourceVirtualServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_standby_se_tag": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ACTIVE_STANDBY_SE_1",
		},
		"analytics_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAnalyticsPolicySchema(),
		},
		"analytics_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"apic_contract_graph": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"application_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"bulk_sync_kvcache": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"client_auth": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPClientAuthenticationParamsSchema(),
		},
		"close_client_conn_on_config_update": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CLOUD_NONE",
		},
		"connections_rate_limit": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"content_rewrite": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceContentRewriteProfileSchema(),
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"delay_fairness": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dns_info": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema(),
		},
		"dns_policies": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsPoliciesSchema(),
		},
		"east_west_placement": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_autogw": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"enable_rhi": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"enable_rhi_snat": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"error_page_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"flow_dist": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LOAD_AWARE",
		},
		"flow_label_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "NO_LABEL",
		},
		"fqdn": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_name_xlate": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"http_policies": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHTTPPoliciesSchema(),
		},
		"ign_pool_net_reach": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"l4_policies": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL4PoliciesSchema(),
		},
		"limit_doser": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"max_cps_per_client": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"microservice_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"min_pools_up": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"network_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"network_security_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nsx_securitygroup": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"performance_limits": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcePerformanceLimitsSchema(),
		},
		"pool_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remove_listening_port_on_vs_down": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"requests_rate_limit": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"scaleout_ecmp": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"server_network_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_metadata": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"service_pool_select": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServicePoolSelectorSchema(),
		},
		"services": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServiceSchema(),
		},
		"sideband_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSidebandProfileSchema(),
		},
		"snat_ip": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"ssl_key_and_certificate_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssl_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_sess_cache_avg_size": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1024,
		},
		"sso_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSSOPolicySchema(),
		},
		"static_dns_records": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRecordSchema(),
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"traffic_clone_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"traffic_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VS_TYPE_NORMAL",
		},
		"use_bridge_ip_as_vip": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"use_vip_as_snat": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vh_domain_name": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"vh_parent_vs_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"vip": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVipSchema(),
		},
		"vrf_context_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_datascripts": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVSDataScriptsSchema(),
		},
		"vsvip_cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"vsvip_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"weight": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
	}
}

func resourceAviVirtualService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVirtualServiceCreate,
		Read:   ResourceAviVirtualServiceRead,
		Update: resourceAviVirtualServiceUpdate,
		Delete: resourceAviVirtualServiceDelete,
		Schema: ResourceVirtualServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVirtualServiceImporter,
		},
	}
}

func ResourceVirtualServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVirtualServiceSchema()
	return ResourceImporter(d, m, "virtualservice", s)
}

func ResourceAviVirtualServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	err := ApiRead(d, meta, "virtualservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	err := ApiCreateOrUpdate(d, meta, "virtualservice", s)
	if err == nil {
		err = ResourceAviVirtualServiceRead(d, meta)
	}
	return err
}

func resourceAviVirtualServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	var err error
	var existingvirtualservice interface{}
	var apiResponse interface{}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	virtualservicepath := "api/virtualservice/" + uuid
	err = client.AviSession.Get(virtualservicepath, &existingvirtualservice)
	if err == nil {
		//adding default values to api_response before it overwrites the d (local state).
		//Before GO lang sets zero value to fields which are absent in api response
		//setting those fields to schema default and then overwritting d (local state)
		if localData, err := SchemaToAviData(d, s); err == nil {
			apiResponse, err = SetDefaultsInAPIRes(existingvirtualservice, localData, s)
		} else {
			log.Printf("[ERROR] resourceAviVirtualServiceUpdate in SchemaToAviData: %v\n", err)
		}
		if virtualserviceobj, err := ApiDataToSchema(apiResponse, nil, nil); err == nil {
			objs := virtualserviceobj.(*schema.Set).List()
			for obj := 0; obj < len(objs); obj++ {
				vsvipref := objs[obj].(map[string]interface{})["vsvip_ref"]
				err = d.Set("vsvip_ref", vsvipref.(string))
				if err != nil {
					log.Printf("[ERROR] resourceAviVirtualServiceUpdate in Setting vsvip ref: %v\n", err)
				}
				vipob := objs[obj].(map[string]interface{})["vip"]
				err = d.Set("vip", vipob)
				if err != nil {
					log.Printf("[ERROR] resourceAviVirtualServiceUpdate in Setting vip: %v\n", err)
				}
			}
		} else {
			log.Printf("[ERROR] resourceAviVirtualServiceUpdate in ApiDataToSchema: %v\n", err)
		}
	} else {
		log.Printf("[ERROR] resourceAviVirtualServiceUpdate in GET: %v\n", err)
	}
	err = ApiCreateOrUpdate(d, meta, "virtualservice", s)
	if err == nil {
		err = ResourceAviVirtualServiceRead(d, meta)
	}
	return err
}

func resourceAviVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "virtualservice"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviVirtualServiceDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
