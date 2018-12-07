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

func ResourceVsVipSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_info": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema(),
		},
		"east_west_placement": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"vsvip_cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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
	err := ApiRead(d, meta, "vsvip", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVsVipCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	err := ApiCreateOrUpdate(d, meta, "vsvip", s)
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
			apiResponse, err = SetDefaultsInAPIRes(existingvsvip, localData, s)
		} else {
			log.Printf("[ERROR] resourceAviVsVipUpdate in SchemaToAviData: %v\n", err)
		}
		if vsvipobj, err := ApiDataToSchema(apiResponse, nil, nil); err == nil {
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
			log.Printf("[ERROR] resourceAviVsVipUpdate in ApiDataToSchema: %v\n", err)
		}
	} else {
		log.Printf("[ERROR] resourceAviVsVipUpdate in GET: %v\n", err)
	}
	err = ApiCreateOrUpdate(d, meta, "vsvip", s)
	if err == nil {
		err = ResourceAviVsVipRead(d, meta)
	}
	return err
}

func resourceAviVsVipDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vsvip"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviVsVipDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
