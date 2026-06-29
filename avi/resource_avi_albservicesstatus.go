// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceALBServicesStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_details": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceALBServicesAssetDetailsSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"connected_at": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTimeStampSchema(),
		},
		"connectivity_status": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ALBSERVICES_CONNECTIVITY_UNKNOWN",
		},
		"error": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"registration_status": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ALBSERVICES_REGISTRATION_UNKNOWN",
		},
		"services_health": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServiceHealthSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePulseServicesTenantStatusSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviALBServicesStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesStatusCreate,
		Read:   ResourceAviALBServicesStatusRead,
		Update: resourceAviALBServicesStatusUpdate,
		Delete: resourceAviALBServicesStatusDelete,
		Schema: ResourceALBServicesStatusSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesStatusImporter,
		},
	}
}

func ResourceALBServicesStatusImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesStatusSchema()
	return ResourceImporter(d, m, "albservicesstatus", s)
}

func ResourceAviALBServicesStatusRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesStatusSchema()
	err := APIRead(d, meta, "albservicesstatus", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesStatusCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesStatusSchema()
	err := APICreate(d, meta, "albservicesstatus", s)
	if err == nil {
		err = ResourceAviALBServicesStatusRead(d, meta)
	}
	return err
}

func resourceAviALBServicesStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesStatusSchema()
	var err error
	err = APIUpdate(d, meta, "albservicesstatus", s)
	if err == nil {
		err = ResourceAviALBServicesStatusRead(d, meta)
	}
	return err
}

func resourceAviALBServicesStatusDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "albservicesstatus")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
