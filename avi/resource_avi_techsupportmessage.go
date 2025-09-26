// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTechSupportMessageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"status_code": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tech_support_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTechSupportMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTechSupportMessageCreate,
		Read:   ResourceAviTechSupportMessageRead,
		Update: resourceAviTechSupportMessageUpdate,
		Delete: resourceAviTechSupportMessageDelete,
		Schema: ResourceTechSupportMessageSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTechSupportMessageImporter,
		},
	}
}

func ResourceTechSupportMessageImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTechSupportMessageSchema()
	return ResourceImporter(d, m, "techsupportmessage", s)
}

func ResourceAviTechSupportMessageRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportMessageSchema()
	err := APIRead(d, meta, "techsupportmessage", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTechSupportMessageCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportMessageSchema()
	err := APICreate(d, meta, "techsupportmessage", s)
	if err == nil {
		err = ResourceAviTechSupportMessageRead(d, meta)
	}
	return err
}

func resourceAviTechSupportMessageUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportMessageSchema()
	var err error
	err = APIUpdate(d, meta, "techsupportmessage", s)
	if err == nil {
		err = ResourceAviTechSupportMessageRead(d, meta)
	}
	return err
}

func resourceAviTechSupportMessageDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "techsupportmessage")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
