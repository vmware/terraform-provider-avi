// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceClfProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clf_pools": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceClfPoolSchema(),
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
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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
		"replicate": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func resourceAviClfProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviClfProfileCreate,
		Read:   ResourceAviClfProfileRead,
		Update: resourceAviClfProfileUpdate,
		Delete: resourceAviClfProfileDelete,
		Schema: ResourceClfProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceClfProfileImporter,
		},
	}
}

func ResourceClfProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceClfProfileSchema()
	return ResourceImporter(d, m, "clfprofile", s)
}

func ResourceAviClfProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClfProfileSchema()
	err := APIRead(d, meta, "clfprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviClfProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClfProfileSchema()
	err := APICreate(d, meta, "clfprofile", s)
	if err == nil {
		err = ResourceAviClfProfileRead(d, meta)
	}
	return err
}

func resourceAviClfProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClfProfileSchema()
	var err error
	err = APIUpdate(d, meta, "clfprofile", s)
	if err == nil {
		err = ResourceAviClfProfileRead(d, meta)
	}
	return err
}

func resourceAviClfProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "clfprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
