// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceFileObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_spec_detail": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiSpecDetailSchema(),
		},
		"checksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"child_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"completed_events": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"compressed": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ForceNew:     true,
			ValidateFunc: validateBool,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"crl_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCRLSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"end_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceFileObjectEventMapSchema(),
		},
		"expires_at": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gslb_geodb_format": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"has_parent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ForceNew:     true,
			ValidateFunc: validateBool,
		},
		"history": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTaskEventHistorySchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ForceNew:     true,
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
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
			ForceNew:     true,
			ValidateFunc: validateBool,
		},
		"size": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"start_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFileObjectStateSchema(),
		},
		"task_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTaskEventMapSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"total_events": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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
