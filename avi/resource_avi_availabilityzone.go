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

func ResourceAvailabilityZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"vcenter_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func resourceAviAvailabilityZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAvailabilityZoneCreate,
		Read:   ResourceAviAvailabilityZoneRead,
		Update: resourceAviAvailabilityZoneUpdate,
		Delete: resourceAviAvailabilityZoneDelete,
		Schema: ResourceAvailabilityZoneSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAvailabilityZoneImporter,
		},
	}
}

func ResourceAvailabilityZoneImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAvailabilityZoneSchema()
	return ResourceImporter(d, m, "availabilityzone", s)
}

func ResourceAviAvailabilityZoneRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	err := ApiRead(d, meta, "availabilityzone", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAvailabilityZoneCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	err := ApiCreateOrUpdate(d, meta, "availabilityzone", s)
	if err == nil {
		err = ResourceAviAvailabilityZoneRead(d, meta)
	}
	return err
}

func resourceAviAvailabilityZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "availabilityzone", s)
	if err == nil {
		err = ResourceAviAvailabilityZoneRead(d, meta)
	}
	return err
}

func resourceAviAvailabilityZoneDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "availabilityzone"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAvailabilityZoneDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
