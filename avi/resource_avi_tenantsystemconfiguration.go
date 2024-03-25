// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTenantSystemConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"dns_virtualservice_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviTenantSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantSystemConfigurationCreate,
		Read:   ResourceAviTenantSystemConfigurationRead,
		Update: resourceAviTenantSystemConfigurationUpdate,
		Delete: resourceAviTenantSystemConfigurationDelete,
		Schema: ResourceTenantSystemConfigurationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTenantSystemConfigurationImporter,
		},
	}
}

func ResourceTenantSystemConfigurationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTenantSystemConfigurationSchema()
	return ResourceImporter(d, m, "tenantsystemconfiguration", s)
}

func ResourceAviTenantSystemConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSystemConfigurationSchema()
	err := APIRead(d, meta, "tenantsystemconfiguration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTenantSystemConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSystemConfigurationSchema()
	err := APICreate(d, meta, "tenantsystemconfiguration", s)
	if err == nil {
		err = ResourceAviTenantSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviTenantSystemConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSystemConfigurationSchema()
	var err error
	err = APIUpdate(d, meta, "tenantsystemconfiguration", s)
	if err == nil {
		err = ResourceAviTenantSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviTenantSystemConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "tenantsystemconfiguration")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
