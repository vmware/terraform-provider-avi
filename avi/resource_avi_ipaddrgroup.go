// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

// nolint
func ResourceIpAddrGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"country_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPortSchema(),
		},
		"marathon_app_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"marathon_service_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
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
		"prefixes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrRangeSchema(),
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

// nolint
func resourceAviIpAddrGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIpAddrGroupCreate,
		Read:   ResourceAviIpAddrGroupRead,
		Update: resourceAviIpAddrGroupUpdate,
		Delete: resourceAviIpAddrGroupDelete,
		Schema: ResourceIpAddrGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIpAddrGroupImporter,
		},
	}
}

// nolint
func ResourceIpAddrGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIpAddrGroupSchema()
	return ResourceImporter(d, m, "ipaddrgroup", s)
}

// nolint
func ResourceAviIpAddrGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := APIRead(d, meta, "ipaddrgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

// nolint
func resourceAviIpAddrGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	err := APICreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	return err
}

// nolint
func resourceAviIpAddrGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpAddrGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "ipaddrgroup", s)
	if err == nil {
		err = ResourceAviIpAddrGroupRead(d, meta)
	}
	return err
}

// nolint
func resourceAviIpAddrGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "ipaddrgroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
