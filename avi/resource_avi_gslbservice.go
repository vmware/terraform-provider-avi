/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceGslbServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_persistence_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"controller_health_status_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"domain_names": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"down_response": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGslbServiceDownResponseSchema(),
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"groups": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceGslbPoolSchema(),
		},
		"health_monitor_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"health_monitor_scope": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS",
		},
		"hm_off": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"min_members": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"num_dns_ip": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"pool_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_ALGORITHM_PRIORITY",
		},
		"resolve_cname": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"site_persistence_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"use_edns_client_subnet": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"wildcard_match": {
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
	err := APIRead(d, meta, "gslbservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := APICreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslbservice"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
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
