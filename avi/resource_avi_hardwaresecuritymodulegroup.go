// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceHardwareSecurityModuleGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"hsm": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceHardwareSecurityModuleSchema(),
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
	}
}

func resourceAviHardwareSecurityModuleGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHardwareSecurityModuleGroupCreate,
		Read:   ResourceAviHardwareSecurityModuleGroupRead,
		Update: resourceAviHardwareSecurityModuleGroupUpdate,
		Delete: resourceAviHardwareSecurityModuleGroupDelete,
		Schema: ResourceHardwareSecurityModuleGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHardwareSecurityModuleGroupImporter,
		},
	}
}

func ResourceHardwareSecurityModuleGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHardwareSecurityModuleGroupSchema()
	return ResourceImporter(d, m, "hardwaresecuritymodulegroup", s)
}

func ResourceAviHardwareSecurityModuleGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := APIRead(d, meta, "hardwaresecuritymodulegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := APICreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "hardwaresecuritymodulegroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
