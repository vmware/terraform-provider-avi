// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceFederationCheckpointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"date": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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

func resourceAviFederationCheckpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviFederationCheckpointCreate,
		Read:   ResourceAviFederationCheckpointRead,
		Update: resourceAviFederationCheckpointUpdate,
		Delete: resourceAviFederationCheckpointDelete,
		Schema: ResourceFederationCheckpointSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFederationCheckpointImporter,
		},
	}
}

func ResourceFederationCheckpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFederationCheckpointSchema()
	return ResourceImporter(d, m, "federationcheckpoint", s)
}

func ResourceAviFederationCheckpointRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	err := APIRead(d, meta, "federationcheckpoint", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviFederationCheckpointCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	err := APICreate(d, meta, "federationcheckpoint", s)
	if err == nil {
		err = ResourceAviFederationCheckpointRead(d, meta)
	}
	return err
}

func resourceAviFederationCheckpointUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	var err error
	err = APIUpdate(d, meta, "federationcheckpoint", s)
	if err == nil {
		err = ResourceAviFederationCheckpointRead(d, meta)
	}
	return err
}

func resourceAviFederationCheckpointDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "federationcheckpoint")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
