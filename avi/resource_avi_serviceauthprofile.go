// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceServiceAuthProfileSchema() map[string]*schema.Schema {
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
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"service_oauth_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceServiceOAuthSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
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

func resourceAviServiceAuthProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceAuthProfileCreate,
		Read:   ResourceAviServiceAuthProfileRead,
		Update: resourceAviServiceAuthProfileUpdate,
		Delete: resourceAviServiceAuthProfileDelete,
		Schema: ResourceServiceAuthProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceServiceAuthProfileImporter,
		},
	}
}

func ResourceServiceAuthProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServiceAuthProfileSchema()
	return ResourceImporter(d, m, "serviceauthprofile", s)
}

func ResourceAviServiceAuthProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceAuthProfileSchema()
	err := APIRead(d, meta, "serviceauthprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServiceAuthProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceAuthProfileSchema()
	err := APICreateOrUpdate(d, meta, "serviceauthprofile", s)
	if err == nil {
		err = ResourceAviServiceAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviServiceAuthProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceAuthProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "serviceauthprofile", s)
	if err == nil {
		err = ResourceAviServiceAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviServiceAuthProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "serviceauthprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
