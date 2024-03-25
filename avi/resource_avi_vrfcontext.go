// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceVrfContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"bfd_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBfdProfileSchema(),
		},
		"bgp_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBgpProfileSchema(),
		},
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
		"debugvrfcontext": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDebugVrfContextSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gateway_mon": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGatewayMonitorSchema(),
		},
		"internal_gateway_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceInternalGatewayMonitorSchema(),
		},
		"lldp_enable": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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
		"static_routes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceStaticRouteSchema(),
		},
		"system_default": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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

func resourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVrfContextCreate,
		Read:   ResourceAviVrfContextRead,
		Update: resourceAviVrfContextUpdate,
		Delete: resourceAviVrfContextDelete,
		Schema: ResourceVrfContextSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVrfContextImporter,
		},
	}
}

func ResourceVrfContextImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVrfContextSchema()
	return ResourceImporter(d, m, "vrfcontext", s)
}

func ResourceAviVrfContextRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := APIRead(d, meta, "vrfcontext", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVrfContextCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := APICreate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	var err error
	err = APIUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "vrfcontext")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
