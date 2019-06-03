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
		"addrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"apic_epg_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"country_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ip_ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPortSchema(),
		},
		"marathon_app_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"marathon_service_port": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"prefixes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrRangeSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
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
		Importer: &schema.ResourceImporter{
			State: ResourceIpAddrGroupImporter,
		},
	}
}

func ResourceIpAddrGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIpAddrGroupSchema()
	return ResourceImporter(d, m, "ipaddrgroup", s)
}

func ResourceAviIpAddrGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := ApiRead(d, meta, "ipaddrgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
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
	var err error
	err = ApiCreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	return err
}

func resourceAviIpAddrGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipaddrgroup"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviIpAddrGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
