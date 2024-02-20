// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceUserAccountProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_lock_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"credentials_timeout_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "180",
			ValidateFunc: validateInteger,
		},
		"login_failure_count_expiry_window": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"max_concurrent_sessions": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"max_login_failure_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"max_password_history_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"name": {
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

func resourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserAccountProfileCreate,
		Read:   ResourceAviUserAccountProfileRead,
		Update: resourceAviUserAccountProfileUpdate,
		Delete: resourceAviUserAccountProfileDelete,
		Schema: ResourceUserAccountProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUserAccountProfileImporter,
		},
	}
}

func ResourceUserAccountProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUserAccountProfileSchema()
	return ResourceImporter(d, m, "useraccountprofile", s)
}

func ResourceAviUserAccountProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := APIRead(d, meta, "useraccountprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUserAccountProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := APICreate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	var err error
	err = APIUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "useraccountprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
