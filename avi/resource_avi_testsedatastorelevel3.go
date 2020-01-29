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

func ResourceTestSeDatastoreLevel3Schema() map[string]*schema.Schema {
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
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTestSeDatastoreLevel3() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTestSeDatastoreLevel3Create,
		Read:   ResourceAviTestSeDatastoreLevel3Read,
		Update: resourceAviTestSeDatastoreLevel3Update,
		Delete: resourceAviTestSeDatastoreLevel3Delete,
		Schema: ResourceTestSeDatastoreLevel3Schema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTestSeDatastoreLevel3Importer,
		},
	}
}

func ResourceTestSeDatastoreLevel3Importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTestSeDatastoreLevel3Schema()
	return ResourceImporter(d, m, "testsedatastorelevel3", s)
}

func ResourceAviTestSeDatastoreLevel3Read(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel3Schema()
	err := ApiRead(d, meta, "testsedatastorelevel3", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTestSeDatastoreLevel3Create(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel3Schema()
	err := ApiCreateOrUpdate(d, meta, "testsedatastorelevel3", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel3Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel3Update(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel3Schema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "testsedatastorelevel3", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel3Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel3Delete(d *schema.ResourceData, meta interface{}) error {
	objType := "testsedatastorelevel3"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTestSeDatastoreLevel3Delete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
