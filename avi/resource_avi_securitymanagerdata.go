// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSecurityManagerDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_learning_info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDbAppLearningInfoSchema(),
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
	}
}

func resourceAviSecurityManagerData() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSecurityManagerDataCreate,
		Read:   ResourceAviSecurityManagerDataRead,
		Update: resourceAviSecurityManagerDataUpdate,
		Delete: resourceAviSecurityManagerDataDelete,
		Schema: ResourceSecurityManagerDataSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSecurityManagerDataImporter,
		},
	}
}

func ResourceSecurityManagerDataImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSecurityManagerDataSchema()
	return ResourceImporter(d, m, "securitymanagerdata", s)
}

func ResourceAviSecurityManagerDataRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	err := APIRead(d, meta, "securitymanagerdata", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSecurityManagerDataCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	err := APICreate(d, meta, "securitymanagerdata", s)
	if err == nil {
		err = ResourceAviSecurityManagerDataRead(d, meta)
	}
	return err
}

func resourceAviSecurityManagerDataUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	var err error
	err = APIUpdate(d, meta, "securitymanagerdata", s)
	if err == nil {
		err = ResourceAviSecurityManagerDataRead(d, meta)
	}
	return err
}

func resourceAviSecurityManagerDataDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "securitymanagerdata")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
