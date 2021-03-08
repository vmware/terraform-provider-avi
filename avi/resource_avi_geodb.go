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

func ResourceGeoDBSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"files": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceGeoDBFileSchema(),
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"mappings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGeoDBMappingSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviGeoDB() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGeoDBCreate,
		Read:   ResourceAviGeoDBRead,
		Update: resourceAviGeoDBUpdate,
		Delete: resourceAviGeoDBDelete,
		Schema: ResourceGeoDBSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGeoDBImporter,
		},
	}
}

func ResourceGeoDBImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGeoDBSchema()
	return ResourceImporter(d, m, "geodb", s)
}

func ResourceAviGeoDBRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGeoDBSchema()
	err := APIRead(d, meta, "geodb", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGeoDBCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGeoDBSchema()
	err := APICreateOrUpdate(d, meta, "geodb", s)
	if err == nil {
		err = ResourceAviGeoDBRead(d, meta)
	}
	return err
}

func resourceAviGeoDBUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGeoDBSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "geodb", s)
	if err == nil {
		err = ResourceAviGeoDBRead(d, meta)
	}
	return err
}

func resourceAviGeoDBDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "geodb"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviGeoDBDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
