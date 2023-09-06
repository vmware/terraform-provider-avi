// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApplicationPersistenceProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_cookie_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAppCookiePersistenceProfileSchema(),
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
		"hdr_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHdrPersistenceProfileSchema(),
		},
		"http_cookie_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHttpCookiePersistenceProfileSchema(),
		},
		"ip_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIPPersistenceProfileSchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
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
		"persistence_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"server_hm_down_recovery": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HM_DOWN_PICK_NEW_SERVER",
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

func resourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationPersistenceProfileCreate,
		Read:   ResourceAviApplicationPersistenceProfileRead,
		Update: resourceAviApplicationPersistenceProfileUpdate,
		Delete: resourceAviApplicationPersistenceProfileDelete,
		Schema: ResourceApplicationPersistenceProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApplicationPersistenceProfileImporter,
		},
	}
}

func ResourceApplicationPersistenceProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApplicationPersistenceProfileSchema()
	return ResourceImporter(d, m, "applicationpersistenceprofile", s)
}

func ResourceAviApplicationPersistenceProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := APIRead(d, meta, "applicationpersistenceprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApplicationPersistenceProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := APICreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "applicationpersistenceprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
