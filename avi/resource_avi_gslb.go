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

func ResourceGslbSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clear_on_max_retries": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"client_ip_addr_group": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGslbClientIpAddrGroupSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dns_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDNSConfigSchema(),
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"leader_cluster_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"maintenance_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"send_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"sites": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbSiteSchema(),
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"third_party_sites": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbThirdPartySiteSchema(),
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"view_id": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}

func resourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbCreate,
		Read:   ResourceAviGslbRead,
		Update: resourceAviGslbUpdate,
		Delete: resourceAviGslbDelete,
		Schema: ResourceGslbSchema(),
	}
}

func ResourceAviGslbRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	log.Printf("[INFO] ResourceAviGslbRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/gslb/" + uuid.(string)
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
	log.Printf("ResourceAviGslbRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviGslbRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviGslbRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviGslbRead Updated %v\n", d)
	return nil
}

func resourceAviGslbCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := ApiCreateOrUpdate(d, meta, "gslb", s)
	log.Printf("[DEBUG] created object %v: %v", "gslb", d)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "gslb", d)
	return err
}

func resourceAviGslbUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := ApiCreateOrUpdate(d, meta, "gslb", s)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "gslb", d)
	return err
}

func resourceAviGslbDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslb"
	log.Println("[INFO] ResourceAviGslbRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviGslbDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
