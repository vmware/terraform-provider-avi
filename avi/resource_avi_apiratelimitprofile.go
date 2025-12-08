// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApiRateLimitProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rate_limit_configuration_refs": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviApiRateLimitProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApiRateLimitProfileCreate,
		Read:   ResourceAviApiRateLimitProfileRead,
		Update: resourceAviApiRateLimitProfileUpdate,
		Delete: resourceAviApiRateLimitProfileDelete,
		Schema: ResourceApiRateLimitProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApiRateLimitProfileImporter,
		},
	}
}

func ResourceApiRateLimitProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApiRateLimitProfileSchema()
	return ResourceImporter(d, m, "apiratelimitprofile", s)
}

func ResourceAviApiRateLimitProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiRateLimitProfileSchema()
	err := APIRead(d, meta, "apiratelimitprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApiRateLimitProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiRateLimitProfileSchema()
	err := APICreate(d, meta, "apiratelimitprofile", s)
	if err == nil {
		err = ResourceAviApiRateLimitProfileRead(d, meta)
	}
	return err
}

func resourceAviApiRateLimitProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiRateLimitProfileSchema()
	var err error
	err = APIUpdate(d, meta, "apiratelimitprofile", s)
	if err == nil {
		err = ResourceAviApiRateLimitProfileRead(d, meta)
	}
	return err
}

func resourceAviApiRateLimitProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "apiratelimitprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
