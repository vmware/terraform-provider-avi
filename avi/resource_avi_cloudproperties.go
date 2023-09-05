// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudPropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_props": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCC_PropertiesSchema(),
		},
		"cc_vtypes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"hyp_props": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHypervisor_PropertiesSchema(),
		},
		"info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloudInfoSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviCloudProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudPropertiesCreate,
		Read:   ResourceAviCloudPropertiesRead,
		Update: resourceAviCloudPropertiesUpdate,
		Delete: resourceAviCloudPropertiesDelete,
		Schema: ResourceCloudPropertiesSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCloudPropertiesImporter,
		},
	}
}

func ResourceCloudPropertiesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCloudPropertiesSchema()
	return ResourceImporter(d, m, "cloudproperties", s)
}

func ResourceAviCloudPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := APIRead(d, meta, "cloudproperties", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCloudPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := APICreateOrUpdate(d, meta, "cloudproperties", s)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	return err
}

func resourceAviCloudPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "cloudproperties", s)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	return err
}

func resourceAviCloudPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "cloudproperties")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
