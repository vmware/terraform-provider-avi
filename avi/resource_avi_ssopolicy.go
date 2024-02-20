// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSSOPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authentication_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthenticationPolicySchema(),
		},
		"authorization_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthorizationPolicySchema(),
		},
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
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSO_TYPE_SAML",
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSSOPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSSOPolicyCreate,
		Read:   ResourceAviSSOPolicyRead,
		Update: resourceAviSSOPolicyUpdate,
		Delete: resourceAviSSOPolicyDelete,
		Schema: ResourceSSOPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSSOPolicyImporter,
		},
	}
}

func ResourceSSOPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSSOPolicySchema()
	return ResourceImporter(d, m, "ssopolicy", s)
}

func ResourceAviSSOPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	err := APIRead(d, meta, "ssopolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSSOPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	err := APICreate(d, meta, "ssopolicy", s)
	if err == nil {
		err = ResourceAviSSOPolicyRead(d, meta)
	}
	return err
}

func resourceAviSSOPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	var err error
	err = APIUpdate(d, meta, "ssopolicy", s)
	if err == nil {
		err = ResourceAviSSOPolicyRead(d, meta)
	}
	return err
}

func resourceAviSSOPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "ssopolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
