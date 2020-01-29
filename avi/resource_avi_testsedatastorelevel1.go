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

func ResourceTestSeDatastoreLevel1Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_se_datastore_level_2_ref": {
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

func resourceAviTestSeDatastoreLevel1() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTestSeDatastoreLevel1Create,
		Read:   ResourceAviTestSeDatastoreLevel1Read,
		Update: resourceAviTestSeDatastoreLevel1Update,
		Delete: resourceAviTestSeDatastoreLevel1Delete,
		Schema: ResourceTestSeDatastoreLevel1Schema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTestSeDatastoreLevel1Importer,
		},
	}
}

func ResourceTestSeDatastoreLevel1Importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTestSeDatastoreLevel1Schema()
	return ResourceImporter(d, m, "testsedatastorelevel1", s)
}

func ResourceAviTestSeDatastoreLevel1Read(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel1Schema()
	err := ApiRead(d, meta, "testsedatastorelevel1", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Create(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel1Schema()
	err := ApiCreateOrUpdate(d, meta, "testsedatastorelevel1", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel1Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Update(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel1Schema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "testsedatastorelevel1", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel1Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Delete(d *schema.ResourceData, meta interface{}) error {
	objType := "testsedatastorelevel1"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTestSeDatastoreLevel1Delete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
