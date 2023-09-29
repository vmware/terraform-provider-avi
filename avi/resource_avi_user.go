// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceUserRoleSchema(),
		},
		"default_tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"email": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_superuser": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"local": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"user_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"username": {
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

func resourceAviUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserCreate,
		Read:   ResourceAviUserRead,
		Update: resourceAviUserUpdate,
		Delete: resourceAviUserDelete,
		Schema: ResourceUserSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUserImporter,
		},
	}
}

func ResourceUserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUserSchema()
	return ResourceImporter(d, m, "user", s)
}

func ResourceAviUserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	err := APIRead(d, meta, "user", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	err := APICreateOrUpdate(d, meta, "user", s)
	if err == nil {
		err = ResourceAviUserRead(d, meta)
	}
	return err
}

func resourceAviUserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "user", s)
	if err == nil {
		err = ResourceAviUserRead(d, meta)
	}
	return err
}

func resourceAviUserDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "user")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
