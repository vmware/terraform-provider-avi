// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceWebhookSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"callback_url": {
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
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
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
		"verification_token": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWebhookCreate,
		Read:   ResourceAviWebhookRead,
		Update: resourceAviWebhookUpdate,
		Delete: resourceAviWebhookDelete,
		Schema: ResourceWebhookSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWebhookImporter,
		},
	}
}

func ResourceWebhookImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWebhookSchema()
	return ResourceImporter(d, m, "webhook", s)
}

func ResourceAviWebhookRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := APIRead(d, meta, "webhook", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := APICreate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	var err error
	err = APIUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "webhook")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
