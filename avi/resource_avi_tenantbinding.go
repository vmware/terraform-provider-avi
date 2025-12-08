// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTenantBindingSchema() map[string]*schema.Schema {
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
		"se_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"shared_tenant_ref": {
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

func resourceAviTenantBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantBindingCreate,
		Read:   ResourceAviTenantBindingRead,
		Update: resourceAviTenantBindingUpdate,
		Delete: resourceAviTenantBindingDelete,
		Schema: ResourceTenantBindingSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTenantBindingImporter,
		},
	}
}

func ResourceTenantBindingImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTenantBindingSchema()
	return ResourceImporter(d, m, "tenantbinding", s)
}

func ResourceAviTenantBindingRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantBindingSchema()
	err := APIRead(d, meta, "tenantbinding", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTenantBindingCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantBindingSchema()
	err := APICreate(d, meta, "tenantbinding", s)
	if err == nil {
		err = ResourceAviTenantBindingRead(d, meta)
	}
	return err
}

func resourceAviTenantBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantBindingSchema()
	var err error
	err = APIUpdate(d, meta, "tenantbinding", s)
	if err == nil {
		err = ResourceAviTenantBindingRead(d, meta)
	}
	return err
}

func resourceAviTenantBindingDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "tenantbinding")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
