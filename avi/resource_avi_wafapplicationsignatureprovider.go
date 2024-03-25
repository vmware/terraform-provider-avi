// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceWafApplicationSignatureProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"available_applications": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafApplicationSignatureAppVersionSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"filter_rules_on_import": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ruleset_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWebApplicationSignatureServiceStatusSchema(),
		},
		"signatures": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleSchema(),
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

func resourceAviWafApplicationSignatureProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafApplicationSignatureProviderCreate,
		Read:   ResourceAviWafApplicationSignatureProviderRead,
		Update: resourceAviWafApplicationSignatureProviderUpdate,
		Delete: resourceAviWafApplicationSignatureProviderDelete,
		Schema: ResourceWafApplicationSignatureProviderSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafApplicationSignatureProviderImporter,
		},
	}
}

func ResourceWafApplicationSignatureProviderImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafApplicationSignatureProviderSchema()
	return ResourceImporter(d, m, "wafapplicationsignatureprovider", s)
}

func ResourceAviWafApplicationSignatureProviderRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := APIRead(d, meta, "wafapplicationsignatureprovider", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := APICreate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	var err error
	err = APIUpdate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "wafapplicationsignatureprovider")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
