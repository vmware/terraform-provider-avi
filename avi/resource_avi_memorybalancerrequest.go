// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceMemoryBalancerRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"controller_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerInfoSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"node_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"process_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProcessInfoSchema(),
		},
		"process_instance": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"timestamp": {
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

func resourceAviMemoryBalancerRequest() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviMemoryBalancerRequestCreate,
		Read:   ResourceAviMemoryBalancerRequestRead,
		Update: resourceAviMemoryBalancerRequestUpdate,
		Delete: resourceAviMemoryBalancerRequestDelete,
		Schema: ResourceMemoryBalancerRequestSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceMemoryBalancerRequestImporter,
		},
	}
}

func ResourceMemoryBalancerRequestImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceMemoryBalancerRequestSchema()
	return ResourceImporter(d, m, "memorybalancerrequest", s)
}

func ResourceAviMemoryBalancerRequestRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMemoryBalancerRequestSchema()
	err := APIRead(d, meta, "memorybalancerrequest", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviMemoryBalancerRequestCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMemoryBalancerRequestSchema()
	err := APICreate(d, meta, "memorybalancerrequest", s)
	if err == nil {
		err = ResourceAviMemoryBalancerRequestRead(d, meta)
	}
	return err
}

func resourceAviMemoryBalancerRequestUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMemoryBalancerRequestSchema()
	var err error
	err = APIUpdate(d, meta, "memorybalancerrequest", s)
	if err == nil {
		err = ResourceAviMemoryBalancerRequestRead(d, meta)
	}
	return err
}

func resourceAviMemoryBalancerRequestDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "memorybalancerrequest")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
