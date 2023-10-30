// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIPReputationDBSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base_file_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
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
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
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
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "ipreputationdb")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
