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

func ResourceIpAddrGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"apic_epg_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"country_codes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ip_ports": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPortSchema(),
		},
		"marathon_app_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"marathon_service_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"prefixes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"ranges": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrRangeSchema(),
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
	}
}

func resourceAviIpAddrGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIpAddrGroupCreate,
		Read:   ResourceAviIpAddrGroupRead,
		Update: resourceAviIpAddrGroupUpdate,
		Delete: resourceAviIpAddrGroupDelete,
		Schema: ResourceIpAddrGroupSchema(),
	}
}

func ResourceAviIpAddrGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	log.Printf("[INFO] ResourceAviIpAddrGroupRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/ipaddrgroup/" + uuid.(string)
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
	log.Printf("ResourceAviIpAddrGroupRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviIpAddrGroupRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviIpAddrGroupRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviIpAddrGroupRead Updated %v\n", d)
	return nil
}

func resourceAviIpAddrGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "ipaddrgroup", s)
	log.Printf("[DEBUG] created object %v: %v", "ipaddrgroup", d)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "ipaddrgroup", d)
	return err
}

func resourceAviIpAddrGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "ipaddrgroup", d)
	return err
}

func resourceAviIpAddrGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipaddrgroup"
	log.Println("[INFO] ResourceAviIpAddrGroupRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviIpAddrGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
