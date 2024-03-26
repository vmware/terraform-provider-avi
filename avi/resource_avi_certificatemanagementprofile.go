// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceCertificateManagementProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
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
		"run_script_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"script_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
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

func resourceAviCertificateManagementProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCertificateManagementProfileCreate,
		Read:   ResourceAviCertificateManagementProfileRead,
		Update: resourceAviCertificateManagementProfileUpdate,
		Delete: resourceAviCertificateManagementProfileDelete,
		Schema: ResourceCertificateManagementProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCertificateManagementProfileImporter,
		},
	}
}

func ResourceCertificateManagementProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCertificateManagementProfileSchema()
	return ResourceImporter(d, m, "certificatemanagementprofile", s)
}

func ResourceAviCertificateManagementProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	err := APIRead(d, meta, "certificatemanagementprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCertificateManagementProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	err := APICreate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	var err error
	err = APIUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "certificatemanagementprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
