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

func ResourceGslbServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_persistence_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"controller_health_status_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"domain_names": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"down_response": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGslbServiceDownResponseSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbPoolSchema(),
		},
		"health_monitor_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"health_monitor_scope": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS",
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"min_members": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"num_dns_ip": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"pool_algorithm": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_ALGORITHM_PRIORITY",
		},
		"site_persistence_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ttl": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"use_edns_client_subnet": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"wildcard_match": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourceAviGslbService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbServiceCreate,
		Read:   ResourceAviGslbServiceRead,
		Update: resourceAviGslbServiceUpdate,
		Delete: resourceAviGslbServiceDelete,
		Schema: ResourceGslbServiceSchema(),
	}
}

func ResourceAviGslbServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	log.Printf("[INFO] ResourceAviGslbServiceRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/gslbservice/" + uuid.(string)
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
	log.Printf("ResourceAviGslbServiceRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviGslbServiceRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviGslbServiceRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviGslbServiceRead Updated %v\n", d)
	return nil
}

func resourceAviGslbServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := ApiCreateOrUpdate(d, meta, "gslbservice", s)
	log.Printf("[DEBUG] created object %v: %v", "gslbservice", d)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "gslbservice", d)
	return err
}

func resourceAviGslbServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := ApiCreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "gslbservice", d)
	return err
}

func resourceAviGslbServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslbservice"
	log.Println("[INFO] ResourceAviGslbServiceRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviGslbServiceDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
