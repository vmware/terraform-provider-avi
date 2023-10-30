// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceVSDataScriptSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datascript": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVSDataScriptSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"geo_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_reputation_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ipgroup_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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
		"pki_profile_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pool_group_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pool_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"protocol_parser_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"rate_limiters": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRateLimiterSchema(),
		},
		"ssl_key_certificate_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssl_profile_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"string_group_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviVSDataScriptSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVSDataScriptSetCreate,
		Read:   ResourceAviVSDataScriptSetRead,
		Update: resourceAviVSDataScriptSetUpdate,
		Delete: resourceAviVSDataScriptSetDelete,
		Schema: ResourceVSDataScriptSetSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVSDataScriptSetImporter,
		},
	}
}

func ResourceVSDataScriptSetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVSDataScriptSetSchema()
	return ResourceImporter(d, m, "vsdatascriptset", s)
}

func ResourceAviVSDataScriptSetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	err := APIRead(d, meta, "vsdatascriptset", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVSDataScriptSetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	err := APICreateOrUpdate(d, meta, "vsdatascriptset", s)
	if err == nil {
		err = ResourceAviVSDataScriptSetRead(d, meta)
	}
	return err
}

func resourceAviVSDataScriptSetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "vsdatascriptset", s)
	if err == nil {
		err = ResourceAviVSDataScriptSetRead(d, meta)
	}
	return err
}

func resourceAviVSDataScriptSetDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "vsdatascriptset")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
