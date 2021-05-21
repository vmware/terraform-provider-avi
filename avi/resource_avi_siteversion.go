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

func ResourceSiteVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datetime": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"prev_target_version": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"replication_state": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"target_timeline": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"target_version": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"timeline": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"version_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSiteVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSiteVersionCreate,
		Read:   ResourceAviSiteVersionRead,
		Update: resourceAviSiteVersionUpdate,
		Delete: resourceAviSiteVersionDelete,
		Schema: ResourceSiteVersionSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSiteVersionImporter,
		},
	}
}

func ResourceSiteVersionImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSiteVersionSchema()
	return ResourceImporter(d, m, "siteversion", s)
}

func ResourceAviSiteVersionRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	err := APIRead(d, meta, "siteversion", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSiteVersionCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	err := APICreateOrUpdate(d, meta, "siteversion", s)
	if err == nil {
		err = ResourceAviSiteVersionRead(d, meta)
	}
	return err
}

func resourceAviSiteVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "siteversion", s)
	if err == nil {
		err = ResourceAviSiteVersionRead(d, meta)
	}
	return err
}

func resourceAviSiteVersionDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "siteversion"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSiteVersionDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
