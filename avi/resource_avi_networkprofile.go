// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceNetworkProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"connection_mirror": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
		"profile": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceNetworkProfileUnionSchema(),
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

func resourceAviNetworkProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkProfileCreate,
		Read:   ResourceAviNetworkProfileRead,
		Update: resourceAviNetworkProfileUpdate,
		Delete: resourceAviNetworkProfileDelete,
		Schema: ResourceNetworkProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkProfileImporter,
		},
	}
}

func ResourceNetworkProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkProfileSchema()
	return ResourceImporter(d, m, "networkprofile", s)
}

func ResourceAviNetworkProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := APIRead(d, meta, "networkprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := APICreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "networkprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
