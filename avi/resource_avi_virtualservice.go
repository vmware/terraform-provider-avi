/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceVirtualServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_standby_se_tag": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ACTIVE_STANDBY_SE_1",
		},
		"allow_invalid_client_cert": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"analytics_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAnalyticsPolicySchema(),
		},
		"analytics_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"apic_contract_graph": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"application_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"bulk_sync_kvcache": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"client_auth": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPClientAuthenticationParamsSchema(),
		},
		"close_client_conn_on_config_update": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		"cloud_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CLOUD_NONE",
		},
		"connections_rate_limit": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"content_rewrite": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceContentRewriteProfileSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"delay_fairness": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema(),
		},
		"dns_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsPoliciesSchema(),
		},
		"east_west_placement": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_autogw": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"enable_rhi": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"enable_rhi_snat": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"error_page_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"flow_dist": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LOAD_AWARE",
		},
		"flow_label_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "NO_LABEL",
		},
		"fqdn": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_name_xlate": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"http_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHTTPPoliciesSchema(),
		},
		"ign_pool_net_reach": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"l4_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL4PoliciesSchema(),
		},
		"limit_doser": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"max_cps_per_client": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"microservice_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"min_pools_up": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"network_security_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nsx_securitygroup": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"performance_limits": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePerformanceLimitsSchema(),
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
		"remove_listening_port_on_vs_down": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"requests_rate_limit": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"saml_sp_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSAMLSPConfigSchema(),
		},
		"scaleout_ecmp": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"server_network_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_pool_select": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServicePoolSelectorSchema(),
		},
		"services": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServiceSchema(),
		},
		"sideband_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSidebandProfileSchema(),
		},
		"snat_ip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"ssl_key_and_certificate_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssl_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_profile_selectors": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSSLProfileSelectorSchema(),
		},
		"ssl_sess_cache_avg_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1024,
		},
		"sso_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"static_dns_records": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRecordSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_se_datastore_level_1_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"topology_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsPoliciesSchema(),
		},
		"traffic_clone_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"traffic_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VS_TYPE_NORMAL",
		},
		"use_bridge_ip_as_vip": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"use_vip_as_snat": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vh_domain_name": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"vh_parent_vs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVipSchema(),
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_datascripts": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVSDataScriptsSchema(),
		},
		"vsvip_cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vsvip_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"weight": {
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
	err := APIRead(d, meta, "virtualservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	err := APICreateOrUpdate(d, meta, "virtualservice", s)
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
			apiResponse, _ = SetDefaultsInAPIRes(existingvirtualservice, localData, s)
		} else {
			log.Printf("[ERROR] resourceAviVirtualServiceUpdate in SchemaToAviData: %v\n", err)
		}
		if virtualserviceobj, err := APIDataToSchema(apiResponse, nil, nil); err == nil {
			objs := virtualserviceobj.(*schema.Set).List()
			for obj := 0; obj < len(objs); obj++ {
				vsvipref, isVsVip := objs[obj].(map[string]interface{})["vsvip_ref"]
				if isVsVip {
					err = d.Set("vsvip_ref", vsvipref.(string))
					if err != nil {
						log.Printf("[ERROR] resourceAviVirtualServiceUpdate in Setting vsvip ref: %v\n", err)
					}
				}
				vipob, isVip := objs[obj].(map[string]interface{})["vip"]
				if isVip {
					err = d.Set("vip", vipob)
					if err != nil {
						log.Printf("[ERROR] resourceAviVirtualServiceUpdate in Setting vip: %v\n", err)
					}
				}
			}
		} else {
			log.Printf("[ERROR] resourceAviVirtualServiceUpdate in APIDataToSchema: %v\n", err)
		}
	} else {
		log.Printf("[ERROR] resourceAviVirtualServiceUpdate in GET: %v\n", err)
	}
	err = APICreateOrUpdate(d, meta, "virtualservice", s)
	if err == nil {
		err = ResourceAviVirtualServiceRead(d, meta)
	}
	return err
}

func resourceAviVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "virtualservice"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
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
