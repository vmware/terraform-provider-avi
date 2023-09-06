// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_unlabelled_access": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"filters": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"privileges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePermissionSchema(),
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

func resourceAviRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRoleCreate,
		Read:   ResourceAviRoleRead,
		Update: resourceAviRoleUpdate,
		Delete: resourceAviRoleDelete,
		Schema: ResourceRoleSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRoleImporter,
		},
	}
}

func ResourceRoleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRoleSchema()
	return ResourceImporter(d, m, "role", s)
}

func ResourceAviRoleRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := APIRead(d, meta, "role", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRoleCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := APICreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "role")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
