// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTrafficCloneProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clone_servers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloneServerSchema(),
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"preserve_client_ip": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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

func resourceAviTrafficCloneProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTrafficCloneProfileCreate,
		Read:   ResourceAviTrafficCloneProfileRead,
		Update: resourceAviTrafficCloneProfileUpdate,
		Delete: resourceAviTrafficCloneProfileDelete,
		Schema: ResourceTrafficCloneProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTrafficCloneProfileImporter,
		},
	}
}

func ResourceTrafficCloneProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTrafficCloneProfileSchema()
	return ResourceImporter(d, m, "trafficcloneprofile", s)
}

func ResourceAviTrafficCloneProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := APIRead(d, meta, "trafficcloneprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTrafficCloneProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := APICreate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	var err error
	err = APIUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "trafficcloneprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
