// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
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
			Optional: true,
			Elem:     ResourceGeoDBFileSchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "geodb")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
