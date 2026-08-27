// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApiPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_api_labels": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiLabelsSchema(),
		},
		"api_spec_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiSpecInfoSchema(),
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
		"file_object_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"label_mappings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceApiPolicyLabelActionMappingSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"non_api_url_labels": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiLabelsSchema(),
		},
		"orphan_api_classification_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOrphanApiClassificationSettingsSchema(),
		},
		"orphan_api_labels": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiLabelsSchema(),
		},
		"path_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"routing_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiRoutingInfoSchema(),
		},
		"server_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiServerInfoSchema(),
		},
		"shadow_api_labels": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiLabelsSchema(),
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
		"validation_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiValidationSettingsSchema(),
		},
		"zombie_api_classification_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceZombieApiClassificationSettingsSchema(),
		},
		"zombie_api_labels": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiLabelsSchema(),
		},
	}
}

func resourceAviApiPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApiPolicyCreate,
		Read:   ResourceAviApiPolicyRead,
		Update: resourceAviApiPolicyUpdate,
		Delete: resourceAviApiPolicyDelete,
		Schema: ResourceApiPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApiPolicyImporter,
		},
	}
}

func ResourceApiPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApiPolicySchema()
	return ResourceImporter(d, m, "apipolicy", s)
}

func ResourceAviApiPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPolicySchema()
	err := APIRead(d, meta, "apipolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApiPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPolicySchema()
	err := APICreate(d, meta, "apipolicy", s)
	if err == nil {
		err = ResourceAviApiPolicyRead(d, meta)
	}
	return err
}

func resourceAviApiPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPolicySchema()
	var err error
	err = APIUpdate(d, meta, "apipolicy", s)
	if err == nil {
		err = ResourceAviApiPolicyRead(d, meta)
	}
	return err
}

func resourceAviApiPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "apipolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
