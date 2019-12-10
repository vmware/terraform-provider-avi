/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func ResourceGslbSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"async_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"clear_on_max_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"client_ip_addr_group": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGslbClientIpAddrGroupSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDNSConfigSchema(),
		},
		"error_resync_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"leader_cluster_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"maintenance_mode": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"send_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"send_interval_prior_to_maintenance_mode": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"sites": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbSiteSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"third_party_sites": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbThirdPartySiteSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"view_id": {
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
		Importer: &schema.ResourceImporter{
			State: ResourceGslbImporter,
		},
	}
}

func ResourceGslbImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbSchema()
	return ResourceImporter(d, m, "gslb", s)
}

func ResourceAviGslbRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := ApiRead(d, meta, "gslb", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := ApiCreateOrUpdate(d, meta, "gslb", s)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	return err
}

func resourceAviGslbUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "gslb", s)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	return err
}

func resourceAviGslbDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslb"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviGslbDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
