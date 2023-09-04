// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
	"log"
)

func ResourceVsVipSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bgp_peer_labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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
		"dns_info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema(),
		},
		"east_west_placement": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"ipam_selector": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSelectorSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
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
		"tier1_lr": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_standard_alb": {
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
		"vsvip_cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviVsVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVsVipCreate,
		Read:   ResourceAviVsVipRead,
		Update: resourceAviVsVipUpdate,
		Delete: resourceAviVsVipDelete,
		Schema: ResourceVsVipSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVsVipImporter,
		},
	}
}

func ResourceVsVipImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVsVipSchema()
	return ResourceImporter(d, m, "vsvip", s)
}

func ResourceAviVsVipRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	err := APIRead(d, meta, "vsvip", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVsVipCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	err := APICreateOrUpdate(d, meta, "vsvip", s)
	if err == nil {
		err = ResourceAviVsVipRead(d, meta)
	}
	return err
}

func resourceAviVsVipUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	var err error
	var existingvsvip interface{}
	var apiResponse interface{}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	vsvippath := "api/vsvip/" + uuid
	err = client.AviSession.Get(vsvippath, &existingvsvip)
	var vipobjs []interface{}
	autoAllocFlag := false
	if err == nil {
		//adding default values to api_response before it overwrites the d (local state).
		//Before GO lang sets zero value to fields which are absent in api response
		//setting those fields to schema default and then overwritting d (local state)
		if localData, err := SchemaToAviData(d, s); err == nil {
			apiResponse, _ = SetDefaultsInAPIRes(existingvsvip, localData, s)
		} else {
			log.Printf("[ERROR] resourceAviVsVipUpdate in SchemaToAviData: %v\n", err)
		}
		if vsvipobj, err := APIDataToSchema(apiResponse, nil, nil); err == nil {
			objs := vsvipobj.(*schema.Set).List()
			for obj := 0; obj < len(objs); obj++ {
				vipobjs = objs[obj].(map[string]interface{})["vip"].([]interface{})
				for ob := 0; ob < len(vipobjs); ob++ {
					vipob := vipobjs[ob].(map[string]interface{})
					if value, ok := vipob["auto_allocate_ip"].(bool); ok && value {
						autoAllocFlag = true
						break
					}
				}
			}
			if autoAllocFlag {
				err = d.Set("vip", vipobjs)
				if err != nil {
					log.Printf("[ERROR] resourceAviVSVIPUpdate in Setting vip: %v\n", err)
				}
			}
		} else {
			log.Printf("[ERROR] resourceAviVsVipUpdate in APIDataToSchema: %v\n", err)
		}
	} else {
		log.Printf("[ERROR] resourceAviVsVipUpdate in GET: %v\n", err)
	}
	err = APICreateOrUpdate(d, meta, "vsvip", s)
	if err == nil {
		err = ResourceAviVsVipRead(d, meta)
	}
	return err
}

func resourceAviVsVipDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "vsvip")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
