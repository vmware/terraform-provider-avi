// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceErrorPageBodySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"error_page_body": {
			Type:     schema.TypeString,
			Required: true,
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ERROR_PAGE_FORMAT_HTML",
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

func resourceAviErrorPageBody() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviErrorPageBodyCreate,
		Read:   ResourceAviErrorPageBodyRead,
		Update: resourceAviErrorPageBodyUpdate,
		Delete: resourceAviErrorPageBodyDelete,
		Schema: ResourceErrorPageBodySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceErrorPageBodyImporter,
		},
	}
}

func ResourceErrorPageBodyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceErrorPageBodySchema()
	return ResourceImporter(d, m, "errorpagebody", s)
}

func ResourceAviErrorPageBodyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := APIRead(d, meta, "errorpagebody", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviErrorPageBodyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := APICreate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	var err error
	err = APIUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "errorpagebody")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
