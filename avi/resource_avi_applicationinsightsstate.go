// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApplicationInsightsStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_insights_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"application_sampling_runtime": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApplicationSamplingRuntimeSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviApplicationInsightsState() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationInsightsStateCreate,
		Read:   ResourceAviApplicationInsightsStateRead,
		Update: resourceAviApplicationInsightsStateUpdate,
		Delete: resourceAviApplicationInsightsStateDelete,
		Schema: ResourceApplicationInsightsStateSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApplicationInsightsStateImporter,
		},
	}
}

func ResourceApplicationInsightsStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApplicationInsightsStateSchema()
	return ResourceImporter(d, m, "applicationinsightsstate", s)
}

func ResourceAviApplicationInsightsStateRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsStateSchema()
	err := APIRead(d, meta, "applicationinsightsstate", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApplicationInsightsStateCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsStateSchema()
	err := APICreate(d, meta, "applicationinsightsstate", s)
	if err == nil {
		err = ResourceAviApplicationInsightsStateRead(d, meta)
	}
	return err
}

func resourceAviApplicationInsightsStateUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationInsightsStateSchema()
	var err error
	err = APIUpdate(d, meta, "applicationinsightsstate", s)
	if err == nil {
		err = ResourceAviApplicationInsightsStateRead(d, meta)
	}
	return err
}

func resourceAviApplicationInsightsStateDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "applicationinsightsstate")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
