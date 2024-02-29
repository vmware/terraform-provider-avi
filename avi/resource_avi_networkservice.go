// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceNetworkServiceSchema() map[string]*schema.Schema {
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
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"routing_service": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRoutingServiceSchema(),
		},
		"se_group_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"service_type": {
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
		"vrf_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviNetworkService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkServiceCreate,
		Read:   ResourceAviNetworkServiceRead,
		Update: resourceAviNetworkServiceUpdate,
		Delete: resourceAviNetworkServiceDelete,
		Schema: ResourceNetworkServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkServiceImporter,
		},
	}
}

func ResourceNetworkServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkServiceSchema()
	return ResourceImporter(d, m, "networkservice", s)
}

func ResourceAviNetworkServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	err := APIRead(d, meta, "networkservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	err := APICreate(d, meta, "networkservice", s)
	if err == nil {
		err = ResourceAviNetworkServiceRead(d, meta)
	}
	return err
}

func resourceAviNetworkServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	var err error
	err = APIUpdate(d, meta, "networkservice", s)
	if err == nil {
		err = ResourceAviNetworkServiceRead(d, meta)
	}
	return err
}

func resourceAviNetworkServiceDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "networkservice")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
