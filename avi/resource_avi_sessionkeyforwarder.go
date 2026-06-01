// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSessionKeyForwarderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"enable": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"ip_ports": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceIpAddrPortSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pki_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_key_and_certificate_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_mgmt": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSessionKeyForwarder() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSessionKeyForwarderCreate,
		Read:   ResourceAviSessionKeyForwarderRead,
		Update: resourceAviSessionKeyForwarderUpdate,
		Delete: resourceAviSessionKeyForwarderDelete,
		Schema: ResourceSessionKeyForwarderSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSessionKeyForwarderImporter,
		},
	}
}

func ResourceSessionKeyForwarderImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSessionKeyForwarderSchema()
	return ResourceImporter(d, m, "sessionkeyforwarder", s)
}

func ResourceAviSessionKeyForwarderRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSessionKeyForwarderSchema()
	err := APIRead(d, meta, "sessionkeyforwarder", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSessionKeyForwarderCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSessionKeyForwarderSchema()
	err := APICreate(d, meta, "sessionkeyforwarder", s)
	if err == nil {
		err = ResourceAviSessionKeyForwarderRead(d, meta)
	}
	return err
}

func resourceAviSessionKeyForwarderUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSessionKeyForwarderSchema()
	var err error
	err = APIUpdate(d, meta, "sessionkeyforwarder", s)
	if err == nil {
		err = ResourceAviSessionKeyForwarderRead(d, meta)
	}
	return err
}

func resourceAviSessionKeyForwarderDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "sessionkeyforwarder")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
