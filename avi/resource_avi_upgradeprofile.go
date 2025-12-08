// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceUpgradeProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"controller_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerParamsSchema(),
		},
		"dry_run": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDryRunParamsSchema(),
		},
		"image": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceImageParamsSchema(),
		},
		"pre_checks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePreChecksParamsSchema(),
		},
		"service_engine": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceServiceEngineParamsSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviUpgradeProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUpgradeProfileCreate,
		Read:   ResourceAviUpgradeProfileRead,
		Update: resourceAviUpgradeProfileUpdate,
		Delete: resourceAviUpgradeProfileDelete,
		Schema: ResourceUpgradeProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUpgradeProfileImporter,
		},
	}
}

func ResourceUpgradeProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUpgradeProfileSchema()
	return ResourceImporter(d, m, "upgradeprofile", s)
}

func ResourceAviUpgradeProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeProfileSchema()
	err := APIRead(d, meta, "upgradeprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUpgradeProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeProfileSchema()
	err := APICreate(d, meta, "upgradeprofile", s)
	if err == nil {
		err = ResourceAviUpgradeProfileRead(d, meta)
	}
	return err
}

func resourceAviUpgradeProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeProfileSchema()
	var err error
	err = APIUpdate(d, meta, "upgradeprofile", s)
	if err == nil {
		err = ResourceAviUpgradeProfileRead(d, meta)
	}
	return err
}

func resourceAviUpgradeProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "upgradeprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
