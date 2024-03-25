// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceAlertEmailConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_emails": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"to_emails": {
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

func resourceAviAlertEmailConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertEmailConfigCreate,
		Read:   ResourceAviAlertEmailConfigRead,
		Update: resourceAviAlertEmailConfigUpdate,
		Delete: resourceAviAlertEmailConfigDelete,
		Schema: ResourceAlertEmailConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAlertEmailConfigImporter,
		},
	}
}

func ResourceAlertEmailConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAlertEmailConfigSchema()
	return ResourceImporter(d, m, "alertemailconfig", s)
}

func ResourceAviAlertEmailConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	err := APIRead(d, meta, "alertemailconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAlertEmailConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	err := APICreate(d, meta, "alertemailconfig", s)
	if err == nil {
		err = ResourceAviAlertEmailConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertEmailConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	var err error
	err = APIUpdate(d, meta, "alertemailconfig", s)
	if err == nil {
		err = ResourceAviAlertEmailConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertEmailConfigDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "alertemailconfig")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
