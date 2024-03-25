// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"config_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTenantConfigurationSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enforce_label_group": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"label_group_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"local": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantCreate,
		Read:   ResourceAviTenantRead,
		Update: resourceAviTenantUpdate,
		Delete: resourceAviTenantDelete,
		Schema: ResourceTenantSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTenantImporter,
		},
	}
}

func ResourceTenantImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTenantSchema()
	return ResourceImporter(d, m, "tenant", s)
}

func ResourceAviTenantRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := APIRead(d, meta, "tenant", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTenantCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := APICreate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	var err error
	err = APIUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "tenant")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
