/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIPReputationDBSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base_file_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"incremental_file_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"service_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIPReputationServiceStatusSchema(),
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
		"vendor": {
			Type:     schema.TypeString,
			Required: true,
		},
		"version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviIPReputationDB() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIPReputationDBCreate,
		Read:   ResourceAviIPReputationDBRead,
		Update: resourceAviIPReputationDBUpdate,
		Delete: resourceAviIPReputationDBDelete,
		Schema: ResourceIPReputationDBSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIPReputationDBImporter,
		},
	}
}

func ResourceIPReputationDBImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIPReputationDBSchema()
	return ResourceImporter(d, m, "ipreputationdb", s)
}

func ResourceAviIPReputationDBRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIPReputationDBSchema()
	err := APIRead(d, meta, "ipreputationdb", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviIPReputationDBCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIPReputationDBSchema()
	err := APICreateOrUpdate(d, meta, "ipreputationdb", s)
	if err == nil {
		err = ResourceAviIPReputationDBRead(d, meta)
	}
	return err
}

func resourceAviIPReputationDBUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIPReputationDBSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "ipreputationdb", s)
	if err == nil {
		err = ResourceAviIPReputationDBRead(d, meta)
	}
	return err
}

func resourceAviIPReputationDBDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipreputationdb"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviIPReputationDBDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
