// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceFileObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"checksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"compressed": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"created": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"expires_at": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"read_only": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"restrict_download": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"size": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviFileObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviFileObjectCreate,
		Read:   ResourceAviFileObjectRead,
		Update: resourceAviFileObjectUpdate,
		Delete: resourceAviFileObjectDelete,
		Schema: ResourceFileObjectSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFileObjectImporter,
		},
	}
}

func ResourceFileObjectImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFileObjectSchema()
	return ResourceImporter(d, m, "fileobject", s)
}

func ResourceAviFileObjectRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	err := APIRead(d, meta, "fileobject", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviFileObjectCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	err := APICreate(d, meta, "fileobject", s)
	if err == nil {
		err = ResourceAviFileObjectRead(d, meta)
	}
	return err
}

func resourceAviFileObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	var err error
	err = APIUpdate(d, meta, "fileobject", s)
	if err == nil {
		err = ResourceAviFileObjectRead(d, meta)
	}
	return err
}

func resourceAviFileObjectDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "fileobject")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
