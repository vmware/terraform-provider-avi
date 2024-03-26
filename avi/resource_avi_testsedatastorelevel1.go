// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTestSeDatastoreLevel1Schema() map[string]*schema.Schema {
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
	err := APIRead(d, meta, "testsedatastorelevel1", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Create(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel1Schema()
	err := APICreate(d, meta, "testsedatastorelevel1", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel1Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Update(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTestSeDatastoreLevel1Schema()
	var err error
	err = APIUpdate(d, meta, "testsedatastorelevel1", s)
	if err == nil {
		err = ResourceAviTestSeDatastoreLevel1Read(d, meta)
	}
	return err
}

func resourceAviTestSeDatastoreLevel1Delete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "testsedatastorelevel1")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
