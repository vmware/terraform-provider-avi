// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceBotIPReputationTypeMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_reputation_mappings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIPReputationTypeMappingSchema(),
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

func resourceAviBotIPReputationTypeMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBotIPReputationTypeMappingCreate,
		Read:   ResourceAviBotIPReputationTypeMappingRead,
		Update: resourceAviBotIPReputationTypeMappingUpdate,
		Delete: resourceAviBotIPReputationTypeMappingDelete,
		Schema: ResourceBotIPReputationTypeMappingSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBotIPReputationTypeMappingImporter,
		},
	}
}

func ResourceBotIPReputationTypeMappingImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBotIPReputationTypeMappingSchema()
	return ResourceImporter(d, m, "botipreputationtypemapping", s)
}

func ResourceAviBotIPReputationTypeMappingRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	err := APIRead(d, meta, "botipreputationtypemapping", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	err := APICreateOrUpdate(d, meta, "botipreputationtypemapping", s)
	if err == nil {
		err = ResourceAviBotIPReputationTypeMappingRead(d, meta)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "botipreputationtypemapping", s)
	if err == nil {
		err = ResourceAviBotIPReputationTypeMappingRead(d, meta)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "botipreputationtypemapping")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
