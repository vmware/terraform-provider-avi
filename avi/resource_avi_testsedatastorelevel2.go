// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceTestSeDatastoreLevel2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
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
		"test_se_datastore_level_3_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTestSeDatastoreLevel2() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTestSeDatastoreLevel2Create,
		Read:   ResourceAviTestSeDatastoreLevel2Read,
		Update: resourceAviTestSeDatastoreLevel2Update,
		Delete: resourceAviTestSeDatastoreLevel2Delete,
		Schema: ResourceTestSeDatastoreLevel2Schema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTestSeDatastoreLevel2Importer,
		},
	}
}

func ResourceTestSeDatastoreLevel2Importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTestSeDatastoreLevel2Schema()
	return ResourceImporter(d, m, "testsedatastorelevel2", s)
}

func ResourceAviTestSeDatastoreLevel2Read(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel2Schema()
	err := APIRead(d, meta, "testsedatastorelevel2", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTestSeDatastoreLevel2Create(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel2Schema()
	err := APICreateOrUpdate(d, meta, "testsedatastorelevel2", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel2Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel2Update(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel2Schema()
	var err error
	err = APICreateOrUpdate(d, meta, "testsedatastorelevel2", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel2Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel2Delete(d *schema.ResourceData, meta interface{}) error {
	objType := "testsedatastorelevel2"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTestSeDatastoreLevel2Delete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
