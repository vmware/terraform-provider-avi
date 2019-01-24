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
			Computed: true,
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
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"health_monitor_scope": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS",
		},
		"hm_off": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
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
			Computed: true,
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
		Importer: &schema.ResourceImporter{
			State: ResourceGslbServiceImporter,
		},
	}
}

func ResourceGslbServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbServiceSchema()
	return ResourceImporter(d, m, "gslbservice", s)
}

func ResourceAviGslbServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := ApiRead(d, meta, "gslbservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := ApiCreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslbservice"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviGslbServiceDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
