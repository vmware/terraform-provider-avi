// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePriorityLabelsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"equivalent_labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEquivalentLabelsSchema(),
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
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPriorityLabels() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPriorityLabelsCreate,
		Read:   ResourceAviPriorityLabelsRead,
		Update: resourceAviPriorityLabelsUpdate,
		Delete: resourceAviPriorityLabelsDelete,
		Schema: ResourcePriorityLabelsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePriorityLabelsImporter,
		},
	}
}

func ResourcePriorityLabelsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePriorityLabelsSchema()
	return ResourceImporter(d, m, "prioritylabels", s)
}

func ResourceAviPriorityLabelsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := APIRead(d, meta, "prioritylabels", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPriorityLabelsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := APICreate(d, meta, "prioritylabels", s)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	return err
}

func resourceAviPriorityLabelsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	var err error
	err = APIUpdate(d, meta, "prioritylabels", s)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	return err
}

func resourceAviPriorityLabelsDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "prioritylabels")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
