// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

// nolint
func ResourceCustomIpamDnsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"script_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_uri": {
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

// nolint
func resourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCustomIpamDnsProfileCreate,
		Read:   ResourceAviCustomIpamDnsProfileRead,
		Update: resourceAviCustomIpamDnsProfileUpdate,
		Delete: resourceAviCustomIpamDnsProfileDelete,
		Schema: ResourceCustomIpamDnsProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCustomIpamDnsProfileImporter,
		},
	}
}

// nolint
func ResourceCustomIpamDnsProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCustomIpamDnsProfileSchema()
	return ResourceImporter(d, m, "customipamdnsprofile", s)
}

// nolint
func ResourceAviCustomIpamDnsProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := APIRead(d, meta, "customipamdnsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

// nolint
func resourceAviCustomIpamDnsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := APICreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

// nolint
func resourceAviCustomIpamDnsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

// nolint
func resourceAviCustomIpamDnsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "customipamdnsprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
