// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApiSchemaSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"additional_object_key_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "API_ACTION_INHERIT_FROM_API_POLICY",
		},
		"additional_properties_schema": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
		},
		"allow_additional_properties": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"array_item_type": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
		},
		"composite_types": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
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
		"discriminator": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDiscriminatorDescriptionSchema(),
		},
		"max_items": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"min_items": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"object_properties": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceApiObjectPropertiesSchema(),
		},
		"source": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"unique_items": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviApiSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApiSchemaCreate,
		Read:   ResourceAviApiSchemaRead,
		Update: resourceAviApiSchemaUpdate,
		Delete: resourceAviApiSchemaDelete,
		Schema: ResourceApiSchemaSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApiSchemaImporter,
		},
	}
}

func ResourceApiSchemaImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApiSchemaSchema()
	return ResourceImporter(d, m, "apischema", s)
}

func ResourceAviApiSchemaRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSchemaSchema()
	err := APIRead(d, meta, "apischema", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApiSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSchemaSchema()
	err := APICreate(d, meta, "apischema", s)
	if err == nil {
		err = ResourceAviApiSchemaRead(d, meta)
	}
	return err
}

func resourceAviApiSchemaUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSchemaSchema()
	var err error
	err = APIUpdate(d, meta, "apischema", s)
	if err == nil {
		err = ResourceAviApiSchemaRead(d, meta)
	}
	return err
}

func resourceAviApiSchemaDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "apischema")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
