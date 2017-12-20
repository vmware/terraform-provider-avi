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
			Default:  "/api/cloud?name=Default-Cloud",
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
	}
}

func ResourceAviVsVipRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	log.Printf("[INFO] ResourceAviVsVipRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/vsvip/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviVsVipRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviVsVipRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviVsVipRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviVsVipRead Updated %v\n", d)
	return nil
}

func resourceAviVsVipCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	err := ApiCreateOrUpdate(d, meta, "vsvip", s)
	log.Printf("[DEBUG] created object %v: %v", "vsvip", d)
	if err == nil {
		err = ResourceAviVsVipRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "vsvip", d)
	return err
}

func resourceAviVsVipUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsVipSchema()
	err := ApiCreateOrUpdate(d, meta, "vsvip", s)
	if err == nil {
		err = ResourceAviVsVipRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "vsvip", d)
	return err
}

func resourceAviVsVipDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vsvip"
	log.Println("[INFO] ResourceAviVsVipRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviVsVipDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
