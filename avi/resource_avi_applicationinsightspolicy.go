// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApplicationInsightsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_insights_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApplicationInsightsParamsSchema(),
		},
		"application_sampling_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApplicationSamplingConfigSchema(),
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
		"enable_application_insights": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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

func resourceAviApplicationInsightsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationInsightsPolicyCreate,
		Read:   ResourceAviApplicationInsightsPolicyRead,
		Update: resourceAviApplicationInsightsPolicyUpdate,
		Delete: resourceAviApplicationInsightsPolicyDelete,
		Schema: ResourceApplicationInsightsPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApplicationInsightsPolicyImporter,
		},
	}
}

func ResourceApplicationInsightsPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApplicationInsightsPolicySchema()
	return ResourceImporter(d, m, "applicationinsightspolicy", s)
}

func ResourceAviApplicationInsightsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsPolicySchema()
	err := APIRead(d, meta, "applicationinsightspolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApplicationInsightsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsPolicySchema()
	err := APICreate(d, meta, "applicationinsightspolicy", s)
	if err == nil {
		err = ResourceAviApplicationInsightsPolicyRead(d, meta)
	}
	return err
}

func resourceAviApplicationInsightsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsPolicySchema()
	var err error
	err = APIUpdate(d, meta, "applicationinsightspolicy", s)
	if err == nil {
		err = ResourceAviApplicationInsightsPolicyRead(d, meta)
	}
	return err
}

func resourceAviApplicationInsightsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "applicationinsightspolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
