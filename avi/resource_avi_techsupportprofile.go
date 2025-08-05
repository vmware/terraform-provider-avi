// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTechSupportProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"archive_rules": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceArchiveRulesSchema(),
		},
		"event_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTechSupportEventParamsSchema(),
		},
		"file_size_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "128",
			ValidateFunc: validateInteger,
		},
		"max_disk_size_percent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"min_free_disk_required": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"no_of_techsupport_retentions": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"simultaneous_invocations": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"task_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "180",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTechSupportProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTechSupportProfileCreate,
		Read:   ResourceAviTechSupportProfileRead,
		Update: resourceAviTechSupportProfileUpdate,
		Delete: resourceAviTechSupportProfileDelete,
		Schema: ResourceTechSupportProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTechSupportProfileImporter,
		},
	}
}

func ResourceTechSupportProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTechSupportProfileSchema()
	return ResourceImporter(d, m, "techsupportprofile", s)
}

func ResourceAviTechSupportProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportProfileSchema()
	err := APIRead(d, meta, "techsupportprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTechSupportProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportProfileSchema()
	err := APICreate(d, meta, "techsupportprofile", s)
	if err == nil {
		err = ResourceAviTechSupportProfileRead(d, meta)
	}
	return err
}

func resourceAviTechSupportProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportProfileSchema()
	var err error
	err = APIUpdate(d, meta, "techsupportprofile", s)
	if err == nil {
		err = ResourceAviTechSupportProfileRead(d, meta)
	}
	return err
}

func resourceAviTechSupportProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "techsupportprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
