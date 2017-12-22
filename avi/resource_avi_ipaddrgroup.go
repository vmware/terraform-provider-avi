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
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviIpAddrGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	return err
}

func resourceAviIpAddrGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	return err
}

func resourceAviIpAddrGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipaddrgroup"
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
